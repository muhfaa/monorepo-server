
# Fetch Commodity

This application is used by the user to access the commodity resources that we have. Here there are 2 APIs with 2 different access rights conditions. The Get Commodity endpoint can be accessed by all existing roles, but the Get Commodity Report can only be accessed by the admin.

## Requirement

- go 1.16+

## Running App

`make run`


## URL Fetch

<br>

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

***Response:***

```
{
    "code": "success",
    "message": "success",
    "data": [
        {
            "UUID": "f64cb9b6-91e0-46e1-a889-fda2d49f842b",
            "Komoditas": "Cupang Merah",
            "AreaProvinsi": "DKI JAKARTA",
            "AreaKota": "JAKARTA SEELATAN",
            "Size": "198",
            "PriceIDR": "200000",
            "PriceUSD": "+Inf",
            "TglParsed": "2021-01-21T08:50:33.523Z",
            "Timestamp": "1611219033524"
        },
        {
            "UUID": "0b455684-19b3-401e-a57f-e50139756c97",
            "Komoditas": "Gabus",
            "AreaProvinsi": "SUMATERA BARAT",
            "AreaKota": "PADANG PARIAMAN",
            "Size": "10",
            "PriceIDR": "50000",
            "PriceUSD": "+Inf",
            "TglParsed": "2020-08-27T09:04:32.431Z",
            "Timestamp": "1598519072431"
        },
        ...
        ]
}

```


### 2. Get Commodity Report



***Endpoint:***

```bash
Method: GET
Type: 
URL: localhost:8000/v1/report/summary
```

***Response:***

```
{
    "code": "success",
    "message": "success",
    "data": [
        {
            "AreaProvinsi": "ACEH",
            "Min": 30,
            "Max": 190,
            "Median": 12,
            "Average": 42.083333333333336
        },
        {
            "AreaProvinsi": "BALI",
            "Min": 30,
            "Max": 190,
            "Median": 19,
            "Average": 47.63157894736842
        },
        ...
        ]
}
```



---
[Back to top](#efishery-test)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-10-19 22:27:42 by [docgen](https://github.com/thedevsaddam/docgen)
