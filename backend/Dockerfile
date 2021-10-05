FROM golang:1.17.1 AS builder

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -a -installsuffix cgo -o main cmd/server/main.go

FROM busybox

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080

ENTRYPOINT [ "./main" ]