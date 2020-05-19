#!/bin/bash
#
# Exit on first error, print all commands
echo "===================================================================================================================="
echo ""
export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}/artifacts
export CHANNEL_NAME=supplychainchannel
# remove previous crypto material and config transactions
rm -fr ../artifacts/network
echo ""

echo "===================================== Generating crypto material ==================================================="
echo ""
# generate crypto material
cryptogen generate --config=./artifacts/crypto-config.yaml --output=./artifacts/network/crypto-config
if [ "$?" -ne 0 ]; then
    echo "Failed to generate crypto material!"
    exit 1
fi
echo ""

echo "======================================= Creating genesis block ====================================================="
echo ""
# generate genesis block for orderer
configtxgen -profile TraceOrdererGenesis -outputBlock ./artifacts/network/genesis.block
if [ "$?" -ne 0 ]; then
    echo "Failed to generate orderer genesis block!"
    exit 1
fi
echo ""

echo "============================= Generating channel configuration transaction ========================================="
echo ""
# generate channel configuration transaction
configtxgen -profile TraceOrgsChannel -outputCreateChannelTx ./artifacts/network/channel.tx -channelID $CHANNEL_NAME
if [ "$?" -ne 0 ]; then
    echo "Failed to generate channel configuration transaction!"
    exit 1
fi
echo ""

echo "============================= Generating anchor peer update for ManufacturerMSP ===================================="
echo ""
# generate anchor peer transaction
configtxgen -profile TraceOrgsChannel -outputAnchorPeersUpdate ./artifacts/network/ManufacturerMSPanchors.tx -channelID $CHANNEL_NAME -asOrg ManufacturerMSP
if [ "$?" -ne 0 ]; then
    echo "Failed to generate anchor peer update for ManufacturerMSP!"
    exit 1
fi
echo ""

echo "============================= Generating anchor peer update for MiddleMenMSP ======================================"
echo ""
configtxgen -profile TraceOrgsChannel -outputAnchorPeersUpdate ./artifacts/network/MiddleMenMSPanchors.tx -channelID $CHANNEL_NAME -asOrg MiddleMenMSP
if [ "$?" -ne 0 ]; then
    echo "Failed to generate anchor peer update for MiddleMenMSP!"
    exit 1
fi
echo ""

echo "============================= Generating anchor peer update for ConsumerMSP ========================================"
echo ""
configtxgen -profile TraceOrgsChannel -outputAnchorPeersUpdate ./artifacts/network/ConsumerMSPanchors.tx -channelID $CHANNEL_NAME -asOrg ConsumerMSP
if [ "$?" -ne 0 ]; then
    echo "Failed to generate anchor peer update for ConsumerMSP!"
    exit 1
fi
echo ""

echo "===================================== Ready to up the network ======================================================"
