package state

import (
	"github.com/celestiaorg/optimint/store"
	"testing"

	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmstate "github.com/tendermint/tendermint/proto/tendermint/state"

	"github.com/stretchr/testify/assert"
)

/*
func TestRestart(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	kv := store.NewDefaultInMemoryKVStore()
	s1 := New(kv)
	expectedHeight := uint64(10)
	//block := getRandomBlock(expectedHeight, 10)
	//err := s1.SaveBlock(block, &types.Commit{Height: block.Header.Height, HeaderHash: block.Header.Hash()})
	err := s1.UpdateState(State{
		LastBlockHeight: int64(expectedHeight),
	})
	assert.NoError(err)

	s2 := New(kv)
	_, err = s2.LoadState()
	assert.NoError(err)

	assert.Equal(expectedHeight, s2.Height())
}
*/

func TestBlockResponses(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	kv := store.NewDefaultInMemoryKVStore()
	s := New(kv)

	expected := &tmstate.ABCIResponses{
		BeginBlock: &abcitypes.ResponseBeginBlock{
			Events: []abcitypes.Event{{
				Type: "test",
				Attributes: []abcitypes.EventAttribute{{
					Key:   []byte("foo"),
					Value: []byte("bar"),
					Index: false,
				}},
			}},
		},
		DeliverTxs: nil,
		EndBlock: &abcitypes.ResponseEndBlock{
			ValidatorUpdates: nil,
			ConsensusParamUpdates: &abcitypes.ConsensusParams{
				Block: &abcitypes.BlockParams{
					MaxBytes: 12345,
					MaxGas:   678909876,
				},
			},
		},
	}

	err := s.SaveBlockResponses(1, expected)
	assert.NoError(err)

	resp, err := s.LoadBlockResponses(123)
	assert.Error(err)
	assert.Nil(resp)

	resp, err = s.LoadBlockResponses(1)
	assert.NoError(err)
	assert.NotNil(resp)
	assert.Equal(expected, resp)
}
