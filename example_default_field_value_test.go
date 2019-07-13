package glog_test

import (
	"github.com/Zhenhanyijiu/glog"
	"os"
)

type DefaultFieldHook struct {
	GetValue func() string
}

func (h *DefaultFieldHook) Levels() []glog.Level {
	return glog.AllLevels
}

func (h *DefaultFieldHook) Fire(e *glog.Entry) error {
	e.Data["aDefaultField"] = h.GetValue()
	return nil
}

func ExampleDefaultField() {
	l := glog.New()
	l.Out = os.Stdout
	l.Formatter = &glog.TextFormatter{DisableTimestamp: true, DisableColors: true}

	l.AddHook(&DefaultFieldHook{GetValue: func() string { return "with its default value" }})
	l.Info("first log")
	// Output:
	// level=info msg="first log" aDefaultField="with its default value"
}
