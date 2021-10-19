const express = require("express")
const { port } = require('./config.json')
const { verifytoken, register, login, profile} = require('./users/controller')
const errorHandler = require('./middleware/error-handler');

const app = express()
app.use(express.json());

// api route
app.post("/users/register/:name/:phone/:role", register)
app.post("/users/login/:phone/:password", login)
app.post("/users/profile", verifytoken, profile)


// global error handler
app.use(errorHandler)

app.listen(port, (req, res) => {
    console.log('server started on posrt 3000')
})
