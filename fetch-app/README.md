
# Fetch Commodity

This application is used by the user to access the commodity resources that we have. Here there are 2 APIs with 2 different access rights conditions. The Get Commodity endpoint can be accessed by all existing roles, but the Get Commodity Report can only be accessed by the admin.

## Requirement

- go 1.16+

## Running App

`make run`


| Key | Value | Type |
| --- | ------|-------------|
| url_fetch | http://127.0.0.1:8000/ |  |

<br>


## Indices

* [fetch api](#fetch-api)

  * [Get Commodity](#1-get-commodity)
  * [Get Commodity Report](#2-get-commodity-report)


--------


## fetch api



### 1. Get Commodity



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8000/v1/commodity
```



### 2. Get Commodity Report



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8000/v1/report/summary
```



---
[Back to top](#efishery-test)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-10-19 22:27:42 by [docgen](https://github.com/thedevsaddam/docgen)
