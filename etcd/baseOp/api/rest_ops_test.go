package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	mvccpb "go.etcd.io/etcd/api/v3/mvccpb"
	"io/ioutil"
	"net/http"
	"testing"
)

/*
*
curl http://localhost:2379/v3/kv/range -X POST -d '{"key": "key_name"}'
curl http://localhost:2379/v3/kv/put -X POST -d '{"key": "key_name", "value": "key_value"}'
curl http://localhost:2379/v3/kv/deleterange -X POST -d '{"key": "key_name"}'
curl http://localhost:2379/v3/kv/put -X POST -d '{"key": "dir_name/", "value": ""}'
*/
func TestRestPut(t *testing.T) {

	// 准备要发送的数据
	data := `{"key": "example_key","value":"example_value"}`
	base64Data := base64.StdEncoding.EncodeToString([]byte(data))

	// 准备请求
	url := "http://192.168.2.44:2379/v3/kv/put"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(base64Data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 输出响应
	fmt.Println("Response:", string(respBody))
}

const prefix = `http://192.168.2.44:2379/v3/kv/`

func TestRestGetKey(t *testing.T) {
	// 准备要发送的数据
	key := `example_key`
	base64Key := base64.StdEncoding.EncodeToString([]byte(key))
	data := fmt.Sprintf(`{"key": "%s"}`, base64Key)

	// 准备请求
	url := prefix + "range"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 输出响应
	response := etcdserverpb.RangeResponse{}
	json.Unmarshal(respBody, &response)
	prinfKvs(response.Kvs)
}

func prinfKvs(kvs []*mvccpb.KeyValue) {
	for _, kv := range kvs {
		fmt.Println(tojsonString(kv))
	}
}

func tojsonString(i interface{}) string {
	marshal, _ := json.Marshal(i)
	return string(marshal)
}
