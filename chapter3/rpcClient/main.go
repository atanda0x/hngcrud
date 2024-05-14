package main

import (
	"log"
	"net/rpc"
)

// This client func dial to the server and get the remote func remote executed.

type Args struct{}

func main() {
	var reply int64
	args := Args{}

	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	err = client.Call("TimeServer.GiverServerTime", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	log.Printf("%d", reply)
}
