FROM golang:latest
WORKDIR /build
COPY *.go /build/
COPY *.sh /build/
RUN go build -o app *.go
CMD ["./app"]
