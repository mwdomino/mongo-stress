FROM golang:1.21 as builder

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/app/app .
CMD ["./app"]
