# Decentralized Escrow

## Usage

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```

### Anvil

```shell
$ anvil
```

### Deploy

```shell
$ forge script script/DeployContracts.s.sol --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Cast

```shell
$ cast send <deployed_marketplace_contract_address> "addProduct(uint256,uint256)" <price> <stock>

$ cast send <deployed_marketplace_contract_address> "updateStock(uint256,uint256)" <product_id> <new_stock>

$ cast send <deployed_marketplace_contract_address> "updatePrice(uint256,uint256)" <product_id> <new_price>

$ cast call <deployed_marketplace_contract_address> "getProduct(uint256)" <product_id>

$ cast send <deployed_escrow_contract_address> "createEscrow(uint256,uint256)" <product_id> <qty> --value <total_value>

$ cast send <deployed_escrow_contract_address> "receivalConfirmation(uint256)" <escrow_id>

$ cast send <deployed_escrow_contract_address> "resolve(uint256)" <escrow_id>

$ cast send <deployed_escrow_contract_address> "refund(uint256)" <escrow_id>

$ cast send <deployed_escrow_contract_address> "getEscrow(uint256)" <escrow_id>

```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```
