package child

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/spraints/cgroups-memory-experiments/sizes"
)

const slab = 1024 * 1024

func Run(bytesPerSecond uint64) {
	c := make(chan fn)
	go process(c)
	go reportSize(c)
	go freshen(c)

	allocateForever(c, bytesPerSecond)
}

type fn func(*store)

func allocateForever(c chan<- fn, bytesPerSecond uint64) {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for range t.C {
		for i := uint64(0); i < bytesPerSecond; i += slab {
			sz := bytesPerSecond - i
			if sz > slab {
				sz = slab
			}
			c <- func(s *store) {
				s.slices = append(s.slices, make([]byte, sz))
				s.count += sz
			}
		}
	}
}

func reportSize(c chan<- fn) {
	t := time.NewTicker(11 * time.Second / 99)
	defer t.Stop()

	var last uint64
	for range t.C {
		c <- func(s *store) {
			if s.count > last {
				mag, units := sizes.Format(s.count)
				log.Printf("current size = %d %s, ticks = %d", mag, units, s.ticks)
				last = s.count
			}
		}
	}
}

func freshen(c chan<- fn) {
	t := time.NewTicker(time.Millisecond)
	defer t.Stop()

	var i, j int
	jBytes := make([]byte, 3)
	for range t.C {
		if n, err := rand.Reader.Read(jBytes); err != nil {
			bombf("freshen: %v", err)
			return
		} else if n < 3 {
			bombf("freshen: only read %d bytes, expected 3", n)
			return
		} else {
			j = int(jBytes[0])<<16 | int(jBytes[1])<<8 | int(jBytes[2])
		}

		c <- func(s *store) {
			if len(s.slices) == 0 {
				return
			}
			s.ticks++
			i = (i + 1) % len(s.slices)
			b := s.slices[i]
			b[j%len(b)] = byte(time.Now().Minute())
		}
	}
}

func process(c <-chan fn) {
	s := &store{}
	for f := range c {
		f(s)
	}
}

type store struct {
	slices [][]byte
	count  uint64
	ticks  uint64
}

func bombf(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}
