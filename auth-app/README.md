
# Authentication API

This application is used to register new users, login users and to get JWT so that users can access commodity systems.

## Requirement

- express
- jsonwebtoken
- sequelize
- sqlite3

## Config

In this app using json for configuration environment variable

## Running App

`node index.js`



| Key | Value | Type |
| --- | ------|-------------|
| url_auth | http://127.0.0.1:3000/ |  |
| token | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoyLCJuYW1lIjoiRmF6ciIsInBob25lIjoiMDgxMjM0NTY3MTAiLCJyb2xlIjoidXNlciIsInBhc3N3b3JkIjoidnFyOCJ9LCJpYXQiOjE2MzQ2NTQ1MjAsImV4cCI6MTYzNDY1ODEyMH0.lt8A-ijvFg8G0Cjb_QX8d9NK_cfa5ctrWXryVeDOECw |  |

<br>


## Indices

* [auth api](#auth-api)

  * [Login](#1-login)
  * [Profile](#2-profile)
  * [Registration](#3-registration)


--------


## auth api



### 1. Login



***Endpoint:***

```bash
Method: POST
Type: 
URL: localhost:3000/users/login/:phone/:password
```



### 2. Profile



***Endpoint:***

```bash
Method: POST
Type: 
URL: localhost:3000/users/profile
```



### 3. Registration



***Endpoint:***

```bash
Method: POST
Type: 
URL: localhost:3000/users/register/:name/:phone/:role
```

---
[Back to top](#efishery-test)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-10-19 22:27:42 by [docgen](https://github.com/thedevsaddam/docgen)
