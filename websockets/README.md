# Websocket example

## Test


- `go run server.go`
- The server will start on `localhost:8085`. You can connect to it from a WebSocket client by targeting `ws://localhost:8085/ws`.

- `curl --include --no-buffer --header "Connection: Upgrade" --header "Upgrade: websocket" --header "Host: localhost:8081" --header "Origin: http://localhost:8085" --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" --header "Sec-WebSocket-Version: 13" http://localhost:8085/ws`
  - causing errors as curl version may not support websocket


- Maybe use websocat to test websocket connections
    - `websocat ws://localhost:8085/ws` (this one works)
