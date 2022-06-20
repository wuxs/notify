FROM --platform=$TARGETPLATFORM  golang:alpine as builder

WORKDIR /app

COPY ./* /app/

RUN GOPROXY="https://goproxy.cn,direct" go build -o /app/notify .

FROM --platform=$TARGETPLATFORM scratch

COPY --from=builder /app/notify /notify

CMD ["/notify"]



