package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/oracle/types"
)

// NewQuerier creates a new querier for nameservice clients.
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2

		case types.QueryExchangeRate:
			return queryExchangeRate(ctx, k, path[1], legacyQuerierCdc)

		case types.QueryAllExchangeRates:
			return queryAllExchangeRates(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown oracle query endpoint")
		}
	}
}

// queryExchangeRate - returns the exchange rate
func queryExchangeRate(ctx sdk.Context, keeper Keeper, pair string, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	exchangeRate, _ := keeper.GetExchangeRate(ctx, pair)

	//if !found {
	//	return nil, types.ErrNoValidatorFound
	//}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, exchangeRate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// queryExchangeRate - returns the exchange rate
func queryAllExchangeRates(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	exchangeRates := keeper.GetAllExchangeRates(ctx)

	//if !found {
	//	return nil, types.ErrNoValidatorFound
	//}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, exchangeRates)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
