# superset

Go 实践项目集合（pnpm 风格的 Go workspace monorepo），持续沉淀为一个完整的作品。工作区由根目录 `go.work` 统一管理：`apps/` 放可运行应用，共享库以后放 `packages/`。

- **`apps/cli`** — 基于 [cobra](https://github.com/spf13/cobra) 的 CLI 小工具，提供单词格式转换与时间计算
- **`apps/blog`** — 基于 [Gin](https://github.com/gin-gonic/gin) 的博客系统（教程实践中）

一次构建全部应用：

```bash
go build github.com/luobochuanqi/superset/apps/...
```

## 功能

`word` 子命令支持 5 种转换模式：

| 模式 | 说明 | 示例 |
|------|------|------|
| 1 | 全部转大写 | `hello` → `HELLO` |
| 2 | 全部转小写 | `HELLO` → `hello` |
| 3 | 下划线转大写驼峰 | `hello_world` → `HelloWorld` |
| 4 | 下划线转小写驼峰 | `hello_world` → `helloWorld` |
| 5 | 驼峰转下划线 | `HelloWorld` → `hello_world` |

## 构建与使用

### CLI

```bash
go build -o superset ./apps/cli
```

单词转换：

```bash
superset word -s hello_world -m 3
# 转换结果：HelloWorld
```

参数说明：

- `-s, --string`：需要转换的字符串
- `-m, --mode`：转换模式（1-5）

时间计算：

```bash
superset time now
# 当前时间：2026-01-02 15:04:05, 1735824245

superset time calc -t "2026-01-01 10:00" -d 1h30m
# 计算结果：2026-01-01 11:30, 1767267000
```

`calc` 在基准时间（`-t` 省略时为当前时间）上加上时间间隔 `-d`，`-t` 支持三种格式：`2006-01-02`、`2006-01-02 15:04`、`2006-01-02 15:04:05`。

查看帮助：

```bash
superset --help
```

### blog

Gin 博客系统（实践进行中），独立子模块：

```bash
cd apps/blog
go run .                          # 启动 http://localhost:8080
curl http://localhost:8080/ping   # {"message":"pong"}
```
