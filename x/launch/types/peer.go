package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewPeerTunnel(name, address string) Peer {
	return Peer{
		Connection: &Peer_HttpTunnel{
			HttpTunnel: &Peer_HTTPTunnel{
				Name:    name,
				Address: address,
			},
		},
	}
}

func NewPeerConn(address string) Peer {
	return Peer{
		Connection: &Peer_TcpAddress{
			TcpAddress: address,
		},
	}
}

// Validate check the Peer object
func (m Peer) Validate() error {
	switch conn := m.Connection.(type) {
	case *Peer_TcpAddress:
		if conn.TcpAddress == "" {
			return sdkerrors.Wrap(ErrInvalidPeer, "empty peer")
		}
	case *Peer_HttpTunnel:
		if conn.HttpTunnel.Name == "" {
			return sdkerrors.Wrap(ErrInvalidPeer, "empty http tunnel peer name")
		}
		if conn.HttpTunnel.Address == "" {
			return sdkerrors.Wrap(ErrInvalidPeer, "empty http tunnel peer address")
		}
	default:
		return sdkerrors.Wrap(ErrInvalidPeer, "invalid peer connection")
	}
	return nil
}
