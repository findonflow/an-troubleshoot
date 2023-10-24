package main

import (
	"context"
	"time"

	"github.com/bjartek/overflow"
	"go.uber.org/zap"
)

func main() {
	start := time.Now()
	o := overflow.Overflow(overflow.WithNetwork("mainnet"))

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	ctx := context.Background()

	// latest, _ := o.GetLatestBlock(ctx)
	height := uint64(63723171)

	block, err := o.GetBlockAtHeight(ctx, height)
	if err != nil {
		logger.Fatal("did not get block", zap.Error(err))
	}

	tx, txR, err := o.Flowkit.GetTransactionsByBlockID(ctx, block.ID)
	if err != nil {
		logger.Fatal("did not get tx results", zap.Error(err))
	}

	taken := time.Since(start)
	logger.Info("done", zap.Int("tx", len(tx)), zap.Int("txr", len(txR)), zap.Duration("time", taken))
}
