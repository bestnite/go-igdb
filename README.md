# go-igdb

A Go client library for the IGDB (Internet Game Database) API v4. This library provides a simple and efficient way to interact with IGDB's protobuf-based API.

## Features

- Full support for IGDB API v4
- Protobuf-based communication for better performance
- Automatic token management for Twitch authentication
- Built-in retry mechanism for failed requests
- Optional Cloudflare bypass support via FlareSolverr
- All endpoints are supported

## Installation

```bash
go get github.com/bestnite/go-igdb
```

## Quick Start

```go
// Create a new IGDB client
client := igdb.New("your-client-id", "your-client-secret")

// Get a game by ID
game, err := client.GetGameByID(1942)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Game: %s\n", game.Name)

// Search games with custom query
games, err := client.GetGames("search \"Zelda\"; fields name,rating,release_dates.*;")
if err != nil {
    log.Fatal(err)
}
```

## Advanced Usage

### Query Format

The library uses IGDB's query syntax. For example:

```go
// Get games released in 2023
games, err := client.GetGames("where release_dates.y = 2023; fields name,rating;")

// Get specific fields for a company
companies, err := client.GetCompanies("where id = 1234; fields name,description,country;")
```

## Requirements

- Go 1.24 or higher
- IGDB/Twitch API credentials (Client ID and Client Secret)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
