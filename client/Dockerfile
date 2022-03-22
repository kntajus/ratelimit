# syntax=docker/dockerfile:1

FROM golang:1.17.8-alpine3.15 as build
ENV CGO_ENABLED=0
WORKDIR /src
COPY main.go go.* ./
RUN go build -o client .

FROM scratch
COPY --from=build /src/client /client
ENTRYPOINT ["/client"]
