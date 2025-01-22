package merkle_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	cmtcrypto "github.com/cometbft/cometbft/api/cometbft/crypto/v1"
	"github.com/cometbft/cometbft/crypto/merkle"
)

type MockProofOperator struct {
	Key  []byte
	Root []byte
}

func (m *MockProofOperator) Run(leaves [][]byte) ([][]byte, error) {
	if len(leaves) == 0 {
		return nil, errors.New("leaves cannot be empty")
	}
	hash := append(leaves[0], m.Key...)
	return [][]byte{hash}, nil
}

func (m *MockProofOperator) GetKey() []byte {
	return m.Key
}

func (m *MockProofOperator) ProofOp() cmtcrypto.ProofOp {
	return cmtcrypto.ProofOp{
		Type: "mock",
		Data: m.Key,
	}
}

func TestVerifyValue(t *testing.T) {
	prt := merkle.NewProofRuntime()
	prt.RegisterOpDecoder("mock", func(op cmtcrypto.ProofOp) (merkle.ProofOperator, error) {
		return &MockProofOperator{
			Key: op.Data,
		}, nil
	})

	key := []byte("test-key")
	value := []byte("test-value")
	expectedHash := append(value, key...)

	root := expectedHash
	keypath := "/test-key"
	mockProofOp := &MockProofOperator{Key: key}

	proofOps := &cmtcrypto.ProofOps{
		Ops: []cmtcrypto.ProofOp{
			mockProofOp.ProofOp(),
		},
	}

	err := prt.VerifyValue(proofOps, root, keypath, value)
	require.NoError(t, err, "Verify Failed")
}

func TestProofOperators_Verify(t *testing.T) {
	op1 := &MockProofOperator{Key: []byte("key1")}
	op2 := &MockProofOperator{Key: []byte("key2")}

	proofOperators := merkle.ProofOperators{op1, op2}

	initialValue := []byte("initial-value")
	intermediateHash := append(initialValue, []byte("key1")...)
	expectedRoot := append(intermediateHash, []byte("key2")...)

	root := expectedRoot
	keypath := "/key2/key1"
	args := [][]byte{initialValue}

	err := proofOperators.Verify(root, keypath, args)

	require.NoError(t, err, "Verify Failed")
}

func TestProofOperatorsFromKeys(t *testing.T) {
	var err error

	// ProofRuntime setup
	// TODO test this somehow.

	// ProofOperators setup
	op1 := merkle.NewDominoOp("KEY1", "INPUT1", "INPUT2")
	op2 := merkle.NewDominoOp("KEY%2", "INPUT2", "INPUT3")
	op3 := merkle.NewDominoOp("", "INPUT3", "INPUT4")
	op4 := merkle.NewDominoOp("KEY/4", "INPUT4", "OUTPUT4")

	// add characters to the keys that would otherwise result in bad keypath if
	// processed
	keys1 := [][]byte{bz("KEY/4"), bz("KEY%2"), bz("KEY1")}
	badkeys1 := [][]byte{bz("WrongKey"), bz("KEY%2"), bz("KEY1")}
	keys2 := [][]byte{bz("KEY3"), bz("KEY%2"), bz("KEY1")}
	keys3 := [][]byte{bz("KEY2"), bz("KEY1")}

	// Good
	popz := merkle.ProofOperators([]merkle.ProofOperator{op1, op2, op3, op4})
	err = popz.VerifyFromKeys(bz("OUTPUT4"), keys1, [][]byte{bz("INPUT1")})
	require.NoError(t, err)

	// BAD INPUT
	err = popz.VerifyFromKeys(bz("OUTPUT4"), keys1, [][]byte{bz("INPUT1_WRONG")})
	require.Error(t, err)

	// BAD KEY 1
	err = popz.VerifyFromKeys(bz("OUTPUT4"), keys2, [][]byte{bz("INPUT1")})
	require.Error(t, err)

	// BAD KEY 2
	err = popz.VerifyFromKeys(bz("OUTPUT4"), badkeys1, [][]byte{bz("INPUT1")})
	require.Error(t, err)

	// BAD KEY 5
	err = popz.VerifyFromKeys(bz("OUTPUT4"), keys3, [][]byte{bz("INPUT1")})
	require.Error(t, err)

	// BAD OUTPUT 1
	err = popz.VerifyFromKeys(bz("OUTPUT4_WRONG"), keys1, [][]byte{bz("INPUT1")})
	require.Error(t, err)

	// BAD OUTPUT 2
	err = popz.VerifyFromKeys(bz(""), keys1, [][]byte{bz("INPUT1")})
	require.Error(t, err)

	// BAD POPZ 1
	popz = merkle.ProofOperators([]merkle.ProofOperator{op1, op2, op4})
	err = popz.VerifyFromKeys(bz("OUTPUT4"), keys1, [][]byte{bz("INPUT1")})
	require.Error(t, err)

	// BAD POPZ 2
	popz = merkle.ProofOperators([]merkle.ProofOperator{op4, op3, op2, op1})
	err = popz.VerifyFromKeys(bz("OUTPUT4"), keys1, [][]byte{bz("INPUT1")})
	require.Error(t, err)

	// BAD POPZ 3
	popz = merkle.ProofOperators([]merkle.ProofOperator{})
	err = popz.VerifyFromKeys(bz("OUTPUT4"), keys1, [][]byte{bz("INPUT1")})
	require.Error(t, err)
}

func bz(s string) []byte {
	return []byte(s)
}
