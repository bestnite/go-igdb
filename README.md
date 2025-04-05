# go-igdb

A Go client library for the IGDB (Internet Game Database) API v4. This library provides a convenient way to interact with IGDB's protobuf-based API endpoints.

## Features

- Full support for IGDB API v4 endpoints
- Protobuf-based communication for efficient data transfer
- Rate limiting support
- Automatic token management for Twitch authentication
- Retry mechanism for failed requests
- Optional FlareSolverr integration for handling Cloudflare protection

## Installation

```bash
go get github.com/bestnite/go-igdb
```

## Quick Start

```go
package main

import (
	"log"

	"github.com/bestnite/go-igdb"
	pb "github.com/bestnite/go-igdb/proto"
)

func Test1(c *igdb.Client) {
	game, err := igdb.GetItemByID[pb.Game](1942, c.Games.Query)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Name of game %d: %s\n", 1942, game.Name)
}

func Test2(c *igdb.Client) {
	games, err := igdb.GetItemsByIDs[pb.Game]([]uint64{119171, 119133}, c.Games.Query)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Names of games %d and %d: %s and %s\n", 119171, 119133, games[0].Name, games[1].Name)
}

func Test3(c *igdb.Client) {
	total, err := igdb.GetItemsLength[pb.Game](c.Games.Query)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total number of games: %d\n", total)
}

func Test4(c *igdb.Client) {
	games, err := igdb.GetItemsPagniated[pb.Game](0, 10, c.Games.Query)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Names of ids 0 to 10 games:\n")
	for _, game := range games {
		log.Println(game.Name)
	}
}

func Test5(c *igdb.Client) {
	game, err := igdb.AssertSingle[pb.Game](c.Games.Query("fields name,rating; sort rating desc; limit 1;"))
	if err != nil {
		log.Fatalf("failed to get game: %s", err)
	}
	log.Printf("Name of first game with highest rating: %s\n", game.Name)
}

func Test6(c *igdb.Client) {
	games, err := igdb.AssertSlice[pb.Game](c.Games.Query("fields *; where rating > 70; limit 10;"))
	if err != nil {
		panic(err)
	}
	log.Printf("Names of games with rating > 70 limit 10:\n")
	for _, game := range games {
		log.Println(game.Name)
	}
}

func main() {
	client := igdb.New("your-client-id", "your-client-secret")
	Test1(client)
	Test2(client)
	Test3(client)
	Test4(client)
	Test5(client)
	Test6(client)
}
```

## Advanced Usage

### Using with FlareSolverr

```go
import "github.com/bestnite/go-flaresolverr"

flaresolverr := flaresolverr.New("http://localhost:8191")
client := igdb.NewWithFlaresolverr("your-client-id", "your-client-secret", flaresolverr)
```

### Rate Limiting

The client automatically handles rate limiting with a default of 4 requests per second. This helps prevent hitting IGDB's rate limits.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
