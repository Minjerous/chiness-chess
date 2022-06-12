# Chess
[![Build Status](https://img.shields.io/badge/build-1.01-brightgreen)](https://travis-ci.org/pibigstar/go-todo)

> 红岩后端考核
>
>

## 1.项目结构

<details>
<summary>展开查看</summary>
<pre><code>
├─build
│  ├─action
│  └─user
├─common
│  ├─cryptx
│  ├─httpcode
│  │  └─statuscode
│  ├─jwt
│  ├─mq
│  ├─pprof                                                             
│  ├─rep                                                               
│  ├─rest                                                              
│  │  └─errdef                                                         
│  └─tool
├─services
│  ├─action
│  │  ├─chess-client
│  │  │  ├─chess
│  │  │  ├─resource
│  │  │  └─wscoon
│  │  ├─cmd
│  │  │  └─api
│  │  │      └─internal
│  │  │          ├─config
│  │  │          ├─dao
│  │  │          └─handler
│  │  │              ├─room
│  │  │              └─ws
│  │  ├─model
│  └─ws
├─sql
│
├─docker-compose.yaml
│
├─go.work


</pre></code>
</details>

## 2. 使用技术
- [x] 微服务架构(为了熟悉微服务并没有考虑使用goctl)
- [x] option设计模式(common组件和gorm)
- [x] go work 工作模式
- [x] docker-compose 配置
- [x] 自用组件(common)
- [x] pprof 性能测试
- [ ] mq消息推送
- [ ] etcd 注册发现服务
- [ ] redis 令牌桶限流and禁用用户
- [ ] redis 使用旁路缓存策略服务
- [ ] 热跟新机制组件

### 2.1 后端框架
基于`Gin`框架, golang 版本: `1.8.3`

### 2.2 接口文档


##### 配置文件

`/user/cmd/api/config/user.yml`
`/user/cmd/rpc/config/user.yml`
`/action/cmd/api/config/action.yml`

### 2.3 安全性

- [x] 密码加盐加密
- [x] 防止xxs注入 
- [x] JWT 
- [x] grpc CA认证
- [x] grpc 密码自定义设置和拦截器 


## 其他
  >其实这次考核挺可惜的,最关键的象棋部分没写完.由于一开始是用的微服务架构也没有用类似好快又好用的go-zero。导致了项目结构庞大，以至于浪费了很多时间，象棋逻辑就没有很充足的时间。写了一晚上还是有很多bug,索性调库！但是时间已经所剩不多了。熬了到了最后的12点也没有完成。
  >由于太累下午睡了一交加上晚上有重要活动。所以现在才写md。
  