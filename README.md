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
cd $GOPATH/src/github.com/snassr/talk-0001-rpc
present
```

#### Run Demo
```bash
# run server
cd $GOPATH/src/github.com/snassr/talk-0001-rpc/server && go run server.go
# run client
cd $GOPATH/src/github.com/snassr/talk-0001-rpc/client && go run client.go
```
#### Deploy Presentation
```bash
# find and kill existing process
sudo kill -9 $(lsof -t -i:80)
# go to directory
cd $GOPATH/src/github.com/snassr/talk-0001-rpc/
# run in background
nohup present -http ":80" -play=false &
```