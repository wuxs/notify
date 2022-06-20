FROM --platform=$TARGETPLATFORM  golang:1.17 as builder

COPY ./* /app/

RUN go build -o /app/notify /app/notify.go

FROM --platform=$TARGETPLATFORM scratch

COPY --from=builder /app/notify /notify

CMD ["/notify"]



