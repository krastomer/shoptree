FROM golang AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/web/build ./web/build
EXPOSE 8080
ENTRYPOINT [ "./main" ] 