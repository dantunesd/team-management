FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ./api ./cmd/main.go

FROM alpine:latest as app

WORKDIR /app

COPY --from=build /app/api ./api
COPY --from=build /app/.env ./.env

CMD ["./api"]