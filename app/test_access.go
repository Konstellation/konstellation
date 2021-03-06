package app

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	ibctransferkeeper "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/keeper"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type TestSupport struct {
	t   *testing.T
	app *KonstellationApp
}

func NewTestSupport(t *testing.T, app *KonstellationApp) *TestSupport {
	return &TestSupport{t: t, app: app}
}

func (s TestSupport) IBCKeeper() ibckeeper.Keeper {
	return *s.app.ibcKeeper
}

//func (s TestSupport) WasmKeeper() wasm.Keeper {
//	return s.app.wasmKeeper
//}

func (s TestSupport) AppCodec() codec.Marshaler {
	return s.app.appCodec
}

func (s TestSupport) ScopeIBCKeeper() capabilitykeeper.ScopedKeeper {
	return s.app.scopedIBCKeeper
}

func (s TestSupport) ScopedTransferKeeper() capabilitykeeper.ScopedKeeper {
	return s.app.scopedTransferKeeper
}

func (s TestSupport) StakingKeeper() stakingkeeper.Keeper {
	return s.app.stakingKeeper
}

func (s TestSupport) BankKeeper() bankkeeper.Keeper {
	return s.app.bankKeeper
}

func (s TestSupport) TransferKeeper() ibctransferkeeper.Keeper {
	return s.app.transferKeeper
}
