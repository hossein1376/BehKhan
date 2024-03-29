# Catalogue

## Generate protobuf and gRPC files

```shell
protoc internal/interface/grpc/pb/catalogue.proto --go_out internal/interface/grpc/pb/
protoc internal/interface/grpc/pb/catalogue.proto --go-grpc_out internal/interface/grpc/pb/
```
