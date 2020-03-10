##sisyphus

**blog test**

### features

- gin
- gorm

### note
- mysql allow visit 
mysql -uroot -pneon
use mysql;
show tables;
select host from user;
update user set host='%' where user = 'root'

- build

set GOOS=linux 
set GOARCH=amd64 
set CGO_ENABLED=0 
go build


- docker rm exit containers

docker rm -v $(docker ps -aq -f status=exited)

- run exam

docker run --rm --name='sisyphus' -p 8080:8080 
-v /home/phoenix/workspace/gowork/sisyphus/test:/app/static 
-v  /home/phoenix/workspace/gowork/sisyphus/conf/app.ini:/app/conf/app.ini   
phoenix/sisyphus:v1.0

### 