package api

import (
	"fmt"
	"testing"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

func TestClientInit(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()
}

func TestPut(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	_, err = cli.Put("example_key", "example_value", 0, requestTimeout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	_, err = cli.Put("example_key", "example_value", 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	res, err := cli.Get("example_key", requestTimeout)
	if err != nil || res != "example_value" {
		t.Fatal(err)
	}
}

func TestGetWithRevision(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	presp, err := cli.Put("example_key", "example_value", 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	res, err := cli.GetWithRevision("example_key", presp.Header.Revision, requestTimeout)
	if err != nil || res != "example_value" {
		t.Fatal(err)
	}
}

func TestGetSortedPrefix(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	for i := 0; i < 3; i++ {
		_, err = cli.Put(fmt.Sprintf("example_key_%d", i), fmt.Sprintf("example_value_%d", i), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	resp, err := cli.GetSortedPrefix("example_key", requestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		if resp[fmt.Sprintf("example_key_%d", i)] != fmt.Sprintf("example_value_%d", i) {
			t.Fatal("TestGetSortedPrefix Error")
		}
	}
}

func TestDelete(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	_, err = cli.Put("example_key", "example_value", 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	res, err := cli.Delete("example_key", requestTimeout)
	if err != nil || res != "example_value" {
		t.Fatal(err)
	}
}

func TestDeletePrefix(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	for i := 0; i < 3; i++ {
		_, err = cli.Put(fmt.Sprintf("example_key_%d", i), fmt.Sprintf("example_value_%d", i), 0, 0)
		if err != nil {
			t.Fatal(err)
		}
	}

	resp, err := cli.DeletePrefix("example_key", requestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		if resp[fmt.Sprintf("example_key_%d", i)] != fmt.Sprintf("example_value_%d", i) {
			t.Fatal("TestGetSortedPrefix Error")
		}
	}
}

func TestEtcdClienter_Txn(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	kv := map[string]string{
		"example_key_0": "example_value_0",
		"example_key_1": "example_value_1",
		"example_key_2": "example_value_2",
	}

	err = cli.Txn(kv, 0, requestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	defer cli.DeletePrefix("example_key", requestTimeout)

	resp, err := cli.GetSortedPrefix("example_key", requestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		if resp[fmt.Sprintf("example_key_%d", i)] != fmt.Sprintf("example_value_%d", i) {
			t.Fatal("TestGetSortedPrefix Error")
		}
	}
}

func TestEtcdClienter_Txn1(t *testing.T) {
	cli, err := ClientInit(dialTimeout)
	if err != nil {
		t.Fatal(err)
	}

	defer cli.Close()

	kv := map[string]string{
		"example_key_0": "example_value_0",
		"example_key_1": "example_value_1",
		"example_key_2": "example_value_2",
	}

	leaseID, _ := cli.Grant(3)
	err = cli.Txn(kv, leaseID, requestTimeout)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.GetSortedPrefix("example_key", requestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		if resp[fmt.Sprintf("example_key_%d", i)] != fmt.Sprintf("example_value_%d", i) {
			t.Fatal("TestGetSortedPrefix Error")
		}
	}
	time.Sleep(time.Second * 5)
}
