const userService = require('./service');
const { validatePhone, validateName, validateRole, validatePassword } = require('../utils/validator')

function verifytoken(req, res, next) {
    const bearerHeader = req.headers['authorization']
    if (!bearerHeader) return res.status(403).json({message: "Unauthorized"})
        
    const beareToken = bearerHeader.split(' ')[1]
    req.token = beareToken
    next()
    
}

function register(req, res, next) {
    if (!validateName(req.params.name) || !validatePhone(req.params.phone) || !validateRole(req.params.role)) {
        res.status(403).json({message: "Bad Request"})
    }
    
    userService.CreateUser(req.params)
    .then(data => res.status(200).json({
        message: "Registration successful", 
        user: data.user}))
    .catch(err => res.status(500).send(err))
}

function login(req, res, next) {
    if (!validatePhone(req.params.phone)|| validatePassword(req.params.password)) {
        res.status(403).json({message: "Bad Request"})
    }
   userService.Auth(req.params)
    .then(token => res.status(200).json(token))
    .catch(err => {
        if (err.message == "user not found") {
            res.status(404).send({message: err.message})
        }
        res.status(500).send(err)
    })
}

function profile(req, res, next) {
    userService.GetUser(req.headers)
    .then(data => res.status(200).json(data))
    .catch(err => {

        if (err.message == "invalid token") {
            res.status(403).json({message: err.message})
        }

        res.status(500).send(err)
    })
}

module.exports = {
    verifytoken,
    register,
    login,
    profile
}
