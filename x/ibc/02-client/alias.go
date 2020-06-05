package client

// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/02-client/types
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/02-client/keeper

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
)

const (
	AttributeKeyClientID   = types.AttributeKeyClientID
	AttributeKeyClientType = types.AttributeKeyClientType
	SubModuleName          = types.SubModuleName
	RouterKey              = types.RouterKey
	QuerierRoute           = types.QuerierRoute
	QueryAllClients        = types.QueryAllClients
	QueryClientState       = types.QueryClientState
	QueryConsensusState    = types.QueryConsensusState
)

var (
	// functions aliases
	RegisterCodec             = types.RegisterCodec
	SetSubModuleCodec         = types.SetSubModuleCodec
	NewClientConsensusStates  = types.NewClientConsensusStates
	NewGenesisState           = types.NewGenesisState
	DefaultGenesisState       = types.DefaultGenesisState
	NewQueryAllClientsParams  = types.NewQueryAllClientsParams
	NewClientStateResponse    = types.NewClientStateResponse
	NewConsensusStateResponse = types.NewConsensusStateResponse
	NewKeeper                 = keeper.NewKeeper
	QuerierClients            = keeper.QuerierClients

	// variable aliases
	SubModuleCdc                              = types.SubModuleCdc
	ErrClientExists                           = types.ErrClientExists
	ErrClientNotFound                         = types.ErrClientNotFound
	ErrClientFrozen                           = types.ErrClientFrozen
	ErrConsensusStateNotFound                 = types.ErrConsensusStateNotFound
	ErrInvalidConsensus                       = types.ErrInvalidConsensus
	ErrClientTypeNotFound                     = types.ErrClientTypeNotFound
	ErrInvalidClientType                      = types.ErrInvalidClientType
	ErrRootNotFound                           = types.ErrRootNotFound
	ErrInvalidHeader                          = types.ErrInvalidHeader
	ErrInvalidEvidence                        = types.ErrInvalidEvidence
	ErrFailedClientConsensusStateVerification = types.ErrFailedClientConsensusStateVerification
	ErrFailedConnectionStateVerification      = types.ErrFailedConnectionStateVerification
	ErrFailedChannelStateVerification         = types.ErrFailedChannelStateVerification
	ErrFailedPacketCommitmentVerification     = types.ErrFailedPacketCommitmentVerification
	ErrFailedPacketAckVerification            = types.ErrFailedPacketAckVerification
	ErrFailedPacketAckAbsenceVerification     = types.ErrFailedPacketAckAbsenceVerification
	ErrFailedNextSeqRecvVerification          = types.ErrFailedNextSeqRecvVerification
	ErrSelfConsensusStateNotFound             = types.ErrSelfConsensusStateNotFound
	EventTypeCreateClient                     = types.EventTypeCreateClient
	EventTypeUpdateClient                     = types.EventTypeUpdateClient
	EventTypeSubmitMisbehaviour               = types.EventTypeSubmitMisbehaviour
	AttributeValueCategory                    = types.AttributeValueCategory
)

type (
	StakingKeeper          = types.StakingKeeper
	ConsensusStates        = types.ClientConsensusStates
	GenesisState           = types.GenesisState
	QueryAllClientsParams  = types.QueryAllClientsParams
	StateResponse          = types.StateResponse
	ConsensusStateResponse = types.ConsensusStateResponse
	Keeper                 = keeper.Keeper
)