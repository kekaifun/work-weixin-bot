# Build the manager binary
FROM golang:1.23 as builder
WORKDIR /app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY . /app
RUN \
  ls &&\
  export GO111MODULE=on &&\
  export GOPROXY="https://goproxy.woa.com,direct" &&\
  make build

FROM golang:1.23

EXPOSE 80
WORKDIR /app
COPY --from=builder /app/bin/weixin-bot /app/weixin-bot

ENTRYPOINT ["/app/weixin-bot", "--config", "/app/config/config.yaml"]