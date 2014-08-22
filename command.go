package main

import (
	"github.com/aurlaw/aurlog"
	"os/exec"
	"strings"
	"sync"
)

type Command struct {
	Log *aurlog.AurLog
}

func (c *Command) ExeCmd(cmd string, wg *sync.WaitGroup) {
	c.Log.Infoln("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		c.Log.Infof("%s", err)
	}
	c.Log.Infof("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done

}
