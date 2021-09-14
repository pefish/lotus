
```shell
protoc -I proto/distribute-prover/ -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis proto/distribute-prover/distribute_prover.proto --go_out=proto/distribute-prover --go-grpc_out=require_unimplemented_servers=false:proto/distribute-prover

```
