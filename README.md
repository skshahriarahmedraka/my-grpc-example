
## Golang grpc :

### install :
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```
### produce code from protocol buffer file :
```
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

## for Golang support 
```
 sudo apt install protobuf-compiler
 ```
 ```
 sudo apt install golang-goprotobuf-dev
```
```
 go get github.com/golang/protobuf
 ```
 ```
 go get github.com/golang/protobuf/proto
```
```
 export PATH=$PATH:$GOPATH/bin
 ```
 example :
```
protoc --go_out=. *.proto
```

```
protoc -I=src/ --go_out=src/ src/simple/simple.proto
```
grpc (golang) : 
```
protoc -I=src/ --go_out=plugins=grpc:src/ src/simple/simple.proto
```

```
protoc -I proto/ proto/currency.proto --go_out=plugins=grpc:proto/
```