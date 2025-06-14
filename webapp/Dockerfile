FROM golang:1.24-bookworm AS builder

WORKDIR /src
COPY . .
RUN go build webapp.go

FROM docker.io/almalinux/9-micro:9.5

EXPOSE 8080
WORKDIR /app
COPY --from=builder /src/webapp .

CMD ["/app/webapp"]

