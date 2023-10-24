FROM golang:1.20-alpine as buildbase

WORKDIR /go/src/github.com/rarimo/questn-integraion-svc
RUN apk add build-base
COPY vendor .
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/questn-integraion-svc github.com/rarimo/questn-integraion-svc

###
FROM alpine:3.9
COPY --from=buildbase /usr/local/bin/questn-integraion-svc /usr/local/bin/questn-integraion-svc

ENTRYPOINT ["questn-integraion-svc"]
