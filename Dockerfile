FROM golang:1.22-alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:3.19 as production

WORKDIR /app
COPY --from=build /app/app .
COPY --from=build /app/public/ public/

CMD ["./app"]
