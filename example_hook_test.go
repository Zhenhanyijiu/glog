// +build !windows

package glog_test

import (
	"log/syslog"
	"os"

	slhooks "github.com/Zhenhanyijiu/glog/hooks/syslog"
	"github.com/sirupsen/glog"
)

// An example on how to use a hook
func Example_hook() {
	var log = glog.New()
	log.Formatter = new(glog.TextFormatter)                     // default
	log.Formatter.(*glog.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*glog.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	if sl, err := slhooks.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, ""); err == nil {
		log.Hooks.Add(sl)
	}
	log.Out = os.Stdout

	log.WithFields(glog.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(glog.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(glog.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")

	// Output:
	// level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
	// level=warning msg="The group's number increased tremendously!" number=122 omg=true
	// level=error msg="The ice breaks!" number=100 omg=true
}
