package rpcserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

/**
go test -run=TestRpcServer
go test -run=TestRpcClient
*/

func TestRpcServer(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen Error:", e)
	}
	go http.Serve(l, nil)
	select {}

}

func TestRpcClient(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dailing:", err)
	}
	args := &Args{105, 10}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith:%d*%d=%d\n", args.A, args.B, reply)

	quo := new(Quotient)
	call := client.Go("Arith.Divide", args, &quo, nil)
	replyCall := <-call.Done
	fmt.Printf("divide:%d/%d=%d,%d%%%d=%d,%v\n", args.A, args.B, quo.Quo, args.A, args.B, quo.Rem, replyCall)
}
