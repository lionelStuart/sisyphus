package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	Hello "sisyphus/common/rpc/proto"
	"testing"
	"time"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 普通rpc
func TestTcpRpc(t *testing.T) {
	go func() {
		rpc.RegisterName("HelloService", &HelloService{})
		listener, err := net.Listen("tcp", ":2234")
		if err != nil {
			t.Fatal("Accept err:", err)
		}
		t.Log("start listen ..")
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal("fail server:", err)
		}
		rpc.ServeConn(conn)
	}()

	time.Sleep(time.Second * 3)
	client, err := rpc.Dial("tcp", "localhost:2234")
	if err != nil {
		t.Fatal("fail client", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		t.Fatal("fail client", err)
	}
	t.Log(reply)

}

type HelloServiceV2 struct {
}

func (p *HelloServiceV2) Hello(request string, reply *string) error {
	*reply = "hello v2: " + request
	return nil
}

// 封装
func TestSaferRpc(t *testing.T) {
	// server
	go func() {
		RegisterHelloService(&HelloServiceV2{})

		listener, err := net.Listen("tcp", ":2234")
		if err != nil {
			t.Fatal("Accept err:", err)
		}
		t.Log("start listen ..")
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal("fail server:", err)
		}
		rpc.ServeConn(conn)
	}()

	// cli
	client, err := DialHelloService("tcp", ":2234")
	if err != nil {
		t.Fatal("fail client", err)
	}

	var reply string
	err = client.Hello("world", &reply)
	if err != nil {
		t.Fatal("fail client", err)
	}
	t.Log(reply)
}

//json-rpc
func TestJsonRpc(t *testing.T) {
	go func() {
		rpc.RegisterName("HelloService", &HelloService{})
		listener, err := net.Listen("tcp", ":2234")
		if err != nil {
			t.Fatal("Accept err:", err)
		}
		t.Log("start listen ..")

		//for - loop
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal("fail accept", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}()

	time.Sleep(time.Second * 3)
	conn, err := net.Dial("tcp", ":2234")
	if err != nil {
		t.Fatal("fail dial cli")
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		t.Fatal("fail client", err)
	}
	t.Log(reply)
}

func TestHttpRpc(t *testing.T) {
	rpc.RegisterName("HelloService", &HelloService{})

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":2234", nil)
}

type HelloSvcImpl struct {
}

func (p *HelloSvcImpl) Hello(ctx context.Context, args *Hello.String) (*Hello.String, error) {
	reply := &Hello.String{Value: "hello: " + args.GetValue()}
	return reply, nil
}

func (p *HelloSvcImpl) Channel(stream Hello.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &Hello.String{Value: "hello: " + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func TestGRpc(t *testing.T) {
	go func() {
		server := grpc.NewServer()
		Hello.RegisterHelloServiceServer(server, &HelloSvcImpl{})
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

	cli := Hello.NewHelloServiceClient(conn)
	rep, err := cli.Hello(context.Background(), &Hello.String{Value: "grpc world"})
	if err != nil {
		t.Fatal("cli fail", err)
	}
	t.Log(rep.GetValue())
}

func TestGpcStream(t *testing.T) {

	go func() {
		server := grpc.NewServer()
		Hello.RegisterHelloServiceServer(server, &HelloSvcImpl{})
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

	cli := Hello.NewHelloServiceClient(conn)
	stream, err := cli.Channel(context.Background())
	if err != nil {
		t.Fatal("cli fail", err)
	}

	go func() {
		for count := 0; ; count++ {
			if err := stream.Send(&Hello.String{
				Value: fmt.Sprintf("count %d", count)}); err != nil {
				t.Fatal("fail send", err)
			}
			time.Sleep(time.Second)
		}
	}()

	ch := make(chan Hello.String, 2)
	go func(ch chan<- Hello.String) {
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					ch <- Hello.String{Value: "Done"}
					break
				}
				t.Fatal("fail recv", err)
			}
			ch <- *reply
		}
	}(ch)

	for count := 0; count != 10; count++ {
		msg := <-ch
		t.Log("recv: ", msg.GetValue())
	}
}

func TestPubSub(t *testing.T) {
	//server
	go func() {
		server := grpc.NewServer()
		svc := NewPubSubService()
		Hello.RegisterPubSubServiceServer(server, svc)
		lis, err := net.Listen("tcp", ":2234")
		if err != nil {
			t.Fatal("fail ", err)
		}
		server.Serve(lis)

	}()

	time.Sleep(time.Second)
	//publisher
	go func() {
		conn, err := grpc.Dial(":2234", grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		cli := Hello.NewPubSubServiceClient(conn)
		for i := 0; i != 3; i++ {
			_, err := cli.Publish(
				context.Background(), &Hello.String{Value: fmt.Sprintf("golang %d", i)})
			if err != nil {
				t.Fatal(err)
			}

			_, err = cli.Publish(
				context.Background(), &Hello.String{Value: fmt.Sprintf("docker %d", i)})
			if err != nil {
				t.Fatal(err)
			}
		}
	}()

	// subscriber
	go func() {
		conn, err := grpc.Dial(":2234", grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		cli := Hello.NewPubSubServiceClient(conn)
		stream, err := cli.Subscribe(context.Background(), &Hello.String{Value: "golang"})
		if err != nil {
			t.Fatal(err)
		}
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				t.Fatal(err)
			}
			t.Log(reply.GetValue())
		}
	}()

	time.Sleep(time.Second * 15)
}
