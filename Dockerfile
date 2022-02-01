FROM golang:1.17.6 as builder
WORKDIR /go/src
# for cache
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/main main.go

FROM scratch as runner
COPY --from=builder /go/bin/main /app/main
ENTRYPOINT ["/app/main"]