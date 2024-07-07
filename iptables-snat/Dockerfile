FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -y iptables iproute2 && \
    apt-get clean

COPY start.sh /start.sh
RUN chmod +x /start.sh

CMD ["/start.sh"]
