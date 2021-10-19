
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

<br>

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

***Response:***

```
{
    "status": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJuYW1lIjoiRmF6cmkiLCJwaG9uZSI6IjA4MTIzNDU2NzExIiwicm9sZSI6ImFkbWluIiwicGFzc3dvcmQiOiJmeW84In0sImlhdCI6MTYzNDY1OTQ5NywiZXhwIjoxNjM0NjYzMDk3fQ.ni8SZ0bh5gkqO0Roz78ZsAMsEmW2Wp76x4RmO3atHHs"
}
```

### 2. Profile

***Authorization:***

```
Type: Bearer Token
Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoyLCJuYW1lIjoiRmF6ciIsInBob25lIjoiMDgxMjM0NTY3MTAiLCJyb2xlIjoidXNlciIsInBhc3N3b3JkIjoidnFyOCJ9LCJpYXQiOjE2MzQ2NTQ1MjAsImV4cCI6MTYzNDY1ODEyMH0.lt8A-ijvFg8G0Cjb_QX8d9NK_cfa5ctrWXryVeDOECw
```

***Endpoint:***

```bash
Method: POST
URL: localhost:3000/users/profile
```

***Response:***

```
{
    "status": "success",
    "data": {
        "user": {
            "id": 2,
            "name": "Fazr",
            "phone": "08123456710",
            "role": "user",
            "password": "vqr8"
        },
        "iat": 1634654520,
        "exp": 1634658120
    }
}
```


### 3. Registration



***Endpoint:***

```bash
Method: POST
URL: localhost:3000/users/register/:name/:phone/:role
```

***Response:***

```
{
    "message": "Registration successful",
    "user": {
        "phone": "08123456711",
        "password": "hgv1"
    }
}
```

---
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam)
