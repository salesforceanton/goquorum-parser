# GoQuorum

## **Конфигурация и запись сети goQuorum**

1. *Необходимо сгенерировать ключи для нод, а так же genesis-file при помощи утилиты npx quorum-genesis-tool. Ключевые моменты*:
- genesis-file включает в себя начальную конфигурацию сети которая определяется в момент инициализации (генезис-блок).
- тип консенсуса (рекомендуемо qbft)
- chainId - айдишка нужна для добавления конфигурации сети в chainlist, а так же для добавления сети по индентификатору в метамаск
- время генерации блока (так же можно задать время генерации блока без транзакций)
- количество нод - для продакшн конфигурации рекомендуется использовать 2 бутноды - они нужны для синхронизации между остальными нодами и обеспечивают дополнительную отказоустойчивость, запускалась сеть со следующей конфигурацией - 1 rpc-нода (сгенерирована была 1 доп мембер нода и указана в docker-compose как рпс нода), 3 валидатора, 3 мембер-ноды (так как модель приватность транзакций включена в конфигурацию то каждая мембер-нода имеет в паре tesseta private transactions manager
- enableGasPriceBlock - можно включить газ за транзакции, по дефолту сеть без газа за транзакции, конфигурация настраивается в зависимости от выбранного типа консенсуса
- genesis предоставляет возможность задать балансы нативной валюты для любых кошельков прописанных в файле, а так же возможность предеплоя контрактов
- так же quorum предоставляет возможность включать фичи с определнного блока (нужно проверять и уточнять отдельные фичи)

Пример genesis файла

```json
{
  "nonce": "0x0",
  "timestamp": "0x6703D915",
  "extraData": "0xf865a00000000000000000000000000000000000000000000000000000000000000000f83f94bd931d33e38215d520a737da70c766731f1e62f7941b654ed84f1ef425bd838af3e812972c5d8e300394ecb33915f4ad3a581fc25356f901b7de946bc699c080c0",
  "gasLimit": "0xffff",
  "gasUsed": "0x0",
  "number": "0x0",
  "difficulty": "0x1",
  "coinbase": "0x0000000000000000000000000000000000000000",
  "mixHash": "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "config": {
    "chainId": 1338,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "muirGlacierBlock": 0,
    "londonBlock": 0,
    "txnSizeLimit": 64,
    "arrowGlacierBlock": 0,
    "grayGlacierBlock": 0,
    "privacyEnhancementsBlock": 0,
    "privacyPrecompileBlock": 0,
    "qbft": {
      "epochLength": 30000,
      "blockPeriodSeconds": 5
    },
    "transitions": [
      {
        "block": 0,
        "transactionSizeLimit": 64,
        "contractSizeLimit": 64,
        "privacyEnhancementsEnabled": true
      }
    ],
    "isQuorum": true
  },
  "alloc": {
    "0x51c9ade0a2a39e680b48f53931f61217ee952866": {
      "balance": "1000000000000000000000000000"
    },
    "0x97c3f95a6883fa7a314bea7e54c74f9fa2f81e1a": {
      "balance": "1000000000000000000000000000"
    }
  }
}
```

2. *Приватность*

- Мы создаём 2 генезиса, один из них используется с приватной конфигурацией
- в docker-compose необходимо добавить для каждой member-ноды по ptm необходимо прописать конфигурацию ptm в docker-entrypoiint.sh. Подробнее в пункте приватность ниже.

3. Урлы всех нод необходимо обновить в сгенеренных permissioned-nodes.json и static-nodes.json
```json
[
  "enode://ea698122b5928859c36b31fd8eb1fec95a006f98b273e7dd2b922cbb20e44711b0bae6de5d85cfb6b69de027c4edd41e7c00ff1796dd8a9aabcb7b1162be523f@172.16.239.11:30303?discport=0&raftport=53000",
  "enode://1cce8a29aed9c1cde8d8ce483e1f34a33b7a7eb4e48620450e421f51ec9a0277e5efc75a389ee8c4104c4471b5b96a1b7fa295b6388492489dc176c4878d9217@172.16.239.12:30303?discport=0&raftport=53000",
  "enode://cb3bc1abca4dfc0c3ad5b541490ee01a8f3369c55fe445dab76e58cb93f70bf5dd91b72aea6d0f793cb8232215435d6a56f265975785f5737be50f38613dd5e6@172.16.239.13:30303?discport=0&raftport=53000",
  "enode://738ce27c421f18a497814428f1f0c4ac44aaf50629f15973b0e20875c884dd1c19c6d7f4b48e56186e8c36ac5a82032b47dd053b75089d6bff5ea1cadc2b31a0@172.16.239.14:30303?discport=0&raftport=53000",
  "enode://079f1d0a6549e57d3a10d8f91532c65acda4125b1ba652ce566be51e4d77f3c0181d872aea4d673e18af7d6c7cc8bed9a1a899e28eed173a0f17888efdaae024@172.16.239.15:30303?discport=0&raftport=53000",
  "enode://c8bb51ec6b4e91b8dadc714c948b1031289e1bf1870dec8236a7a11d514a176d69b6f8f65f1873a6d3a12298e483cddbea0b4862d7f6da5c3d82cd1eef9d3662@172.16.239.16:30303?discport=0&raftport=53000",
  "enode://b6826ab3c7472a9e5f711bca7b4fd7060524462a44b044402c6c37b1f205dd6e7576f45b850c9ac1940a0743306ec2812effae8858dc9ba78aa8718bbefd777b@172.16.239.17:30303?discport=0&raftport=53000",
  "enode://1ff5c6880aa6ab48276ac791b91e1aab4b5bd3d6bb237b9468203661e7c1aa833d096fac7ad5213eceeff0c5b3fcca294156b5a68162bc2ff0ac2e3816466807@172.16.239.18:30303?discport=0&raftport=53000",
  "enode://3d8332104b5a930f3926e7c714007740962e72729fae290917f16c706fa634fb33aa60c8119a9004509133cd5e6c5cd6c466314624c118bc4b038df80338ceea@172.16.239.19:30303?discport=0&raftport=53000"
]
```

4. Создать docker-compose включая все необходимые ноды, сервисы для мониторинга, блок-експлорер blockscout

5. в docker-entrypoint.sh запускается команда geth которая инициализирует из генезис файла сеть и запускает её с опредёлнными параметрами указанными в самой команде, такие как включение ptm, порты rpc, консенсус и так далее

> **Для конфигурации сети в продакшен рекомендуется использовать **bootnodes** - ноды необходимые для синхронизации и повышенной отказоустойчивости, для этого необходимо создать ноду и добавить её урл в лист нод а так же в docker-entrypoint.sh. Рекомендуется ипользовать 2 **bootnode**
> ```sh
> --bootnodes=enode://738ce27c421f18a497814428f1f0c4ac44aaf50629f15973b0e20875c884dd1c19c6d7f4b48e56186e8c36ac5a82032b47dd053b75089d6bff5ea1cadc2b31a0@172.16.239.14:30303
> ```
>{.is-info}



## **Permissioning**

- Для контроля за добавлением/удалением нод а так же для ограничения прав пользователей (например кошелёк без прав не может выполнить транзакцию на смарт-контракте в сети) внутри сети - quorum предоставляет набор смарт-контрактов.
- Набор контрактов находится в репозитории, а так же набор скриптов для деплоя. После деплоя происходит инициализация контрактов (в скрипте необходимо указать набор ключей аккаунтов которые будут обладать правами супер-админа в сети и будут иметь право на добавление организаций, нод, а так же на подписание транзакций смарт-контрактов - по дефолту это аккаунты всех нод)

![image](https://github.com/user-attachments/assets/e5153d49-74ea-4e7b-bc7b-6d77e74d8eb7)

## **Деплой  и взаимодействие со смарт-контрактами**

- Для деплоя контрактов в репозитории предоставлены скрипты. Необходимо так же использовать опредёлнную смарт-контрактом версию компилятора. Для теста был задплоен смарт-контракт usdt, а так же произведены транзакции с сервиса-парсера и парсинг евентов.
- Приватные смарт-контракты позволяют использовать приватные транзакции которые в своё время нужны для ограничения области видимости данных БЧ там где это необходимо. Для того чтобы сделать смарт-контракт приватным - необходимо его задеплоить в quorum следующим скриптом:


```go package main

import (
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/usdt_abi"
)

const (
	member1tesseraPublicKey = "ZTYLEqk5gzZ60uvLJR0zIwPCzzRoRxdyxauV+QxMMRI="
	member3tesseraPublicKey = "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao="
	gasLimit                = 47000000
	ptmUrl                  = "http://localhost:9081"
	rpcUrl                  = "http://localhost:20000"
	ethKey                  = "{\"version\":3,\"id\":\"ee96211f-182c-4ddb-ac9d-9ae69a82bfc3\",\"address\":\"111b0e42105a3412eee350a06a46edf986ce9172\",\"crypto\":{\"ciphertext\":\"0d1b72f1281d8730ada1b4035b611be40e0f74ea9a8788679701c13379e03dc0\",\"cipherparams\":{\"iv\":\"0cccf28477c3f6c7750c2bf60876c58c\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"72992a5b3e1d18e5fc064da2d5f978dffdd9cff2ab6d1baec98cb949c1fc05a0\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"6c581f3ecafa73a2ba60d08827845d9ecd15925cddd000e96d77b1eaf68a1c58\"}}"
)

func main() {
	// connect to member1
	rpcClient, err := rpc.DialHTTP(rpcUrl)
	if err != nil {
		println(err.Error())
		return
	}

	trOpts, err := bind.NewTransactor(strings.NewReader(ethKey), "")
	if err != nil {
		println(err.Error())
		return
	}
	trOpts.GasLimit = gasLimit
	trOpts.PrivateFor = []string{member3tesseraPublicKey}
	trOpts.PrivateFrom = member1tesseraPublicKey

	ethClient, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(ptmUrl)
	if err != nil {
		println(err.Error())
		return
	}

	parsedAbi, err := abi.JSON(strings.NewReader(usdt_abi.UsdtABI))
	if err != nil {
		println(err.Error())
		return
	}

	contractBytecode, err := os.ReadFile("./TetherToken.bin")
	if err != nil {
		println(err.Error())
		return
	}

	// params
	var initialSupply = big.NewInt(100000000000000)
	const name = "TetherToken"
	const symbol = "USDT"
	var decimals = big.NewInt(6)

	initialize := []interface{}{initialSupply, name, symbol, decimals}

	address, _, _, err := bind.DeployContract(trOpts, parsedAbi, common.FromHex(string(contractBytecode)), ethClient, initialize...)
	if err != nil {
		println(err.Error())
		return
	}

	println("contractAddress: " + address.Hex())
}
```
> **Здесь ключевым моментом является создание приватной транзакции деплоя смарт-контракта, в которой определена область видимости контракта при помощи атрибутов **privateFrom**, **privateFor**. Так же используется клиент с раширением **WithPrivateTransactionManager** что необходимо для взаимодействия c Tessera
{.is-info}


## **Приватность**

Для использования приватности транзакций внутри кворума необходимо для каждой ноды (для которых планируется использовать модель приватных транзакций и смарт-контрактов) добавить ноду-компаньон tessera. Tessera позволяет совершать приватные транзакции и деплоить приватные смарт-контракты путём сохранения приватной части информации во внутреннем сторадже.

Пример ноды с tessera:
```yml
member1tessera:
    << : *tessera-def
    ports:
      - 9081:9080
    volumes:
      - ./config/nodes/member1:/config/keys
      - ./logs/tessera:/var/log/tessera/
    networks:
      quorum-network:
        ipv4_address: 172.16.239.26

  member1quorum:
    << : *quorum-def
    ports:
      - 20000:8545/tcp
      - 20001:8546/tcp
      - 30303
      - 9545
    depends_on:
      - member1tessera
    environment:
      - GOQUORUM_CONS_ALGO=${GOQUORUM_CONS_ALGO}
      - GOQUORUM_GENESIS_MODE=enhanced
      - QUORUM_PTM=member1tessera
    volumes:
      - ./config/nodes/member1:/config/keys
      - ./logs/quorum:/var/log/quorum/
      - ./config/permissions:/permissions
    networks:
      quorum-network:
        ipv4_address: 172.16.239.16
```


Quorum предоставляет 3 уровня настройки приватности транзакций:

1. Internal private transaction - По-дефолту мы можем определить область видимости *логов* транзакции задав параметр privateFor при создании транзакции
2. Privacy enhancements чтобы задать уровень приватность при отправки транзакций при помощи параметра privateFlag который определяет мод этой фичи, более подробно https://docs.goquorum.consensys.io/concepts/privacy/privacy-enhancements. Чтобы включить privacy enhancements необоходимо:
- Добавить в генезиз файл параметр privacyEnhancementsEnabled
```js
		"transitions": [
      {
        "block": 0,
        "transactionSizeLimit": 64,
        "contractSizeLimit": 64,
        "privacyEnhancementsEnabled": true
      }
    ],
 ```

3. Privacy transaction marker как более приватный уровень (скрыто to, nonce) работает при соответсвующих настройках и включённом privacy enhancements. Для включения ptm необходимо:
- В генезис файл добавить параметр. Он должен быть одинаковым для всех нод
```js
	 "privacyPrecompileBlock": 0,
```
- В docker-entrypoint.sh в команду geth init добавить следующие флаги
```sh
	--ptm.timeout 5 --ptm.url http://${QUORUM_PTM}:9101 
  --ptm.http.writebuffersize 4096 
  --ptm.http.readbuffersize 4096 
  --ptm.tls.mode off
  --privacymarker.enable
```
- Privacy enhancements должны быть включены

Пример отправки транзакции ptm
```go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/salesforceanton/goquorum-parser/internal/converters"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/usdt_abi"
)

const (
	testAmount                 = "100000000"
	member1tesseraPublicKey    = "ZTYLEqk5gzZ60uvLJR0zIwPCzzRoRxdyxauV+QxMMRI="
	member3tesseraPublicKey    = "1453wQJ/btTDFlK3DrfVdGglnuPg+RAUXBZMJVzA9Ao="
	gasLimit                   = 47000000
	usdtContractAddressPrivate = "0xebaf93b936e40dcd4f5fc78c15d01e91d09e085d"
	usdtContractAddressPublic  = "0xe2889ac4b39a4d4ff9c81d372d6a682c34bda1e9"
	ptmUrl                     = "http://localhost:9081"
	rpcUrl                     = "http://localhost:20000"
	ethKey                     = "{\"version\":3,\"id\":\"ee96211f-182c-4ddb-ac9d-9ae69a82bfc3\",\"address\":\"111b0e42105a3412eee350a06a46edf986ce9172\",\"crypto\":{\"ciphertext\":\"0d1b72f1281d8730ada1b4035b611be40e0f74ea9a8788679701c13379e03dc0\",\"cipherparams\":{\"iv\":\"0cccf28477c3f6c7750c2bf60876c58c\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"72992a5b3e1d18e5fc064da2d5f978dffdd9cff2ab6d1baec98cb949c1fc05a0\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"6c581f3ecafa73a2ba60d08827845d9ecd15925cddd000e96d77b1eaf68a1c58\"}}"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	rpcClient, err := rpc.DialHTTPWithClient(
		rpcUrl,
		&http.Client{
			Timeout: time.Second * 15,
		},
	)
	if err != nil {
		logger.Debug("cryptogateUtils.InitTools", "err", err)
		return
	}

	provider, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(ptmUrl)
	if err != nil {
		logger.Error(fmt.Sprintf("ethclient.NewClient(rpcClient).WithPrivateTransactionManager(tessera1Url): %s", err.Error()))
		return
	}

	addressContract := common.HexToAddress(usdtContractAddressPrivate)

	gasPrice, err := provider.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("Provider.SuggestGasPrice", "err", err)
		return
	}

	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		logger.Debug("tools.Provider.ChainID", "err", err)
		return
	}

	trOpts, err := bind.NewTransactorWithChainID(strings.NewReader(ethKey), "", chainID)
	if err != nil {
		logger.Debug("bind.NewTransactor", "err", err)
		return
	}

	trOpts.PrivateFrom = member1tesseraPublicKey
	trOpts.PrivateFor = []string{member3tesseraPublicKey}
	trOpts.GasLimit = gasLimit
	trOpts.IsUsingPrivacyPrecompile = true

	var transaction *types.Transaction

	contractTransactor, err := usdt_abi.NewUsdtAbiTransactor(
		addressContract,
		provider,
	)
	if err != nil {
		logger.Error("usdt_abi.NewUsdtAbiTransactor", "err", err)
		return
	}

	amount, err := converters.StringToBigInt(testAmount)
	if err != nil {
		logger.Error("converters.StringToBigInt", "err", err)
		return
	}

	transaction, err = contractTransactor.Issue(trOpts, amount)
	if err != nil {
		logger.Error("contractTransactor.Issue", "err", err)
		return
	}

	logger.Debug("Transaction sent successfully",
		"function", "issue",
		"to", addressContract.Hex(),
		"gasPrice", gasPrice.String(),
		"txHash", transaction.Hash().Hex(),
	)
}

```


## **Особенности парсера для Quorum**
- для корректной работы с quorum так как это форк go-ethereum - нужно зареплейсить пакет в go.mod 
```go
replace github.com/ethereum/go-ethereum => github.com/salesforceanton/quorum/v24 v24.4.5
```
В данном случае я сделал форк от quorum потому что невожно симпортить последнюю версию

- При инициализации http-клиентов - необходимо использовать екстеншн для работы с приватными транзами:
```go
	rpcClient, err := rpc.DialHTTP(rpcUrl)
	if err != nil {
		println(err.Error())
		return
	}

	ethClient, err := ethclient.NewClient(rpcClient).WithPrivateTransactionManager(ptmUrl)
	if err != nil {
		println(err.Error())
		return
	}
```

- Необходимо использоваться методы контрактов из генерации abi так как метод transact под капотом если приватный менеджер инициализирован - отправляет транзу ещё и на tessera (это ожидаемое поведение)
```go
switch req.TypeSmartContract {
	case cryptogate.SmartContractTypeUSDT:
		contractTransactor, err := usdt_abi.NewUsdtAbiTransactor(
			addressContract,
			tools.Provider,
		)
		if err != nil {
			c.logger.Error("usdt_abi.NewUsdtAbiTransactor", "err", err)
			return cryptogatemessages.SendTransactionResponse{}, err
		}

		switch req.NameFunction {
		case "issue":
			amount, err := converters.StringToBigInt(req.Amount)
			if err != nil {
				c.logger.Error("converters.StringToBigInt", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}

			transaction, err = contractTransactor.Issue(trOpts, amount)
			if err != nil {
				c.logger.Error("contractTransactor.Issue", "err", err)
				return cryptogatemessages.SendTransactionResponse{}, err
			}
    }
	}
```

> **При использовании вебсокета логи приватных транзакций приходят с некорректным хешем транзакций 
> я в качестве воркэарунда - отлавливал ошибку некорректного хеша и получал логи блока
> ```go
> func (c *Cryptogate) processLogs() {
> 	for log := range c.ethLogs {
> 		c.updateHeight(log.BlockNumber)
> 
> 		err := c.parseEvent(log)
> 		// get logs from http filter logs here if err with getting private tx receipt appeared
> 		if err != nil && errors.Is(err, ErrWithGetTransactionReceipt) {
> 			c.ParseBlockEvents(log)
> 		}
> 	}
> }
> 
> func (c *Cryptogate) ParseBlockEvents(log types.Log) {
> 	query := ethereum.FilterQuery{
> 		Addresses: []common.Address{log.Address},
> 		FromBlock: big.NewInt(int64(log.BlockNumber)),
> 		ToBlock:   big.NewInt(int64(log.BlockNumber)),
> 	}
> 
> 	logs, err := c.httpProvider.FilterLogs(context.Background(), query)
> 	if err != nil {
> 		c.logger.Error("httpProvider.FilterLogs", "err", err)
> 	}
> 
> 	for _, log := range logs {
> 		c.parseEvent(log)
> 	}
> }
> ```
>{.is-warning}

## **Добавление ноды**
![add_node](https://github.com/user-attachments/assets/4706be9b-5c16-4550-b20e-fb25ec86b614)

##### Добавление ноды-валидатора
1.  Создать директорию для новой ноды-валидатора и сгенерить файлы с ключами
```bash
npx quorum-genesis-tool \
--validators 1 \
--members 0 \
--bootnodes 0 \
--outputPath artifacts
```
Полученный набор ключей:

```bash
  ├── validator0                    
  │   └──  nodekey                      # the node private key
  │   └──  nodekey.pub                  # the node's public key which is used in the enode
  │   └──  address                      # the node's address which is used to vote the validator in/out
  │   └──  accountAddress               # GoQuorum only - the accountAddress
  │   └──  accountKeystore              # GoQuorum only - the account's v3 keystore
  │   └──  accountPassword              # GoQuorum only - the account's password (you would have supplied this)
  │   └──  accountPrivateKey            # GoQuorum only - the account's private key
```
2. Все ноды сети должны использовать один и тот же файл genesis. Для добавления новой ноды уже к существующей сети - необходимо предоставить genesis файл хосту.
> Для консенсуса qbft - ноды которые изначально сконфигурированы - закодированы в атрибуте extraData
> ```json
> "extraData": "0xf865a00000000000000000000000000000000000000000000000000000000000000000f83f94bd931d33e38215d520a737da70c766731f1e62f7941b654ed84f1ef425bd838af3e812972c5d8e300394ecb33915f4ad3a581fc25356f901b7de946bc699c080c0",
> ```
>{.is-info}

3. Необходимо добавить в static-nodes.json и permissioned-nodes.json url ноды который состоит из ключа nodekey.pub и ip-адреса ноды *enode://{nodekey.pub}@ip-addr?discport=0&raftport=53000*
```json
[
  ...
  "enode://22daba3793589381e4d90032dce94c28816c44b377e27e247e8fa236dcb36d614cc52d637aebd6372af6d78cc3ca3c917838dab379fced79255d2ae82d77bd35@172.16.239.20:30303?discport=0&raftport=53000"
]
```
4. Необходимо запустить ноду, это можно сделать как командой geth, так и сервис в докере
Пример docker-сервиса:
```bash
validator4:
    << : *quorum-def
    ports:
      - 21006:8545/tcp
      - 30303
      - 9545
    environment:
      - GOQUORUM_CONS_ALGO=${GOQUORUM_CONS_ALGO}
      - GOQUORUM_GENESIS_MODE=standard
    volumes:
      - ./config/nodes/validator4:/config/keys
      - ./logs/quorum:/var/log/quorum/
      - ./config/permissions:/permissions
    networks:
      quorum-network:
        ipv4_address: 172.16.239.20
```
Пример команды geth:
```bash
export ADDRESS=$(grep -o '"address": *"[^"]*"' ./data/keystore/accountKeystore | grep -o '"[^"]*"$' | sed 's/"//g')
export PRIVATE_CONFIG=ignore
geth --datadir data \
    --networkid 1337 --nodiscover --verbosity 5 \
    --syncmode full --nousb \
    --istanbul.blockperiod 5 --mine --miner.threads 1 --miner.gasprice 0 --emitcheckpoints \
    --http --http.addr 127.0.0.1 --http.port 22005 --http.corsdomain "*" --http.vhosts "*" \
    --ws --ws.addr 127.0.0.1 --ws.port 32005 --ws.origins "*" \
    --http.api admin,trace,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul,qbft \
    --ws.api admin,trace,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul,qbft \
    --unlock ${ADDRESS} --allow-insecure-unlock --password ./data/keystore/accountPassword \
    --port 30305
```
5. Необходимо апрувнуть добавление ноды валидатора другими нодами-валидаторами
- подключаемся к ноде-валидатору
```bash
geth attach ./data/geth.ipc
```
- предлагаем нового валидатора
```bash
istanbul.propose("0x" + {NEW_VALIDATOR_ADDRESS}, true);
```
- проделываем для всех валидаторов шаги выше, а затем чтобы убердиться что новый валидатор добавлен  на подключаемся и проверяем
```bash
geth attach ./data/geth.ipc
istanbul.isValidator();
```

##### Добавление ноды мембера
Добавление ноды которая не определна как валидатор аналогчно флоу описанному выше за исключением пункта 5. 
Так же если мы добавляем ноду с tessera-компаньоном - необходимо добавить url tessera-компаньона в лист пиров. Лист пиров tessera должен быть одинаков для всех tessera-нод, то есть его нужно предоставить хосту новой ноды
```json
{
  "peers": [
    {
      "url": "http://existingpeer1.com:8080"
    }
  ]
}
```

> **Важно!!!** При использовании docker как средства развёртывания необходимо обеспечить маппинг объёмов таким образом чтобы можно было динамически изменять следующие файлы без пересборки образов:
>```
> permisssions-nodes.json
> static-nodes.json
> genesis-file.json
> tessera-config-template.json
>```
>{.is-warning}


## **Полезные ссылки**
- репозиторий который был создан для ресёрча, включает сконфигурированную сеть и парсер-бекенд сервис для парсинга евентов и взаимодействия со смартами https://github.com/salesforceanton/goquorum-parser
- ссылка но официальную доку кворума https://docs.goquorum.consensys.io/
