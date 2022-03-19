FROM alpine:3.11

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
     && echo "Asia/Shanghai" > /etc/timezone
WORKDIR /workdir/client

COPY release/bin/client /workdir/client/client

ENTRYPOINT [ "/workdir/client/client"]

