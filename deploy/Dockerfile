FROM alpine:latest AS builder
RUN apk add --no-cache go ca-certificates make git alpine-sdk

WORKDIR /project
COPY / /project

RUN make generate certificate
RUN make build current-platform

FROM scratch as runner
WORKDIR /var/run
COPY --from=builder /project/.build/bin/fizzbuzz /var/run/fizzbuzz
COPY --from=builder /project/.build/certs 		 /var/run/certs
ENTRYPOINT [ "/var/run/fizzbuzz" ]