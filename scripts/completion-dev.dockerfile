FROM ubuntu:latest

RUN apt update && apt install -y bash-completion && apt clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*
RUN echo "source /etc/bash_completion" >> /root/.bashrc
ADD usacloud-linux /usr/local/bin/usacloud
RUN usacloud completion >> /root/.bashrc
