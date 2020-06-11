package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/tendermint/iavl/common"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	tmcli "github.com/tendermint/tendermint/libs/cli"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"github.com/konstellation/kn-sdk/types"
	"github.com/konstellation/konstellation/common/utils"
)

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(
	ctx *server.Context,
	cdc *codec.Codec,
	mbm module.BasicManager,
	gus types.GenesisUpdaters,
	defaultNodeHome string,
) *cobra.Command { // nolint: golint
	initCmd := genutilcli.InitCmd(ctx, cdc, mbm, defaultNodeHome)

	cmd := &cobra.Command{
		Use:   initCmd.Use,
		Short: initCmd.Short,
		Long:  initCmd.Long,
		Args:  initCmd.Args,
		RunE: func(c *cobra.Command, args []string) error {
			err := initCmd.RunE(c, args)
			if err != nil {
				return err
			}

			config := ctx.Config
			config.SetRoot(viper.GetString(tmcli.HomeFlag))

			chainID := viper.GetString(flags.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
			}

			nodeID, _, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			genFile := config.GenesisFile()
			genDoc, err := tmtypes.GenesisDocFromFile(genFile)
			if err != nil {
				return err
			}

			var appState map[string]json.RawMessage
			err = cdc.UnmarshalJSON(genDoc.AppState, &appState)
			if err != nil {
				return err
			}

			// Update default genesis
			for _, gu := range gus {
				gu.UpdateGenesis(cdc, appState)
			}

			if err = mbm.ValidateGenesis(appState); err != nil {
				return fmt.Errorf("error validating genesis: %s", err.Error())
			}

			genDoc.AppState = cdc.MustMarshalJSON(appState)
			if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return errors.Wrap(err, "Failed to export genesis file")
			}

			toPrint := utils.NewPrintInfo(config.Moniker, chainID, nodeID, "", genDoc.AppState)
			return utils.DisplayInfo(cdc, toPrint)
		},
	}

	initCmd.Flags().VisitAll(func(f *pflag.Flag) {
		cmd.Flags().String(f.Name, f.Value.String(), f.Usage)
	})

	return cmd
}
