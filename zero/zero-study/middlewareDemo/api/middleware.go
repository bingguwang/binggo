package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

func middlewareDemoFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("请求处理前 ... ")
		next(w, r)
		logx.Info("请求处理后 ... ")
	}
}
