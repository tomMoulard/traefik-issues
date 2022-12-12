FROM alpine:3.15

RUN apk --no-cache add tor
COPY ./torrc /torrc

USER tor

ENTRYPOINT tor -f /torrc
