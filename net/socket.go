package net

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	mctypes "github.com/recraft/recraft-lib/types"

	servertypes "github.com/recraft/recraft-server/types"
)

type RecraftServer struct {
	socket    net.Listener
	connected bool
	Mode      string
	Address   string
	State     string
}

/*Listen opens a tcp server connection to allow Minecraft clients to connect.
Note: The server will not listen for any connection if the event handler is not called.
*/
func (connection RecraftServer) Listen() error {
	var ipType = "testz"
	if connection.Mode == "4" {
		ipType = "tcp4"
	} else if connection.Mode == "6" {
		ipType = "tcp6"
	} else {
		ipType = "tcp"
	}
	var err error = nil
	connection.socket, err = net.Listen(ipType, connection.Address)
	if err != nil {
		return err
	}

	connection.connected = true
	fmt.Println("Server opened!")

	connection.ServerHandler()
	return nil

}

func (connection *RecraftServer) ServerHandler() {
	if !connection.connected {
		log.Panic("The tcp server is not online!")
	}
	for {
		clientsocket, err := connection.socket.Accept()
		if err != nil {
			log.Panicln("An error occured while accepting a tcp connection: ", err)
		}
		client := servertypes.ClientInfo{Status: servertypes.HandshakeState, Connection: &clientsocket}
		go connection.ClientHandler(&client)
	}

}

func (connection *RecraftServer) ClientHandler(client *servertypes.ClientInfo) {
	Connection := *client.Connection
	connection.State = servertypes.HandshakeState
	Connection.SetDeadline(time.Time{})
	reader := bufio.NewReader(*client.Connection)
	for {

		lenght := mctypes.VarInt(0)
		err := lenght.Read(reader)
		if err != nil {
			fmt.Println("An error occured while reading lenght: ", err)

			continue
		}

	}
}

/*code needs to reworked


StartHandling enables the ability to handle connections from clients.
Note: Listen needs to be called before executing this function, otherwise it will return an error.

func (connection RecraftServer) StartHandling() error {
	if !connection.connected {
		return errors.New("The tcp server is not online!")
	}
	for {
		//Wait for a connection to open, and then start a separate thread for it
		client, err := connection.socket.Accept()
		if err != nil {
			//just debug
			fmt.Println("An error occured while accepting a tcp connection: ", err)
			connection.connected = false
			connection.socket.Close()
			return err
		}
		go clientHandler(client)

	}
}
func clientHandler(connection net.Conn) {
	//wip
	reader := bufio.NewReader(connection)
	var currentQueue string = ""
	for {

		lenght, err := utils.GetVarIntFromBuffer(reader)
		if err != nil {
			fmt.Println("An error occured while reading lenght: ", err)
			connection.Close()

			return
		}

		buffer := make([]byte, lenght)
		_, err = reader.Read(buffer)
		if err == io.EOF {
			connection.Close()
			return
		}
		if err != nil {
			connection.Close()
			return

		}

		utils.VarInt(0).GetBytes()
		PacketID, _ := CountBuffer.GetVarInt()

		if PacketID == 0 && currentQueue == "" {

			connectionInfo := types.ClientHandshake{}
			err = CountBuffer.BinaryToStruct(&connectionInfo)
			if err != nil {
				fmt.Println("An error occured! ", err)
				connection.Close()
				return
			}
			if connectionInfo.NextState == 1 {
				currentQueue = "waitforping"
			} else {
				connection.Close()
			}

		} else if PacketID == 0 && currentQueue == "waitforping" {

		
				serverInfo := jsontypes.ServerInfo{Description: "§b§lRecraft", Players: jsontypes.Players{Max: 1, Online: 5}, Version: jsontypes.Version{Name: "§b§test", Protocol: 1}}

			
			info, _ := json.Marshal(serverInfo)
			bte, err := utils.StructToBinary(&types.ServerListPingResponse{JSON: types.String(info)}, 0)
			if err != nil {
				fmt.Println("An error occured! ", err)
				connection.Close()
				return
			}
			connection.Write(bte)

			currentQueue = "waitforping2"

		} else if PacketID == 1 && currentQueue == "waitforping2" {

			ssf, err := CountBuffer.ReadLong()
			if err != nil {
				fmt.Println("An error occured! ", err)
				connection.Close()
				return
			}
			bte, err := utils.StructToBinary(&types.Pong{Payload: ssf}, 1)

			connection.Write(bte)
			connection.Close()

		} else {

			fmt.Println("PacketID not supported! Received: ", PacketID)
			connection.Close()
			return
		}
	}
}
*/
