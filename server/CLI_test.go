package server_test

import (
	"fmt"
	"go-app/server"
	"strings"
	"testing"
	"time"
)

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

type expectedAlert struct {
	expectedScheduledTime time.Duration
	expectedAmount        int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Carl wins\n")
		playerStore := &server.StubPlayerStore{}

		cli := server.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		server.AssertPlayerWin(t, playerStore, "Carl")
	})

	t.Run("record Kyle win from user input", func(t *testing.T) {
		in := strings.NewReader("Kyle wins\n")
		playerStore := &server.StubPlayerStore{}

		cli := server.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		server.AssertPlayerWin(t, playerStore, "Kyle")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &server.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := server.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) < 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &server.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := server.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []expectedAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduledTime), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled for %v", i, blindAlerter.alerts)
				}
				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, c)
			})
		}
	})
}

func assertScheduledAlert(t *testing.T, got scheduledAlert, want expectedAlert) {

	if got.amount != want.expectedAmount {
		t.Errorf("got amount %d, want %d", got.amount, want.expectedAmount)
	}

	if got.at != want.expectedScheduledTime {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.expectedScheduledTime)
	}
}
