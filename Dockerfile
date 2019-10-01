# Build using golang:1.12 base image
FROM golang:1.12 as base
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o hamara .

# Final image
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=base /build/hamara /app/
WORKDIR /app
CMD ["./hamara"]