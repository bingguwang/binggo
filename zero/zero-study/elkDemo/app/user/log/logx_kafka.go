package log

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

// KafkaWriter 实现了Writer接口
type KafkaWriter struct {
	Pusher *kq.Pusher
}

func NewKafkaWriter(pusher *kq.Pusher) *KafkaWriter {
	return &KafkaWriter{
		Pusher: pusher,
	}
}

func (w *KafkaWriter) Write(p []byte) (n int, err error) {
	// writing logDemo with newlines, trim them.
	if err := w.Pusher.Push(strings.TrimSpace(string(p))); err != nil {
		return 0, err
	}

	return len(p), nil
}

func LogxKafka() *logx.Writer {
	/**
		为什么filebeat可以直接读取日志文件了，这里还要手动推送日志给kafka？
	这样灵活性好，可以细粒度的控制，而且实时性高
	*/
	pusher := kq.NewPusher([]string{"192.168.2.130:9094"}, "looklook-logDemo")
	defer pusher.Close()

	writer := logx.NewWriter(NewKafkaWriter(pusher))
	return &writer
}
