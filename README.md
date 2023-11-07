# Beh Khan

Beh Khan (meaning 'better read') is a social platform for finding new books to read, rate and write reviews, get
suggestions and more.  
It is built on microservice architecture; and serves as a opportunity to learn and explore.

## Generate protobuf and gRPC files

### Catalogue

```shell
protoc services/structure/cataloguePB/catalogue.proto --go_out services/structure/
protoc services/structure/cataloguePB/catalogue.proto --go-grpc_out services/structure/
```

### Review

```shell
protoc services/structure/reviewPB/review.proto --go_out services/structure/ 
protoc services/structure/reviewPB/review.proto --go-grpc_out services/structure/
```