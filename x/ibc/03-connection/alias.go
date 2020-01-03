package connection

// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types"
)

const (
	UNINITIALIZED                    = types.UNINITIALIZED
	INIT                             = types.INIT
	TRYOPEN                          = types.TRYOPEN
	OPEN                             = types.OPEN
	StateUninitialized               = types.StateUninitialized
	StateInit                        = types.StateInit
	StateTryOpen                     = types.StateTryOpen
	StateOpen                        = types.StateOpen
	AttributeKeyConnectionID         = types.AttributeKeyConnectionID
	AttributeKeyCounterpartyClientID = types.AttributeKeyCounterpartyClientID
	SubModuleName                    = types.SubModuleName
	StoreKey                         = types.StoreKey
	RouterKey                        = types.RouterKey
	QuerierRoute                     = types.QuerierRoute
	QueryAllConnections              = types.QueryAllConnections
	QueryConnection                  = types.QueryConnection
	QueryClientConnections           = types.QueryClientConnections
)

var (
	// functions aliases
	NewKeeper                        = keeper.NewKeeper
	QuerierConnections               = keeper.QuerierConnections
	QuerierConnection                = keeper.QuerierConnection
	QuerierClientConnections         = keeper.QuerierClientConnections
	RegisterCodec                    = types.RegisterCodec
	NewConnectionEnd                 = types.NewConnectionEnd
	NewCounterparty                  = types.NewCounterparty
	StateFromString                  = types.StateFromString
	ErrConnectionExists              = types.ErrConnectionExists
	ErrConnectionNotFound            = types.ErrConnectionNotFound
	ErrClientConnectionPathsNotFound = types.ErrClientConnectionPathsNotFound
	ErrConnectionPath                = types.ErrConnectionPath
	ErrInvalidConnectionState        = types.ErrInvalidConnectionState
	ErrInvalidCounterparty           = types.ErrInvalidCounterparty
	ConnectionPath                   = types.ConnectionPath
	ClientConnectionsPath            = types.ClientConnectionsPath
	KeyConnection                    = types.KeyConnection
	KeyClientConnections             = types.KeyClientConnections
	NewMsgConnectionOpenInit         = types.NewMsgConnectionOpenInit
	NewMsgConnectionOpenTry          = types.NewMsgConnectionOpenTry
	NewMsgConnectionOpenAck          = types.NewMsgConnectionOpenAck
	NewMsgConnectionOpenConfirm      = types.NewMsgConnectionOpenConfirm
	NewConnectionResponse            = types.NewConnectionResponse
	NewQueryConnectionParams         = types.NewQueryConnectionParams
	NewClientConnectionsResponse     = types.NewClientConnectionsResponse
	NewQueryClientConnectionsParams  = types.NewQueryClientConnectionsParams
	GetCompatibleVersions            = types.GetCompatibleVersions
	LatestVersion                    = types.LatestVersion
	PickVersion                      = types.PickVersion

	// variable aliases
	SubModuleCdc                   = types.SubModuleCdc
	EventTypeConnectionOpenInit    = types.EventTypeConnectionOpenInit
	EventTypeConnectionOpenTry     = types.EventTypeConnectionOpenTry
	EventTypeConnectionOpenAck     = types.EventTypeConnectionOpenAck
	EventTypeConnectionOpenConfirm = types.EventTypeConnectionOpenConfirm
	AttributeValueCategory         = types.AttributeValueCategory
)

type (
	Keeper                       = keeper.Keeper
	ConnectionEnd                = types.ConnectionEnd
	Counterparty                 = types.Counterparty
	State                        = types.State
	ClientKeeper                 = types.ClientKeeper
	MsgConnectionOpenInit        = types.MsgConnectionOpenInit
	MsgConnectionOpenTry         = types.MsgConnectionOpenTry
	MsgConnectionOpenAck         = types.MsgConnectionOpenAck
	MsgConnectionOpenConfirm     = types.MsgConnectionOpenConfirm
	ConnectionResponse           = types.ConnectionResponse
	QueryConnectionParams        = types.QueryConnectionParams
	ClientConnectionsResponse    = types.ClientConnectionsResponse
	QueryClientConnectionsParams = types.QueryClientConnectionsParams
)
