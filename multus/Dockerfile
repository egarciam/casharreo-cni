FROM ubuntu:latest

RUN apt-get update 
RUN apt-get install -y iptables rsyslog 
RUN apt-get clean

# Configure rsyslog to log iptables output to a specific file
RUN echo ':msg,contains,"EGRESS: " /var/log/iptables.log' > /etc/rsyslog.d/iptables.conf && \
    echo '& stop' >> /etc/rsyslog.d/iptables.conf

CMD ["sh", "-c", "rsyslogd; tail -f /var/log/iptables.log"]


