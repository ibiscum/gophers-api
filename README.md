# gophers-api

This simple API handle a list of Gophers.

## How to install 

### prerequisites

Install Go in 1.16 version minimum.  
Install [Taskfile](https://taskfile.dev/#/installation) (optional)


### Build 

``` 
$ go build -o bin/gophers-api internal/main.go

// or 

$ task build
```

### Run app 

``` 
$ go run internal/main.go

// or 

$ task run
```

### Serve Swagger UI 

This will open you browser on Swagger UI

``` 
$ task swagger:serve
```

### Test the API

* Get all Gophers:

```bash
$ curl localhost:8080/gophers
```

* Get all Gophers with the input name

```bash
$ curl localhost:8080/gophers?name=5th-element
```

* Add a new Gopher

```
$ curl -X POST localhost:8080/gopher \
   -H "Content-Type: application/json" \
   -d '{"name":"yoda-gopher","path":"yoda-gopher.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}'  
```

## Notes

This API use [go-swagger](https://goswagger.io/install.html)