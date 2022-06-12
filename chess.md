---
title: 象棋 v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.11"

---

# 象棋

> v1.0.0

# User

## POST 注册

POST /chess/user/register

> Body 请求参数

```yaml
email: 1725014728@qq.com
username: mj
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 是 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 登录

POST /chess/user/login

> Body 请求参数

```yaml
account: mj
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» account|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# action

## POST 进入房间

POST /chess/room/enter

> Body 请求参数

```yaml
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjUsIlR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NTUwNzQyMTF9.HyKrzQl75WvDvectLjX-7ZhPP-xHY1fCe1WDe9b-OLM

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|room_id|query|string| 是 |none|
|body|body|object| 否 |none|
|» token|body|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 开房间

POST /chess/room/create

> Body 请求参数

```yaml
password: string
room_name: lmj
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjUsIlR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJUaW1lIjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJleHAiOjE2NTUxMTcxNjV9.7El7XNH4NFJ7h4e9pOcKCDr1jq8Z1aklyOrGRvCxLDE
room_id: "66888"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» password|body|string| 否 |none|
|» room_name|body|string| 否 |none|
|» token|body|string| 是 |none|
|» room_id|body|string| 是 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

