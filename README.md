# Go Talk - RPC
## Executing Functions Across A Network (aka RPC) and Why?

### Instructions
#### Install
```bash
# install go present tool
go get -u -v golang.org/x/tools/cmd/present
```
#### Run Presentation
```bash
cd $GOPATH/src/github.com/snassr/talk-0001-RPC
present
```

#### Run Demo
```bash
# run server
cd $GOPATH/src/github.com/snassr/talk-0001-RPC/server && go run server.go
# run client
cd $GOPATH/src/github.com/snassr/talk-0001-RPC/client && go run client.go
```