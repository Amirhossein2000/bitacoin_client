### Configuration
Items | Explanations
------------ | -------------
Hosts | It's a list of the address of miners

### Commands:

#### Generate key pair of a wallet:

```bash
./bitacoin_client wallet -dir <wallet_path>
````

#### Get the balance of a public_key from nodes:

```bash
./bitacoin_client balance -pub <public_key> -config defaultConfig.json
````

#### Generate and signature new transaction:

```bash
./bitacoin_client transaction -pr <src private_key> -pub <src public_key> -to <dst public_key> -a <amount> -config defaultConfig.json
```