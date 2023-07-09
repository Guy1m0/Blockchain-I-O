# Blockchain-I-O
In this work, we enumerate a set of additional desirable functionalities required as building blocks for versatile cross-chain commerce, and describe a modular technology stack, BlockChain I/O (see Figure 1) that achieves these desiderata.


# PIEChain

We develop a general cross-chain communication framework that uses the Kafka protocol for secure interaction.

![alt text](blockarch.png)

## Entities
* The Kafka network provider, who maintains the Kafka network.
* Developers, who develops cross-chain services (CC-SVCs) using an event-
driven approach.
* End users, who deploy smart contracts and use the CC-SVCs for cross- chain operations.

## Setup

### Kafka
Start Kafka with `docker-compose`.
```bash
cd kafka
docker-compose up -d

# to stop after experiments
docker-compose down
```

### Fabric
Install Hyperledger Fabric.
```bash
cd fabric
./install.sh
```

### Ethereum
Install go-ethereum and quorum.
```bash
cd ethereum
sudo ./install.sh
```

## Cross-chain Auction

1. Run Fabric.
```bash
cd fabric-samples/test-network
./network.sh up createChannel -ca
```
Rename cert file name
```bash
cd ../..
./rename_cert.sh
```

2. Run Ethereum POA.
```bash
cd ethereum/poa
./remove.sh
./init.sh
./start.sh
```

3. Run Quorum Raft.
```bash
cd ethereum/raft
./remove.sh
./init.sh
./start.sh
```

4. Deploy `fabric_asset` on `fabric`.
```bash
sudo ./network.sh deployCC -ccn asset -ccp ../../contracts/fabric_asset/chaincode -ccl go
```

5. Run `relayer` crosschain service.
```bash
cd examples/auction/relayer
go build .
./relayer
```

6. Run `signer` crosschain services.

```bash
cd examples/auction/signer
go build .

# on different terminals
./signer -t -p ethereum -eth localhost:8545 -key ../../keys/key1 -id 1
./signer -p ethereum -eth localhost:8545 -key ../../keys/key2 -id 2
./signer -p quorum -eth localhost:8546 -key ../../keys/key1 -id 1
./signer -p quorum -eth localhost:8546 -key ../../keys/key2 -id 2
```

7. Run the `scenario` script.
```bash
cd examples/auction/scenario
go run .
```

The scenario script will do the following steps.
1. Add a new asset on the fabric `asset` contract.
3. Deploy auction contracts on `ethereum` and `quorum`.
2. Create a new auction for this asset on on fabric.
3. Bid correspondingly on both `ethereum` and `quorum`.
4. End auction on `fabric` and print out the winner info and final asset owner on `fabric`.