# Go Micro Web Framework
<hr/>
Easter is a web framework written in Go and Gin. It features a API. If you need performance and good productivity, you will love it.
<br/>
Easter's key features are well encapsulated:<br/>

- router
- controller
- service
- model
- log
- db

# Getting started
<hr/>

## install

```
go mod tidy
```

## running

```
go run main.go -e dev -p 8080
```

Then visit 0.0.0.0:8080/api/test in your browser to see the response!
```
{
    "code": 0,
    "content": {
        "data": "hello world!"
    },
    "message": ""
}
```



