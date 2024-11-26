package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600     //会话的生命长度
	sessionSecret := "pwd123" //密钥
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))

	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}

func Captcha(c *gin.Context, length ...int) { //生成验证码
	l := captcha.DefaultLen
	w, h := 107, 36 //验证码的大小
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)    //生成验证码的id,生成的验证码长度是l
	session := sessions.Default(c)    //获取session
	session.Set("captcha", captchaId) //验证码id设置到session中
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func CaptchaVerify(c *gin.Context, code string) bool { //校验传入的验证码code是否正确
	session := sessions.Default(c)                             //获取上下文的session
	if captchaId := session.Get("captcha"); captchaId != nil { //从session中取出之前生成验证码的时候设置到session中的验证码id
		fmt.Println(captchaId)
		fmt.Println(code)
		session.Delete("captcha") //取到了验证码id就删除
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) { //校验验证码
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./*.html")
	router.Use(Session("my_session"))
	router.GET("/captcha", func(c *gin.Context) {
		Captcha(c, 4)
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/captcha/verify/:value", func(c *gin.Context) {
		value := c.Param("value")
		if CaptchaVerify(c, value) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})
	router.Run(":8088")
}
