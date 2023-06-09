FROM alpine:edge AS builder
RUN apk add --no-cache go ca-certificates make git alpine-sdk

WORKDIR /project
COPY . .

RUN echo "nobody:x:65534:65534:Nobody:/:" > .build/passwd
RUN make generate certificate
RUN make build current-platform

FROM scratch AS runner

WORKDIR /var/run
COPY --from=builder --chown=nobody:nobody /project/.build/bin/fizzbuzz /var/run/fizzbuzz
COPY --from=builder --chown=nobody:nobody /project/.build/passwd 	   /etc/passwd
COPY --from=builder --chown=nobody:nobody /project/.build/certs 	   /etc/ssl/certs
USER nobody

HEALTHCHECK --interval=1m --timeout=30s --retries=3 CMD curl --fail http://localhost/api/v1/status/healthcheck || exit 1
EXPOSE 80 443

ENTRYPOINT [ "/var/run/fizzbuzz" ]