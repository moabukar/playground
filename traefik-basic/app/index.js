const app = require("express")();
const containerNumber = process.env.CONTAINER_NUMBER || 'no number';

app.get("/", (req, res) => {
    res.send(`Hello from a lightweight container ${containerNumber}`)
})

app.listen(9999, ()=>console.log("Listening on 9999"))
