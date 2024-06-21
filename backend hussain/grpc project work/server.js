const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

// Load the protocol buffer definition
const packageDef = protoLoader.loadSync("todo.proto", {});
const grpcObject = grpc.loadPackageDefinition(packageDef);
const todoPackage = grpcObject.todoPackage;

// Create a gRPC server
const server = new grpc.Server();

// Bind the server to a port
server.bindAsync("0.0.0.0:8080", grpc.ServerCredentials.createInsecure(), (err, port) => {
    if (err) {
        console.error("Server bind failed:", err);
        return;
    }
    console.log(`Server is listening on port ${port}`);
});

// Add service to the server
server.addService(todoPackage.Todo.service, {
    "createTodo": createTodo,
    "readTodos": readTodos,
    "readTodosStream": readTodosStream,
});

// Start the server
// server.start();

// Service method implementations
const todos = []
function createTodo(call, callback) {
    // Dummy implementation, you can replace it with your logic
    console.log("Received createTodo request:", call.request);
    
    // request will be object of id and name
    const todoItem = {
        "id": todos.length + 1,
        "text": call.request.text
    }

    todos.push(todoItem)

    // Invoke the callback with null (no error) and the response message
    // 
    callback(null, todoItem);
}

function readTodos(call, callback) {
    // Dummy implementation, you can replace it with your logic
    console.log("Received readTodos request");
    
    // i just need to send back the result to user

    // schema has to be exactly matched

    // here i was using unneccessary object at first.
    callback(null, { items: todos} );
}

// This is the beautifull thing // streaming.
function readTodosStream(call, callback) {
   // now we can directly write to call
   todos.forEach(todo => call.write(todo))
   call.end()
}
