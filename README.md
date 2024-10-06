# Protoc-Go-My-Plugin

## Setup
**Step-1:-** Clone the repository
```
git clone https://github.com/CodeAnk2829/protoc-go-my-plugin
```
**Step-2:-** Go to myPlugin root directory
```
cd myPlugin && go mod tidy
``` 
**Step-3:-** Go to the actual file where plugin logic has been written and build the plugin
```
cd cmd/my-plugin && go build -o $GOPATH/bin/protoc-gen-go-my-plugin
```
> NOTE: Check build by running `which protoc-gen-go-plugin`

**Step-4:-** Go back to the root directory i.e. `myPlugin` and compile the proto file by running the following command.
```
protoc --go_out=. --go_opt=paths=source_relative --go-my-plugin_out=. \
--go-my-plugin_opt=paths=source_relative,font=doh example/example.proto
```
> After this you will see two files under example directory has been created

**Step-5:-** Run the main file `go run .`