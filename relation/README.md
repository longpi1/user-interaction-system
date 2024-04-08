# relation

## relation-service

## 1.项目介绍

用golang实现一个关注模块




## 2.基础功能模块

### 关注/取关操作

1. 关注操作:
   - 根据传递的参数判断是否为关注行为。
   - 插入relation表记录,插入超时或失败则报错返回。
   - 如果relation插入成功,则删除相关缓存记录。
2. 取关操作:
   - 根据传递的参数判断是否为取关行为。
   - 如果是取关行为查询数据库relation表是否存在对应记录，如果存在则进行删除，不存在则无需处理；
   - 删除成功后进行缓存更新

### 查询关注关系

1. 查询我的关注列表:
   - 根据uid查询relation表,分页返回结果。
2. 查询我/资源的粉丝列表:
   - 根据resourceId查询relation表,分页返回结果。
3. 查询A是否关注B:
   - 优先查询Redis,判断A的关注列表是否包含B，不存在则查询relation表。
4. 批量查询A是否关注B、C、D:
   - 优先查询Redis,一次获取A的关注列表,判断是否包含B、C、D不存在则查询relation表。
5. 查询A和B是否互相关注:
   - 先查询A是否关注B,再查询B是否关注A。

### 关注数、粉丝数

1. 查询用户或者资源的关注数
2. 查询粉丝数

### 消息推送

### **其他**

- 使用一致性哈希算法，根据uid计算分表索引，以提高查询性能。
- 使用缓存机制，减少对数据库的访问。
- 定期清理缓存，以保持缓存的有效期。

## 3.架构设计







## 4.存储设计

#### 数据库设计

**采用relation表来存储关注关系:**

```go
type Relation struct {
    ID         int64 `json:"id"`   //主键id
    Source int64 `json:"source"`  //来源
    UID   int64 `json:"uid"`   // 用户id，也就是发起关注行为的用户id
    ResourceID int64 `json:"resource_id"` // 被关注的资源或者人
    Platform int64 `json:"platform"` // 相关的平台
    Status     int   `json:"status"`   // 状态
    Type     int     `gorm:"comment:'类型'"`    // 类型
    CreatedAt  int64 `json:"created_at"` // 发起关注时间
    UpdateAt  int64 `json:"update_at"`
    Ext       string  `json:"ext"` // 额外信息
}
```

分表可以使用一致性哈希算法,根据uid计算分表索引。这样可以提高查询性能,同时也可以应对未来的高并发和海量数据。

unique_key： uid_type_resource_id

**采用relation_counts表来存储关注数/粉丝数关系:**

```go
// RelationCount 用户关注数/粉丝数表
type RelationCount struct {
	gorm.Model
	ResourceId  int64  `json:"resource_id"`  // 资源/用户id
	FansCount   int64  `json:"fans_count"`   // 粉丝数
	FollowCount int64  `json:"follow_count"` // 关注数
	Platform    int64  `json:"platform"`     // 相关的平台
	Type        int64  `json:"type"`         // 资源类型
	Ext         string `json:"ext"`          // 额外信息`
}

```



#### 缓存设计

我们将用户的关注列表缓存到Redis中,使用Hash结构存储。Key为用户ID,Field为被关注的用户ID,Value为null。

- 关注/取关操作时,同步更新Redis缓存。
- 查询关注关系时,优先从Redis中查询,Redis miss时再查询数据库。
- Redis缓存过期时间设置为1天,利用定时任务定期更新过期数据



## 5.可用性设计

针对QPS大的服务与接口，可以采取以下措施来有效缓解热点事件带来的读写压力增加：

1. **缓存技术**：使用缓存技术可以降低数据库访问压力，提高接口响应速度。通过将热点数据缓存在内存中，可以减少对数据库的频繁访问，特别是针对读多写少的场景[4](https://www.51cto.com/article/773129.html)。

2. **分布式缓存**：采用分布式缓存可以将数据分布到多台服务器上，提高了系统整体的读取性能和并发能力。常见的分布式缓存包括Redis、Memcached等。

3. **本地缓存**：对于一些热点数据，可以考虑使用本地缓存技术，将数据缓存在应用服务器的内存中，减少对分布式缓存或数据库的访问[4](https://www.51cto.com/article/773129.html)。

4. **缓存预热**：在系统启动或者更新热点数据时，可以进行缓存预热，将热点数据加载到缓存中，避免在系统运行过程中突然的读写压力增加[4](https://www.51cto.com/article/773129.html)。

5. **读写分离**：针对读多写少的场景，可以考虑使用读写分离技术，将读请求和写请求分发到不同的数据库实例上，从而提高系统整体的读取性能[1](https://www.zhihu.com/question/458352302/answer/2990748762)。

6. **限流与熔断**：对于热点事件，可以通过限流和熔断等手段来控制请求的并发量，避免突发的请求对系统造成过大的压力[1](https://www.zhihu.com/question/458352302/answer/2990748762)。

   综上所述，通过合理运用缓存技术、读写分离、限流与熔断等手段，可以有效缓解热点事件带来的读写压力增加，提高系统的稳定性和性能。

#### Singlefilght

对于热门的主题，如果存在缓存穿透（缓存中没有数据，请求穿透了缓存，直接打到数据库）的情况，会导致大量的同进程、跨进程的数据回源到存储层，可能会引起存储过载的情况，如何只交给同进程内，一个人去做加载存储。

使用归并回源的思路：

[singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight)

同进程只交给一个请求去获取 `mysql` 数据，然后批量返回。同时这个 `lease owner` 投递一个 `kafka` 消息，做 `index cache` 的 `recovery` 操作。这样可以大大减少 `mysql` 的压力，以及大量穿透导致的密集写 `kafka` 的问题。

更进一步的，后续连续的请求，仍然可能会短时 `cache miss` ，我们可以在进程内设置了一个 `short-lived flag`，标记最近有一个请求投递了 `cache rebuild` 的消息，直接 `drop`。

可以看到，这里说明的都是单进程下的解决思路。那么在多进程下，能否使用分布式锁来解决。理论上可以，但是实际操作起来，容易将这个简单问题复杂化，不推荐使用分布式锁。（PS：`redis` 作者不推荐使用 `redis` 实现分布式锁。）

多进程下，也是一样的思想，多个进程会发送多个消息到消息队列中，消费端获取消息的时候，通过单飞的思路，同样处理。

#### 热点

热点分为写热点和读热点。

写操作一般会通过MQ削峰，当大量的请求都集中在 `MQ` 中，不仅仅会影响当前服务，还可能导致下游服务出现异常。这种情况下，可以再进行解耦，增加上游服务的吞吐，将下游服务解耦，不依赖同一个同步逻辑。

流量热点是因为突然热门的主题，被高频次的访问，因为底层的 `cache` 设计，一般是按照主题 `key` 进行一致性 `hash` 来进行分片，但是热点 `key` 一定命中某一个节点，这时 `remote cache` 可能会变成瓶颈。因此做 `cache` 升级 `local cache` 是有必要的，一般使用**单进程自适应发现热点**的思路，附加一个短时的 `ttl local cache`，可以在进程内吞掉大量的读请求。

![image-20231007174023897](https://raw.githubusercontent.com/xiaoyeshiyu/image-hosting-service/main/uPic/2023/10/image-20231007174023897.png)

在内存中使用 `hashmap` 统计每个 `key` 的访问频次，这里可以使用滑动窗口（左角标和右角标一起移动，统计区间内部的数据量）统计，即每个窗口中，维护一个 `hashmap`，之后统计所有未过去的 `bucket`，汇总所有 `key` 的数据。

之后使用小顶堆计算 `TopK` 的数据，自动进行热点识别。



## 6.安全性设计





## 7.功能详细设计

### 接口设计

### 1. 关注/取关操作接口

- **HTTP方法**：POST

- **路径**：`/api/relation/relation`

- 请求体

  ```go
  {
    "source": "用户来源",
    "uid": "发起关注的用户ID",
    "type": "资源类型",
    "platform": "平台",
    "op_type": "操作类型,关注或者取关",
    "resource_id": "被关注的资源或人的ID"
  }
  ```

- 成功响应

  ```go
  {
    "code": 200,
    "message": "Follow/Unfollow operation successful."
  }
  ```

- 错误响应

  ```go
  {
    "code": 400,
    "message": "Error message based on the failure reason."
  }
  ```

### 2. 查询用户或者资源的关注数/粉丝数接口

- **HTTP方法**：GET

- **路径**：`/api/relation/relation_count`

- 请求参数

  - `uid`: 用户ID或资源ID
  - type: 资源类型
  - platform： 平台

- 响应体

  ```go
  {
    "code": 200,
    "followCount": 102,
     "fansCount": 102,  
    "message": "Query successful."
  }
  ```

### 3. 查询我的关注列表接口

- **HTTP方法**：GET

- **路径**：`/api/relation/following`

- 请求参数

  - `uid`: 用户ID
  - `page`: 页码
  - `limit`: 每页数量

- 响应体

  ```go
  {
    "code": 200,
    "data": [
      {
        "resource_id": "被关注资源或人的ID",
        "type": "资源类型",
        "status": "关注状态",
        "created_at": "关注时间"
      }
    ],
    "message": "Query successful."
  }
  ```

### 4. 查询我/资源的粉丝列表接口

- **HTTP方法**：GET

- **路径**：`/api/relation/fans`

- 请求参数

  - `resource_id`: 资源或人的ID
  - `page`: 页码
  - `limit`: 每页数量

- 响应体

  ：

  ```go
  {
    "code": 200,
    "data": [
      {
        "uid": "粉丝的用户ID",
        "type": "资源类型",
        "status": "关注状态",
        "created_at": "关注时间"
      }
    ],
    "message": "Query successful."
  }
  ```

### 5. 查询A是否关注B接口

- **HTTP方法**：GET

- **路径**：`/api/relation/isFollowing`

- 请求参数

  - `uid`: A的用户ID
  - `resource_id`: B的资源或人的ID

- 响应体

  ```go
  {
    "code": 200,
    "isFollowing": true,
    "message": "Query successful."
  }
  ```

### 6. 批量查询A是否关注B、C、D接口

- **HTTP方法**：POST

- **路径**：`/api/relation/isFollowingBatch`

- 请求体

  ```go
  {
    "uid": "A的用户ID",
    "resource_ids": ["B的ID", "C的ID", "D的ID"]
  }
  ```

- 响应体

  ```go
  {
    "code": 200,
    "data": {
      "B的ID": true,
      "C的ID": false,
      "D的ID": true
    },
    "message": "Query successful."
  }
  ```

### 7. 查询A和B是否互相关注接口

- **HTTP方法**：GET

- **路径**：`/api/relation/mutualFollow`

- 请求参数

  - `uid_a`: A的用户ID
  - `uid_b`: B的用户ID

- 响应体

  ```go
  {
    "code": 200,
    "isMutualFollow": true,
    "message": "Query successful."
  }
  ```

关键实现:

1. 使用Redis缓存用户的关注信息,设计合适的数据结构存储关注关系。
2. 关注和取消关注操作先更新缓存,再通过消息队列异步更新数据库,提高响应速度和并发能力。
3. 获取关注列表和粉丝列表时,先查询缓存,缓存未命中再查询数据库,减轻数据库压力。
4. 使用布隆过滤器快速判断用户是否已关注,避免缓存穿透。
5. 对热点用户的关注信息进行分片存储,避免单个Redis节点成为性能瓶颈。
6. 数据库表设计时,使用适当的索引优化查询性能。



## 8.参考链接



### 
