##sisyphus

**blog test**

### features

- gin
- gorm

### note
- mysql allow visit 
```
mysql -uroot -pneon
use mysql;
show tables;
select host from user;
update user set host='%' where user = 'root'
```

- build
```
set GOOS=linux 
set GOARCH=amd64 
set CGO_ENABLED=0 
go build
```

- docker rm exit containers
```
docker rm -v $(docker ps -aq -f status=exited)
```

- docker run redis

``` 
docker run --rm -itd -p 6379 
-v /data/redis:/data 
-v /home/phoenix/workspace/gowork/sisyphus/conf/redis.conf:/etc/redis/redis.conf 
--name sisyphus redis 
redis-server /etc/redis/redis.conf
```

- run exam

```$xslt
docker run --rm --name='sisyphus' -p 8080:8080 
-v /home/phoenix/workspace/gowork/sisyphus/test:/app/static 
-v  /home/phoenix/workspace/gowork/sisyphus/conf/app.ini:/app/conf/app.ini   
phoenix/sisyphus:v1.0

```

### 