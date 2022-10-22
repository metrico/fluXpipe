FROM alpine as builder
RUN apk add --allow-untrusted --update --no-cache curl ca-certificates
WORKDIR /
RUN curl -fsSL github.com/metrico/fluxpipe/releases/latest/download/fluxpipe-server -O && chmod +x fluxpipe-server

FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /fluxpipe-server /fluxpipe-server
CMD ["/fluxpipe-server", "-port", "8086"]
