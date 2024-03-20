# relation

### 接口设计

1. 关注用户:
    - 接口:POST /api/relation/follow
    - 请求参数:
        - userId:当前用户ID
        - followUserId:要关注的用户ID
    - 返回结果:
        - success:是否关注成功
        - message:关注结果信息
2. 获取关注列表:
    - 接口:GET /api/relation/get_follower
    - 请求参数:
        - userId:当前用户ID
        - page:页码
        - pageSize:每页数量
    - 返回结果:
        - userList:关注的用户列表
        - totalCount:关注用户总数
3. 获取粉丝列表:
    - 接口:GET /api/relation/get_fans
    - 请求参数:
        - userId:当前用户ID
        - page:页码
        - pageSize:每页数量
    - 返回结果:
        - userList:粉丝用户列表
        - totalCount:粉丝总数



关键实现:

1. 使用Redis缓存用户的关注信息,设计合适的数据结构存储关注关系。
2. 关注和取消关注操作先更新缓存,再通过消息队列异步更新数据库,提高响应速度和并发能力。
3. 获取关注列表和粉丝列表时,先查询缓存,缓存未命中再查询数据库,减轻数据库压力。
4. 使用布隆过滤器快速判断用户是否已关注,避免缓存穿透。
5. 对热点用户的关注信息进行分片存储,避免单个Redis节点成为性能瓶颈。
6. 数据库表设计时,使用适当的索引优化查询性能。
