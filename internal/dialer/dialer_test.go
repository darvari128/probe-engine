package dialer_test

import (
	"context"
	"net"
	"testing"

	"github.com/apex/log"
	"github.com/ooni/probe-engine/internal/dialer"
)

func TestIntegration(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var d dialer.Dialer
	d = dialer.Base()
	d = dialer.ErrWrapper{Dialer: d}
	events := &dialer.EventsSaver{Dialer: d}
	d = events
	d = dialer.LoggingDialer{Dialer: d, Logger: log.Log}
	d = dialer.ResolvingDialer{Connector: d, Resolver: net.DefaultResolver}
	d = dialer.LoggingDialer{Dialer: d, Logger: log.Log}
	conn, err := d.DialContext(context.Background(), "tcp", "www.facebook.com:80")
	if err != nil {
		t.Fatal(err)
	}
	for _, ev := range events.ReadEvents() {
		t.Logf("%+v", ev)
	}
	conn.Close()
}
