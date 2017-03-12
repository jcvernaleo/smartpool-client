package geth

import (
	"../"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

// transaction pool keeps track of pending transactions
// and acknowledge corresponding channel when a transaction is
// confirmed
type TxWatcher struct {
	tx      *types.Transaction
	geth    *ethereum.GethRPC
	verChan chan bool
}

func (tw *TxWatcher) isVerified() bool {
	return tw.geth.IsVerified(tw.tx.Hash())
}

// loop to check transactions verification
// if a transaction is verified, send it to verChan
func (tw *TxWatcher) loop() {
	for {
		if tw.isVerified() {
			tw.verChan <- true
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func (tw *TxWatcher) Wait() {
	go tw.loop()
	<-tw.verChan
}

func NewTxWatcher(tx *types.Transaction, geth *ethereum.GethRPC) *TxWatcher {
	return &TxWatcher{tx, geth, make(chan bool)}
}
