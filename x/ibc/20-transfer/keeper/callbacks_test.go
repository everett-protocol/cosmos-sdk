package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types"
)

func (suite *KeeperTestSuite) TestOnChanOpenInit() {
	invalidOrder := channel.ORDERED

	counterparty := channel.NewCounterparty(testPort2, testChannel2)
	err := suite.app.TransferKeeper.OnChanOpenInit(suite.ctx, invalidOrder, []string{testConnection}, testPort1, testChannel1, counterparty, "")
	suite.Error(err) // invalid channel order

	err = suite.app.TransferKeeper.OnChanOpenInit(suite.ctx, testChannelOrder, []string{testConnection}, testPort1, testChannel1, counterparty, "")
	suite.Error(err) // invalid counterparty port ID

	counterparty = channel.NewCounterparty(testPort1, testChannel2)
	err = suite.app.TransferKeeper.OnChanOpenInit(suite.ctx, testChannelOrder, []string{testConnection}, testPort1, testChannel1, counterparty, testChannelVersion)
	suite.Error(err) // invalid version

	err = suite.app.TransferKeeper.OnChanOpenInit(suite.ctx, testChannelOrder, []string{testConnection}, testPort1, testChannel1, counterparty, "")
	suite.NoError(err) // successfully executed
}

func (suite *KeeperTestSuite) TestOnChanOpenTry() {
	invalidOrder := channel.ORDERED

	counterparty := channel.NewCounterparty(testPort2, testChannel2)
	err := suite.app.TransferKeeper.OnChanOpenTry(suite.ctx, invalidOrder, []string{testConnection}, testPort1, testChannel1, counterparty, "", "")
	suite.Error(err) // invalid channel order

	err = suite.app.TransferKeeper.OnChanOpenTry(suite.ctx, testChannelOrder, []string{testConnection}, testPort1, testChannel1, counterparty, "", "")
	suite.Error(err) // invalid counterparty port ID

	counterparty = channel.NewCounterparty(testPort1, testChannel2)
	err = suite.app.TransferKeeper.OnChanOpenTry(suite.ctx, testChannelOrder, []string{testConnection}, testPort1, testChannel1, counterparty, testChannelVersion, "")
	suite.Error(err) // invalid version

	err = suite.app.TransferKeeper.OnChanOpenTry(suite.ctx, testChannelOrder, []string{testConnection}, testPort1, testChannel1, counterparty, "", testChannelVersion)
	suite.Error(err) // invalid counterparty version

	err = suite.app.TransferKeeper.OnChanOpenTry(suite.ctx, testChannelOrder, []string{testConnection}, testPort1, testChannel1, counterparty, "", "")
	suite.NoError(err) // successfully executed
}

func (suite *KeeperTestSuite) TestOnChanOpenAck() {
	err := suite.app.TransferKeeper.OnChanOpenAck(suite.ctx, testPort1, testChannel1, testChannelVersion)
	suite.Error(err) // invalid version

	err = suite.app.TransferKeeper.OnChanOpenAck(suite.ctx, testPort1, testChannel1, "")
	suite.NoError(err) // successfully executed
}

func (suite *KeeperTestSuite) TestOnRecvPacket() {
	packetSeq := uint64(1)
	packetTimeout := uint64(100)

	// when the source is true
	source := true

	packetData := types.NewPacketDataTransfer(testPrefixedCoins1, testAddr1, testAddr2, source, packetTimeout)
	packet := channel.NewPacket(packetData, packetSeq, testPort1, testChannel1, testPort2, testChannel2)

	err := suite.app.TransferKeeper.OnRecvPacket(suite.ctx, packet, packetData)
	suite.Error(err) // invalid denom prefix

	packetData = types.NewPacketDataTransfer(testPrefixedCoins2, testAddr1, testAddr2, source, packetTimeout)
	packet = channel.NewPacket(packetData, packetSeq, testPort1, testChannel1, testPort2, testChannel2)

	err = suite.app.TransferKeeper.OnRecvPacket(suite.ctx, packet, packetData)
	suite.NoError(err) // successfully executed

	totalSupply := suite.app.SupplyKeeper.GetSupply(suite.ctx)
	suite.Equal(testPrefixedCoins2, totalSupply.GetTotal()) // supply should be inflated

	receiverCoins := suite.app.BankKeeper.GetCoins(suite.ctx, packetData.Receiver)
	suite.Equal(testPrefixedCoins2, receiverCoins)

	// when the source is false
	source = false

	packetData = types.NewPacketDataTransfer(testPrefixedCoins2, testAddr1, testAddr2, source, packetTimeout)
	packet = channel.NewPacket(packetData, packetSeq, testPort1, testChannel1, testPort2, testChannel2)

	err = suite.app.TransferKeeper.OnRecvPacket(suite.ctx, packet, packetData)
	suite.Error(err) // invalid denom prefix

	packetData = types.NewPacketDataTransfer(testPrefixedCoins1, testAddr1, testAddr2, source, packetTimeout)
	packet = channel.NewPacket(packetData, packetSeq, testPort1, testChannel1, testPort2, testChannel2)

	err = suite.app.TransferKeeper.OnRecvPacket(suite.ctx, packet, packetData)
	suite.Error(err) // insufficient coins in the corresponding escrow account

	escrowAddress := types.GetEscrowAddress(testPort2, testChannel2)
	_ = suite.app.BankKeeper.SetCoins(suite.ctx, escrowAddress, testCoins)
	_ = suite.app.BankKeeper.SetCoins(suite.ctx, packetData.Receiver, sdk.Coins{})
	err = suite.app.TransferKeeper.OnRecvPacket(suite.ctx, packet, packetData)
	suite.NoError(err) // successfully executed

	receiverCoins = suite.app.BankKeeper.GetCoins(suite.ctx, packetData.Receiver)
	suite.Equal(testCoins, receiverCoins)
}

func (suite *KeeperTestSuite) TestOnTimeoutPacket() {
	packetSeq := uint64(1)
	packetTimeout := uint64(100)

	// when the source is true
	source := true

	packetData := types.NewPacketDataTransfer(testPrefixedCoins2, testAddr1, testAddr2, source, packetTimeout)
	packet := channel.NewPacket(packetData, packetSeq, testPort1, testChannel1, testPort2, testChannel2)

	err := suite.app.TransferKeeper.OnTimeoutPacket(suite.ctx, packet, packetData)
	suite.Error(err) // invalid denom prefix

	packetData = types.NewPacketDataTransfer(testPrefixedCoins1, testAddr1, testAddr2, source, packetTimeout)
	packet = channel.NewPacket(packetData, packetSeq, testPort1, testChannel1, testPort2, testChannel2)

	err = suite.app.TransferKeeper.OnTimeoutPacket(suite.ctx, packet, packetData)
	suite.Error(err) // insufficient coins in the corresponding escrow account

	escrowAddress := types.GetEscrowAddress(testPort2, testChannel2)
	_ = suite.app.BankKeeper.SetCoins(suite.ctx, escrowAddress, testCoins)
	err = suite.app.TransferKeeper.OnTimeoutPacket(suite.ctx, packet, packetData)
	suite.NoError(err) // successfully executed

	senderCoins := suite.app.BankKeeper.GetCoins(suite.ctx, packetData.Sender)
	suite.Equal(testCoins, senderCoins)

	escrowCoins := suite.app.BankKeeper.GetCoins(suite.ctx, escrowAddress)
	suite.Equal(sdk.Coins(nil), escrowCoins)

	// when the source is false
	source = false

	packetData = types.NewPacketDataTransfer(testPrefixedCoins1, testAddr1, testAddr2, source, packetTimeout)
	packet = channel.NewPacket(packetData, packetSeq, testPort1, testChannel1, testPort2, testChannel2)

	_ = suite.app.BankKeeper.SetCoins(suite.ctx, packetData.Sender, sdk.Coins{})
	err = suite.app.TransferKeeper.OnTimeoutPacket(suite.ctx, packet, packetData)
	suite.NoError(err) // successfully executed

	totalSupply := suite.app.SupplyKeeper.GetSupply(suite.ctx)
	suite.Equal(testPrefixedCoins1, totalSupply.GetTotal()) // supply should be inflated

	senderCoins = suite.app.BankKeeper.GetCoins(suite.ctx, packetData.Sender)
	suite.Equal(testPrefixedCoins1, senderCoins)
}
