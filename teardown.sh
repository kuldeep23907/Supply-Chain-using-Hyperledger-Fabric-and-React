#!/bin/bash
#
# Exit on first error, print all commands
set -e

echo "======================================== Removing volumes =========================================================="
echo ""
# Shut down the Docker containers for the system tests.
docker-compose -f artifacts/docker-compose.yaml down -v
rm -fr artifacts/network
docker rm $(docker ps -aq)
docker kill $(docker ps -aq)
# remove the local store
echo ""
# Your system is now clean
docker ps
echo ""
echo "===================================== Removed volumes successfully ================================================="