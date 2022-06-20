FROM alpine

COPY notify /app/notify

CMD ["/app/notify"]