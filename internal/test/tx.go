package test

import "github.com/cometbft/cometbft/types"

func MakeNTxs(height, n int64) types.Txs {
	txs := make([]types.Tx, n)
	for i := range txs {
		txs[i] = types.Tx([]byte{byte(height), byte(i / 256), byte(i % 256)})
	}
	return txs
}

func MakeTenTxs(height int64) (txs []types.Tx) {
	return MakeNTxs(height, 10)
}

func MakeData(txs []types.Tx) types.Data {
	return types.Data{
		Txs: txs,
	}
}
