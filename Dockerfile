FROM golang:bookworm AS builder
LABEL maintainer="joona@kuori.org"


ENV GOPATH=/tmp/buildcache
COPY *.go *.mod *.sum /tmp/acme-dns/
WORKDIR /tmp/acme-dns
RUN CGO_ENABLED=1 go build
RUN mkdir /tmp/dir

FROM gcr.io/distroless/base-debian12:nonroot

USER nonroot:nonroot
WORKDIR /root/
COPY --from=builder /tmp/acme-dns/acme-dns /acme-dns
COPY --from=builder /tmp/dir /etc/acme-dns
COPY --from=builder /tmp/dir /var/lib/acme-dns

VOLUME ["/etc/acme-dns", "/var/lib/acme-dns"]
ENTRYPOINT ["/acme-dns"]
EXPOSE 53 80 443
EXPOSE 53/udp
