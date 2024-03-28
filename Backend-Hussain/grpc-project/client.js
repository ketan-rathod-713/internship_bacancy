// it also need to understand all the stuff that it has

// client actually connects and not listens

const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

// Load the protocol buffer definition
const packageDef = protoLoader.loadSync("todo.proto", {});
const grpcObject = grpc.loadPackageDefinition(packageDef);
const todoPackage = grpcObject.todoPackage;

// define where to connect for given service.
const client = new todoPackage.Todo("localhost:8080", grpc.credentials.createInsecure()); // create object of todo service

// client can call methods
// it takes two arguments, one is argument and other is callback which takes errr and response as arguments.
client.createTodo({
    "id": -1,
    "text": "laundry"
}, (err, resp)=> {
    console.log(err, resp)
})

client.readTodos({}, (err, resp)=> {
    console.log(JSON.stringify(resp))

    // resp.items.forEach(a => console.log(a))
})

// it is not a callback anymore
// now we can start listening to events
const call = client.readTodosStream()

call.on("data", item => {
    console.log("received item from server", JSON.stringify(item))
})

call.on("end", e => console.log("server done."))