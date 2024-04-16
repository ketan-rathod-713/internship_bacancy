const express = require("express")

const app = express();
var PORT = 8080;

app.get("/", (req, res)=>{
    res.send("This is overall change and this is further change wow great initial wow amazing")
})

app.listen(PORT, ()=>{
    console.log(`Server is running on port ${PORT}`)
})