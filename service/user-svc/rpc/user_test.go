package rpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"sisyphus/common/redis"
	"sisyphus/common/setting"
	"sisyphus/models"
	proto "sisyphus/proto/user"
	"testing"
	"time"
)

func setupSuite() {
	path := `D:\PHOENIX\Documents\WORKSPACE\Go\GO_WORKSPACE\sisyphus\conf\app.ini`
	setting.Setup(path)
	models.Setup()
	redis.SetUp()
}

func TestService_AddUser(t *testing.T) {
	setupSuite()

	go func() {
		server := grpc.NewServer()
		proto.RegisterUserSvcServer(server, &Service{})
		lis, err := net.Listen("tcp", ":2234")
		if err != nil {
			t.Fatal("fail ", err)
		}
		server.Serve(lis)

	}()

	time.Sleep(time.Second * 3)

	conn, err := grpc.Dial(":2234", grpc.WithInsecure())
	if err != nil {
		t.Fatal("cli fail", err)
	}
	defer conn.Close()

	// cli := Hello.NewHelloServiceClient(conn)
	cli := proto.NewUserSvcClient(conn)
	rep, err := cli.AddUser(context.Background(), &proto.AddUserRequest{
		User: &proto.User{
			Username: "mary",
			Password: "pass123",
			Email:    "jim@126.com",
			Phone:    "12300002222",
			State:    0,
			Profile: &proto.Profile{
				Nickname: "mary nick",
				Age:      11,
				Gender:   "M",
				Address:  "this is address",
			},
		},
	})
	if err != nil {
		t.Fatal("cli fail", err)
	}
	t.Log(rep.GetId())
}
