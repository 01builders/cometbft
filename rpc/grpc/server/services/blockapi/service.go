package blockapi

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	crypto "github.com/cometbft/cometbft/api/cometbft/crypto/v1"
	blockapisvc "github.com/cometbft/cometbft/api/cometbft/services/blockAPI/v1"
	cmttypes "github.com/cometbft/cometbft/api/cometbft/types/v1"
	"github.com/cometbft/cometbft/crypto/encoding"
	"github.com/cometbft/cometbft/internal/rand"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/pubsub"
	"github.com/cometbft/cometbft/rpc/core"
	"github.com/cometbft/cometbft/types"
)

type BlockAPI struct {
	sync.Mutex
	env                  *core.Environment
	heightListeners      map[chan blockapisvc.NewHeightEvent]struct{}
	newBlockSubscription types.Subscription
	eventBus             *types.EventBus
	subscriptionID       string
	subscriptionQuery    pubsub.Query
	logger               log.Logger
}

func New(env *core.Environment) *BlockAPI {
	return &BlockAPI{
		heightListeners:   make(map[chan blockapisvc.NewHeightEvent]struct{}, 1000),
		subscriptionID:    fmt.Sprintf("block-api-subscription-%s", rand.Str(6)),
		subscriptionQuery: types.EventQueryNewBlock,
		eventBus:          env.EventBus,
		env:               env,
	}
}

func (blockAPI *BlockAPI) StartNewBlockEventListener(ctx context.Context) error {
	if blockAPI.newBlockSubscription == nil {
		var err error
		blockAPI.newBlockSubscription, err = blockAPI.eventBus.Subscribe(
			ctx,
			blockAPI.subscriptionID,
			blockAPI.subscriptionQuery,
			500,
		)
		if err != nil {
			blockAPI.logger.Error("Failed to subscribe to new blocks", "err", err)
			return err
		}
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-blockAPI.newBlockSubscription.Canceled():
			blockAPI.logger.Error("cancelled grpc subscription. retrying")
			ok, err := blockAPI.retryNewBlocksSubscription(ctx)
			if err != nil {
				return err
			}
			if !ok {
				// this will happen when the context is done. we can stop here
				return nil
			}
		case event, ok := <-blockAPI.newBlockSubscription.Out():
			if !ok {
				blockAPI.logger.Error("new blocks subscription closed. re-subscribing")
				ok, err := blockAPI.retryNewBlocksSubscription(ctx)
				if err != nil {
					return err
				}
				if !ok {
					// this will happen when the context is done. we can stop here
					return nil
				}
				continue
			}
			newBlockEvent, ok := event.Events()[types.EventTypeKey]
			if !ok || len(newBlockEvent) == 0 || newBlockEvent[0] != types.EventNewBlock {
				continue
			}
			data, ok := event.Data().(types.EventDataNewBlock)
			if !ok {
				blockAPI.logger.Error("couldn't cast event data to new block")
				return fmt.Errorf("couldn't cast event data to new block. Events: %s", event.Events())
			}
			blockAPI.broadcastToListeners(ctx, data.Block.Height, data.Block.Hash())
		}
	}
}

// RetryAttempts the number of retry times when the subscription is closed.
const RetryAttempts = 6

// SubscriptionCapacity the maximum number of pending blocks in the subscription.
const SubscriptionCapacity = 500

func (blockAPI *BlockAPI) retryNewBlocksSubscription(ctx context.Context) (bool, error) {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	blockAPI.Lock()
	defer blockAPI.Unlock()
	for i := 1; i < RetryAttempts; i++ {
		select {
		case <-ctx.Done():
			return false, nil
		case <-ticker.C:
			var err error
			blockAPI.newBlockSubscription, err = blockAPI.eventBus.Subscribe(
				ctx,
				fmt.Sprintf("block-api-subscription-%s", rand.Str(6)),
				blockAPI.subscriptionQuery,
				SubscriptionCapacity,
			)
			if err != nil {
				blockAPI.logger.Error("Failed to subscribe to new blocks. retrying", "err", err, "retry_number", i)
			} else {
				return true, nil
			}
		}
	}
	return false, errors.New("couldn't recover from failed blocks subscription. stopping listeners")
}

func (blockAPI *BlockAPI) broadcastToListeners(ctx context.Context, height int64, hash []byte) {
	blockAPI.Lock()
	defer blockAPI.Unlock()
	for ch := range blockAPI.heightListeners {
		func() {
			defer func() {
				if r := recover(); r != nil {
					// logging the error then removing the heights listener
					blockAPI.logger.Debug("failed to write to heights listener", "err", r)
					blockAPI.removeHeightListener(ch)
				}
			}()
			select {
			case <-ctx.Done():
				return
			case ch <- blockapisvc.NewHeightEvent{Height: height, Hash: hash}:
			}
		}()
	}
}

func (blockAPI *BlockAPI) addHeightListener() chan blockapisvc.NewHeightEvent {
	blockAPI.Lock()
	defer blockAPI.Unlock()
	ch := make(chan blockapisvc.NewHeightEvent, 50)
	blockAPI.heightListeners[ch] = struct{}{}
	return ch
}

func (blockAPI *BlockAPI) removeHeightListener(ch chan blockapisvc.NewHeightEvent) {
	blockAPI.Lock()
	defer blockAPI.Unlock()
	delete(blockAPI.heightListeners, ch)
}

func (blockAPI *BlockAPI) closeAllListeners() {
	blockAPI.Lock()
	defer blockAPI.Unlock()
	if blockAPI.heightListeners == nil {
		// if this is nil, then there is no need to close anything
		return
	}
	for channel := range blockAPI.heightListeners {
		delete(blockAPI.heightListeners, channel)
	}
}

// Stop cleans up the BlockAPI instance by closing all listeners
// and ensuring no further events are processed.
func (blockAPI *BlockAPI) Stop(ctx context.Context) error {
	blockAPI.Lock()
	defer blockAPI.Unlock()

	// close all height listeners
	blockAPI.closeAllListeners()

	var err error
	// stop the events subscription
	if blockAPI.newBlockSubscription != nil {
		err = blockAPI.eventBus.Unsubscribe(ctx, blockAPI.subscriptionID, blockAPI.subscriptionQuery)
		blockAPI.newBlockSubscription = nil
	}

	blockAPI.logger.Info("gRPC streaming API has been stopped")
	return err
}

func (blockAPI *BlockAPI) BlockByHash(req *blockapisvc.BlockByHashRequest, stream blockapisvc.BlockAPI_BlockByHashServer) error {
	blockStore := blockAPI.env.BlockStore
	blockMeta := blockStore.LoadBlockMetaByHash(req.Hash)
	if blockMeta == nil {
		return fmt.Errorf("nil block meta for block hash %d", req.Hash)
	}
	commit := blockStore.LoadBlockCommit(blockMeta.Header.Height)
	if commit == nil {
		return fmt.Errorf("nil commit for block hash %d", req.Hash)
	}
	protoCommit := commit.ToProto()

	validatorSet, err := blockAPI.env.StateStore.LoadValidators(blockMeta.Header.Height)
	if err != nil {
		return err
	}
	protoValidatorSet, err := validatorSet.ToProto()
	if err != nil {
		return err
	}

	for i := 0; i < int(blockMeta.BlockID.PartSetHeader.Total); i++ {
		part, err := blockStore.LoadBlockPart(blockMeta.Header.Height, i).ToProto()
		if err != nil {
			return err
		}
		if part == nil {
			return fmt.Errorf("nil block part %d for block hash %d", i, req.Hash)
		}
		if !req.Prove {
			part.Proof = crypto.Proof{}
		}
		isLastPart := i == int(blockMeta.BlockID.PartSetHeader.Total)-1
		resp := blockapisvc.StreamedBlockByHashResponse{
			BlockPart: part,
			IsLast:    isLastPart,
		}
		if i == 0 {
			resp.ValidatorSet = protoValidatorSet
			resp.Commit = protoCommit
		}
		err = stream.Send(&resp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (blockAPI *BlockAPI) BlockByHeight(req *blockapisvc.BlockByHeightRequest, stream blockapisvc.BlockAPI_BlockByHeightServer) error {
	blockStore := blockAPI.env.BlockStore
	height := req.Height
	if height == 0 {
		height = blockStore.Height()
	}

	blockMeta := blockStore.LoadBlockMeta(height)
	if blockMeta == nil {
		return fmt.Errorf("nil block meta for height %d", height)
	}

	commit := blockStore.LoadSeenCommit(height)
	if commit == nil {
		return fmt.Errorf("nil block commit for height %d", height)
	}
	protoCommit := commit.ToProto()

	validatorSet, err := blockAPI.env.StateStore.LoadValidators(height)
	if err != nil {
		return err
	}
	protoValidatorSet, err := validatorSet.ToProto()
	if err != nil {
		return err
	}

	for i := 0; i < int(blockMeta.BlockID.PartSetHeader.Total); i++ {
		part, err := blockStore.LoadBlockPart(height, i).ToProto()
		if err != nil {
			return err
		}
		if part == nil {
			return fmt.Errorf("nil block part %d for height %d", i, height)
		}
		if !req.Prove {
			part.Proof = crypto.Proof{}
		}
		isLastPart := i == int(blockMeta.BlockID.PartSetHeader.Total)-1
		resp := blockapisvc.StreamedBlockByHeightResponse{
			BlockPart: part,
			IsLast:    isLastPart,
		}
		if i == 0 {
			resp.ValidatorSet = protoValidatorSet
			resp.Commit = protoCommit
		}
		err = stream.Send(&resp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (blockAPI *BlockAPI) Status(_ context.Context, _ *blockapisvc.StatusRequest) (*blockapisvc.StatusResponse, error) {
	status, err := blockAPI.env.Status(nil)
	if err != nil {
		return nil, err
	}

	protoPubKey, err := encoding.PubKeyToProto(status.ValidatorInfo.PubKey)
	if err != nil {
		return nil, err
	}
	return &blockapisvc.StatusResponse{
		NodeInfo: status.NodeInfo.ToProto(),
		SyncInfo: &blockapisvc.SyncInfo{
			LatestBlockHash:     status.SyncInfo.LatestBlockHash,
			LatestAppHash:       status.SyncInfo.LatestAppHash,
			LatestBlockHeight:   status.SyncInfo.LatestBlockHeight,
			LatestBlockTime:     status.SyncInfo.LatestBlockTime,
			EarliestBlockHash:   status.SyncInfo.EarliestBlockHash,
			EarliestAppHash:     status.SyncInfo.EarliestAppHash,
			EarliestBlockHeight: status.SyncInfo.EarliestBlockHeight,
			EarliestBlockTime:   status.SyncInfo.EarliestBlockTime,
			CatchingUp:          status.SyncInfo.CatchingUp,
		},
		ValidatorInfo: &blockapisvc.ValidatorInfo{
			Address:     status.ValidatorInfo.Address,
			PubKey:      &protoPubKey,
			VotingPower: status.ValidatorInfo.VotingPower,
		},
	}, nil
}

func (blockAPI *BlockAPI) Commit(_ context.Context, req *blockapisvc.CommitRequest) (*blockapisvc.CommitResponse, error) {
	blockStore := blockAPI.env.BlockStore
	height := req.Height
	if height == 0 {
		height = blockStore.Height()
	}
	commit := blockStore.LoadSeenCommit(height)
	if commit == nil {
		return nil, fmt.Errorf("nil block commit for height %d", height)
	}
	protoCommit := commit.ToProto()

	return &blockapisvc.CommitResponse{
		Commit: &cmttypes.Commit{
			Height:     protoCommit.Height,
			Round:      protoCommit.Round,
			BlockID:    protoCommit.BlockID,
			Signatures: protoCommit.Signatures,
		},
	}, nil
}

func (blockAPI *BlockAPI) ValidatorSet(_ context.Context, req *blockapisvc.ValidatorSetRequest) (*blockapisvc.ValidatorSetResponse, error) {
	blockStore := blockAPI.env.BlockStore
	height := req.Height
	if height == 0 {
		height = blockStore.Height()
	}
	validatorSet, err := blockAPI.env.StateStore.LoadValidators(height)
	if err != nil {
		return nil, err
	}
	protoValidatorSet, err := validatorSet.ToProto()
	if err != nil {
		return nil, err
	}
	return &blockapisvc.ValidatorSetResponse{
		ValidatorSet: protoValidatorSet,
		Height:       height,
	}, nil
}

func (blockAPI *BlockAPI) SubscribeNewHeights(_ *blockapisvc.SubscribeNewHeightsRequest, stream blockapisvc.BlockAPI_SubscribeNewHeightsServer) error {
	heightListener := blockAPI.addHeightListener()
	defer blockAPI.removeHeightListener(heightListener)

	for {
		select {
		case event, ok := <-heightListener:
			if !ok {
				return errors.New("blocks subscription closed from the service side")
			}
			if err := stream.Send(&event); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}
