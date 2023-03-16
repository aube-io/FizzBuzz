FROM alpine:latest AS builder
RUN apk add --no-cache go ca-certificates make git alpine-sdk

WORKDIR /project
COPY . .

RUN echo "nobody:x:65534:65534:Nobody:/:" > .build/passwd
RUN make generate certificate
RUN make build current-platform

FROM scratch as runner

WORKDIR /var/run
COPY --from=builder --chown=nobody:nobody /project/.build/bin/fizzbuzz /var/run/fizzbuzz
COPY --from=builder --chown=nobody:nobody /project/.build/passwd 	   /etc/passwd
COPY --from=builder --chown=nobody:nobody /project/.build/certs 	   /etc/ssl/certs
USER nobody

ENTRYPOINT [ "/var/run/fizzbuzz" ]