FROM alpine

RUN apk add --no-cache \
    bash \
    curl \
    docker-cli \
    docker \
    openrc

# RUN apk add --update docker openrc
RUN rc-update add docker boot

WORKDIR /docker-agent
ADD . /docker-agent/

ENTRYPOINT ["/docker-agent/entrypoint.sh"]