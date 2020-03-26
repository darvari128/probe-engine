package httptransport_test

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"testing"

	"github.com/apex/log"
	"github.com/ooni/probe-engine/internal/httptransport"
)

type simpleTLSDialer struct{}

func (d simpleTLSDialer) DialTLSContext(
	ctx context.Context, network, address string) (net.Conn, error) {
	// just ignore the context
	return tls.Dial(network, address, nil)
}

func TestIntegration(t *testing.T) {
	var txp httptransport.Transport
	txp = httptransport.NewBase(new(net.Dialer), new(simpleTLSDialer))
	events := &httptransport.EventsSaver{Transport: txp}
	txp = events
	txp = httptransport.ErrWrapper{Transport: txp}
	txp = httptransport.HeaderAdder{Transport: txp}
	saver := &httptransport.SnapshotSaver{Transport: txp}
	txp = saver
	txp = httptransport.Logging{Transport: txp, Logger: log.Log}
	client := &http.Client{Transport: txp}
	resp, err := client.Get("http://facebook.com")
	if err != nil {
		t.Fatal(err)
	}
	for _, ev := range events.ReadEvents() {
		t.Logf("%+v", ev)
	}
	for _, snap := range saver.Snapshots() {
		t.Logf("%+v", snap)
	}
	resp.Body.Close()
}
