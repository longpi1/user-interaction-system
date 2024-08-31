# 点赞系统

## 1.项目介绍

用golang实现一个点赞模块



## 2.基础功能模块

1. 点赞/点踩
2. 查看点赞历史
3. 查看点赞数



## 3.架构设计

#### 相关服务

**Service：**`like-service`



**Job：**`like-job`

用于读取消息队列的数据实现写操作，消息队列的最大用途是削峰处理。当写入请求非常大的时候，通过异步消息队列处理写请求。



#### **点赞的核心逻辑**





## 4.存储设计

#### 数据库表设计

**1. 点赞记录表 (likes)**

```sql
CREATE TABLE likes (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,  -- 自增ID，用于唯一标识每条记录
    uid BIGINT NOT NULL,              -- 用户ID
    business_id BIGINT NOT NULL,           -- 业务ID
    resource_id BIGINT NOT NULL,            -- 被点赞的实体ID
    source VARCHAR(50) NOT NULL,           -- 点赞来源
    op_type   int NOT NULL, -- 操作类型  点赞/点踩
    type   int NOT NULL, -- 点赞类型  
    status int NOT NULL,  -- 状态 0未点赞，1已点赞
    is_delete tinyint not null default '0' comment '是否逻辑删除',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 点赞时间
    update_time timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '更新时间',
    ext     Longtext,  -- 额外信息存储
    INDEX idx_user_message (uid, message_id)  -- 联合索引
);
```

- 数据类型:
  - `uid`: 使用 INT 或 BIGINT 作为用户 ID 类型，避免不必要的存储空间浪费。
  - `resource_id`: 使用 INT 或 BIGINT 作为消息 ID 类型，同样避免存储空间浪费。
  - `source`: 使用 ENUM 类型，定义明确的点赞来源，例如 "app", "web", "social_media" 等。
  - `timestamp`: 使用 DATETIME 类型存储点赞时间，方便后续查询和分析。
- 索引:
  - 除了 `mid` 和 `resource_id` 联合索引外，还可以考虑添加 `timestamp` 索引，方便根据时间范围查询点赞记录。
  - 考虑使用覆盖索引，例如在 `uid` 和 `resource_id` 联合索引上，覆盖 `timestamp` 和 `source` 字段，减少查询时读取额外数据。
- 分表策略:
  - 为了应对高 QPS，可以考虑将 `likes` 表进行分表，例如按照 `uid` 的哈希值进行分表，或者按照 `resource_id` 进行分表。
  - 分表后，可以根据查询条件选择合适的表进行查询，避免全表扫描，进一步提升查询效率。

**2. 点赞数表 (counts)**

```sql
CREATE TABLE counts (
    business_id BIGINT NOT NULL,           -- 业务ID
    resource_id BIGINT NOT NULL,            -- 实体ID
    likes_count INT DEFAULT 0,             -- 点赞数
    dislikes_count INT DEFAULT 0,          -- 点踩数
    PRIMARY KEY (business_id, message_id), -- 复合主键
    INDEX idx_message_id (message_id)      -- 单字段索引
);
```

- 数据类型:
  - `business_id`: 使用 INT 或 BIGINT 作为业务 ID 类型。
  - `resource_id`: 使用 INT 或 BIGINT 作为消息 ID 类型。
  - `like_count`: 使用 INT 或 BIGINT 存储点赞数，根据实际情况选择合适的数据类型。
  - `dislike_count`: 使用 INT 或 BIGINT 存储点踩数。
- 索引:
  - 除了 `resource_id` 索引外，还可以考虑添加 `business_id` 索引，方便根据业务 ID 查询点赞数。
  - 考虑使用覆盖索引，例如在 `resource_id` 索引上，覆盖 `like_count` 和 `dislike_count` 字段，减少查询时读取额外数据。
- 数据更新策略:
  - 考虑使用乐观锁或悲观锁机制，保证数据更新的原子性。
  - 对于点赞数和点踩数的更新，可以考虑使用事务处理，确保数据一致性。
- 分表策略:
  - 类似于 `likes` 表，可以考虑将 `counts` 表进行分表，例如按照 `business_id` 或 `resource_id` 进行分表。

**3. 补充:**

- **缓存:** 使用 Redis 等缓存系统缓存点赞数，避免频繁查询数据库，提升系统性能。
- **异步处理:** 将点赞记录的插入操作异步处理，避免影响主线程的性能。
- **监控和报警:** 监控系统性能，及时发现问题并进行报警，以便及时处理。
- **数据冗余:** 可以考虑在 `likes` 表中增加 `like_count` 和 `dislike_count` 字段，用于快速获取点赞数和点踩数，减少对 `counts` 表的查询。



#### 缓存设计

todo



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




## 8.参考链接

8.1[B站评论系统架构设计](https://www.bilibili.com/read/cv20346888/?spm_id_from=333.999.0.0)

8.2[百亿数据个性化推荐：弹幕工程架构演进]()

8.3[百亿数据百万查询——关系链架构演进](https://www.bilibili.com/read/cv24151036/?spm_id_from=333.999.0.0)

8.4[10Wqps评论中台，如何架构？B站是这么做的！！！](https://www.cnblogs.com/crazymakercircle/p/17197091.html)

8.5[领域驱动点播直播弹幕业务合并设计实践](https://www.bilibili.com/read/cv24830816/?spm_id_from=333.999.0.0)

8.6[【点个赞吧】 - B站千亿级点赞系统服务架构设计](https://www.bilibili.com/read/cv21576373/?spm_id_from=333.999.0.0)
