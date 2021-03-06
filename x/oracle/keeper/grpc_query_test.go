package keeper_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle/types"
)

type OracleTestSuite struct {
	suite.Suite

	app         *app.KonstellationApp
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *OracleTestSuite) SetupTest() {
	simapp := app.Setup(false)
	simapp.Commit()
	ctx := simapp.NewContext(true, tmproto.Header{})
	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, []string{"abc"})
	simapp.GetOracleKeeper().SetExchangeRate(ctx, "abc", coin)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, codectypes.NewInterfaceRegistry())
	types.RegisterQueryServer(queryHelper, simapp.GetOracleKeeper())
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = simapp
	suite.ctx = ctx

	suite.queryClient = queryClient
}

func (suite *OracleTestSuite) TestGRPCExchangeRate() {
	app, ctx, queryClient := suite.app, suite.ctx, suite.queryClient

	exchangeRate, err := queryClient.ExchangeRate(gocontext.Background(), &types.QueryExchangeRateRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(exchangeRate.ExchangeRate, app.GetOracleKeeper().GetExchangeRate(ctx))
}

func TestOracleTestSuite(t *testing.T) {
	suite.Run(t, new(OracleTestSuite))
}
