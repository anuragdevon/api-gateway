package inventory

import (
	"fmt"

	"api-gateway/pkg/config"
	"api-gateway/pkg/inventory/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.InventoryServiceClient
}

func InitServiceClient(c *config.Config) pb.InventoryServiceClient {
	cc, err := grpc.Dial(c.InventorySvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewInventoryServiceClient(cc)
}
