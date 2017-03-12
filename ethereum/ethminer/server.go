package ethminer

import (
	"../../"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
)

type Server struct {
	Port      uint16
	rpcServer *rpc.Server
	server    *http.Server
	output    smartpool.UserOutput
}

func NewRPCServer(output smartpool.UserOutput, port uint16) *Server {
	rpcServer := rpc.NewServer()
	service := SmartPoolService{}
	rpcServer.RegisterName("eth", service)
	return &Server{port, rpcServer, &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: rpcServer,
	}, output}
}

func (s *Server) Start() {
	if SmartPool == nil {
		panic("SmartPool instance must be initialized first.")
	}
	SmartPool.Run()
	s.output.Printf("RPC Server is running...\n")
	s.server.ListenAndServe()
}
