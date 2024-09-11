# Go Proto Cli
```bash
go get github.com/jarium/go-proto-cli
go install github.com/jarium/go-proto-cli
``` 

### Add Proto File
```bash
go-proto-cli add -name=example
go-proto-cli add -name=example -http=true #with http server
``` 

### Generate Go Code From Proto File
```bash
go-proto-cli generate -name=example
``` 