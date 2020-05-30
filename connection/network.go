package connection

import (
	"bufio"
	"bytes"
	"net"
	"strconv"

	"github.com/recraft/recraft-server/packets"

	libpackets "github.com/recraft/recraft-lib/packets"
	"github.com/recraft/recraft-lib/types"
	"github.com/recraft/recraft-server/console"
	packetsutils "github.com/recraft/recraft-server/packets/utils"

	"github.com/recraft/recraft-lib/utils"
)

// ServerNetwork object
// Contains low level networking objects
type ServerNetwork struct {
	socket    *net.TCPListener
	Connected bool
	Address   string
	Port      int16
}

type Info struct {
	Action string
}

// Server object
// Contains some server data
type Server struct {
	Network *ServerNetwork
	Info    chan Info
	Console *console.Console
}

// Listen for clients to connect
func (server Server) Listen(logging *console.Console) error {
	if server.Network.Connected {
		return utils.NewError("Already connected")
	}
	server.Console = logging
	address, err := net.ResolveTCPAddr("tcp", server.Network.Address+":"+strconv.Itoa(int(server.Network.Port)))

	server.Network.socket, err = net.ListenTCP("tcp", address)
	if err != nil {
		server.Info <- Info{Action: "close"}
		return err
	}
	server.Network.Connected = true
	logging.Info("Listening on ", server.Network.Address, ":", server.Network.Port)

	//return nil

	go func() {
		for {
			connection, err := server.Network.socket.AcceptTCP()
			if err != nil {
				logging.Warning("Failed to accept connection: ", err)
				continue
			}
			go server.NewClientConnection(connection)
		}
	}()
	return nil
}

// NewClientConnection handles client connections
func (server Server) NewClientConnection(clientConnection *net.TCPConn) {

	server.Console.Debug("Connection opened from ", clientConnection.RemoteAddr())
	var buffer []byte
	reader := bufio.NewReader(clientConnection)

	player := &packetsutils.PlayerConnection{}
	player.State = libpackets.HANDSHAKE
	player.Console = server.Console
	player.Connection = clientConnection
	player.Packets = &packetsutils.PacketBuffer{PacketList: packets.GenPackets()}
	for {
		buffer = make([]byte, 1024)
		buflen, err := clientConnection.Read(buffer)
		if err != nil && err.Error() == "EOF" {
			server.Console.Debug("[", clientConnection.RemoteAddr(), "] Exited ")
			break
		}
		if err != nil && buflen == 0 {
			break
		}

		newBuffer := bufio.NewReader(bytes.NewReader(buffer))
		len := types.VarInt(0)
		len.Read(newBuffer)
		var fullBytesBuffer []byte
		if int(len) < buflen {
			fullBytesBuffer = buffer
		} else {
			fullBytesBuffer, err = types.ReadBytes(reader, int(len)-(buflen-1))
		}
		if err != nil {
			server.Console.Debug("Failed reading fullBuffer")

		}
		fullBuffer := bufio.NewReader(bytes.NewReader(append(fullBytesBuffer, buffer...)))
		player.Buffer = fullBuffer
		//skip again lenght

		len.Read(fullBuffer)

		packetID := types.VarInt(0)
		err = packetID.Read(fullBuffer)
		if err != nil {
			server.Console.Debug("Failed reading fullBuffer: ", err)

		}

		packet := packets.GetPacket(player.Packets.PacketList, player.State, int32(packetID))
		if packet == nil {
			server.Console.Debug("Packed with id ", packetID, " was not found")
			continue
		}
		packet.Receive(player)
	}
	clientConnection.Close()
	server.Console.Debug("Connection closed from ", clientConnection.RemoteAddr())
}
