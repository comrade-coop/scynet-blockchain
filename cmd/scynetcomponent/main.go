package main

//go:generate protoc -I ../../protocols/ ../../protocols/Component.proto --go_out=plugins=grpc:protobufs
//go:generate protoc -I ../../protocols/ ../../protocols/Hatchery.proto --go_out=plugins=grpc:protobufs
//go:generate protoc -I ../../protocols/ ../../protocols/Shared.proto --go_out=plugins=grpc:protobufs

import (
	"context"
	"fmt"
	"log"
	"net"

	sdkContext "github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	tendermintQuery "github.com/tendermint/tendermint/libs/pubsub/query"
	tendermintTypes "github.com/tendermint/tendermint/types"
	"google.golang.org/grpc"

	app "github.com/comrade-coop/scynet-blockchain"
	protobufs "github.com/comrade-coop/scynet-blockchain/cmd/scynetcomponent/protobufs"
	agentTypes "github.com/comrade-coop/scynet-blockchain/x/agent"
	agent "github.com/comrade-coop/scynet-blockchain/x/agent/client/facade"
)

const (
	storeAgent    = "agent"
	componentType = "external_blockchain"
)

func main() {
	cdc := app.MakeCodec()

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	agentCtx := agent.NewContext(cdc, storeAgent)

	hatcheryConn, err := grpc.Dial("127.0.0.1:9998", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer hatcheryConn.Close()
	hatcheryClient := protobufs.NewHatcheryClient(hatcheryConn)

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

	componentUUID := "252efc75-b125-4053-a611-9c2cc65823e0" // TODO: generate and store in config

	hatcheryClient.RegisterComponent(context.Background(), &protobufs.ComponentRegisterRequest{
		Uuid:       componentUUID,
		Address:    "127.0.0.1:3451", // TODO: make configurable or pick proper bound address
		RunnerType: []string{componentType},
	})

	cliCtx := sdkContext.NewCLIContext()

	transactions := make(chan interface{})
	cliCtx.Client.Subscribe(context.Background(), tendermintTypes.EventTx, tendermintQuery.Empty{}, transactions)

	go func() {
		ctx := context.Background()
		for tx := range transactions {
			tx := tx.(tendermintTypes.Tx)
			var message sdk.Msg
			err := cdc.UnmarshalBinaryBare(tx, &message)
			if err != nil {
				log.Fatalf("failed to unmarshal message: %v", err)
				continue
			}

			switch message.(type) {
			case agentTypes.MsgPublishAgentPrice:
				message := message.(agentTypes.MsgPublishAgentPrice)
				var uuid uuid.UUID = message.UUID
				hatcheryClient.RegisterAgent(ctx, &protobufs.AgentRegisterRequest{
					Agent: &protobufs.Agent{
						Uuid:          uuid.String(),
						ComponentType: componentType,
						ComponentId:   componentUUID,
						// Outputs: from domain,
						// Frequency: from domain,
						// TODO: different types of tokens? FIXME: Also, might overflow here!
						Price: uint32(message.Price.AmountOf("scynet").BigInt().Uint64()),
					},
				})
			case agentTypes.MsgPublishData:
				message := message.(agentTypes.MsgPublishData)
				var uuid uuid.UUID = message.UUID

				var outputs = make([]*protobufs.Shape, len(message.Shapes))
				for i := range message.Shapes {
					dim := make([]uint32, len(message.Shapes[i]))
					for j := range message.Shapes[i] {
						dim[j] = uint32(message.Shapes[i][j])
					}
					outputs[i] = &protobufs.Shape{
						Dimension: dim,
					}
				}

				hatcheryClient.RegisterAgent(ctx, &protobufs.AgentRegisterRequest{
					Agent: &protobufs.Agent{
						Uuid:          uuid.String(),
						ComponentType: componentType,
						ComponentId:   componentUUID,
						Outputs:       outputs,
						// Frequency: uh...,
						// TODO: different types of tokens? FIXME: Also, might overflow here!
						Price: uint32(message.Price.AmountOf("scynet").BigInt().Uint64()),
					},
				})
				// TODO: handle renting
			}
		}
	}()
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
