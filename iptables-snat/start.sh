#!/bin/bash

# Ensure the external IP is provided
if [ -z "$EGRESS_IP" ]; then
  echo "EGRESS_IP environment variable is not set. Exiting."
  exit 1
fi

# Enable IP forwarding
sysctl -w net.ipv4.ip_forward=1

# Clear any existing rules
iptables -t nat -F

# Configure SNAT to the specified IP
iptables -t nat -A POSTROUTING -j SNAT --to-source $EGRESS_IP

# Keep the container running
tail -f /dev/null
