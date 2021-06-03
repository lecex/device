FROM bigrocs/golang-gcc:1.13 as builder

WORKDIR /go/src/github.com/lecex/device
COPY . .

ENV GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -o bin/device

FROM bigrocs/alpine:ca-data

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/src/github.com/lecex/device/bin/device /usr/local/bin/
CMD ["device"]
EXPOSE 8080
