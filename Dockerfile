FROM golang:1.14-alpine AS build

WORKDIR /src

COPY main.go /src

RUN go build -o /bin/app

FROM alpine

COPY --from=build /bin/app /

ENTRYPOINT ["/app"]
