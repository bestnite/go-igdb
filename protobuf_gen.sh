go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

wget https://api.igdb.com/v4/igdbapi.proto -O ./igdbapi.proto
protoc --go_out=. --go_opt=Mproto/igdbapi.proto=/proto ./proto/igdbapi.proto
