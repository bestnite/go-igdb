# go-igdb

[IGDB(Internet Game Database)API v4](https://api-docs.igdb.com) 的 Go 客户端库,基于 protobuf 端点通信,数据传输更高效。

[English](README.md)

## 特性

- 覆盖全部 IGDB API v4 端点(78 个,创建客户端时自动注册)
- 基于 protobuf 通信,传输效率高
- 泛型 `BaseEndpoint[T]` —— 所有端点暴露同一套类型安全 API(`Query` / `GetByID` / `GetByIDs` / `Count` / `Paginated`)
- 内置限速(4 req/s,即 IGDB 官方上限),超出自动排队
- 自动管理 Twitch OAuth token 的获取与刷新
- 网络错误与 `429 Too Many Requests` 自动重试(最多 10 次,间隔 3 秒)

## 安装

```bash
go get git.nite07.com/nite/go-igdb
```

## 快速开始

```go
package main

import (
	"context"
	"log"

	igdb "git.nite07.com/nite/go-igdb"
)

func main() {
	// Client ID / Secret 在 Twitch 开发者后台申请,
	// 库内自动获取并刷新 OAuth token,无需手动处理
	client := igdb.New("your-client-id", "your-client-secret")

	ctx := context.Background()

	// GetByID:按 ID 查单个游戏
	game, err := client.Games.GetByID(ctx, 1942)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("游戏 %d:%s\n", 1942, game.Name)

	// GetByIDs:批量查询,内部自动按 500 个一批分批请求
	games, err := client.Games.GetByIDs(ctx, []uint64{119171, 119133})
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range games {
		log.Printf("游戏 %d:%s\n", g.Id, g.Name)
	}

	// Count:该端点下的记录总数
	total, err := client.Games.Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("游戏总数:%d\n", total)

	// Paginated:按 offset/limit 翻页,
	// 等价于 "offset 0; limit 10; fields *; sort id asc;"
	page, err := client.Games.Paginated(ctx, 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range page {
		log.Println(g.Name)
	}

	// Query:直接写 APICalypse 查询语句,最灵活的方式
	top, err := client.Games.Query(ctx, "fields name,rating; sort rating desc; limit 1;")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("评分最高的游戏:%s\n", top[0].Name)

	// Query + where:带过滤条件
	filtered, err := client.Games.Query(ctx, "fields name; where rating > 70; limit 10;")
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range filtered {
		log.Println(g.Name)
	}

	// 所有端点共享同一套泛型 API,例如按 ID 查封面
	cover, err := client.Covers.GetByID(ctx, 65586)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("封面 image_id:%s\n", cover.ImageId)

	// 搜索端点支持 IGDB 的 search 指令,结果按名称相关度排序
	results, err := client.Search.Query(ctx, `search "zelda"; fields name,game; limit 5;`)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range results {
		log.Printf("搜索结果:%s(游戏 %d)\n", s.Name, s.Game.GetId())
	}

	// Webhooks 端点是独立 API(注册/注销回调),不提供查询能力
	webhooks, err := client.Webhooks.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("已注册 webhook 数:%d\n", len(webhooks))
}
```

全部端点常量见 [endpoint/endpoint_name.go](endpoint/endpoint_name.go)。

## API 说明

- 每个查询型端点都内嵌 `BaseEndpoint[T]`,`T` 为 protobuf 实体类型(如 `client.Games` 是包装了 `BaseEndpoint[pb.Game]` 的 `*endpoint.Games`)。
- `Query` 接受原生 [APICalypse](https://api-docs.igdb.com/#apicalypse-1) 查询语句,与 JSON API 使用同一套查询语言。
- 响应从 IGDB 的 protobuf 端点(`/v4/<端点>.pb`)解码,不走 JSON。
- `New()` 返回的客户端已注册全部端点,通过字段直接访问,如 `client.Games`、`client.Covers`、`client.Search`。

## 依赖

- [go-resty/resty](https://github.com/go-resty/resty) — HTTP 客户端
- [google/protobuf](https://github.com/google/protobuf) — protobuf 运行时
- [golang.org/x/time](https://pkg.go.dev/golang.org/x/time) — 限速

## 参与贡献

欢迎提交 Issue 和 Pull Request。

## 许可证

本项目基于 MIT 许可证开源,详见 [LICENSE](LICENSE)。