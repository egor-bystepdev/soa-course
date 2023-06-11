# Mafia

## Билд докера (но образ есть и в докерхабе)

```
docker build -t dofe71/mafia_client -f client/Dockerfile .

docker build -t dofe71/mafia_server -f server/Dockerfile .

docker build -t dofe71/soa_graphql -f sessions_stat_service/Dockerfile .
```

### Запуск докера

```
docker-compose up --scale client=4
```
Если захочется больше клиентов в одной игре, то нужно в docker-compose поменять MAFIA_PLAYERS_COUNT.

rabbitmq стартует долго, поэтому спам ботов в чат можно будет увидеть через секунд 30.

graphql доступен на localhost:8080

### Сборка:
```
mkdir -p bin && go build -o bin/client client/main.go && go build -o bin/server server/main.go
cd pkg && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
```

Возможно нужно будет поставить protoc, и протобуфы для go, тут есть туториал https://grpc.io/docs/languages/go/quickstart/ (go mod download && go mod verify можно позвать)

Запуск на хосте
```
go mod download && go mod verify
```


```
go run server/main.go
```
Сервер запускается по дефолту на порте 50051 с 4 игроками, но можно установить переменные окружения MAFIA_SERVER_PORT и MAFIA_PLAYERS_COUNT чтобы изменить эти параметры.

Клиент
```
go run client/main.go -p -h
```
-p отвечает за ручной ввод адреса и порта сервера, -h за ручной формат игры (выбираете кого убить, проверить, отправить сообщение). Эти параметры опциональны.

По дефолту подключается к 0.0.0.0:50051, но можно установить переменные окружения MAFIA_SERVER_HOST и MAFIA_SERVER_PORT.

Rabbitmq

```
docker run --rm -it -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

Graphql 

```
go run sessions_stat_service/server.go
```

### Чат

В процессе игры вы увидете возможность отправлять сообщения.

