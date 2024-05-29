# MATHSERVER

Basic client-server application built using gRPC, protobuffers and Gonic API to return fibonacci sequence of given length or check if given number is a prime number.

# Usage

Run `server/server.go` and `api/api.go` in separate terminal sessions and send requests to endpoints below:

1. `http://localhost:8080/fibo/X` where X is length of fibonacci sequence

2. `http://localhost:8080/isprime/X` where X is a number to be verified if it is a prime

# Examples

1. Fibonacci sequence of length X
```bash
$ curl "http://localhost:8080/fibo/10"
{"result":"0 1 1 2 3 5 8 13 21 34"}
```


2. Is X a prime number?
```bash
$ curl "http://localhost:8080/isprime/4"
{"result":"false"}%
$ curl "http://localhost:8080/isprime/3"
{"result":"true"}
```
