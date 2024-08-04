# Use a lightweight Alpine Linux image
FROM alpine:latest

# Install curl and parallel
RUN apk --no-cache add curl parallel

# Copy the script into the container
COPY run_requests.sh /run_requests.sh

# Make the script executable
RUN chmod +x /run_requests.sh

# Run the script by default
CMD ["/run_requests.sh"]
