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