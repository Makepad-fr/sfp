ARG DEBIAN_TAG
FROM golang:1.22.2-bookworm as build

WORKDIR /app

COPY go.mod .
COPY go.work .
COPY go.work.sum .
COPY Makefile .
COPY ./core ./core
COPY ./forward-proxy ./forward-proxy

RUN make build-local

FROM debian:${DEBIAN_TAG}

EXPOSE 9090
WORKDIR /app

COPY --from=build /app/out/forward-proxy ./fp

CMD ["./fp"]