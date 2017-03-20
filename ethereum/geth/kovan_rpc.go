package geth

import (
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"sync"
	"time"
)

var BlockTime = 4 * time.Second

type KovanRPC struct {
	*GethRPC
	mu            sync.Mutex
	lastTimestamp *big.Int
	lastTime      time.Time
}

func (k *KovanRPC) GetWork() *ethereum.Work {
	var h *types.Header
	h = k.GethRPC.GetPendingBlockHeader()
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.lastTimestamp != nil {
		if time.Since(k.lastTime) > BlockTime {
			h.Time.Add(k.lastTimestamp, big.NewInt(1))
		}
	}
	if k.lastTimestamp == nil || k.lastTimestamp.Cmp(h.Time) != 0 {
		k.lastTimestamp = big.NewInt(0)
		k.lastTimestamp.Set(h.Time)
		k.lastTime = time.Now()
	}
	seedHash, err := ethash.GetSeedHash(uint64(h.Number.Int64()))
	if err != nil {
		panic(err)
	}
	seed := common.BytesToHash(seedHash).Hex()
	return ethereum.NewWork(h, h.HashNoNonce().Hex(), seed, k.ShareDifficulty)
}

// never submit solution to the node because in Kovan, miners can't propose blocks
func (k *KovanRPC) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	return false
}

func NewKovanRPC(endpoint, contractAddr, extraData string, diff *big.Int) (*KovanRPC, error) {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	return &KovanRPC{
		&GethRPC{client, common.HexToAddress(contractAddr), []byte(extraData), diff},
		sync.Mutex{},
		nil,
		time.Now(),
	}, nil
}
