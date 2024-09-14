
// grpc
// request response arch is not scallable
// one way to scale is to use remote procedure calls
// directly call functions from client to server
// instead of JSON, user protobuf.
// size of payload is small hence fast communication.

// use cases
// microservices, bloackchaings peers interactions :D

// graphql need to be learned


# GRPC

First of we need to creat a proto file. then using protoc command line tool we will gennerate go source code.

Define and implement service interface inside server code logic and make it as server.

In client side we can directly initialise client and then call the server functions directly.