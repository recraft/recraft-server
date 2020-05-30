package packets

import (
	clientpackets "github.com/recraft/recraft-lib/packets/client"

	libpackets "github.com/recraft/recraft-lib/packets"
	handshakepackets "github.com/recraft/recraft-server/packets/handshake"
	statuspackets "github.com/recraft/recraft-server/packets/status"
	packetutils "github.com/recraft/recraft-server/packets/utils"
)

//GenPackets automatically
func GenPackets() map[libpackets.State]map[int32]func() packetutils.PacketS {

	return map[libpackets.State]map[int32]func() packetutils.PacketS{

		libpackets.HANDSHAKE: {
			0: func() packetutils.PacketS {
				return &handshakepackets.PacketSHandshake{Packet: new(clientpackets.PacketHandshake)}
			},
		},
		libpackets.STATUS: {
			0: func() packetutils.PacketS {
				return &statuspackets.PacketSStatus{Packet: new(clientpackets.PacketStatusRequest)}
			},
			1: func() packetutils.PacketS {
				return &statuspackets.PacketSPing{Packet: new(clientpackets.PacketPing)}
			},
		},
	}

}
func GetPacket(packetList map[libpackets.State]map[int32]func() packetutils.PacketS, state libpackets.State, id int32) packetutils.PacketS {

	packt := packetList[state][id]
	if packt == nil {
		return nil
	}

	return packt()
}
