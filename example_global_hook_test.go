package glog_test

import (
	"github.com/Zhenhanyijiu/glog"
	"os"
)

var (
	mystring string
)

type GlobalHook struct {
}

func (h *GlobalHook) Levels() []glog.Level {
	return glog.AllLevels
}

func (h *GlobalHook) Fire(e *glog.Entry) error {
	e.Data["mystring"] = mystring
	return nil
}

func ExampleGlobalVariableHook() {
	l := glog.New()
	l.Out = os.Stdout
	l.Formatter = &glog.TextFormatter{DisableTimestamp: true, DisableColors: true}
	l.AddHook(&GlobalHook{})
	mystring = "first value"
	l.Info("first log")
	mystring = "another value"
	l.Info("second log")
	// Output:
	// level=info msg="first log" mystring="first value"
	// level=info msg="second log" mystring="another value"
}
