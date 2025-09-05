# builder stage
FROM registry.suse.com/bci/bci-base:15.5 AS builder

# fetched from goreleaser build proccess
COPY freighter /freighter

RUN echo "freighter:x:1001:1001::/home/freighter:" > /etc/passwd \
&& echo "freighter:x:1001:freighter" > /etc/group \
&& mkdir /home/freighter \
&& mkdir /store \
&& mkdir /fileserver \
&& mkdir /registry

# release stage
FROM scratch AS release

COPY --from=builder /var/lib/ca-certificates/ca-bundle.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder --chown=freighter:freighter /home/freighter/. /home/freighter
COPY --from=builder --chown=freighter:freighter /tmp/. /tmp
COPY --from=builder --chown=freighter:freighter /store/. /store
COPY --from=builder --chown=freighter:freighter /registry/. /registry
COPY --from=builder --chown=freighter:freighter /fileserver/. /fileserver
COPY --from=builder --chown=freighter:freighter /freighter /freighter

USER freighter
ENTRYPOINT [ "/freighter" ]

# debug stage
FROM alpine AS debug

COPY --from=builder /var/lib/ca-certificates/ca-bundle.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder --chown=freighter:freighter /home/freighter/. /home/freighter
COPY --from=builder --chown=freighter:freighter /freighter /usr/local/bin/freighter

RUN apk --no-cache add curl

USER freighter
WORKDIR /home/freighter