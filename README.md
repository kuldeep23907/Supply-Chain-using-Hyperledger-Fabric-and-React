# Supply-Chain-HLF-Platform

This projects keeps track record of any product starting from manufacturer to customer.

# Steps to up the network

mkdir channel-artifacts

mkdir crypto-config

## 1. generate certificates

../../bin/cryptogen generate --config=./crypto-config.yaml


## 2. export

export FABRIC_CFG_PATH=$PWD

## 3. generate the genesis block

../../bin/configtxgen -profile TraceOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

## 4. export channel name

export CHANNEL_NAME=supplychainchannel

## 5. generate the channel transaction file

../../bin/configtxgen -profile TraceOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

## 6. update anchor peers

../../bin/configtxgen -profile TraceOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/ManufacturerMSPanchors.tx -channelID $CHANNEL_NAME -asOrg ManufacturerMSP

../../bin/configtxgen -profile TraceOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/MiddleMenMSPanchors.tx -channelID $CHANNEL_NAME -asOrg MiddleMenMSP

../../bin/configtxgen -profile TraceOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/ConsumerMSPanchors.tx -channelID $CHANNEL_NAME -asOrg ConsumerMSP

## 7. up the network

docker-compose -f docker-compose.yaml up -d

## 8. down the network

docker-compose -f docker-compose.yaml down -v

