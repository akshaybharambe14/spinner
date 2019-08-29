package spinner

import (
	"fmt"
	"time"
)

// Spinner -
type Spinner struct {
	variant  Variant
	speed    int
	duration time.Duration
}

// Variant -
type Variant int

const (
	// Default - {`-`, `\`, `|`, `/`}
	Default Variant = iota
)

const (
	// MaxSpeed -
	MaxSpeed = 5
	// NoDuration -
	NoDuration = -1
)

var variants = map[Variant][]string{
	Default: []string{`-`, `\`, `|`, `/`},
}

// New -
func New(variant Variant, speed int, duration time.Duration) *Spinner {
	s := &Spinner{
		variant:  variant,
		speed:    speed,
		duration: duration,
	}
	s.configure()
	return s
}

func (s *Spinner) configure() {
	if s.speed < 0 {
		s.speed = 1
	} else if s.speed > MaxSpeed {
		s.speed = MaxSpeed
	}
	s.speed = (MaxSpeed - s.speed) + 1

	if _, ok := variants[s.variant]; !ok {
		s.variant = Default
	}
}

// Start -
func (s *Spinner) Start(quit chan struct{}) {
	speed := time.Tick(time.Millisecond * time.Duration(s.speed*100))
	duration := time.Tick(s.duration)
	end := len(variants[s.variant]) - 1
	pf := printFrame(end, s.variant)
	for {
		select {
		case <-speed:
			pf()
		case <-quit:
			goto Done
		case <-duration:
			goto Done
		}
	}

Done:
	// exit
}

func printFrame(end int, variant Variant) func() {
	i := 0
	return func() {
		fmt.Printf("\r%s", variants[variant][i])
		if i == end {
			i = 0
		} else {
			i++
		}
	}
}
