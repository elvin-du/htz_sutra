[TOC]

# 概述

区块链浏览器API遵循restful设计，返回格式为JSON。

字段命名方式: 单词使用`_`作分隔，而非驼峰方式。

以下是统一的返回样式:
| 字段名称    | 类型     | 解释   |
| ----------| ----------| --------- |
| code      |   int     | 状态码 |
| msg       | String    | 消息   |
| data      |  任何类型  | 接口成功时返回的具体数据|
| error     |  任何类型  | 接口返回失败时的错误详情|

```
正确是返回:
{
   "code": 200000,
   "msg": "ok",
   "data": {},
   "error": null
}

错误返回:
{
   "code": 401001,
   "msg": "fail",
   "data": null,
   "error": "error details"
}

```

**NOTE:**
**以下接口返回都只写了成功时的data字段。**

分页查询格式如下
| 字段名称    | 类型     | 解释   |
| ----------| ---------| --------- |
| page_index |   int| 页码 |
| page_size  | String| 每页数据量   |
| total_pages|  int64| 总页数|
| total_items|  int64| 总的数据条数|
| items     | array[任意类型]| 具体数据列表 |

Example:
```
正确是返回:
{
   "code": 200000,
   "msg": "ok",
   "data": {
     "page_index": 1,
     "page_size": 20,
     "total_pages": 3,
     "total_items": 56,
     "items": [{a:1}, {a:6}]
   }
}
```

# 链接口

## 查询链状态

URL: /chain/status
Method: GET

返回字段:
| 字段名称       | 类型| 解释 |
| -------------- | -----| ----- |
| last_height    | int64 | 最新区块高度    |
| txs_num        | Int64 | 总交易量       |
| avg_blockTime |       |  平均出块时间    |
| val_num        | Int64 | 生效的验证者人数  |
| jailed_val_num | Int64 | 被关起来的验证者人数 |
| candidate_num  | Int64 | 验证者候选者人数。  总人数 == val_num + jailed_val_num + candidate_num  |
| inflation      | String| 当前年通胀率。浮点型，但是String表示     |


# 区块接口

## 查询区块列表

URL: /blocks/:page_index/:page_size
Method: GET

请求参数:
| 字段名称   | 请求体   |类型|   解释     |
| --------- |--------|-----| --------- |
| page_index | path   | int    | 起始页，从1开始   |
| page_size  | path   | int    | 每页数量(0,1000] |
| proposer   | query  |string  | 非必填, 提议人地址  |

返回字段:
| 字段名称     | 类型   | 解释                   |
| ------------ | ---- | ---------------------- |
| height       | Int64| 块高；index             |
| transactions | Int64| 块包含了多少交易         |
| timestamp    | Int64| 时间戳                  |
| blockHash   | String| 区块hash               |
| proposer     | String| 区块提议人              |
| validators   | Int   | 此块有多少验证人进行签名 |
| raward       | String| 块收益                |

## 区块详情

URL: /block/:heightOrHash
Method: GET
请求参数:
| 字段名称      | 请求体  | 类型           |    解释                    |
| -------------|--------|----------------| ------------------------- |
| heightOrHash | path   | int64 or string| 区块高度或者区块hash        |


返回字段:
| 字段名称     | 类型   | 解释                   |
| ------------ | ---- | ---------------------- |
| height       | Int64| 块高；index             |
| transactions | Int64| 块包含了多少交易         |
| timestamp    | Int64| 时间戳                  |
| blockHash   | String| 区块hash               |
| proposer     | String| 区块提议人              |
| validators   | Int   | 此块有多少验证人进行签名 |
| raward       | String| 块收益                |


# 交易接口

## 查询单个交易

URL: /tx/:hash
Method: GET

请求参数:
| 字段名称      | 请求体  | 类型   |   解释    |
| -------------|--------|--------| --------- |
| hash         | path   | string | 交易hash   |

返回字段:
| 字段名称   | 类型   | 解释                      |
| --------- | ------| --------------------------|
| tx_hash   | String | 交易hash；index           |
| timestamp | Int64| 时间戳                      |
| status    | Int64| 交易是否成功                 |
| block_height | Int64 | 属于那个区块的区块高度；index  |
| fee       | String| 交易手续费                  |
| used_gas  | Int64 | 消耗的gas数量               |
| memo      | String| 交易备注                    |
| from      | String| 交易发起账户地址             |
| txType   | String| 交易类型；值：send           |
| to        | String| 交易接受账户地址             |
| amount    | String| 转账了多少币。不是转账交易，此字段为空;  具体格式：[ { "denom" : "atom", "amount" : "10" } ]   |

## 查询交易列表

URL: /txs/:page_index/:page_size
Method: GET
请求参数:
| 字段名称   | 请求体   |类型|   解释     |
| --------- |--------|-----| --------- |
| page_index| path   | int    | 起始页，从1开始   |
| page_size | path   | int    | 每页数量(0,1000] |
| block_hash| query  | String | 非必填,区块hash；index      |
| from      | query  | String | 非必填,发送address          |
| to        | query  | String | 非必填,接收address          |


返回字段:
| 字段名称   | 类型   | 解释                      |
| --------- | ------| --------------------------|
| tx_hash   | String | 交易hash；index           |
| timestamp | Int64| 时间戳                      |
| status    | Int64| 交易是否成功                 |
| block_height | Int64 | 属于那个区块的区块高度；index  |
| fee       | String| 交易手续费                  |
| used_gas  | Int64 | 消耗的gas数量               |
| memo      | String| 交易备注                    |
| from      | String| 交易发起账户地址             |
| tx_type   | String| 交易类型；值：send           |
| to        | String| 交易接受账户地址             |
| amount    | String| 转账了多少币。不是转账交易，此字段为空;  具体格式：[ { "denom" : "atom", "amount" : "10" } ]   |


# 验证人接口

## 获取验证人详情

URL: /validator/:oper_address
Method: GET
请求参数:
| 字段名称     | 请求体   | 类型  |   解释              |
| ------------|-------- |------ | ------------------|
| oper_address| path   | String| 验证人address        |

返回字段:
| 字段名称          | 类型   | 解释          |
| ---------------- | ----- | ------------- |
| moniker          | String| 验证人名称     |
| operator_address | String| 验证人address  |
| owner_address    | String| 自委托人地址   |
| reward_address   | String| 收益地址      |
| details          | String| 验证人详情     |
| voting_power     | String| 投票权重       |
| website          | String| 验证人的网址    |

# 地址接口

## 获取地址详情

URL: /address/:addr
Method: GET
请求参数:
请求参数:
| 字段名称   | 请求体   | 类型  |   解释             |
| --------- |-------- |------ | ------------------|
| addr      | path    | String| address地址        |

返回字段:
| 字段名称         | 类型   | 解释                       |
| --------------- | ------| ---------------------------|
| address         | String| 地址；index                 |
| balance         | int64 | 余额;                       |
| delegated       | String| 本地址被委托的Token数量;有可能为空； |
| reward_address  | String| 接受收益的地址；有可能为空; 如果address不是validator，此字段为空  |
| transaction_num | int64 | 本地址发送的交易总数量；|