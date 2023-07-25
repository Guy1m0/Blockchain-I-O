#!/bin/sh

set -x

quorum --datadir data init genesis.json
cp nodekey data/geth