#!/bin/bash
#
# Exit on first error, print all commands
set -e
echo ""
# Shut down the Docker containers that might be currently running.
docker-compose -f artifacts/docker-compose.yaml stop

echo ""
echo "===================================== Stopped network successfully ================================================="