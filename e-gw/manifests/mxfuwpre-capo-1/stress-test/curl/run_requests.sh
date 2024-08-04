#!/bin/sh


# Function to perform requests
perform_requests() {
  echo "Starting requests to $URL with $NUM_REQUESTS parallel connections"
  seq $NUM_REQUESTS | xargs -n1 -P0 -I {} sh -c 'curl "$0" 2>/dev/null | tee /responses/response_{}.txt' $URL
  echo "Requests completed."
}

# Function to get the current timestamp
get_timestamp() {
  date "+%Y-%m-%d %H:%M:%S"
}


# Check if URL and NUM_REQUESTS are provided
# Check if URL and NUM_REQUESTS environment variables are set
if [ -z "$URL" ] || [ -z "$NUM_REQUESTS" ]; then
    echo "Environment variables URL and NUM_REQUESTS must be set"
    exit 1
fi


# URL=$1
# NUM_REQUESTS=$2

# Create a directory for logging responses
mkdir -p /responses

# Generate a sequence of numbers up to NUM_REQUESTS and pass them to xargs to run in parallel
# seq $NUM_REQUESTS | xargs -n1 -P0 -I {} sh -c 'curl "$0" > /responses/response_{}.txt' $URL
# seq $NUM_REQUESTS | xargs -n1 -P0 -I {} sh -c 'curl "$0" | tee /responses/response_{}.txt' $URL
# seq $NUM_REQUESTS | xargs -n1 -P0 -I {} sh -c 'curl "$0" 2>/dev/null | tee /responses/response_{}.txt' $URL
# Endless loop that checks the RUN variable
while true; do
  if [ "$RUN" = "true" ]; then
    # Check if URL and NUM_REQUESTS environment variables are set
    if [ -z "$URL" ] || [ -z "$NUM_REQUESTS" ]; then
      echo "Environment variables URL and NUM_REQUESTS must be set"
      exit 1
    fi
    
    # Perform the requests
    perform_requests
  else
    timestamp=$(get_timestamp)
    echo "$timestamp - RUN is not set to 'true'. Sleeping for $SLEEP_INTERVAL seconds..."
  fi

  # Sleep for a certain period before the next iteration (e.g., 1 minute)
  sleep "$SLEEP_INTERVAL"
done