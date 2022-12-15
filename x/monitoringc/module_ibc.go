package monitoringc

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v6/modules/core/exported"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// OnChanOpenInit implements the IBCModule interface
func (am AppModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	_ channeltypes.Counterparty,
	version string,
) (string, error) {
	if order != channeltypes.ORDERED {
		return "", sdkerrors.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s ", channeltypes.ORDERED, order)
	}

	// Require portID is the portID module is bound to
	boundPort := am.keeper.GetPort(ctx)
	if boundPort != portID {
		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if version != types.Version {
		return "", sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected %s", version, types.Version)
	}

	// Claim channel capability passed back by IBC module
	if err := am.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return "", err
	}

	if len(connectionHops) != 1 {
		return "", sdkerrors.Wrap(
			channeltypes.ErrTooManyConnectionHops,
			"must have direct connection to provider chain",
		)
	}

	// Check if the client ID is a verified from MsgCreateClient
	if err := am.keeper.VerifyClientIDFromConnID(ctx, connectionHops[0]); err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidHandshake, err.Error())
	}

	return version, nil
}

// OnChanOpenTry implements the IBCModule interface
func (am AppModule) OnChanOpenTry(
	_ sdk.Context,
	_ channeltypes.Order,
	_ []string,
	_,
	_ string,
	_ *capabilitytypes.Capability,
	_ channeltypes.Counterparty,
	_ string,
) (string, error) {
	return "", sdkerrors.Wrap(types.ErrInvalidHandshake, "IBC handshake must be initiated by the consumer")
}

// OnChanOpenAck implements the IBCModule interface
func (am AppModule) OnChanOpenAck(
	ctx sdk.Context,
	_,
	channelID string,
	_,
	counterpartyVersion string,
) error {
	if counterpartyVersion != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: %s, expected %s", counterpartyVersion, types.Version)
	}

	// register the official client ID for the provider connection
	if err := am.keeper.RegisterProviderClientIDFromChannelID(ctx, channelID); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidHandshake, err.Error())
	}

	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (am AppModule) OnChanOpenConfirm(
	_ sdk.Context,
	_,
	_ string,
) error {
	return sdkerrors.Wrap(types.ErrInvalidHandshake, "IBC handshake must be initiated by the consumer")
}

// OnChanCloseInit implements the IBCModule interface
func (am AppModule) OnChanCloseInit(
	_ sdk.Context,
	_,
	_ string,
) error {
	// Disallow user-initiated channel closing for channels
	return types.ErrCannotCloseChannel
}

// OnChanCloseConfirm implements the IBCModule interface
func (am AppModule) OnChanCloseConfirm(
	_ sdk.Context,
	_,
	_ string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (am AppModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	_ sdk.AccAddress,
) ibcexported.Acknowledgement {
	var ack channeltypes.Acknowledgement

	// this line is used by starport scaffolding # oracle/packet/module/recv

	var modulePacketData spntypes.MonitoringPacketData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return channeltypes.NewErrorAcknowledgement(sdkerrors.Wrap(types.ErrJSONUnmarshal, err.Error()))
	}

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	case *spntypes.MonitoringPacketData_MonitoringPacket:
		packetAck, err := am.keeper.OnRecvMonitoringPacket(ctx, modulePacket, *packet.MonitoringPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err)
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
			if err != nil {
				return channeltypes.NewErrorAcknowledgement(sdkerrors.Wrap(types.ErrJSONMarshal, err.Error()))
			}
			ack = channeltypes.NewResultAcknowledgement(sdk.MustSortJSON(packetAckBytes))
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMonitoringPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
		// this line is used by starport scaffolding # ibc/packet/module/recv
	default:
		err := fmt.Errorf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return ack
}

// OnAcknowledgementPacket implements the IBCModule interface
func (am AppModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	_ sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrap(types.ErrJSONUnmarshal, err.Error())
	}

	// this line is used by starport scaffolding # oracle/packet/module/ack

	var modulePacketData spntypes.MonitoringPacketData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return sdkerrors.Wrap(types.ErrJSONUnmarshal, err.Error())
	}

	var eventType string

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	case *spntypes.MonitoringPacketData_MonitoringPacket:
		err := am.keeper.OnAcknowledgementMonitoringPacket(ctx, modulePacket, *packet.MonitoringPacket, ack)
		if err != nil {
			return err
		}
		eventType = types.EventTypeMonitoringPacket
		// this line is used by starport scaffolding # ibc/packet/module/ack
	default:
		return sdkerrors.Wrapf(types.ErrUnrecognizedPacketType, "packet type: %T", packet)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			eventType,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ack)),
		),
	)

	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
			),
		)
	}

	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (am AppModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	_ sdk.AccAddress,
) error {
	var modulePacketData spntypes.MonitoringPacketData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return sdkerrors.Wrap(types.ErrJSONUnmarshal, err.Error())
	}

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	case *spntypes.MonitoringPacketData_MonitoringPacket:
		err := am.keeper.OnTimeoutMonitoringPacket(ctx, modulePacket, *packet.MonitoringPacket)
		if err != nil {
			return err
		}
		// this line is used by starport scaffolding # ibc/packet/module/timeout
	default:
		return sdkerrors.Wrapf(types.ErrUnrecognizedPacketType, "packet type: %T", packet)
	}

	return nil
}
