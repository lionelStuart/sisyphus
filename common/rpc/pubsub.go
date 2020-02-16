package rpc

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	Hello "sisyphus/common/rpc/proto"
	"strings"
	"time"
)

type PubSubSvc struct {
	pub *pubsub.Publisher
}

func NewPubSubService() *PubSubSvc {
	return &PubSubSvc{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubSubSvc) Publish(ctx context.Context,
	args *Hello.String) (*Hello.String, error) {
	p.pub.Publish(args.GetValue())
	return &Hello.String{}, nil
}

func (p *PubSubSvc) Subscribe(args *Hello.String, stream Hello.PubSubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if k, ok := v.(string); ok {
			if strings.HasPrefix(k, args.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&Hello.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil

}
