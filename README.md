# golang-healthcheck

[![Build Status](https://travis-ci.org/gugahoi/golang-healthcheck.svg?branch=master)](https://travis-ci.org/gugahoi/golang-healthcheck)
[![Coverage Status](https://coveralls.io/repos/github/gugahoi/golang-healthcheck/badge.svg?branch=master)](https://coveralls.io/github/gugahoi/golang-healthcheck?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/6c6ee4069794a0141aa2/maintainability)](https://codeclimate.com/github/gugahoi/golang-healthcheck/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/gugahoi/golang-healthcheck)](https://goreportcard.com/report/github.com/gugahoi/golang-healthcheck)

Basic healthcheck application with a web api style endpoint.

It has the following implemented endpoint:

- A `/healthcheck` which returns basic information about your application: Applications Version, Description, Last Commit SHA.

## Deployments

This project produces docker images on each git tag. To run the image simply convert the git tag to the docker tag, for example to use version `0.0.3`:

```bash
> docker run -it --rm -p 80 gugahoi/golang-healthcheck:0.0.3
2019/02/26 02:13:54 Listening on :80
```

or to test the endpoint with an ephemeral port:

```bash
> ID=$(docker run -it --rm -p 80 -d gugahoi/golang-healthcheck:123)
> curl -v $(docker port ${ID} 80)/healthcheck
*   Trying 0.0.0.0...
* TCP_NODELAY set
* Connected to 0.0.0.0 (127.0.0.1) port 32769 (#0)
> GET /healthcheck HTTP/1.1
> Host: 0.0.0.0:32769
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 26 Feb 2019 02:16:34 GMT
< Content-Length: 108
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 0.0.0.0 left intact
{"version":"123","lastcommitsha":"ff60f240e72cabcd86b1cac852d156abffdee837","description":"Web API for ANZ"}
```
> Note that the application listens on port `80` but if mapping it to different port in the host, access it via the host port.
