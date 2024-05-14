package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// The RPC server send UTC server time back to the RPC client.
type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(arg *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeServer := new(TimeServer)
	rpc.Register(timeServer)
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	lsn, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	http.Serve(lsn, nil)
}
