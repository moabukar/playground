#!/bin/bash
set -e

TOKEN=${GITHUB_RUNNER_TOKEN}
# Configure the runner
./config.sh --url https://github.com/moabukar/playground --token ${TOKEN}

# Run the runner
./run.sh
