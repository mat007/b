package building

import (
	"os"
	"os/exec"
)

type command struct {
	name    string
	success bool
}

func (b *B) Command(name string, args ...string) command {
	c := command{
		name: name,
	}
	if len(args) > 0 {
		c.Run(args...)
	}
	return c
}

func (c command) WithSuccess() command {
	c.success = true
	return c
}

func (c command) Run(args ...string) int {
	Debugln("running", append([]string{c.name}, args...))
	cmd := exec.Command(c.name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	code, err := run(cmd, c.success)
	if err != nil {
		Fatalf("error running %s: %s", c.name, err)
	}
	return code
}
