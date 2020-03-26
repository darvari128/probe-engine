package resolver_test

import (
	"context"
	"testing"

	"github.com/apex/log"
	"github.com/ooni/probe-engine/internal/resolver"
)

func TestIntegration(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	var d resolver.Resolver
	d = resolver.Base()
	d = resolver.ErrWrapper{Resolver: d}
	saver := &resolver.EventsSaver{Resolver: d}
	d = saver
	d = resolver.LoggingResolver{Resolver: d, Logger: log.Log}
	addrs, err := d.LookupHost(context.Background(), "www.facebook.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(addrs)
	for _, ev := range saver.ReadEvents() {
		t.Logf("%+v", ev)
	}
}
