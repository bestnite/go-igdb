# go-igdb

a go library to access IGDB API

## Usage

```go
g := igdb.New("clientID", "clientSecret")
game, err := g.GetGameByID(325594)
if err != nil {
    log.Fatal(err)
}
fmt.Println(game.Name)
```
