package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os/user"
	"strings"
)

var (
	username  string
	password  string
	hostnames string
	keyfile   string
	envstr    string
	envs      []string
	hosts     []string
	verbose   bool
	config    *ssh.ClientConfig
)

func getKeyAuth() (key ssh.Signer) {
	buf, err := ioutil.ReadFile(keyfile)
	if err != nil {
		log.Fatal(err)
	}

	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func runOnHosts(cmd string) {
	results := make(chan string, len(hosts))

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

	for _, e := range envs {
		envKeyVal := strings.Split(e, "=")
		session.Setenv(envKeyVal[0], envKeyVal[1])
	}

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(cmd)

	verboseStr := ""
	if verbose {
		verboseStr = hostname + ":\n"
	}
	return verboseStr + stdoutBuf.String()
}

func init() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&username, "u", user.Username, "The username of the machines")
	flag.StringVar(&password, "p", "", "The password of the machines")
	flag.StringVar(&hostnames, "h", "", "The hosts separated by a comma. Ex. host1,host2,host3")
	flag.StringVar(&keyfile, "k", "", "The public key to connect to the servers with")
	flag.StringVar(&envstr, "e", "", "Environment variables separate by space. Ex. FOO=bar BAR=foo")
	flag.BoolVar(&verbose, "v", false, "Show the server name in the output.")
}

func main() {
	flag.Parse()

	hosts = strings.Split(hostnames, ",")
	if envstr != "" {
		envs = strings.Split(envstr, " ")
	}

	switch {
	case password != "":
		config = &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
		}
	case keyfile != "":
		config = &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(getKeyAuth()),
			},
		}
	}

	runOnHosts(flag.Arg(0))
}
