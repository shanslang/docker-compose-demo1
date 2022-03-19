FROM alpine:3.11

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
     && echo "Asia/Shanghai" > /etc/timezone
WORKDIR /workdir/api

COPY release/bin/api /workdir/api/api
COPY config /workdir/api/config

ENTRYPOINT [ "/workdir/api/api"]

