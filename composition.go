package main
import (
	"bytes"
	"io"
	"net"
)
type recordingConn struct {
        net.Conn
        io.Writer
}
func (c *recordingConn) Write(buf []byte) (int, error) {
        return c.Writer.Write(buf)
}
func main() {
	client, server := net.Pipe()
	var buf bytes.Buffer
	client = &recordingConn {
		Conn: client,
		Writer: io.MultiWriter(client, &buf),
	}
	// ...
}
