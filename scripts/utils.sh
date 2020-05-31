# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
PEER0_MANUFACTURER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/peers/peer0.manufacturer.example.com/tls/ca.crt
PEER0_MIDDLEMEN_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/middlemen.example.com/peers/peer0.middlemen.example.com/tls/ca.crt
PEER1_MIDDLEMEN_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/middlemen.example.com/peers/peer1.middlemen.example.com/tls/ca.crt
PEER2_MIDDLEMEN_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/middlemen.example.com/peers/peer2.middlemen.example.com/tls/ca.crt 
PEER0_CONSUMER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/consumer.example.com/peers/peer0.consumer.example.com/tls/ca.crt 
ORDERER_MSP=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/users/Admin@example.com/msp
MANUFACTURER_MSP=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.example.com/users/Admin@manufacturer.example.com/msp
MIDDLEMEN_MSP=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/middlemen.example.com/users/Admin@middlemen.example.com/msp
CONSUMER_MSP=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/consumer.example.com/users/Admin@consumer.example.com/msp
COLLECTIONS_CONFIG=/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/collections_config.json

# verify the result of the end-to-end test
verifyResult() {
    if [ $1 -ne 0 ]; then
        echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
        echo "========= ERROR !!! FAILED to execute End-2-End Scenario ==========="
        echo
        exit 1
    fi
}

# Set OrdererOrg.Admin globals
setOrdererGlobals() {
    CORE_PEER_LOCALMSPID="OrdererMSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$ORDERER_CA
    CORE_PEER_MSPCONFIGPATH=$ORDERER_MSP
}

# Set Org.Peer globals
setGlobals() {
    PEER=$1
    ORG=$2
    if [ $ORG -eq 1 ]; then
        CORE_PEER_LOCALMSPID="ManufacturerMSP"
        CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_MANUFACTURER_CA
        CORE_PEER_MSPCONFIGPATH=$MANUFACTURER_MSP
        if [ $PEER -eq 0 ]; then
            CORE_PEER_ADDRESS=peer0.manufacturer.example.com:7051
        fi
    elif [ $ORG -eq 2 ]; then
        CORE_PEER_LOCALMSPID="MiddleMenMSP"
        CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_MIDDLEMEN_CA
        CORE_PEER_MSPCONFIGPATH=$MIDDLEMEN_MSP
        if [ $PEER -eq 0 ]; then
            CORE_PEER_ADDRESS=peer0.middlemen.example.com:8051
        elif [ $PEER -eq 1 ]; then
            CORE_PEER_ADDRESS=peer1.middlemen.example.com:9051
        elif [ $PEER -eq 2 ]; then
            CORE_PEER_ADDRESS=peer2.middlemen.example.com:10051
        fi
     elif [ $ORG -eq 3 ]; then
        CORE_PEER_LOCALMSPID="ConsumerMSP"
        CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_CONSUMER_CA
        CORE_PEER_MSPCONFIGPATH=$CONSUMER_MSP
        if [ $PEER -eq 0 ]; then
            CORE_PEER_ADDRESS=peer0.consumer.example.com:11051
        fi
    else
        echo "================== ERROR !!! ORG Unknown =================="
    fi
}

updateAnchorPeers() {
    PEER=$1
    ORG=$2
    setGlobals $PEER $ORG

    if [ -z $CORE_PEER_TLS_ENABLED -o $CORE_PEER_TLS_ENABLED = "false" ]; then
        set -x
        peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx >&log.txt
        res=$?
        set +x
    else
        set -x
        peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
        res=$?
        set +x
    fi
    cat log.txt
    verifyResult $res "Anchor peer update failed"
    echo "===================== Anchor peers updated for org '$CORE_PEER_LOCALMSPID' on channel '$CHANNEL_NAME' ===================== "
    sleep $DELAY
    echo
}

## Sometimes Join takes time hence RETRY at least 5 times
joinChannelWithRetry() {
    PEER=$1
    ORG=$2
    setGlobals $PEER $ORG

    set -x
    peer channel join -b $CHANNEL_NAME.block >&log.txt
    res=$?
    set +x
    cat log.txt
    if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
        COUNTER=$(expr $COUNTER + 1)
        echo "peer failed to join the channel, Retry after $DELAY seconds"
        sleep $DELAY
        joinChannelWithRetry $PEER $ORG
    else
        COUNTER=1
    fi
    verifyResult $res "After $MAX_RETRY attempts, peer has failed to join channel '$CHANNEL_NAME' "
}

installChaincode() {
    PEER=$1
    ORG=$2
    setGlobals $PEER $ORG
    VERSION=${3:-1.0.0}

    set -x
    peer chaincode install -n $CC_NAME -v 1.0 -p $CC_SRC_PATH >&log.txt
    res=$?
    set +x
    cat log.txt
    verifyResult $res "Chaincode installation on peer has failed"
    echo "===================== Chaincode is installed on peer ===================== "
    echo
}

instantiateChaincode() {
    PEER=$1
    ORG=$2
    setGlobals $PEER $ORG
    VERSION=${3:-1.0.0}

    # while 'peer chaincode' command can get the orderer endpoint from the peer
    # (if join was successful), let's supply it directly as we know it using
    # the "-o" option
    if [ -z $CORE_PEER_TLS_ENABLED -o $CORE_PEER_TLS_ENABLED = "false" ]; then
        set -x
        peer chaincode instantiate -o orderer.example.com:7050 -C $CHANNEL_NAME -n $CC_NAME -v 1.0 -c '{"Args":[""]}' -P "OR ('ManufacturerMSP.peer','MiddleMenMSP.peer', 'ConsumerMSP.peer')">&log.txt
        res=$?
        set +x
    else
        set -x
        peer chaincode instantiate -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CC_NAME  -v 1.0 -c '{"Args":[""]}' -P "OR ('ManufacturerMSP.peer','MiddleMenMSP.peer', 'ConsumerMSP.peer')" >&log.txt
        res=$?
        set +x
    fi
    sleep $DELAY
    cat log.txt
    verifyResult $res "Chaincode instantiation on peer on channel '$CHANNEL_NAME' failed"
    echo "===================== Chaincode is instantiated on peer on channel '$CHANNEL_NAME' ===================== "
    echo
}

# parsePeerConnectionParameters $@
parsePeerConnectionParameters() {
    if [ $(($# % 2)) -ne 0 ]; then
        exit 1
    fi

    PEER_CONN_PARMS=""
    PEERS=""
    while [ $# -gt 0 ]; do
        setGlobals $1 $2
        if [ $2 -eq 1 ]; then
            ORG="manager"
            ORG_UPPER="MANAGER"
        else
            ORG="student"
            ORG_UPPER="STUDENT"
        fi

        PEER="peer$1.$ORG"
        PEERS="$PEERS $PEER"
        PEER_CONN_PARMS="$PEER_CONN_PARMS --peerAddresses $CORE_PEER_ADDRESS"
        if [ -z $CORE_PEER_TLS_ENABLED -o $CORE_PEER_TLS_ENABLED = "true" ]; then
            TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER$1_${ORG_UPPER}_CA")
            PEER_CONN_PARMS="$PEER_CONN_PARMS $TLSINFO"
        fi
        shift 2
    done
    # remove leading space for output
    PEERS="$(echo -e "$PEERS" | sed -e 's/^[[:space:]]*//')"
}

# chaincodeInvoke <func> ( <peer> <org> ) ...
chaincodeInvoke() {
    FUNC=$1
    shift
    parsePeerConnectionParameters $@
    res=$?
    verifyResult $res "Invoke transaction failed on channel '$CHANNEL_NAME' due to uneven number of peer and org parameters "

    case $FUNC in
    "manufacturer")
        ARGS='{"Args":["initLedger"]}'
        ;;
    esac

    if [ -z $CORE_PEER_TLS_ENABLED -o $CORE_PEER_TLS_ENABLED = "false" ]; then
        set -x
        peer chaincode invoke -o orderer.example.com:7050 -C $CHANNEL_NAME -n $CC_NAME -c $ARGS >&log.txt
        res=$?
        set +x
    else
        set -x
        peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CC_NAME -c $ARGS >&log.txt
        res=$?
        set +x
    fi
    sleep $DELAY
    cat log.txt
    verifyResult $res "Invoke execution on $PEERS failed "
    echo "===================== Invoke transaction successful on $PEERS on channel '$CHANNEL_NAME' ===================== "
    echo
}