package myid

import (
	"context"

	myid "github.com/ChisTrun/myid/api"
)

func (s *myidServer) TestEnvoy(ctx context.Context, request *myid.TestEnvoyRequest) (*myid.TestEnvoyResponse, error) {
	status := "ok"
	return &myid.TestEnvoyResponse{
		Status: &status,
	}, nil
}
