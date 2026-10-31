```
📦blog
 ┣ 📂configs        // 配置文件
 ┣ 📂docs           // 文档
 ┣ 📂global         // 全局变量
 ┣ 📂internal       // 内部模块
 ┃ ┣ 📂dao          // 数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL
 ┃ ┣ 📂middleware   // HTTP中间件
 ┃ ┣ 📂model        // 模型层，用于存放 model 对象
 ┃ ┣ 📂routers      // 路由逻辑
 ┃ ┗ 📂service      // 核心业务
 ┣ 📂pkg            // 项目相关的模块包
 ┣ 📂scripts        // 各类构建，安装，分析等操作的脚本
 ┣ 📂storage        // 临时生成的文件
 ┣ 📂third_party    // 第三方的资源工具，例如 Swagger UI
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜main.go
 ┗ 📜README.md
 ```

## 数据库设计
文章表: blog_article
标签表: blog_tag
文章标签关联表: blog_article_tag

> 一篇文章有多个标签、一个标签属于多篇文章——这叫多对多，两张表直接外键是表达不了的，必须拆出中间表 blog_article_tag，每行 = "文章 1 ↔ 标签 Go" 这样一条配对。

> GET：读取/检索动作。
> POST：新增/新建动作。
> PUT：更新动作，用于更新一个完整的资源，要求为幂等。
> PATCH：更新动作，用于更新某一个资源的一个组成部分，也就是只需要更新该资源的某一项，就应该使用 PATCH 而不是 PUT，可以不幂等。
> DELETE：删除动作。

## 快速开始

```bash
podman compose up -d          # 首次启动自动执行 scripts/schema.sql 建表
```

- 数据库 `blog`，root 密码 `root`，端口 `3306`
- 改了 `scripts/schema.sql` 想重来：`podman compose down -v && podman compose up -d`（`-v` 删数据卷，触发重新建表）