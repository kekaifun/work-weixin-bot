# Build the manager binary
FROM golang:1.18 as builder
WORKDIR /workspace
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY . /workspace
RUN \
  ls &&\
  export GO111MODULE=on &&\
  export GOPROXY="https://goproxy.woa.com,direct" &&\
  make build

FROM golang:1.18

EXPOSE 8099
WORKDIR /workspace
COPY --from=builder /workspace/bin/weixin-bot /workspace/weixin-bot

ENTRYPOINT ["/workspace/weixin-bot"]