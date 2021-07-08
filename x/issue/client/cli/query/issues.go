package query

import (
	"github.com/konstellation/konstellation/x/issue/query"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

const (
	flagOwner = "owner"
	flagLimit = "limit"
)

// getCmdQueryIssues implements the query issue command.
func getQueryCmdIssues(cdc *codec.LegacyAmino) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query issue list",
		Long:  "Query all or one of the account issue list, the limit default is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := client.Context{}
			cliCtx := ctx.WithLegacyAmino(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(flagOwner))
			if err != nil {
				return err
			}
			qp := types.NewIssuesParams(
				address.String(),
				int32(viper.GetInt(flagLimit)),
			)

			bz, err := cliCtx.LegacyAmino.MarshalJSON(qp)
			if err != nil {
				return err
			}

			// Query the issues
			res, _, err := cliCtx.QueryWithData(query.PathQueryIssues(), bz)
			if err != nil {
				return err
			}

			var issues types.CoinIssues
			cdc.MustUnmarshalJSON(res, &issues)
			return cliCtx.PrintObjectLegacy(issues)
		},
	}

	cmd.Flags().String(flagOwner, "", "Token owner address")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")
	return cmd
}
