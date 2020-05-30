package handshakepackets

import (
	"github.com/recraft/recraft-lib/packets"
	clientpackets "github.com/recraft/recraft-lib/packets/client"
	"github.com/recraft/recraft-lib/utils"
	packetsutils "github.com/recraft/recraft-server/packets/utils"
)

type PacketSHandshake struct {
	Packet *clientpackets.PacketHandshake
}

// Receive packet data
func (packetS *PacketSHandshake) Receive(player *packetsutils.PlayerConnection) {
	if packetS.Packet == nil {
		packetS.Packet = &clientpackets.PacketHandshake{}
	}

	err := utils.BinaryToStruct(packetS.Packet, player.Buffer)
	if err != nil {
		player.Console.Debug("Failed reading data from ", player.Connection.RemoteAddr(), ": ", err)
	}
	if packetS.Packet.NextState == 1 {
		player.State = packets.STATUS
		return
	}
	if packetS.Packet.NextState == 2 {
		//	player.State = packets.LOGIN - not currently supported
		player.Connection.Close()
	}

	player.Console.Debug("Invalid nextState received")
}

// ID of packet
func (packetS *PacketSHandshake) ID() int32 {
	if packetS.Packet == nil {
		packetS.Packet = &clientpackets.PacketHandshake{}
	}
	return packetS.Packet.ID()
}
