FROM zeyanlin/mod AS builder

FROM alpine:3.8
RUN apk add --no-cache ca-certificates tzdata

ENV NOTIFIER_DEBUG=false
ENV NOTIFIER_PORT=8080
ENV NOTIFIER_ENABLEADMIN=false
ENV NOTIFIER_PARSEMODE=Markdown
WORKDIR /home
COPY --from=builder /go/bin/notifier /usr/local/bin/notifier
ENTRYPOINT ["notifier"]
