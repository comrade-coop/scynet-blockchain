package facade

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"

	"github.com/comrade-coop/scynet/blockchain/x/agent"
)

// Context stores common parameters
type Context struct {
	cdc       *codec.Codec
	txBldr    authtxb.TxBuilder
	cliCtx    context.CLIContext
	storeName string
}

// NewContext creates the context for this facade
func NewContext(cdc *codec.Codec, storeName string) Context {
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
	return Context{
		cdc,
		txBldr,
		cliCtx,
		storeName,
	}
}

// PublishAgentPriceTx creates an publish agent transaction
func (c Context) PublishAgentPriceTx(id [16]byte, price sdk.Coins, status agent.AgentStatus) error {

	msg := agent.MsgPublishAgentPrice{
		c.cliCtx.GetFromAddress(), // ?!
		id,
		agent.Subscription,
		price,
	}

	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	return utils.CompleteAndBroadcastTxCLI(c.txBldr, c.cliCtx, []sdk.Msg{msg})
}
