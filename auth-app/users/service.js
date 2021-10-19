const User = require('./models') 
const jwt = require("jsonwebtoken")
const { secret } = require('../config.json')
const sequelize = require('../database')

sequelize.sync({ force: true }).then(() => console.log('db is ready'))

async function CreateUser (req) {

    try {
        let user = await User.findOne({where: {phone: req.phone}})
        if (!user) {
            req.password = Math.random().toString(36).slice(-4);
            user = await User.create(req)
        }

        let data = {
            phone: user.phone,
            password: user.password
        }

        return {status: "success", user: data}
    } catch (error) {
        throw error
    }
    
}

async function GetUser (req) {

    try {
        const beareToken = req['authorization'].split(' ')[1]
        const data = jwt.verify(beareToken, secret)
        
        return {status: "success", data: data}
    } catch (error) {
        throw error
    }
}

async function Auth (req) {

    try {
        const user = await User.findOne({where: {phone: req.phone, password: req.password}})
        if (!user) {
            throw {status: "error", message: "user not found"} 
        }

        const token = jwt.sign({user: user}, secret, { expiresIn: '1h' })
        
        return {status: "success", token: token}
    } catch (error) {
        throw error
    }

}

module.exports = {
    CreateUser,
    GetUser,
    Auth
}
