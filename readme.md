# Go Proto Cli
```bash
go build main.go
``` 

### Add Proto File
Added proto files will be located in **proto/gen** folder as **example_pb/example.proto**
```bash
./main add -name=example
./main add -name=example -http=true #with http server
``` 

### Generate Go Code From Proto File
```bash
./main generate -file=example
``` 