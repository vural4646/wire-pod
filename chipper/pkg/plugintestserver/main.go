package main

import (
	"context"
	"fmt"
	"net"

	plugcomms "github.com/kercre123/chipper/pkg/wirepod/plugincomms"
	"google.golang.org/grpc"
)

type PluginServer struct {
	plugcomms.UnimplementedPluginServiceServer
}

func NewPluginServer() *PluginServer {
	return &PluginServer{}
}

func (s *PluginServer) ProcessPlugin(ctx context.Context, req *plugcomms.PluginRequest) (*plugcomms.PluginResponse, error) {
	fmt.Println("Incoming request: ")
	fmt.Println(req)
	resp := &plugcomms.PluginResponse{
		IsRunning:        false,
		Intent:           "intent_greeting_goodmorning",
		IntentParams:     map[string]string{"": ""},
		IsKnowledgeGraph: false,
		SpeechOutput:     "",
	}
	return resp, nil
}

func main() {
	plugServ := NewPluginServer()
	srv := grpc.NewServer()
	plugcomms.RegisterPluginServiceServer(srv, plugServ)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("serving")
	srv.Serve(listener)
}
