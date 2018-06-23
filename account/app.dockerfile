FROM golang:1.10.2-alpine3.7 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /Users/christianvestre/src/github.com/spidey/account
COPY vendor ../vendor
COPY account ./
RUN go build -o /Users/christianvestre/bin/app ./cmd/account/main.go

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /Users/christianvestre/bin .
EXPOSE 8080
CMD ["app"]