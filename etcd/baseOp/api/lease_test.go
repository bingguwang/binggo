package api

import (
	"testing"
	"time"
)

func TestGrant(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	// minimum lease TTL is 5-second
	leaseID, err := cli.Grant(50)
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.Put("example_key", "example_value", leaseID, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRevoke(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	// minimum lease TTL is 5-second
	leaseID, err := cli.Grant(5)
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.Put("example_key", "example_value", leaseID, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = cli.Revoke(leaseID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeepAlive(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	leaseID, err := cli.Grant(5)
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.Put("example_key", "example_value", leaseID, 0)
	if err != nil {
		t.Fatal(err)
	}

	// the key will be kept forever
	ch, err := cli.KeepAlive(leaseID)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 6)

	ka := <-ch
	t.Log("keep alive ttl:", ka.TTL)
	// Output: ttl: 5
}

func TestKeepAliveOnce(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	leaseID, err := cli.Grant(5)
	if err != nil {
		t.Fatal(err)
	}

	_, err = cli.Put("example_key", "example_value", leaseID, 0)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 3)

	// to renew the lease only once
	ka, err := cli.KeepAliveOnce(leaseID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("keep alive once ttl:", ka.TTL)
	// Output: ttl: 4
}
