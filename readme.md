# Go Proto Cli
A tool for both generating .proto files and go codes from them (including http (using <a href="https://github.com/jarium/protoc-gen-http" target="_blank">plugin</a>), grpc server codes).
```bash
go install github.com/jarium/go-proto-cli@latest
``` 

### Init
* Install Google Protocol Buffers compiler (protoc) on your system to be able to run protoc cli commands.
* Then run command to initiate the tool:
```bash
go-proto-cli initiate
```

### Add Proto File
```bash
go-proto-cli add -name=example
go-proto-cli add -name=example -http=true #with http server
``` 

### Generate Go Code From Proto File
```bash
go-proto-cli generate -name=example
go-proto-cli generate -name=example -lib=gin #generate http code using gin library
``` 