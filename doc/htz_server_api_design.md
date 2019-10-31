黄庭禅经典听读系统后端API设计文档
===
[TOC]

# 1. 介绍
因为用户量不大，为了加快开发速度，并易于以后的维护，所以没有采用微服务架构，而是采用把所有服务集成在一个程序中。

* 数据库： MongoDB

* 整个程序从逻辑上分为以下三个主要服务：
	+ 主逻辑服务： 负责提供系统的主要逻辑。
	+ 文件服务：负责文件的存储和读取。
	+ 搜索服务：负责提供系统的搜索服务。

* 接口约定
	+ 所有的API都是通过POST方法进行调用。
	+ 除了文件服务之外，其余所有的入参和出参都是采用JSON格式。
	+ 通过在URI的第一个参数来表示动作。
		- get：获取；例如：/get/user（获取用户信息）
		- delete：删除；例如：/delete/user（删除用户）
		- post：添加；例如：/post/user（添加用户）
		- put：更新；例如：/put/user（更新用户信息）
	+ 字段命名方式: 单词使用`_`作分隔，而非驼峰方式。
	+ 

# 2. 公共数据结构


* **Response** 统一的返回样式。**以下接口返回都只写了成功时的data字段。**
```
{
	int code - 执行结果。200 - 成功，其余为失败
	string err - 执行错误结果描述
	string data - json数据结构的执行结果
}
```
* **SutraInfo** 经典信息
```
{
	string id - 唯一标识符
	string name - 经典名称
	string cover - 封面文件ID
	string description - 经典简介
	uint64 played_count - 总共播放次数
	uint64 item_total - 一共有多少集
	string created_at - 创建时间
}
```
* **SutraItem** 每集经典的详情
```
{
	string id - 唯一标识符
	string sutra_id - 经典专辑唯一标识符
	string title - 标题
	string description - 介绍
	string original - 经典原文，存储在数据库，而不是在文件中
	string audio_id - 音频文件id
	string lyric_id - 歌词文件id，包含张讲师所有的话
	uint64 lesson - 第几集
	uint64 played_count - 播放次数
	uint64 duration - 时长
	string created_at
}
```
* **Recommendation** 推荐的经典专辑
```
{
	string sutra_id - 经典唯一标识符
	string sutra_name - 经典名称
	string sutra_desc - 经典简介
	string sutra_cover - 经典封面文件的id 
	uint64 sort - 排序位置 
}
```
* **ListenHistory** 收听记录
```
{
  string sutra_id - 经典唯一标识符
  string sutra_name - 经典名称
  string sutra_cover - 经典专辑封面文件id
  string sutra_item_id - 经典专辑条目id 
  string sutra_item_title - 经典专辑条目标题
  int64 last_position - 上次听到哪里，单位：秒。小于零表示已经听完。
}
```
# 3. 主逻辑服务

## 3.1. 用户
### 3.1.1. 注册
### 3.1.2. 登录
### 3.1.3. 登出
### 3.1.4. 收听记录
#### 3.1.4.1. 添加收听记录
* URI：/post/listen/history
* Inputs
```
ListenHistory - 参阅公共数据结构
```
* Outputs
```
Response - 参阅公共数据结构
```
#### 3.1.4.2. 获取收听记录
* URI: /get/listen/histories
* Inputs
```
无
```
* Outputs
```
{
	uint64 total - 总数量
	ListenHistory[] histories - 参阅公共数据结构
}
```

## 3.2. 经典
### 3.2.1. 推荐列表
* URI： /get/recommendations
* Inputs
```
空
```
* Outputs
```
{
	uint64 total - 总推荐数量
	Recommendation[] recommendations - json数组，推荐的经典信息，详细信息请参阅公共数据结构
}
```

### 3.2.2. 经典专辑列表
* URI：/get/sutra/items
* Inputs
```
{
	string sutra_id - 经典专辑ID 
	uint64 num - 每页显示多少条
	uint64 page - 第几页，从0开始。
}
```
* Outputs
```
{
	SutraItem[] items - json数组，专辑条目，详细信息请参阅公共数据结构
	uint64 total - 一共多少条目
	uint64 cur_page - 当前是第几页，从0开始。
	uint64 cur_num - 返回的数量
}
```

### 3.2.3. 经典详情
* URI：/get/sutra/info
* Inputs
```
{
	string id - 经典专辑唯一标识符
}
```
* Outputs
```
SutraInfo - 经典专辑的信息，详细信息请参阅公共数据结构
```

### 3.2.4. 列出所有经典
* URI：/get/sutras/all
* Inputs
```
无
```
* Outputs
```
{
	uint64 total - 总数量
	SutraInfo[] sutras - json数组，详细信息请参阅公共数据结构
}
```

## 3.3. 用户通知
### 3.3.1. 获取我的通知信息
* URI： /get/user/notifications
* Inputs
```
无
```

* Outputs
```
{
	uint64 total - 总数量
	notifications: [{
		string user_id - 用户ID
		bool is_read - 是否已读
		string title - 通知标题
		string msg - 通知内容
		string created_at - 产生时间
	}] - json数组，按照created_at倒序排列
}
```

### 3.3.2. 设置我的通知信息状态为已读
* URI： /put/user/notifications
* Inputs
```
{
	string[] notification_ids - json数组，可以同时设置多条信息为已读状态
}
```

* Outputs
```
Response - 参阅公共数据结构
```

# 4. 文件服务
## 4.1. 上传
* URI：/post/file/upload/{file_hash}/{mime}
	+ mime: 文件类型。请按照http协议规定的文件类型名填写
	+ file_hash: 文件sha1值，160位，20字节。十六进制表示，用于检查文件是否完整。如果有已经有相同hash值的文件，不会重复存储，会引用同一个文件。
* Inputs
```
文件字节流
```

* Outputs
```
{
	string file_id - 文件ID。id为空的话，表示上传失败
}
```

## 4.2. 下载
* URI：/get/file/download/{height}/{width}
	+ 下载图片时可以指定高度和宽度。TODO：第一期先不做，也就说只会下载原图
	+ TODO： 对常用大小的图片会进行缓存。
* Inputs
```
{
	string file_id - 文件ID
}
```

* Outputs
```
文件字节流
```

# 5. 搜索服务
## 5.1. sample
* URI：
* Inputs
```

```
* Outputs
```

```
