# Packages

## Struct generation

```bash
protoc \
--go_out=pb \
--go_opt=paths=source_relative \
--go-grpc_out=pb \
--go-grpc_opt=paths=source_relative \
protos/product_service.proto
```
