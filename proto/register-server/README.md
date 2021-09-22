
```shell
protoc -I proto/register-server/ -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis proto/register-server/register_server.proto --go_out=proto/register-server --go-grpc_out=require_unimplemented_servers=false:proto/register-server

```
