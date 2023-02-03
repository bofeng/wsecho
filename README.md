# ws echo

Simple websocket echo server built with [go's fiber framework](https://gofiber.io)

Run: `go run main.go`

Test:
```
❯ wscat --connect localhost:6001/ws
Connected (press CTRL+C to quit)
> hello
< hello
```