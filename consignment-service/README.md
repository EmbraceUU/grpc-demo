# grpc server

## comment

使用protobuf和grpc创建一个简易的server和client

## options

### generate pb file

```text
protoc -I /Projects/caesar-go/consignment-service/proto/consignment/consignment.proto --go_out=plugins=grpc:C://Users/Administrator/Nida/Projects/caesar-go/consignment-service/proto/
```