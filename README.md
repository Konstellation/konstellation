[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/Konstellation/konstellation)

Konstellation is the blockchain built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk). Konstellation will be interact with other sovereign blockchains using a protocol called [IBC](https://github.com/cosmos/ics/tree/master/ibc) that enables Inter-Blockchain Communication.

# Konstellation network

# *Readme needs to be updated

## Testnet Full Node Quick Start
With each version of the Konstellation Hub, the chain is restarted from a new Genesis state. We are currently on knstlhub-2.

Get testnet config [here](https://github.com/Konstellation/testnet)

### Build from code

This assumes that you're running Linux or MacOS and have installed [Go 1.14+](https://golang.org/dl/).  This guide helps you:

Build, Install, and Name your Node
```bash
# Clone Konstellation from the latest release found here: https://github.com/konstellation/konstellation/releases
git clone -b <latest_release> https://github.com/konstellation/konstellation
# Enter the folder Konstellation was cloned into
cd konstellation
# Compile and install Konstellation
make install
```

### Using binaries
```bash
# linux
wget https://github.com/Konstellation/konstellation/releases/download/v0.1.30/linux_amd64.tar.gz
tar -xvzf linux_amd64.tar.gz
sudo cp ./linux_amd64/* /usr/local/bin
# macos
wget https://github.com/Konstellation/konstellation/releases/download/v0.1.30/darwin_amd64.tar.gz
tar -xvzf darwin_amd64.tar.gz
sudo cp ./darwin_amd64/* /usr/local/bin
```
* NOTE: For Windows download archive by [link](https://github.com/Konstellation/konstellation/releases/download/v0.1.30/windows_amd64.tar.gz) , untar archive using 7z and move files to "C:\\Users\user" folder. If you choose a different folder, make sure that it is added to the PATH env variable. [How to add to the PATH on Windows 10](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/)

Replace `v0.1.30` with the latest version or the one that you need

### To join testnet follow this steps

* NOTE: If you are using Windows, add flag *--home D:\\.konstellation* or another folder you choose for all commands.  This folder will be used as home directory and will contain blockchain history and it takes up a lot of disk space

#### Initialize data and folders
```bash
konstellation unsafe-reset-all
```

Remember to add flag *--home D:\\.konstellation* for Windows

#### Genesis & Seeds
Download [genesis.json](https://raw.githubusercontent.com/Konstellation/testnet/master/knstlhub-1/genesis.json)
```
wget -O $HOME/.konstellation/config/genesis.json https://raw.githubusercontent.com/Konstellation/testnet/master/knstlhub-1/genesis.json
```
Download [config.toml](https://raw.githubusercontent.com/Konstellation/testnet/master/knstlhub-1/config.toml) with predefined seeds and persistent peers
```
wget -O $HOME/.konstellation/config/config.toml https://raw.githubusercontent.com/Konstellation/testnet/master/knstlhub-1/config.toml
```

Replace `knstlhub-1` with the latest chain-id
* NOTE: See [testnet repo](https://github.com/Konstellation/testnet) for the latest testnet info.

* NOTE: For Windows open links in browser -> Save As -> Choose "D:\\.konstellation\config" as path

Alternatively enter persistent peers to config.toml provided [here](https://github.com/Konstellation/testnet/tree/master/knstlhub-1)

1) Open ~/.konstellation/config/config.toml with text editor (D:\\.konstellation/config/config.toml for Windows). Alternatively you can use cli editor, like nano ``` nano ~/.konstellation/config/config.toml ```
2) Scroll down to persistant peers in `config.toml`, and add the persistant peers as a comma-separated list

#### Setting Up a New Node
Name your node. Moniker defaults to the machine name
```
konstellation config set moniker <your_moniker>
```

You can edit this moniker later, in the ~/.konstellation/config/config.toml file:
```bash
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

You can edit the ~/.konstellation/config/app.toml file in order to enable the anti spam mechanism and reject incoming transactions with less than the minimum gas prices:
```
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10udarc).

minimum-gas-prices = ""
```
Your full node has been initialized!

#### Run a full node
```
# Start Konstellation
konstellation start
# Check your node's status with konstellationcli
konstellationcli status
```

### Create a key
Add new
``` bash
konstellationcli keys add <key_name>
```

Or import via mnemonic
```bash
konstellationcli keys add <key_name> -i
```

As a result, you got
```bash
- name: <key_name>
  type: local
  address: <key_address>
  pubkey: <key_pubkey>
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

<key_mnemonic>
```

### To become a validator follow this steps
Before setting up your validator node, make sure you've already gone through the [Full Node Setup](https://github.com/Konstellation/konstellation#to-join-testnet-follow-this-steps)

#### What is a Validator?
[Validators](https://hub.cosmos.network/master/validators/overview.html) are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable or sign blocks at the same height.
Please read about [Sentry Node Architecture](https://hub.cosmos.network/master/validators/validator-faq.html#how-can-validators-protect-themselves-from-denial-of-service-attacks) to protect your node from DDOS attacks and to ensure high-availability.

#### Create Your Validator

Your `darcvalconspub` can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

```bash
konstellation tendermint show-validator
```

To create your validator, just use the following command:
 
Don't use more `udarc` than you have! 

```bash
konstellationcli tx staking create-validator \
  --amount=100000000000udarc \
  --pubkey=$(konstellation tendermint show-validator) \
  --moniker=<choose a moniker> \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=<key_name>
```

* NOTE: If you have troubles with \'\\\' symbol, run the command in a single line like `konstellationcli tx staking create-validator --amount=100000000000udarc --pubkey=$(konstellation tendermint show-validator) ...`

When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.

`Min-self-delegation` is a strictly positive integer that represents the minimum amount of self-delegated voting power your validator must always have. A `min-self-delegation` of 1 means your validator will never have a self-delegation lower than `1000000darc`

You can check that you are in the validator set by using a third party explorer or using cli tool
```bash
konstellationcli q staking validators --chain-id <chain_id>
```

* Note: You can edit the params after, by running command `konstellationcli tx staking edit-validator ... —from <key_name> --chain-id=<chain_id>` with the necessary options

## How to init chain

Add key to your keyring
```knstld keys add key1```

Initialize genesis and config files 
```knstld init node-knstld-13 --chain-id darcmatter```

Add genesis account
```knstld add-genesis-account key1 200000000000udarc``` - 200000darc

Create genesis transaction
```knstld gentx --name key1 100000000000udarc``` - create CreateValidator transaction

Collect all of gentxs
```knstld collect-gentxs```

Run network
```knstld start```



## Dockerized

We provide a docker image to help with test setups. There are two modes to use it

Build: ```docker build -t knstld:latest .```  or pull from dockerhub ```kirdb/knstld:latest```

### Dev server
Bring up a local node with a test account containing tokens

This is just designed for local testing/CI - do not use these scripts in production. Very likely you will assign tokens to accounts whose mnemonics are public on github.
Prepend `VOLTYPE=vol|b` if you want to bind mount or volume into container as storage

#### Set IMAGE env variable
```shell script
export IMAGE=kirdb/knstld:0.2.0
```
#### Set other env variable
```shell script
export CHAIN_ID=darchub
export MONIKER=<YOUR_MONIKER>
```

**You can git clone and run script or call docker commands directly****
**In a case of using volume (not bind mount) as containe volume, change `-v ~/.knstld:/root/.knstld` to `--mount type=volume,source="$VOLUME",target=/root`**

#### Init
Initialize blockchain folder
```shell script
# Using script
./docker/start.sh init

# Using docker cmd
docker run --rm -it \
  -e MONIKER="$MONIKER" \
  -e CHAIN_ID="$CHAIN_ID" \
  -v ~/.knstld:/root/.knstld "$IMAGE" /opt/init.sh
```

#### Moniker
Change moniker
```shell script
export MONIKER=moniker 

# Using script
./docker/start.sh config

# Using docker cmd
docker run --rm -it \
  -e MONIKER="$MONIKER" \
  -v ~/.knstld:/root/.knstld "$IMAGE" /opt/config.sh
```

#### Setup
**Omit `KEY_NAME`, `KEY_PASSWORD`, `KEY_MNEMONIC` if you want to create a new identity.**
Setup genaccs, gentxs, collectGentxs

```shell script
export KEY_PASSWORD="..."
export KEY_NAME="..."
export KEY_MNEMONIC="..."

# Using script
./docker/start.sh setup

# Using docker cmd
docker run --rm -it \
  -e KEY_PASSWORD="$KEY_PASSWORD" \
  -e KEY_NAME="$KEY_NAME" \
  -e KEY_MNEMONIC="$KEY_MNEMONIC" \
  -v ~/.knstld:/root/.knstld "$IMAGE" /opt/setup.sh
```

#### Run 
Run blockchain node in container
```shell script
# Using script
./docker/start.sh run

# Using docker cmd
docker run --rm -it \
  -p 26657:26657 \
  -p 26656:26656 \
  -p 1317:1317 \
  -p 9090:9090 \
  -v ~/.knstld:/root/.knstld  "$IMAGE" /opt/run.sh
```

### Localnet
```shell script
export IMAGE=kirdb/knstld:0.2.0
```
#### Create network files for specified node configs
```shell script
export CHAIN_ID=darchub
./scripts/localnet.sh create
```
#### Run network
```shell script
 ./scripts/localnet.sh run
```

#### With docker-compose
```shell script
docker-compose up
```

#### Connect to network
**You can run docker commands directly without cloning repository. Take a look at the chapter above**

```shell script
./docker/start.sh init
```
```shell script
export MONIKER=<YOUR_MONIKER>
./docker/start.sh config
```

```shell script
./scripts/localnet.sh copy
```

```shell script
./docker/start.sh run
```
### Resolving errors

#### Missing ziphash
```shell script
go get -u go.opencensus.io
go get gopkg.in/fsnotify/fsnotify.v1
github.com/fsnotify/fsnotify v1.4.8
```
