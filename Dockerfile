FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
# Verify correct gpg key
RUN apt-key fingerprint 0EBFCD88
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
RUN apt-get update && apt-get install docker-ce-cli

WORKDIR /hackathon
ADD . /hackathon/

ENTRYPOINT [ "/hackathon/entrypoint.sh"]
