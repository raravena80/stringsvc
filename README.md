# stringsvc

Sample count and uppercase service using go-kit. Includes instrumentation and metrics.

From: [go-kit examples](https://github.com/go-kit/kit/tree/master/examples)

## Run

```
go run *.go
```

## Use

```
$ curl -XPOST http://localhost:8080/count -d '{"s": "my happy string"}'
{"v":15}
curl -XPOST http://localhost:8080/uppercase -d '{"s": "my happy string"}'
{"v":"MY HAPPY STRING"}
curl -XPOST http://localhost:8080/uppercase -d '{"s": "MY HAPPY STRING"}'
{"v":"my happy string"}
```
