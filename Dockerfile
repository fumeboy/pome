FROM alpine:3.5
RUN apk add --no-cache iptables
RUN mkdir /pome
COPY ./pome /pome/sidecar