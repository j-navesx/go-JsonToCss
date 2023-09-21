
build:
	go build -o main main.go
	go build -o client ./client/client.go
	go build -o server ./server/server.go
	go build ./backend/jsoncss.go

run: build
	./main

exec:
	./main

c:
	./client/client

s: build
	./server/server

clean:
	go clean .
	go clean ./server
	go clean ./client
	go clean ./backend

#	rm ${BINARY_NAME}