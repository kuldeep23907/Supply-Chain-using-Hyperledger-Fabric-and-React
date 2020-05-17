#!/bin/bash
#
# Exit on first error, print all commands
set -e

echo "===================================== Attempting to start network =================================================="
echo ""
# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATH_CONV=1
echo ""

echo "====================================== Removing existing volumes ==================================================="
echo ""
docker-compose -f artifacts/docker-compose.yaml down

export MANUFACTURER_CA_PRIVATE_KEY=$(cd ./artifacts/network/crypto-config/peerOrganizations/manufacturer.example.com/ca && ls *_sk)
export MIDDLEMEN_CA_PRIVATE_KEY=$(cd ./artifacts/network/crypto-config/peerOrganizations/middlemen.example.com/ca && ls *_sk)
export CONSUMER_CA_PRIVATE_KEY=$(cd ./artifacts/network/crypto-config/peerOrganizations/consumer.example.com/ca && ls *_sk)
echo ""

echo "======================================= Creating peers and orgs ===================================================="
echo ""
docker-compose -f artifacts/docker-compose.yaml up -d

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=5
# echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}
echo ""

echo "===================================== Started network successfully ================================================="
