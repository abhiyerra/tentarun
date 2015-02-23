// Based off of:
// http://kukuruku.co/hub/golang/ssh-commands-execution-on-hundreds-of-servers-via-go
// http://golang-basic.blogspot.com/2014/06/step-by-step-guide-to-ssh-using-go.html

// Connect to a few machines and tail a file
// multitail "cmdtorun /location/of/file" machine1 machine2 machine3...

package main

import (
	"bytes"
	"fmt"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"os/user"
)

var (
	config *ssh.ClientConfig
)

func executeCmd(cmd, hostname string) string {
	conn, err1 := ssh.Dial("tcp", hostname+":22", config)
	if err1 != nil {
		log.Fatal(err1)
	}

	session, err2 := conn.NewSession()
	if err2 != nil {
		log.Fatal(err2)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(cmd)

	return hostname + ": " + stdoutBuf.String()
}

func main() {
	user, _ := user.Current()

	fmt.Printf("%s Password: ", user.Username)
	pass := gopass.GetPasswd()

	config = &ssh.ClientConfig{
		User: user.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(string(pass)),
		},
	}

	cmd := os.Args[1]
	hosts := os.Args[2:]

	results := make(chan string, 10)
	for _, hostname := range hosts {
		go func(h string) {
			results <- executeCmd(cmd, h)
		}(hostname)
	}

	for i := 0; i < len(hosts); i++ {
		select {
		case res := <-results:
			fmt.Print(res)
		}
	}
}
