package main
import (
	"github.com/nats-io/go-nats"
)

func main() {

	nc,_ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Publish("foo",[]byte("hello"))
}
