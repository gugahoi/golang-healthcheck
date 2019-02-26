FROM golang:1.11 as tester
WORKDIR /go/src/github.com/gugahoi/golang-healthcheck
ADD ./ ./
RUN make test

FROM golang:1.11 as builder
WORKDIR /go/src/github.com/gugahoi/golang-healthcheck
ADD ./ ./
ARG TAG
RUN make build

FROM scratch
COPY --from=builder /go/src/github.com/gugahoi/golang-healthcheck/build/golang-healthcheck /app
EXPOSE 80
ENTRYPOINT [ "/app" ]
