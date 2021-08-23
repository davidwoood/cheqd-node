# `cheqd-testsnet` docker image

## Description

Debian based docker image with the `cheqd-nonded` executable inside preconfigured to run a network of 2 nodes. Intended for use in CI pipelines.

## Prerequisites

- Build `cheqd-node` image first. See the [instruction](../cheqd_node/README.md).

## Building

To build the image:

- Go to the repository root
- Run `docker build -f docker/single_image_testnet/Dockerfile -t cheqd-testnet .`

## Running

- Run `docker run -it --rm -p "26657:26657" -p "26659:26659" cheqd-testnet`
- RPC apis are exposed on the folowing ports:
  - node_0: `26657`
  - node_1: `26659`
- Try to connect to any node in your browser, for instance: `http://localhost:26657/`