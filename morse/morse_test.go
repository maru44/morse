package morse_test

import (
	"testing"
	"time"

	"github.com/maru44/morse/morse"
)

type send struct {
	durationAfter time.Duration // millisecond
	ping          string
}

func TestMorse(t *testing.T) {
	m := morse.NewMorse(morse.IntervalDuration(100))
	tests := []struct {
		name        string
		sends       []send
		want        string
		wantDecoded string
	}{
		{
			name: "OK",
			sends: []send{
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
				{
					durationAfter: 50,
					ping:          m.DitPing,
				},
				{
					durationAfter: 50,
					ping:          m.DitPing,
				},
				{
					durationAfter: 150,
					ping:          m.DahPing,
				},
				{
					durationAfter: 0,
					ping:          m.DahPing,
				},
				{
					durationAfter: 0,
					ping:          m.DahPing,
				},
				{
					durationAfter: 350,
					ping:          m.DitPing,
				},
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
			},
			want:        "... ---   ...",
			wantDecoded: "SOS",
		},
		{
			name: "OK with spaces",
			sends: []send{
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
				{
					durationAfter: 50,
					ping:          m.DitPing,
				},
				{
					durationAfter: 50,
					ping:          m.DitPing,
				},
				{
					durationAfter: 750,
					ping:          m.DahPing,
				},
				{
					durationAfter: 0,
					ping:          m.DahPing,
				},
				{
					durationAfter: 0,
					ping:          m.DahPing,
				},
				{
					durationAfter: 650,
					ping:          m.DitPing,
				},
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
			},
			want:        "...       ---      ...",
			wantDecoded: "S OS",
		},
		{
			name: "OK with spaces and not ping",
			sends: []send{
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
				{
					durationAfter: 50,
					ping:          m.DitPing,
				},
				{
					durationAfter: 50,
					ping:          m.DitPing,
				},
				{
					durationAfter: 750,
					ping:          m.DahPing,
				},
				{
					durationAfter: 0,
					ping:          m.DahPing,
				},
				{
					durationAfter: 0,
					ping:          m.DahPing,
				},
				{
					durationAfter: 680,
					ping:          "u",
				},
				{
					durationAfter: 50,
					ping:          m.DitPing,
				},
				{
					durationAfter: 0,
					ping:          m.DitPing,
				},
			},
			want:        "...       ---      ..",
			wantDecoded: "S OI",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan string)
			got := ""
			m.SetSend(func(m *morse.Morse, ch chan string) {
				for _, s := range tt.sends {
					time.Sleep(s.durationAfter * time.Millisecond)
					ch <- s.ping
				}
				time.Sleep(50 * time.Millisecond)
				ch <- m.QuitPing
			})
			m.SetRecieve(func(m *morse.Morse, ch chan string, ret *string) {
				morse.BaseReceive(m, ch, &got, false)
			})
			go m.Recieve(ch, &got)
			m.Send(ch)

			if got != tt.want {
				t.Fatalf("expected is `%s`, but got `%s`", tt.want, got)
			}
			gotDecodedByte := m.ConvertCode(got)
			if string(gotDecodedByte) != tt.wantDecoded {
				t.Fatalf("expected decoded is `%s`, but got `%s`", tt.wantDecoded, string(gotDecodedByte))
			}
		})
	}
}
