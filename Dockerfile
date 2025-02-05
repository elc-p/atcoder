FROM ubuntu:20.04

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y gcc && \
    ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apt-get install -y python3.9 && \
    apt-get install -y software-properties-common && \
    add-apt-repository -y ppa:longsleep/golang-backports && \
    apt-get update -y && \
    apt-get install -y golang-go

RUN echo 'alias python="python3.9"' >> ~/.bashrc

ENTRYPOINT [ "bash" ]