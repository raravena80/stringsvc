# stringsvc

Sample count and uppercase service using go-kit. Includes instrumentation and metrics.

From: [go-kit examples](https://github.com/go-kit/kit/tree/master/examples)

## Run

Basic usage:

```
go run *.go
```

Specify port:

```
go run *.go -listen:8020
```

Specify proxy:

```
go run *go -listen:8021 -proxy=localhost:8020
```

## Use

```
$ curl -XPOST http://localhost:8080/count -d '{"s": "my happy string"}'
{"v":15}
curl -XPOST http://localhost:8080/uppercase -d '{"s": "my happy string"}'
{"v":"MY HAPPY STRING"}
curl -XPOST http://localhost:8080/downcase -d '{"s": "MY HAPPY STRING"}'
{"v":"my happy string"}
$ curl -XPOST http://localhost:8080/palindrome -d '{"s": "MY HAPPY STRING"}'
{"v":false}
$ curl -XPOST http://localhost:8080/palindrome -d '{"s": "ana"}'
{"v":true}
```
