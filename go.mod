module sisyphus

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/docker v1.13.1
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.52.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.3.3
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.13.0
	github.com/jinzhu/gorm v1.9.12
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.4.1 // indirect
	github.com/prometheus/common v0.9.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/unknwon/com v1.0.1
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200210222208-86ce3cb69678 // indirect
	golang.org/x/mod v0.2.0 // indirect
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	golang.org/x/sys v0.0.0-20200212091648-12a6c2dcc1e4 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200211205636-11eff242d136 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20200212174721-66ed5ce911ce
	google.golang.org/grpc v1.27.1
	gopkg.in/ini.v1 v1.52.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
