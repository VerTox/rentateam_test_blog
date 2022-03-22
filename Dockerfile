##
## Build
##
FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /rentateam_test_blog

##
## Deploy
##
FROM alpine:latest

ENV TZ=Europe/Moscow

RUN apk update && apk add --update  tzdata && rm -rf /var/cache/apk/*

WORKDIR /

COPY --from=build /rentateam_test_blog /rentateam_test_blog

EXPOSE 9090

CMD ["/rentateam_test_blog"]