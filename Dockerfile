FROM containerize/dep AS builder
WORKDIR /go/src/github.com/gotoolkit/notifier
COPY . .
RUN go install .

FROM alpine:3.8
RUN apk add --no-cache ca-certificates tzdata

ENV NOTIFIER_DEBUG=false
ENV NOTIFIER_PORT=8080
ENV NOTIFIER_ENABLEADMIN=false

WORKDIR /home
COPY --from=builder /go/bin/notifier /usr/local/bin/notifier
ENTRYPOINT ["notifier"]