const { Sequelize } = require('sequelize');

const sequelize = new Sequelize('efishery', 'user', 'pass', {
    dialect: 'sqlite',
    host: './dev.sqlite'
})

module.exports = sequelize
