FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /var/www

ADD . .

RUN go build

FROM alpine

ENV TZ Asia/Shanghai

WORKDIR /app/project

COPY --from=builder /var/www ./

EXPOSE 8002

RUN chmod +x /app/project/activity

CMD ["./activity"]
