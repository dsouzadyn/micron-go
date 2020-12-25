# :microscope: micron-go

micron-go is a realy tiny microservice template.
It comes with a bare minimum server and a config file.
Does not come a database library :).

## Running

```sh
$ go build -o ./bin/server
$ ./bin/server
```

## Testing
```sh
$ go test ./... # The 3 '.'s are not a typo
```

## Warning

Incase you decide to change the cloned repo name please update it in these files...
```
go.mod
server.go
```