package evmstore

import (
	"github.com/copyco6628/lachesis-base/kvdb"

	"github.com/copyco6628ra/opera/genesis"
)

// ApplyGenesis writes initial state.
func (s *Store) ApplyGenesis(g genesis.Genesis) (err error) {
	batch := s.EvmDb.NewBatch()
	defer batch.Reset()
	g.RawEvmItems.ForEach(func(key, value []byte) bool {
		if err != nil {
			return false
		}
		err = batch.Put(key, value)
		if err != nil {
			return false
		}
		if batch.ValueSize() > kvdb.IdealBatchSize {
			err = batch.Write()
			if err != nil {
				return false
			}
			batch.Reset()
		}
		return true
	})
	if err != nil {
		return err
	}
	return batch.Write()
}
