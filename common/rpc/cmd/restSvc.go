package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	Rest "sisyphus/common/rpc/proto"
)

type RestSvcImpl struct {
}

func (r *RestSvcImpl) Get(ctx context.Context, message *Rest.StringMessage) (*Rest.StringMessage, error) {
	rep := &Rest.StringMessage{Value: "rest: " + message.GetValue()}
	return rep, nil
}

func (r *RestSvcImpl) Post(ctx context.Context, message *Rest.StringMessage) (*Rest.StringMessage, error) {
	rep := &Rest.StringMessage{Value: "rest post: " + message.GetValue()}
	return rep, nil
}

func main() {
	server := grpc.NewServer()
	Rest.RegisterRestServiceServer(server, &RestSvcImpl{})
	lis, err := net.Listen("tcp", ":2234")
	if err != nil {
		panic(err)
	}
	server.Serve(lis)
}
