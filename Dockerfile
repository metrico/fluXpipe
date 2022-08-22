FROM alpine as builder
RUN apk add --allow-untrusted --update --no-cache curl
WORKDIR /
RUN curl -fsSL github.com/metrico/fluxpipe/releases/latest/download/fluxpipe-server -O && chmod +x fluxpipe-server

FROM scratch
COPY --from=builder /fluxpipe-server /fluxpipe-server
CMD ["/fluxpipe-server", "-port", "8086"]
