server : HelloServer.go
	go build HelloServer.go
client : HelloClient.go
	go build HelloClient.go
all : HelloServer.go HelloClient.go
	go build HelloClient.go
	go build HelloServer.go
