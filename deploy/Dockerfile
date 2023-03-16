FROM golang:alpine as builder

RUN apk update && apk upgrade
RUN apk add make git

WORKDIR /project
COPY / /project

RUN make build current-platform

FROM scratch as runner
COPY --from=builder /project/.build/fizzbuzz /fizzbuzz
ENTRYPOINT [ "/fizzbuzz" ]