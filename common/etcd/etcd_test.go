package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/mvcc/mvccpb"
	. "github.com/smartystreets/goconvey/convey"

	"testing"
	"time"
)

func setupEtcd() (cli *clientv3.Client, kv clientv3.KV, err error) {

	config := clientv3.Config{
		Endpoints:   []string{"192.168.0.240:2379"},
		DialTimeout: 5 * time.Second,
	}

	if cli, err = clientv3.New(config); err != nil {
		panic(err)
	}

	kv = clientv3.NewKV(cli)
	return
}

func TestEtcdKv(t *testing.T) {

	Convey("setup", t, func() {
		cli, kv, err := setupEtcd()
		So(err, ShouldBeNil)

		Convey("test set&get", func() {
			kv.Put(context.TODO(), "foo", "bar")
			kv.Put(context.TODO(), "foo1", "bar1")

			getResp, err := kv.Get(context.TODO(), "foo", clientv3.WithPrefix())
			So(err, ShouldBeNil)
			t.Log("get ", getResp.Kvs)
			_, err = kv.Delete(context.TODO(), "foo", clientv3.WithPrefix())
			So(err, ShouldBeNil)
		})

		Reset(func() {
			cli.Close()
		})

	})
}

func TestEtcdWatch(t *testing.T) {
	Convey("setup", t, func() {
		cli, kv, err := setupEtcd()
		So(err, ShouldBeNil)

		Convey("watch", func() {
			go func() {
				for i := 0; i != 10; i++ {
					kv.Put(context.TODO(), "job", fmt.Sprintf("guard %d", i))
					kv.Delete(context.TODO(), "job")
					time.Sleep(1 * time.Second)
				}
			}()

			watcher := clientv3.NewWatcher(cli)
			ctx, cancelFunc := context.WithCancel(context.TODO())
			time.AfterFunc(10*time.Second, func() {
				cancelFunc()
			})
			var watchRev int64
			watchRespChan := watcher.Watch(ctx, "job", clientv3.WithRev(watchRev))
			for waatchResp := range watchRespChan {
				for _, event := range waatchResp.Events {
					switch event.Type {
					case mvccpb.PUT:
						t.Log("modify as ", string(event.Kv.Value))
					case mvccpb.DELETE:
						t.Log("delete ", string(event.Kv.Key))
					}

				}
			}
		})

		Reset(func() {
			cli.Close()
		})
	})
}

func TestOp(t *testing.T) {
	Convey("testup", t, func() {
		cli, kv, err := setupEtcd()
		So(err, ShouldBeNil)

		Convey("op", func() {
			putOp := clientv3.OpPut("foo", "replace")
			opResp, err := kv.Do(context.TODO(), putOp)
			So(err, ShouldBeNil)
			t.Log("put", opResp.Put())
			getOp := clientv3.OpGet("foo")
			opResp, err = kv.Do(context.TODO(), getOp)
			So(err, ShouldBeNil)
			t.Log("get", opResp.Get().Kvs)
		})

		Reset(func() {
			cli.Close()
		})
	})
}

func TestEtcdLock(t *testing.T) {
	Convey("setup", t, func() {
		cli, _, err := setupEtcd()
		So(err, ShouldBeNil)

		Convey("test lock", func() {
			s1, err := concurrency.NewSession(cli)
			So(err, ShouldBeNil)
			defer s1.Close()

			m1 := concurrency.NewMutex(s1, "/my-lock")
			// s1.Orphan()

			s2, err := concurrency.NewSession(cli)
			So(err, ShouldBeNil)
			defer s2.Close()

			m2 := concurrency.NewMutex(s2, "/my-lock")
			//s2.Orphan()

			err = m1.Lock(context.TODO())
			So(err, ShouldBeNil)

			m2Locked := make(chan struct{})
			go func() {
				defer close(m2Locked)
				ctxTtl, _ := context.WithTimeout(context.Background(), time.Second*5)
				t.Log("m2 wait for lock")

				err := m2.Lock(ctxTtl)
				t.Log("m2 lock", err)
				if err != nil {
					t.Log(err)
				}
			}()

			// time.Sleep(time.Second*10)

			err = m1.Unlock(context.TODO())
			So(err, ShouldBeNil)
			t.Log("release lock for s1")

			<-m2Locked
			t.Log("acquire lock for s2")
		})

		Reset(func() {
			cli.Close()
		})
	})
}
