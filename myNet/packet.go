package myNet

import "TCPGameServer/iface"

type Option func(s *Server)

func AddPacket(pack iface.IDataPack) Option {
	return func(s *Server) {
		s.packet = pack
	}
}
