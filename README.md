# go-igdb

A Go client library for the [IGDB (Internet Game Database) API v4](https://api-docs.igdb.com), built on protobuf endpoints for efficient data transfer.

[中文文档](README.zh.md)

## Features

- Full support for all IGDB API v4 endpoints (78 endpoints, auto-registered on the client)
- Protobuf-based communication for efficient data transfer
- Generic `BaseEndpoint[T]` — every endpoint exposes the same typed API (`Query` / `GetByID` / `GetByIDs` / `Count` / `Paginated`)
- Rate limiting (4 req/s, the IGDB limit)
- Automatic Twitch OAuth token management and refresh
- Automatic retry on network errors and `429 Too Many Requests` (up to 10 attempts, 3s interval)

## Installation

```bash
go get git.nite07.com/nite/go-igdb
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	igdb "git.nite07.com/nite/go-igdb"
	"git.nite07.com/nite/go-igdb/endpoint"
)

// 每 4 个请求限速 4 req/s,超出的自动排队等待
// 限流 4 req/s(IGDB 上限),超出部分自动等待
func main() {
	// Client ID / Secret 来自 Twitch Developer Portal
	// Twitch 开发者后台申请,库内自动获取并刷新 OAuth token
	client := igdb.New("your-client-id", "your-client-secret")

	ctx := context.Background()

	// GetByID: 按 ID 查单个游戏
	// Get a single game by ID
	game, err := client.Games.GetByID(ctx, 1942)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Game %d: %s\n", 1942, game.Name)

	// GetByIDs: 批量查询,自动按 500 个一批分批请求
	// Batch query by IDs (auto-batched, 500 per request)
	games, err := client.Games.GetByIDs(ctx, []uint64{119171, 119133})
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range games {
		log.Printf("游戏 %d:%s\n", g.Id, g.Name)
	}

	// Count: 该端点下的记录总数
	// Total number of records on the endpoint
	total, err := client.Games.Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total games: %d\n", total)

	// Paginated: 按 offset/limit 翻页,等价于 "offset 0; limit 10; fields *; sort id asc;"
	// Paginated fetch, equivalent to "offset 0; limit 10; fields *; sort id asc;"
	page, err := client.Games.Paginated(ctx, 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range page {
		log.Println(g.Name)
	}

	// Query: 直接写 APICalypse 查询语句,最灵活的方式
	// Raw APICalypse query for full flexibility
	top, err := client.Games.Query(ctx, "fields name,rating; sort rating desc; limit 1;")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Highest rated: %s\n", top[0].Name)

	// Query + where: 过滤条件
	// Query with a where filter
	filtered, err := client.Games.Query(ctx, "fields name; where rating > 70; limit 10;")
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range filtered {
		log.Println(g.Name)
	}

	// 每个端点都是同一套泛型 API,例如按封面 ID 查封面
	// Every endpoint shares the same generic API
	cover, err := client.Covers.GetByID(ctx, 65586)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cover image_id: %s\n", cover.ImageId)

	// 搜索端点支持 IGDB 的 search 指令,按名称相关度排序
	// The Search endpoint supports IGDB's `search` directive
	results, err := client.Search.Query(ctx, `search "zelda"; fields name,game; limit 5;`)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range results {
		log.Printf("Result: %s (game %d)\n", s.Name, s.Game.GetId())
	}

	// Webhooks 端点是独立 API(注册/注销回调),不属于查询型端点
	// The Webhooks endpoint has its own dedicated API
	webhooks, err := client.Webhooks.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered webhooks: %d\n", len(webhooks))

	_ = endpoint.EPGames // 所有端点常量见 endpoint/endpoint_name.go
	_ = client
}
```

All endpoint names are listed in [endpoint/endpoint_name.go](endpoint/endpoint_name.go).

## API Notes

- Every query-style endpoint embeds `BaseEndpoint[T]` where `T` is the protobuf entity (e.g. `client.Games` is `*endpoint.Games` wrapping `BaseEndpoint[pb.Game]`).
- `Query` accepts raw [APICalypse](https://api-docs.igdb.com/#apicalypse-1) — the same query language used by the JSON API.
- Responses are decoded from IGDB's protobuf endpoints (`/v4/<endpoint>.pb`), not JSON.
- `New()` returns a client with all endpoints registered; access them via fields like `client.Games`, `client.Covers`, `client.Search`.

## Dependencies

- [go-resty/resty](https://github.com/go-resty/resty) — HTTP client
- [google/protobuf](https://github.com/google/protobuf) — protobuf runtime
- [golang.org/x/time](https://pkg.go.dev/golang.org/x/time) — rate limiting

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.