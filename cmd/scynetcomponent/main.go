package main

//go:generate protoc -I ../../../protocols/ ../../../protocols/Component.proto --go_out=plugins=grpc:protobufs
//go:generate protoc -I ../../../protocols/ ../../../protocols/Hatchery.proto --go_out=plugins=grpc:protobufs
//go:generate protoc -I ../../../protocols/ ../../../protocols/Shared.proto --go_out=plugins=grpc:protobufs

import (
	"os"
	"fmt"
	"log"
	"context"
	"net"

	"google.golang.org/grpc"
	"github.com/google/uuid"

	sdk "github.com/cosmos/cosmos-sdk/types"

	app "github.com/comrade-coop/scynet/blockchain"
	agentTypes "github.com/comrade-coop/scynet/blockchain/x/agent"
	agent "github.com/comrade-coop/scynet/blockchain/x/agent/client/facade"
	protobufs "github.com/comrade-coop/scynet/blockchain/cmd/scynetcomponent/protobufs"
)

const (
	storeAcc   = "acc"
	storeAgent = "agent"
)

var defaultCLIHome = os.ExpandEnv("$HOME/.nscli")

func main() {
	cdc := app.MakeCodec()

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	agentCtx := agent.NewContext(cdc, storeAgent)

	server := componentServer {
		agentCtx,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3451))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protobufs.RegisterComponentServer(grpcServer, &server)
	grpcServer.Serve(lis)
}

type componentServer struct {
	agentCtx agent.Context
}


// Register a new input to the component

func (s *componentServer) RegisterInput(c context.Context, request *protobufs.RegisterInputRequest) (*protobufs.Void, error) {

	id, err := uuid.Parse(request.Input.Uuid)
	if err != nil {
		return nil, err
	}

	s.agentCtx.PublishAgentPriceTx(
		id,
		[]sdk.Coin{sdk.NewCoin("scynet", sdk.NewInt(int64(request.Input.Price)))},
		agentTypes.Sale,
	)
	return new(protobufs.Void), nil
}
// Start running a particular agent

func (s *componentServer) AgentStart(context.Context, *protobufs.AgentStartRequest) (*protobufs.Void, error) {
	return new(protobufs.Void), nil
}
// Stop that agent

func (s *componentServer) AgentStop(context.Context, *protobufs.AgentRequest) (*protobufs.Void, error) {
	return new(protobufs.Void), nil
}
// Check the status of an agent.

func (s *componentServer) AgentStatus(context.Context, *protobufs.AgentRequest) (*protobufs.AgentStatusResponse, error) {
	res := new(protobufs.AgentStatusResponse)
	res.Running = false
	return res, nil
}
// Retrieve a list of running agents.

func (s *componentServer) AgentList(context.Context, *protobufs.AgentQuery) (*protobufs.ListOfAgents, error) {
	return new(protobufs.ListOfAgents), nil
}
