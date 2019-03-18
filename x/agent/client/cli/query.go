package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/comrade-coop/scynet-blockchain/x/agent"
)

// GetCmdGetAgent queries information about an agent
func GetCmdGetAgent(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-agent [agent]",
		Short: "get-agent agent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			id, err := uuid.Parse(args[0])
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get_agent/%s", queryRoute, id), nil)
			if err != nil {
				return err
			}

			var out agent.Agent
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
