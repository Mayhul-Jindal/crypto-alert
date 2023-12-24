# builder stage
FROM golang:1.21.5-alpine3.19 as builder

WORKDIR /app

# these layers can be reused because of caching mechanism
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o alert-service .
RUN chmod +x alert-service

# Final stage
FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/alert-service .

EXPOSE 3000

CMD [ "./alert-service" ]