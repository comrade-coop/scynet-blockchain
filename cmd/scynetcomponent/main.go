package main

//go:generate protoc -I ../../protocols/ ../../protocols/Component.proto --go_out=plugins=grpc:protobufs
//go:generate protoc -I ../../protocols/ ../../protocols/Hatchery.proto --go_out=plugins=grpc:protobufs
//go:generate protoc -I ../../protocols/ ../../protocols/Shared.proto --go_out=plugins=grpc:protobufs

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	sdk "github.com/cosmos/cosmos-sdk/types"

	app "github.com/comrade-coop/scynet-blockchain"
	protobufs "github.com/comrade-coop/scynet-blockchain/cmd/scynetcomponent/protobufs"
	agentTypes "github.com/comrade-coop/scynet-blockchain/x/agent"
	agent "github.com/comrade-coop/scynet-blockchain/x/agent/client/facade"
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

	server := componentServer{
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

// RegisterInput - Register a new input to the component
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

// AgentStart - Start running a particular agent
func (s *componentServer) AgentStart(context.Context, *protobufs.AgentStartRequest) (*protobufs.Void, error) {
	return new(protobufs.Void), nil
}

// AgentStop - Stop that agent
func (s *componentServer) AgentStop(context.Context, *protobufs.AgentRequest) (*protobufs.Void, error) {
	return new(protobufs.Void), nil
}

// AgentStatus - Check the status of an agent.
func (s *componentServer) AgentStatus(context.Context, *protobufs.AgentRequest) (*protobufs.AgentStatusResponse, error) {
	res := new(protobufs.AgentStatusResponse)
	res.Running = false
	return res, nil
}

// AgentList - Retrieve a list of running agents.
func (s *componentServer) AgentList(context.Context, *protobufs.AgentQuery) (*protobufs.ListOfAgents, error) {
	return new(protobufs.ListOfAgents), nil
}
