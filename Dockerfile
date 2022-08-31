FROM golang:1.17-alpine as builder
WORKDIR /root/
COPY . .
RUN apk add --update --no-cache make git

RUN make build

FROM alpine:3.15
WORKDIR /root/

COPY --from=builder /root/bin/ .

RUN chmod +x rpsls-api

EXPOSE 8080
CMD [ "./rpsls-api", "run" ]
