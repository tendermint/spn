package types

import (
	"errors"
)

func NewPeerTunnel(id, name, address string) Peer {
	return Peer{
		Id: id,
		Connection: &Peer_HttpTunnel{
			HttpTunnel: &Peer_HTTPTunnel{
				Name:    name,
				Address: address,
			},
		},
	}
}

func NewPeerConn(id, address string) Peer {
	return Peer{
		Id: id,
		Connection: &Peer_TcpAddress{
			TcpAddress: address,
		},
	}
}

func NewPeerEmpty(id string) Peer {
	return Peer{Id: id, Connection: &Peer_None{None: &Peer_EmptyConnection{}}}
}

// Validate check the Peer object
func (m Peer) Validate() error {
	if m.Id == "" {
		return errors.New("empty peer id")
	}
	switch conn := m.Connection.(type) {
	case *Peer_TcpAddress:
		if conn.TcpAddress == "" {
			return errors.New("empty peer tcp address")
		}
	case *Peer_HttpTunnel:
		if conn.HttpTunnel.Name == "" {
			return errors.New("empty http tunnel peer name")
		}
		if conn.HttpTunnel.Address == "" {
			return errors.New("empty http tunnel peer address")
		}
	case *Peer_None:
		return nil
	default:
		return errors.New("invalid peer connection")
	}
	return nil
}
