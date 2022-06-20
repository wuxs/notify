FROM --platform=$BUILDPLATFORM  golang:alpine as builder

WORKDIR /app

COPY ./* /app/

RUN GOPROXY="https://goproxy.cn,direct" go build -ldflags="-s -w"  -o notify .
RUN chmod +x /app/notify

FROM scratch

COPY --from=builder /app/notify /

CMD ["/notify"]



