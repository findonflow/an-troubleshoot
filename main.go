package main

import (
	"context"
	"time"

	"github.com/bjartek/overflow"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func main() {
	o := overflow.Overflow(overflow.WithNetwork("mainnet"))

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	ctx := context.Background()

	// latest, _ := o.GetLatestBlock(ctx)
	height := uint64(63800000)

	rng := lo.RangeFrom(height, 10000)

	for _, h := range rng {
		start := time.Now()
		block, err := o.GetBlockAtHeight(ctx, h)
		if err != nil {
			logger.Fatal("did not get block", zap.Error(err))
		}

		tx, txR, err := o.Flowkit.GetTransactionsByBlockID(ctx, block.ID)
		if err != nil {
			logger.Fatal("did not get tx results", zap.Error(err))
		}

		taken := time.Since(start)
		logger.Info("get block", zap.Uint64("height", h), zap.Int("tx", len(tx)), zap.Int("txr", len(txR)), zap.Duration("time", taken))
	}
	logger.Info("done")
}
