package server_test

import (
	"go-app/server"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Carl wins\n")
		playerStore := &server.StubPlayerStore{}

		cli := server.NewCLI(playerStore, in)
		cli.PlayPoker()

		server.AssertPlayerWin(t, playerStore, "Carl")
	})

	t.Run("record Kyle win from user input", func(t *testing.T) {
		in := strings.NewReader("Kyle wins\n")
		playerStore := &server.StubPlayerStore{}

		cli := server.NewCLI(playerStore, in)
		cli.PlayPoker()

		server.AssertPlayerWin(t, playerStore, "Kyle")
	})
}
