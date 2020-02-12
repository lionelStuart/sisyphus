package rpc

import (
	"context"
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
