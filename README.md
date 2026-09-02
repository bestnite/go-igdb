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
	"log"

	igdb "git.nite07.com/nite/go-igdb"
)

func main() {
	// Client ID / Secret come from the Twitch Developer Portal;
	// the library fetches and refreshes the OAuth token automatically
	client := igdb.New("your-client-id", "your-client-secret")

	ctx := context.Background()

	// GetByID: fetch a single game by ID
	game, err := client.Games.GetByID(ctx, 1942)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Game %d: %s\n", 1942, game.Name)

	// GetByIDs: batch query (auto-batched, 500 IDs per request)
	games, err := client.Games.GetByIDs(ctx, []uint64{119171, 119133})
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range games {
		log.Printf("Game %d: %s\n", g.Id, g.Name)
	}

	// Count: total number of records on the endpoint
	total, err := client.Games.Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total games: %d\n", total)

	// Paginated: offset/limit pagination,
	// equivalent to "offset 0; limit 10; fields *; sort id asc;"
	page, err := client.Games.Paginated(ctx, 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range page {
		log.Println(g.Name)
	}

	// Query: raw APICalypse query, the most flexible option
	top, err := client.Games.Query(ctx, "fields name,rating; sort rating desc; limit 1;")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Highest rated: %s\n", top[0].Name)

	// Query with a where filter
	filtered, err := client.Games.Query(ctx, "fields name; where rating > 70; limit 10;")
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range filtered {
		log.Println(g.Name)
	}

	// Every endpoint shares the same generic API, e.g. covers by ID
	cover, err := client.Covers.GetByID(ctx, 65586)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cover image_id: %s\n", cover.ImageId)

	// The Search endpoint supports IGDB's `search` directive,
	// results are ranked by name similarity
	results, err := client.Search.Query(ctx, `search "zelda"; fields name,game; limit 5;`)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range results {
		log.Printf("Result: %s (game %d)\n", s.Name, s.Game.GetId())
	}

	// The Webhooks endpoint has its own dedicated API (register/unregister),
	// it is not a query-style endpoint
	webhooks, err := client.Webhooks.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered webhooks: %d\n", len(webhooks))
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