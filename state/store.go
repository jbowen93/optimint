package state

import (
	"encoding/binary"
	"encoding/json"
	"github.com/celestiaorg/optimint/store"
	tmstate "github.com/tendermint/tendermint/proto/tendermint/state"
	"sync"
)

var (
	statePrefix     = [1]byte{4}
	responsesPrefix = [1]byte{5}
)

type Store interface {
	// SaveBlockResponses saves block responses (events, tx responses, validator set updates, etc) in Store.
	SaveBlockResponses(height uint64, responses *tmstate.ABCIResponses) error

	// LoadBlockResponses returns block results at given height, or error if it's not found in Store.
	LoadBlockResponses(height uint64) (*tmstate.ABCIResponses, error)

	// UpdateState updates state saved in Store. Only one State is stored.
	// If there is no State in Store, state will be saved.
	UpdateState(state State) error
	// LoadState returns last state saved with UpdateState.
	LoadState() (State, error)
}

type DefaultStore struct {
	db store.KVStore

	height uint64

	// mtx ensures that db is in sync with height
	mtx sync.RWMutex
}

var _ Store = &DefaultStore{}

// New returns new, default store.
func New(kv store.KVStore) Store {
	return &DefaultStore{
		db: kv,
	}
}

// Height returns height of the highest block saved in the Store.
func (s *DefaultStore) Height() uint64 {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.height
}

// SaveBlockResponses saves block responses (events, tx responses, validator set updates, etc) in Store.
func (s *DefaultStore) SaveBlockResponses(height uint64, responses *tmstate.ABCIResponses) error {
	data, err := responses.Marshal()
	if err != nil {
		return err
	}
	return s.db.Set(getResponsesKey(height), data)
}

// LoadBlockResponses returns block results at given height, or error if it's not found in Store.
func (s *DefaultStore) LoadBlockResponses(height uint64) (*tmstate.ABCIResponses, error) {
	data, err := s.db.Get(getResponsesKey(height))
	if err != nil {
		return nil, err
	}
	var responses tmstate.ABCIResponses
	err = responses.Unmarshal(data)
	return &responses, err
}

// UpdateState updates state saved in Store. Only one State is stored.
// If there is no State in Store, state will be saved.
func (s *DefaultStore) UpdateState(state State) error {
	blob, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return s.db.Set(getStateKey(), blob)
}

// LoadState returns last state saved with UpdateState.
func (s *DefaultStore) LoadState() (State, error) {
	var state State

	blob, err := s.db.Get(getStateKey())
	if err != nil {
		return state, err
	}

	err = json.Unmarshal(blob, &state)
	s.mtx.Lock()
	s.height = uint64(state.LastBlockHeight)
	s.mtx.Unlock()
	return state, err
}

func getStateKey() []byte {
	return statePrefix[:]
}

func getResponsesKey(height uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, height)
	return append(responsesPrefix[:], buf[:]...)
}
