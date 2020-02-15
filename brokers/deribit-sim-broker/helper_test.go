package deribit_sim_broker

import (
	"github.com/coinrust/gotrader/models"
	"testing"
)

func TestCalcPnl(t *testing.T) {
	size := 50.0
	entryPrice := 10351.5
	exitPrice := 10348.5
	pnl, pnlUsd := CalcPnl(models.Buy, size, entryPrice, exitPrice)
	t.Logf("pnl: %.8f", pnl)
	t.Logf("pnlUsd: %.8f", pnlUsd)
}
