package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"

	agentCli "github.com/comrade-coop/scynet-blockchain/x/agent/client/cli"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

// NewModuleClient creates a new ModuleClient
func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	agentQueryCmd := &cobra.Command{
		Use:   "agent",
		Short: "Querying commands for the agent module",
	}

	agentQueryCmd.AddCommand(client.GetCommands(
		agentCli.GetCmdGetAgent(mc.storeKey, mc.cdc),
	)...)

	return agentQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	agentTxCmd := &cobra.Command{
		Use:   "agent",
		Short: "Agent transactions subcommands",
	}

	agentTxCmd.AddCommand(client.PostCommands(
		agentCli.GetCmdPublishAgentPrice(mc.cdc),
	)...)

	return agentTxCmd
}
