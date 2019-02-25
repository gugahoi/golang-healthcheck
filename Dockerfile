FROM golang:1.11 as builder
WORKDIR /go/src/github.com/gugahoi/golang-healthcheck
ADD ./ ./
RUN CGO_ENABLED=0 go build -o /build/app .

FROM scratch
COPY --from=builder /build/app /app
ENTRYPOINT [ "/app" ]
