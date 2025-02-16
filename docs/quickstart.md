# Quick Start

A brief installation guide. Find more details and explanation of the configuration settings in [configure.md](configure.md).

_"Proxeus is a platform for quick and convenient document digitalization, signing, processing, and distribution. It allows users to keep their important documents secure and registered on the blockchain. Proxeus empowers anyone to create blockchain applications and is available for free as an open-source project." --[S-Pro](https://s-pro.io/)_



## Source Code

You can access the source code of this application on the [Proxeus GitHub repository](https://github.com/ProxeusApp).

**By installing and using the Proxeus software you agree with the terms of the [Proxeus License Agreement](LICENSE).**

## Install docker and docker-compose

The quickest way to try Proxeus is to use `docker-compose`.

1. [Install Docker Engine](https://docs.docker.com/install/)
2. [Install docker-compose](https://docs.docker.com/compose/install/)

## Get API Keys for Infura and SparkPost

The Proxeus platform depends on [Infura](https://infura.io/) and [SparkPost](https://www.sparkpost.com/)
for Ethereum and email integration respectively.

Please create an account on those platform and get an API Keys.

## Proxeus Demo Smart Contract

For your convenience, a demo smart contract is deployed on several networks at the following addresses.

You can switch between the networks using the environment configuration `PROXEUS_BLOCKCHAIN_NET`

### Ethereum: Sepolia

```
0x9bc518Fd070BE3DBB07399261688015744a7FB02
```
[Verify on Etherscan](https://sepolia.etherscan.io/address/0x9bc518Fd070BE3DBB07399261688015744a7FB02#code)

`PROXEUS_BLOCKCHAIN_NET=sepolia`

### Ethereum: Goerli

```
0x66FF4FBF80D4a3C85a54974446309a2858221689
```
[Verify on Etherscan](https://goerli.etherscan.io/address/0x66FF4FBF80D4a3C85a54974446309a2858221689#code)

`PROXEUS_BLOCKCHAIN_NET=goerli`

### Ethereum: Mainnet

```
0xf63e471d8cbc57517c37c39c35381a385628e012
```
[Verify on Etherscan](https://etherscan.io/address/0xf63e471d8cbc57517c37c39c35381a385628e012)

`PROXEUS_BLOCKCHAIN_NET=main`

### Polygon: mumbai

```
0x00119d8C02bbC4c1231D054BB2813792B4411Ed5
```
[Verify on Etherscan](https://mumbai.polygonscan.com/address/0x00119d8C02bbC4c1231D054BB2813792B4411Ed5)

`PROXEUS_BLOCKCHAIN_NET=polygon-mumbai`

### Polygon: main

```
0x60970BeFda93464A105DD21Dc6a30B69C5B5c6e4
```
[Verify on Etherscan](https://polygonscan.com/address/0x60970BeFda93464A105DD21Dc6a30B69C5B5c6e4)

`PROXEUS_BLOCKCHAIN_NET=polygon`

## Create a docker-compose.yml file

**Note: Please make sure that you always pull Docker images from the official `proxeus` DockerHub repository and that you are using the latest version.**

User the example below as your `docker-compose.yml` file:

```
---
version: '3.7'

networks:
  xes-platform-network:
    name: xes-platform-network

services:
  platform:
    image: proxeus/proxeus-core:latest
    container_name: xes-platform
    depends_on:
      - document-service
    networks:
      - xes-platform-network
    restart: unless-stopped
    environment:
      TZ: Europe/Zurich
      PROXEUS_PLATFORM_DOMAIN: "${PROXEUS_PLATFORM_DOMAIN:-http://xes-platform:1323}"
      PROXEUS_DOCUMENT_SERVICE_URL: "http://document-service:2115/"
      PROXEUS_ENCRYPTION_SECRET_KEY: "${PROXEUS_ENCRYPTION_SECRET_KEY}"
      PROXEUS_BLOCKCHAIN_CONTRACT_ADDRESS: "${PROXEUS_BLOCKCHAIN_CONTRACT_ADDRESS}"
      PROXEUS_INFURA_API_KEY: "${PROXEUS_INFURA_API_KEY}"
      PROXEUS_SPARKPOST_API_KEY: "${PROXEUS_SPARKPOST_API_KEY}"
      PROXEUS_EMAIL_FROM: "${PROXEUS_EMAIL_FROM:-no-reply@example.com}"
      PROXEUS_AIRDROP_WALLET_FILE: "${PROXEUS_AIRDROP_WALLET_FILE:-/root/.proxeus/settings/airdropwallet.json}"
      PROXEUS_AIRDROP_WALLET_KEY: "${PROXEUS_AIRDROP_WALLET_KEY:-/root/.proxeus/settings/airdropwallet.key}"
      PROXEUS_DATABASE_ENGINE: "${PROXEUS_DATABASE_ENGINE:-storm}"
      PROXEUS_DATABASE_URI: "${PROXEUS_DATABASE_URI:-mongodb://root:root@mongo:27017}"
      PROXEUS_TEST_MODE: "${PROXEUS_TEST_MODE:-false}"
      PROXEUS_ALLOW_HTTP: "${PROXEUS_ALLOW_HTTP:-false}"
    ports:
      - "1323:1323"
    volumes:
      - ${PROXEUS_DATA_DIR:-./data}/proxeus-platform/data:/data/hosted
      - ${PROXEUS_DATA_DIR:-./data}/proxeus-platform/settings:/root/.proxeus/settings

  document-service:
    image: proxeus/document-service:latest
    container_name: xes_document_service
    networks:
      - xes-platform-network
    restart: unless-stopped
    environment:
      TZ: Europe/Zurich
    ports:
      - "2115:2115"
    volumes:
      - ${PROXEUS_DATA_DIR:-./data}/document-service/logs:/document-service/logs
      - ${PROXEUS_DATA_DIR:-./data}/document-service/fonts:/document-service/fonts
```

## Start Proxeus

Run the following command in the directory containing your `docker-compose.yml` file (Linux and OSX):
```
export PROXEUS_EMAIL_FROM=<Your valid Sender Email Address>
export PROXEUS_INFURA_API_KEY=<Your Infura project ID>
export PROXEUS_SPARKPOST_API_KEY=<Your SparkPost API Key>
export PROXEUS_BLOCKCHAIN_CONTRACT_ADDRESS=0x1d3e5c81bf4bc60d41a8fbbb3d1bae6f03a75f71
export PROXEUS_ALLOW_HTTP=true
docker-compose up
```

Proxeus should be available at http://localhost:1323

**Note:** that your system configuration at this point will be reflected in the local configuration database under `data/proxeus-platform/settings/main.json`. Any future changes to the configuration must be made here - the environment variables will not be propagated, unless you delete this file to reset the deployment.

The next step is to [configure](configure.md) your instance for the first time.
