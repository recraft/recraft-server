package packetsutils

import (
	"bufio"
	"net"

	libpackets "github.com/recraft/recraft-lib/packets"

	"github.com/recraft/recraft-server/console"
)

//PacketS -tream
type PacketS interface {
	ID() int32
	Receive(*PlayerConnection)
}

// PlayerConnection object
// Contains players info
type PlayerConnection struct {
	Joined     bool
	Auth       bool
	Buffer     *bufio.Reader
	Connection *net.TCPConn
	Console    *console.Console
	State      libpackets.State
	Packets    *PacketBuffer
}

type PacketBuffer struct {
	PacketList map[libpackets.State]map[int32]func() PacketS
}
