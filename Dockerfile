FROM golang:alpine3.10 AS build

LABEL maintainer="Kevin Zheng <kevinzhenng@gmail.com>"

WORKDIR /pong

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o pong .

RUN pwd

FROM alpine:3.10

RUN apk add ca-certificates

COPY --from=build /pong/pong .

CMD ["./pong"]
