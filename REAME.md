# GoPlayground
Working on building a boilerplate Go microservice in here. Stay tuned.

# Getting started for development
1. Ensure you are using Go version >= 1.13
`go version`

2. Install the protoc compiler version >= 3
    ```
    apt install -y protobuf-compiler
    protoc --version
    ```

3. Install the gRPC for Go
    ```
    Add to your RC file
    export GO111MODULE=on  # Enable module mode (https://grpc.io/docs/languages/go/quickstart/)
    export PATH="$PATH:/home/douglas/go/go1.14.2/bin/"
    
    $ go get github.com/golang/protobuf/protoc-gen-go
    ```

4. Install the following additional packages
    ```
    $ go get -u google.golang.org/grpc
    $ go get github.com/golang/protobuf/protoc-gen-go@v1.3
    ```

## GO111MODULE versus GOPATH
GOPATH was the original mechanism for storing third part source files using `go get`. Go get
would get the source code and then store it in $GOPATH/src. Go Modules later got introduced in
Go 1.11 and instead of using GOPATH to store a single git checkout of every package, Go modules
stored and tagged with `go.mod` keep track of every package and their specific version. Here since
we are using `GO111MODULE` we are requiring Go to use Modules instead of GOPATH which means that
`go get` should be populating the `go.mod` file in the root of my project structure.

## gRPC and protocol buffers
https://grpc.io/docs/what-is-grpc/introduction/

RPC (Remote procedure calls) enable a client to call a method on a server

In gRPC, a service is defined with methods signatures that can be called remotely specified. The clients will
offer this interface using a gRPC stuff, while the server will actually implement the interface.

gRPCs use protocol buffers which is a method for serializing structured data (like JSON). First, you define the
structure to be serialized in a `.proto` file and then use the `protoc` compile to generate data access classes
in an allowed language in order to obtain getters, setters, and serialization methods for that class.

Protobuff releases: https://github.com/protocolbuffers/protobuf/releases

gRPC services are also defined in `.proto` files, and then `protoc` with a special gRPC plugin can be used to
generate server and client code from your proto file along with the standard getters, setters, and serialization
methods mentioned above.

## Resources Used
https://ewanvalentine.io/microservices-in-golang-part-x/ where 1<= x <= 10 
