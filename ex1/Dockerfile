FROM golang:alpine3.7 as artifacts

RUN apk update && apk --update add --no-cache git

COPY . /go/src/
WORKDIR /go/src/

RUN go build -o app main.go

# Release
FROM alpine:3.7 as release
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=artifacts /go/src/app /usr/bin/app

WORKDIR /app

CMD ["app"]
