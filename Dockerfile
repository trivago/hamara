FROM golang:1.15-alpine as build
COPY . /src/
WORKDIR /src
RUN set -x \
  && apk add --no-cache git \
  && go build -o hamara .

FROM alpine
WORKDIR /app
COPY --from=build /src/hamara /app/
RUN set -x \
  && addgroup -g 5000 -S hamara \
  && adduser -u 5000 -H -D -S -s /sbin/nologin -G hamara hamara
USER hamara
ENTRYPOINT ["/app/hamara"]
CMD ["-h"]
