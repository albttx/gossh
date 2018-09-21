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
// pass can empty when ssh keys are in the ssh-agent
func Prompt(user, pass, host, port string) error {
	session, err := connect(user, pass, host, port)
	if err != nil {
		return err
	}
	fd, oldState, err := handleKeys(session)
	defer term.RestoreTerminal(fd, oldState)

	if err := session.Shell(); err != nil {
		return err
	}
	if err := session.Wait(); err != nil {
		return err
	}
	return nil
}

// Exec run a command over the ssh connection
// pass can empty when ssh keys are in the ssh-agent
func Exec(user, pass, host, port, command string) error {
	session, err := connect(user, pass, host, port)
	if err != nil {
		return err
	}
	fd, oldState, err := handleKeys(session)
	defer term.RestoreTerminal(fd, oldState)

	if err := session.Run(command); err != nil {
		return err
	}
	return nil
}

func connect(user, pass, host, port string) (*ssh.Session, error) {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	session, err := conn.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	return session, nil
}

func handleKeys(session *ssh.Session) (uintptr, *term.State, error) {
	var (
		oldState = &term.State{}
		err      error
	)

	fd := os.Stdin.Fd()
	if term.IsTerminal(fd) {
		oldState, err = term.MakeRaw(fd)
		if err != nil {
			return 0, oldState, err
		}
	} else {
		return 0, nil, fmt.Errorf("Error: File Descriptor isn't a terminal")
	}

	winsize, err := term.GetWinsize(fd)
	if err != nil || winsize == nil {
		winsize = &term.Winsize{
			Width: 80, Height: 24,
		}
	}
	err = session.RequestPty("xterm", int(winsize.Height), int(winsize.Width), ssh.TerminalModes{
		ssh.ECHO: 1,
	})
	if err != nil {
		return 0, oldState, err
	}
	return fd, oldState, nil
}
