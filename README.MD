## subot
a server, multiple language server

---
> feature 正在构思中...

1. `github` 自动贴砖定时任务




---
#### 目录
- proto     : pb 文件定义
- gen       : 由*.proto 生成的公共模板/基类依赖文件
- script    : 部分脚本
- luna      : GO服务器
  - docs    : 文档
  - dev     : 环境准备和生成
  - cmd     : 项目入口
  - internal: 具体的server代码
    - api   : 接口api
    - service
    - db
    - errx
    - models
    - mocks
    - ent
    - utils
    - config
  - migrations ： 数据库变动 
- lan  : java 服务器
- web  : web 相关
- app  : app 相关
