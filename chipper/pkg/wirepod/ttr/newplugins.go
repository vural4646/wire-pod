package wirepod_ttr

import (
	"context"

	"github.com/kercre123/chipper/pkg/logger"
	plugcomms "github.com/kercre123/chipper/pkg/wirepod/plugincomms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Target string = "localhost:8081"

func TestConnection() {
	cc, err := grpc.Dial(Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Println(err)
		return
	}
	client := plugcomms.NewPluginServiceClient(cc)
	req := &plugcomms.PluginRequest{
		Intent:       "intenttest",
		IntentParams: map[string]string{"test": "test"},
		Esn:          "testesn",
		Ip:           "192",
		Guid:         "guidtest",
	}
	resp, err := client.ProcessPlugin(context.Background(), req)
	if err != nil {
		logger.Println(err)
		return
	}
	logger.Println("response from plugin server:")
	logger.Println(resp)
}
