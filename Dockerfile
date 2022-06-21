FROM --platform=$BUILDPLATFORM  golang:1.17 as builder

WORKDIR /app

COPY ./* /app/

RUN GOPROXY="https://goproxy.cn,direct" CGO_ENABLED=0 go build -ldflags="-s -w" -o notify .
RUN chmod +x /app/notify

FROM scratch

COPY --from=builder /app/notify /

CMD ["/notify"]



