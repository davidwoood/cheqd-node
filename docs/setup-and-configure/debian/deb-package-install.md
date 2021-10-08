# Installing a cheqd node using Debian package releases

## Context

This document provides guidance on how to install and configure a node for the cheqd testnet.

* Our [guide to Debian packages for `cheqd-node`](readme.md) provides an overview of what configuration actions are carried out by the installer.
* The [node setup guide provides pre-requisites](../readme.md) needed before the steps below should be attempted.
* Separate instructions apply if you're looking to [upgrade an existing cheqd node](deb-package-upgrade.md).

## Installation steps for `cheqd-node` .deb

1. **Download [the latest release of `cheqd-node` .deb](https://github.com/cheqd/cheqd-node/releases/latest) package**

    ```bash
    wget https://github.com/cheqd/cheqd-node/releases/download/v0.2.3/cheqd-node_0.2.3_amd64.deb
    ```

2. **Install the package**

   ```bash
   sudo dpkg -i cheqd-node_0.2.3_amd64.deb
   ```

3. **Switch to the `cheqd` system user**

   Always switch to `cheqd` user before managing node. By default, the node stores configuration files in the home directory for the user that initialises the node. By switching to the special `cheqd` system user, you can ensure that the configuration files are stored in the [system app data directories](readme.md) configured by the Debian package.

   ```bash
   sudo su cheqd
   ```

4. **Initialise the node configuration files**

   ```bash
   cheqd-noded init <your-node-name>
   ```

5. **Download the genesis file for a persistent chain, such as the cheqd testnet**

   Download the `genesis.json` file [corresponding a persistent chain](https://github.com/cheqd/cheqd-node/tree/main/persistent_chains/testnet) and put it in the `/etc/cheqd-node/` directory.

   For cheqd testnet:

   ```bash
   wget -O /etc/cheqd-node/genesis.json https://raw.githubusercontent.com/cheqd/cheqd-node/main/persistent_chains/testnet/genesis.json
   ```

6. **Define the seed configuration for populating the list of peers known by a node**

   Search for the `seeds` parameter in the node configuration file `/etc/cheqd-node/config.toml` and set its value to a comma separated list of seed node addresses specified in `seeds.txt` for [persistent chains](https://github.com/cheqd/cheqd-node/tree/main/persistent_chains/testnet).

   For cheqd testnet, executing the following commands will set this up correctly:

   ```bash
   $ SEEDS=$(wget -qO- https://raw.githubusercontent.com/cheqd/cheqd-node/main/persistent_chains/testnet/seeds.txt)

   $ echo $SEEDS
   # Comma separated list should be printed

   $ sed -i.bak 's/seeds = ""/seeds = "'$SEEDS'"/g' /etc/cheqd-node/config.toml
   ```

7. **Set gas prices accepted by the node**

   Search for the `minimum-gas-prices` parameter in the node configuration file `/etc/cheqd-node/config.toml` and set it to a non-empty value. The recommended value is `25ncheq`.

   For cheqd testnet, executing the following command will set this up correctly:

   ```bash
   sed -i.bak 's/minimum-gas-prices = ""/minimum-gas-prices = "25ncheq"/g' /etc/cheqd-node/app.toml
   ```

8. **Make the RPC endpoint available externally** (optional)

   This step is necessary only if you want to allow incoming client application connections to your node. Otherwise, the node will be accessible only locally. Further details about the RPC endpoints is available in the [cheqd node setup guide](../readme.md).

   Search for the `laddr` parameter under the `RPC Server Configuration Options` section in the node configuration file `/etc/cheqd-node/config.toml` and replace its value to `0.0.0.0:26657`

   For cheqd testnet, executing the following commands will set this up correctly:

   ```bash
   sed -i.bak 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /etc/cheqd-node/config.toml
   ```

9. **Enable and start the `cheqd-noded` system service**

   If you are prompted for a password for the `cheqd` user, type `exit` to logout and then attempt to execute this as a privileged user (with `sudo` or as root user).

   ```bash
   systemctl enable cheqd-noded
   ```

   ```bash
   systemctl start cheqd-noded
   ```

   Check that the `cheqd-noded` service is running. If successfully started, the status output should return `Active: active (running)`

   ```bash
   systemctl status cheqd-noded
   ```

## Post-installation checks

Once the `cheqd-noded` daemon is active and running, check that the node is connected to the cheqd testnet and catching up with the latest updates on the ledger.

### Checking node status via terminal

```bash
cheqd-noded status
```

In the output, look for the text `latest_block_height` and note the value. Execute the status command above a few times and make sure the value of `latest_block_height` has increased each time.

The node is fully caught up when the parameter `catching_up` returns the output `false`.

### Checking node status via the RPC endpoint

An alternative method to check a node's status is via the RPC interface, if it has been configured.

* Remotely via the RPC interface: `cheqd-noded status --node <rpc-address>`
* By opening the JSONRPC over HTTP status page through a web browser: `<node-address:rpc-port>/status`


## Next steps

At this stage, your node would be connected to the cheqd testnet as an observer node. Learn [how to configure your node as a validator node](../configure-new-validator.md) to participate in staking rewards, block creation, and governance.