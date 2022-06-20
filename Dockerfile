FROM --platform=$TARGETPLATFORM  golang:1.17 as builder

WORKDIR /app

COPY ./* /app/

RUN GOPROXY="https://goproxy.cn" go build -o /app/notify .

FROM --platform=$TARGETPLATFORM scratch

COPY --from=builder /app/notify /notify

CMD ["/notify"]



