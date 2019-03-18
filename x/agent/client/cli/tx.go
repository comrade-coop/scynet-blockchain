package cli

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/comrade-coop/scynet/blockchain/x/agent"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdPublishAgentPrice is the CLI command for sending a PublishAgentPrice transaction
func GetCmdPublishAgentPrice(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "publish-price [uuid] [amount]",
		Short: "publish the price for an agent",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			id, err := uuid.Parse(args[0])
			if err != nil {
				return err
			}

			msg := agent.MsgPublishAgentPrice{
				cliCtx.GetFromAddress(),
				id,
				agent.Subscription,
				coins,
			}
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}
