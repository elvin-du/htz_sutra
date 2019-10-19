[TOC]

## 1. 初始化数据库&表

```
//create database and user
use admin
db.createUser(
     {
         user:"mdu",
         pwd:"12345678",
         roles:[{role:"root",db:"admin"}]
    }
)

use sync_mdu
db.createCollection("blockchain_status");
db.createCollection("blocks");
db.createCollection("transactions");
db.createCollection("validators");
db.createCollection("addresses");

db.blocks.createIndex({"height": -1},{"unique": true});
db.transactions.createIndex({"tx_hash": -1}, {"unique": true});
db.transactions.createIndex({"block_height": -1});
db.transactions.createIndex({"from": 1});
db.validators.createIndex({"operator_address": 1}, {"unique": true});
db.addresses.createIndex({"address": 1}, {"unique": true});
```
## 2. 数据库表结构
### 2.1. blockchain_status表

| 字段名称       | 解释                                                         | 类型  |
| -------------- | ------------------------------------------------------------ | ----- |
| last_height    | 最新区块高度                                                 |       |
| txs_num        | 总交易量                                                     | Int64 |
| avg_block_time | 平均出块时间                                                 |       |
| val_num        | 生效的验证者人数                                             | Int64 |
| jailed_val_num | 被关起来的验证者人数                                         | Int64 |
| candidate_num  | 验证者候选者人数。  总人数 == val_num + jailed_val_num + candidate_num | Int64 |
| inflation      | 当前年通胀率。浮点型，但是String表示                                 |  String     |

**last_height**: 调用 [/blocks/latest]  header.height 获取。

**txs_num**: 调用 [/blocks/latest]  header.total_txs 获取。

**avg_block_time**: 调用 [/blocks/latest] 获取最新区块的时间，/blocks/{height} 获取最近100个区块的时间。
 lastestBlock.time - (lastestBock - 100).time / 100

**inflation**: 调用 [/minting/inflation]  获取。

**val_num**：待用【/staking/validators】获取。

**jailed_val_num**：待用【/staking/validators】获取。

**candidate_num**：待用【/staking/validators】获取。

### 2.2. blocks表

| 字段名称     | 解释                     | 类型   |
| ------------ | ------------------------ | ------ |
| height       | 块高；index              | Int64  |
| transactions | 块包含了多少交易         | Int64  |
| timestamp    | 时间戳                   | Int64  |
| block_hash   | 区块hash                 | String |
| proposer     | 区块提议人               | String |
| validators   | 此块有多少验证人进行签名 | Int  |
| raward       | 块收益                   | String       |

### 2.3. transactions表

| 字段名称  | 解释                | 类型   |
| --------- | ------------------- | ------ |
| tx_hash   | 交易hash；index  | String |
| timestamp | 时间戳              | Int64  |
| status    | 交易是否成功        | Int64  |
| block_height | 属于那个区块的区块高度；index | Int64 |
| fee       | 交易手续费          |   String     |
| used_gas  | 消耗的gas数量       | Int64    |
| memo      | 交易备注            | String |
| from      | 交易发起账户地址            | String |
| messages      | 交易消息，json数组           | String |


#### messages交易表

**注意**：此表根据`tx_type`交易类型的不同，存储的字段会完全不同,现在阶段一共有15个类型。

| 字段名称  | 解释                | 类型   |
| --------- | ------------------- | ------ |
| tx_hash | 交易hash；index；此字段也许不可以是index，因为一个tx可以对应多个msg，但是现在阶段都是一个tx对应一个msg | String |
| tx_type     | 交易类型；值：send             | String |
| from     | 交易发起账户地址；                 | String |
| to       | 交易接受账户地址                       | String |
| amount   | 转账了多少币。不是转账交易，此字段为空;  具体格式：[ { "denom" : "atom", "amount" : "10" } ] |  String      |

### 2.4. validators表

| 字段名称         | 解释         | 类型   |
| ---------------- | ------------ | ------ |
| moniker          | 验证人名称   | String |
| operator_address | ；index      | String |
| owner_address    | 自委托人地址   | String |
| reward_address   | 收益地址     | String |
| details          | 验证人详情   | String |
| voting_power     | 投票权重     | String       |
| website          | 验证人的网址 | String |

### 2.5. addresses表

| 字段名称        | 解释                    | 类型   |
| --------------- | ----------------------- | ------ |
| address         | 地址；index             | String |
| balance         | 余额;     | String       |  
| delegated       | 本地址被委托的Token数量;有可能为空； |  String |
| reward_address  | 接受收益的地址；有可能为空; 如果address不是validator，此字段为空    |    String    |
| transaction_num | 本地址发送的交易总数量；  | Int64  |

**transaction_num**: 如果要算上失败的数量，直接根据`transacitions`表的来计算，如果不需要记录失败的tx交易数量，直接用【:1317/txs?message.sender=address】来计算。

**balance**：访问URL 【:1317/auth/accounts/{address} 】获取 。

**delegated**：访问URL：【/staking/validators/{validatorAddr}】获取。