const express = require("express");
const app = express();
const  importData = require("./data.json");
let port = process.env.port || 3000;

app.get("/" , (req,res) => {
    res.send("Hello world");
});

app.get("/questions", (req,res) => {
    res.send(importData);
});

app.listen(port, () => {
    console.log(`Eg app is listening on port http://localhost:3000`);
})