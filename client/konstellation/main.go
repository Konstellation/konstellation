package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/konstellation/kn-sdk/types"
	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/cmd"
)

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	types.RegisterNativeCoinUnits()
	types.RegisterBech32Prefix()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "konstellation",
		Short:             "Konstellation App Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		cmd.InitCmd(ctx, cdc, app.ModuleBasics, app.GenesisUpdaters, app.DefaultNodeHome),
		cmd.ConfigCmd(ctx),
		cmd.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{}, auth.GenesisAccountIterator{}, app.DefaultNodeHome, app.DefaultCLIHome),
		genutilcli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, app.DefaultNodeHome),
		//cmd.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome),
		genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics),
		cmd.AppStatusCmd(ctx, cdc, app.ModuleBasics),
		cmd.AppVersionCmd(ctx, cdc),
	)

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, app.EnvPrefixNode, app.DefaultNodeHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, _ io.Writer) abci.Application {
	return app.NewKonstellationApp(logger, db, uint(1))
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, _ io.Writer, height int64, forZeroHeight bool,
	jailWhiteList []string) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	if height != -1 {
		kApp := app.NewKonstellationApp(logger, db, uint(1))
		err := kApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}

		return kApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	kApp := app.NewKonstellationApp(logger, db, uint(1))

	return kApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
