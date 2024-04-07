#!/bin/bash
set -e

# Configure the runner
./config.sh --url https://github.com/moabukar/playground --token <TOKEN_HERE>

# Run the runner
./run.sh
