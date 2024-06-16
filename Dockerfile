
FROM golang:1.23 AS builder
WORKDIR /app


ENV GOARCH=amd64
ENV GOOS=linux


COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -o server ./cmd/main.go


RUN ls -l /app  # This will show if 'server' is in the /app directory

FROM alpine:latest
WORKDIR /app


COPY --from=builder /app/server .


RUN chmod +x /app/server


EXPOSE 8080

CMD ["/app/server"]
