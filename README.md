# superset

一个基于 [cobra](https://github.com/spf13/cobra) 的命令行小工具，提供单词格式转换与时间计算。

## 功能

`word` 子命令支持 5 种转换模式：

| 模式 | 说明 | 示例 |
|------|------|------|
| 1 | 全部转大写 | `hello` → `HELLO` |
| 2 | 全部转小写 | `HELLO` → `hello` |
| 3 | 下划线转大写驼峰 | `hello_world` → `HelloWorld` |
| 4 | 下划线转小写驼峰 | `hello_world` → `helloWorld` |
| 5 | 驼峰转下划线 | `HelloWorld` → `hello_world` |

## 安装

```bash
go install github.com/luobochuanqi/superset@latest
```

## 使用

### word

```bash
superset word -s hello_world -m 3
# 转换结果：HelloWorld
```

参数说明：

- `-s, --string`：需要转换的字符串
- `-m, --mode`：转换模式（1-5）

### time

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
