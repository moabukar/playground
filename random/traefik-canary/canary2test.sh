#!/bin/bash

# Set the number of requests you want to send
NUM_REQUESTS=100

# Set the URL of your app
URL="http://myapp.local"

# Loop for the specified number of requests
for (( i=1; i<=NUM_REQUESTS; i++ ))
do
  # Send the request
  curl -s $URL > /dev/null

  # Print out which request number is being sent
  echo "Sent request $i"

  # Sleep for a second to give some gap between requests
  sleep 1
done

echo "Completed sending requests to $URL"
