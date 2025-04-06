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
)

func Test1(c *igdb.Client) {
	game, err := c.Games.GetByID(1942)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Name of game %d: %s\n", 1942, game.Name)
}

func Test2(c *igdb.Client) {
	games, err := c.Games.GetByIDs([]uint64{119171, 119133})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Names of games %d and %d: %s and %s\n", 119171, 119133, games[0].Name, games[1].Name)
}

func Test3(c *igdb.Client) {
	total, err := c.Games.GetLastOneId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total number of games: %d\n", total)
}

func Test4(c *igdb.Client) {
	games, err := c.Games.Paginated(0, 10)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Names of ids 0 to 10 games:\n")
	for _, game := range games {
		log.Println(game.Name)
	}
}

func Test5(c *igdb.Client) {
	game, err := c.Games.Query("fields name,rating; sort rating desc; limit 1;")
	if err != nil {
		log.Fatalf("failed to get game: %s", err)
	}
	log.Printf("Name of first game with highest rating: %s\n", game[0].Name)
}

func Test6(c *igdb.Client) {
	games, err := c.Games.Query("fields *; where rating > 70; limit 10;")
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

## Example Projects

- [igdb-database](https://github.com/bestnite/igdb-database)

## Dependencies

- [go-resty/resty](https://github.com/go-resty/resty)
- [google/protobuf](https://github.com/google/protobuf)
- [bestnite/go-flaresolverr](https://github.com/bestnite/go-flaresolverr)
- [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
