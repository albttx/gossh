package gossh

import (
	"fmt"
	"net"
	"os"

	"github.com/albttx/gossh/term"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// Prompt start a ssh connection in your terminal
// pass can empty when ssh keys
func Prompt(user, pass, host, port string) error {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return err
	}
	sshKeys := ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)

	conn, err := ssh.Dial("tcp", net.JoinHostPort(host, port), &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			sshKeys,
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return err
	}

	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("Failed to create session: %s", err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO: 1,
	}
	fd := os.Stdin.Fd()
	winsize := &term.Winsize{}
	if term.IsTerminal(fd) {
		oldState, err := term.MakeRaw(fd)
		if err != nil {
			return err
		}
		defer term.RestoreTerminal(fd, oldState)
		winsize, err = term.GetWinsize(fd)
		if err != nil || winsize == nil {
			winsize = &term.Winsize{
				Width: 80, Height: 24,
			}
		}
	} else {
		return fmt.Errorf("Error: File Descriptor isn't a terminal")
	}

	if err := session.RequestPty("xterm", int(winsize.Height), int(winsize.Width), modes); err != nil {
		return err
	}

	if err := session.Shell(); err != nil {
		return err
	}
	if err := session.Wait(); err != nil {
		return err
	}
	return nil
}
