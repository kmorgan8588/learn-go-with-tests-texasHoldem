package server_test

import (
	"bytes"
	"fmt"
	"go-app/server"
	"strings"
	"testing"
	"time"
)

type SpyBlindAlerter struct {
	alerts []ScheduledAlert
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlert{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &server.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartedWith  int
	FinishedWith string

	StartCalled  bool
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("game starts with 4 players and record Kyle win from user input", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("4\nKyle wins\n")
		game := &GameSpy{}

		cli := server.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, server.PlayerPrompt)
		assertGameStartedWith(t, game, 4)
		assertGameFinishedWith(t, game, "Kyle")
	})

	t.Run("game starts with 8 players and record Lyle win from user input", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("8\nLyle wins\n")
		game := &GameSpy{}

		cli := server.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertGameFinishedWith(t, game, "Lyle")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies")
		game := &GameSpy{}

		cli := server.NewCLI(in, stdout, game)
		cli.PlayPoker()

		asssertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, server.PlayerPrompt, server.BadPlayerInputErrMsg)
	})

	t.Run("it prints a message when an invalid win statement is entered", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("4\nLyle styles")
		game := &GameSpy{}

		cli := server.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, server.PlayerPrompt, server.BadPlayerWinnerInputErrMsg)
		assertGameNotFinished(t, game)
	})
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.FinishCalled {
		t.Errorf("game should not have ended")
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, playerCount int) {
	t.Helper()

	if game.StartedWith != playerCount {
		t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
	}
}

func assertGameFinishedWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	if game.FinishedWith != winner {
		t.Errorf("wanted Finished to be Kyle but got %q", game.FinishedWith)
	}
}

func AssertScheduledAlert(t *testing.T, got ScheduledAlert, want ScheduledAlert) {

	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func asssertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}
