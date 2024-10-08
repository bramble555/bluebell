#### bluebell项目

## /api/v1:
        /signup 注册功能
        /login 登录功能
        /community 获取所有社区详情
        /community/:id 获取某个社区详情
        /post 发布帖子
        /post/:id 获取某个帖子详情
        /posts 获取所有帖子
        /vote 为帖子投票

## 学到的知识:

1. 项目脚手架 (control 层获取参数, logic 层进行逻辑处理,dao 层进行数据处理)
2. 雪花算法 分布式ID
3. 令牌桶 进行限速
4. 初步会编写 Makefilew 文件 (本质就是很多命令简化)
5. 初步会编写 Dockerfile 和d ocker-compose.yaml
6. ab 测压工具  (-n 请求个数 -c 并发量 -t 多少s)
``` 
$ ab -n 1000 -c 50 -t 10 URL
```
7. pprof 性能分析 (cpu memory...)


## to do  未实现
- 返回前端没有进行validator，返回的是英文
- air 热加载