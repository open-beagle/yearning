ARG BASE

FROM $BASE

ARG AUTHOR
ARG VERSION
LABEL maintainer=$AUTHOR version=$VERSION

ARG TARGETOS
ARG TARGETARCH

EXPOSE 8000

COPY dist/yearning-$TARGETOS-$TARGETARCH  /opt/Yearning
COPY .beagle/conf.toml /opt/conf.toml

RUN apk add --no-cache tzdata libc6-compat && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >> /etc/timezone && \
    echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

WORKDIR /opt

CMD /opt/Yearning install && /opt/Yearning run
