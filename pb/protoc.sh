protoc \
  -I . \
  --gofast_out=.. \
  --go-grpc_out=.. \
  *.proto
