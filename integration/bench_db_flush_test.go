package integration

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/copyco6628/lachesis-base/abft"
	"github.com/copyco6628/lachesis-base/hash"
	"github.com/copyco6628/lachesis-base/inter/idx"
	"github.com/copyco6628/lachesis-base/kvdb"
	"github.com/copyco6628/lachesis-base/kvdb/leveldb"
	"github.com/copyco6628/lachesis-base/utils/cachescale"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/copyco6628/go-opera/gossip"
	"github.com/copyco6628/go-opera/integration/makefakegenesis"
	"github.com/copyco6628/go-opera/inter"
	"github.com/copyco6628/go-opera/utils"
	"github.com/copyco6628/go-opera/vecmt"
)

func BenchmarkFlushDBs(b *testing.B) {
	rawProducer, dir := dbProducer("flush_bench")
	defer os.RemoveAll(dir)
	genStore := makefakegenesis.FakeGenesisStore(1, utils.ToFtm(1), utils.ToFtm(1))
	g := genStore.Genesis()
	_, _, store, s2, _ := MakeEngine(rawProducer, &g, Configs{
		Opera:         gossip.DefaultConfig(cachescale.Identity),
		OperaStore:    gossip.DefaultStoreConfig(cachescale.Identity),
		Lachesis:      abft.DefaultConfig(),
		LachesisStore: abft.DefaultStoreConfig(cachescale.Identity),
		VectorClock:   vecmt.DefaultConfig(cachescale.Identity),
	})
	defer store.Close()
	defer s2.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := idx.Block(0)
		randUint32s := func() []uint32 {
			arr := make([]uint32, 128)
			for i := 0; i < len(arr); i++ {
				arr[i] = uint32(i) ^ (uint32(n) << 16) ^ 0xd0ad884e
			}
			return []uint32{uint32(n), uint32(n) + 1, uint32(n) + 2}
		}
		for !store.IsCommitNeeded() {
			store.SetBlock(n, &inter.Block{
				Time:        inter.Timestamp(n << 32),
				Atropos:     hash.Event{},
				Events:      hash.Events{},
				Txs:         []common.Hash{},
				InternalTxs: []common.Hash{},
				SkippedTxs:  randUint32s(),
				GasUsed:     uint64(n) << 24,
				Root:        hash.Hash{},
			})
			n++
		}
		err := store.Commit()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func cache64mb(string) int {
	return 64 * opt.MiB
}

func dbProducer(name string) (kvdb.IterableDBProducer, string) {
	dir, err := ioutil.TempDir("", name)
	if err != nil {
		panic(err)
	}
	return leveldb.NewProducer(dir, cache64mb), dir
}
