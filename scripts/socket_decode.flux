import "socket"

// csv decoder
// socket.from(url: "tcp://127.0.0.1:1234", decoder: "csv")

// line decoder
socket.from(url: "tcp://127.0.0.1:1234", decoder: "line")
