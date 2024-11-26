package case1

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"testing"
	"time"
)

func TestRandomCase1(t *testing.T) {

}

func (r *request) SetRandom() *request {
	// 这里必须用纳秒，用秒的话如果一秒内发起两次请求会报错
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var (
		md5Handler = md5.New()
		rd         = fmt.Sprint(random.Uint64())
		sec        = r.secret + rd
	)
	io.WriteString(md5Handler, sec)
	r.random = rd
	r.randomMD5 = fmt.Sprintf("%x", md5Handler.Sum(nil))
	return r
}

type request struct {
	target    string
	method    string
	path      string
	secret    string
	random    string
	randomMD5 string
	query     url.Values
	data      any
	langcn    bool
	flatten   bool
}
