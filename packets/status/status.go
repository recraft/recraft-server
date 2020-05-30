package statuspackets

import (
	"encoding/json"

	types "github.com/recraft/recraft-lib/types"
	packetsutils "github.com/recraft/recraft-server/packets/utils"

	serverpackets "github.com/recraft/recraft-lib/packets/server"
	jsontypes "github.com/recraft/recraft-lib/types/json"
	"github.com/recraft/recraft-lib/utils"

	clientpackets "github.com/recraft/recraft-lib/packets/client"
)

type PacketSStatus struct {
	Packet *clientpackets.PacketStatusRequest
}

// Receive packet data
func (packetS *PacketSStatus) Receive(player *packetsutils.PlayerConnection) {
	if packetS.Packet == nil {
		packetS.Packet = &clientpackets.PacketStatusRequest{}
	}

	status := &serverpackets.PacketStatus{}
	jsn := &jsontypes.ServerInfo{}
	jsn.Description = "Recraft powered server"
	jsn.Players = jsontypes.Players{Max: 0, Online: 0, Sample: []jsontypes.SamplePlayersList{}}
	jsn.Version = jsontypes.Version{Name: "1.15.2", Protocol: 578}
	bin, err := json.Marshal(jsn)
	if err != nil {
		player.Console.Debug("Failed serializing server info: ", err)
		return
	}

	status.JSON = types.String(bin)
	structbin, err := utils.StructToBinary(status, packetS.ID())
	if err != nil {
		player.Console.Debug("Failed writing binary: ", err)
		return
	}

	_, err = player.Connection.Write(structbin)
	if err != nil {
		player.Console.Debug("Failed writing to buffer: ", err)
	}

}

// ID of packet
func (packetS *PacketSStatus) ID() int32 {
	if packetS.Packet == nil {
		packetS.Packet = &clientpackets.PacketStatusRequest{}
	}
	return packetS.Packet.ID()
}

type PacketSPing struct {
	Packet *clientpackets.PacketPing
}

// Receive packet data
func (packetS *PacketSPing) Receive(player *packetsutils.PlayerConnection) {
	if packetS.Packet == nil {
		packetS.Packet = &clientpackets.PacketPing{}
	}
	err := utils.BinaryToStruct(packetS.Packet, player.Buffer)
	if err != nil {
		player.Console.Debug("Failed reading data from ", player.Connection.RemoteAddr(), ": ", err)
	}

	reply := &serverpackets.PacketPong{Payload: packetS.Packet.Payload}

	structbinary, err := utils.StructToBinary(reply, packetS.ID())
	if err != nil {
		player.Console.Debug("Failed writing binary: ", err)
		return
	}

	_, err = player.Connection.Write(structbinary)
	if err != nil {
		player.Console.Debug("Failed writing to buffer: ", err)
	}
}

// ID of packet
func (packetS *PacketSPing) ID() int32 {
	if packetS.Packet == nil {
		packetS.Packet = &clientpackets.PacketPing{}
	}
	return packetS.Packet.ID()
}
