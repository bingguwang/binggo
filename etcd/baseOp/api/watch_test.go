package api

import (
	"log"
	"testing"
)

func testWatch(cli *EtcdClienter) {
	_, err := cli.Put("example_key", "example_value", 0, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func TestWatch(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	rch := cli.Watch("example_key")
	go testWatch(cli)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			t.Logf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
		break
	}
}

func testWatchWithPrefix(cli *EtcdClienter) {
	_, err := cli.Put("example_key_prefix", "example_value", 0, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func TestWatchWithPrefix(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	rch := cli.WatchWithPrefix("example_key")
	go testWatchWithPrefix(cli)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			t.Logf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
		break
	}
}

// TODO:没太理解这个功能具体用法
func TestWatchWithProgressNotify(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	rch := cli.WatchWithProgressNotify("example_key")
	go testWatch(cli)
	resp := <-rch
	t.Log("resp.IsProgressNotify:", resp.IsProgressNotify())
}
