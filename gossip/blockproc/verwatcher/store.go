package verwatcher

import (
	"github.com/copyco6628/lachesis-base/kvdb"

	"github.com/copyco6628/go-opera/logger"
)

// Store is a node persistent storage working over physical key-value database.
type Store struct {
	mainDB kvdb.Store

	logger.Instance
}

// NewStore creates store over key-value db.
func NewStore(mainDB kvdb.Store) *Store {
	s := &Store{
		mainDB:   mainDB,
		Instance: logger.New("verwatcher-store"),
	}

	return s
}
