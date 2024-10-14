#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

GOQUORUM_CONS_ALGO=`echo "${GOQUORUM_CONS_ALGO:-qbft}" | tr '[:lower:]'`
GENESIS_FILE=${GENESIS_FILE:-"/data/${GOQUORUM_CONS_ALGO}-${GOQUORUM_GENESIS_MODE}-genesis.json"}

cp -R /config/* /data
mkdir -p /data/keystore/

echo "Applying ${GENESIS_FILE} ..."
geth --nousb --verbosity 1 --datadir=/data init ${GENESIS_FILE}; 

cp /config/keys/accountKeystore /data/keystore/key;
cp /config/keys/nodekey /data/geth/nodekey;

if [ "istanbul" == "$GOQUORUM_CONS_ALGO" ];
then
    echo "Using istanbul for consensus algorithm..."
    export CONSENSUS_ARGS="--istanbul.blockperiod 5 --mine --miner.threads 1 --miner.gasprice 0 --emitcheckpoints"
    export QUORUM_API="istanbul"
elif [ "qbft" == "$GOQUORUM_CONS_ALGO" ];
then
    echo "Using qbft for consensus algorithm..."
    export CONSENSUS_ARGS="--mine --miner.threads 1 --miner.gasprice 0 --emitcheckpoints"
    export QUORUM_API="istanbul"
elif [ "raft" == "$GOQUORUM_CONS_ALGO" ];
then
    echo "Using raft for consensus algorithm..."
    export CONSENSUS_ARGS="--raft --raftblocktime 300 --raftport 53000"
    export QUORUM_API="raft"
fi

export ADDRESS=$(grep -o '"address": *"[^"]*"' /config/keys/accountKeystore | grep -o '"[^"]*"$' | sed 's/"//g')

if [[ ! -z ${QUORUM_PTM:-} ]];
then
    echo -n "Checking tessera is up ... "
    
    curl \
        --connect-timeout 5 \
        --max-time 10 \
        --retry 5 \
        --retry-connrefused \
        --retry-delay 0 \
        --retry-max-time 60 \
        --silent \
        --fail \
        "${QUORUM_PTM}:9000/upcheck"
    echo ""

    ADDITIONAL_ARGS="${ADDITIONAL_ARGS:-} --ptm.timeout 5 --ptm.url http://${QUORUM_PTM}:9101 --ptm.http.writebuffersize 4096 --ptm.http.readbuffersize 4096 --ptm.tls.mode off"
fi

touch /var/log/quorum/geth-$(hostname -i).log
cat /proc/1/fd/2 /proc/1/fd/1 > /var/log/quorum/geth-$(hostname -i).log &

if [ -f /permissions/permission-config.json ]; then
  echo "Using Enhanced Permissions: Copying permission-config.json exists to /data ...";
  cp /permissions/permission-config.json /data/permission-config.json
fi

exec geth \
--datadir /data \
--nodiscover \
--permissioned \
--verbosity 3 \
$CONSENSUS_ARGS \
--syncmode full --revertreason \
--metrics --pprof --pprof.addr 0.0.0.0 --pprof.port 9545 \
--networkid ${QUORUM_NETWORK_ID:-1338} \
--http --http.addr 0.0.0.0 --http.port 8545 --http.corsdomain "*" --http.vhosts "*" --http.api admin,eth,debug,miner,net,txpool,personal,web3,$QUORUM_API \
--ws --ws.addr 0.0.0.0 --ws.port 8546 --ws.origins "*" --ws.api admin,eth,debug,miner,net,txpool,personal,web3,$QUORUM_API \
--port 30303 \
--identity ${HOSTNAME}-${GOQUORUM_CONS_ALGO} \
--unlock ${ADDRESS} \
--allow-insecure-unlock \
--password /data/passwords.txt \
--privacymarker.enable \
--bootnodes=enode://738ce27c421f18a497814428f1f0c4ac44aaf50629f15973b0e20875c884dd1c19c6d7f4b48e56186e8c36ac5a82032b47dd053b75089d6bff5ea1cadc2b31a0@172.16.239.14:30303,enode://079f1d0a6549e57d3a10d8f91532c65acda4125b1ba652ce566be51e4d77f3c0181d872aea4d673e18af7d6c7cc8bed9a1a899e28eed173a0f17888efdaae024@172.16.239.15:30303 \
${ADDITIONAL_ARGS:-} \
2>&1