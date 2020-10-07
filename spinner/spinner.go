package spinner

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Spinner struct {
	active 			bool
	mu 	   			*sync.Mutex
	loader 			[]string
	i 				int
	speed			time.Duration
	text			string
	oldText 		string
	time 			time.Time
}

var (
	SimpleLoading = []string{ "|", "/", "-", "\\" }
	BarLoading = []string{"[        ]", "[=       ]", "[===     ]", "[====    ]", "[=====   ]", "[======  ]", "[======= ]", "[========]", "[ =======]", "[  ======]", "[   =====]", "[    ====]", "[     ===]", "[      ==]", "[       =]", "[        ]", "[        ]"}
	SmallBarLoading = []string{"[=     ]", "[ =    ]", "[  =   ]", "[   =  ]", "[    = ]", "[     =]", "[    = ]", "[   =  ]", "[  =   ]", "[ =    ]"}
)

func Init(loader []string, speed time.Duration) *Spinner {
	return &Spinner{
		active: true,
		mu: &sync.Mutex{},
		loader: loader,
		speed: speed,
		time: time.Now(),
	}
}

func (s *Spinner) Start(msg string, args... interface{}) {
	s.Write(msg, args...)
	go func() {
		for {
			s.mu.Lock()
			s.i++
			if !s.active {
				s.mu.Unlock()
				break
			}
			fmt.Printf("\u001B[?25l\r%s ", strings.Repeat(" ", len(s.oldText)+len(s.loader[0])))
			fmt.Printf("\u001B[?25l\r%s %s", s.loader[s.i % len(s.loader)], s.text)
			time.Sleep(s.speed)
			s.mu.Unlock()
		}
	}()
}

func (s *Spinner) Write(msg string, args... interface{}) {
	s.mu.Lock()
	s.oldText = s.text
	s.text = fmt.Sprintf(msg, args...)
	s.mu.Unlock()
}

func (s *Spinner) Stop(error bool, msg string, args... interface{}) {
	var symbol string
	if error {
		symbol = "✗"
	} else {
		symbol = "✓"
	}
	s.Write(msg, args...)
	s.active = false
	s.mu.Lock()
	fmt.Print("\033[?25h")
	fmt.Printf("\r%s ", strings.Repeat(" ", len(s.oldText)+len(s.loader[0])))
	fmt.Printf("\r%s %s\n", symbol, s.text)
	s.mu.Unlock()
}

func (s *Spinner) TimeElapsed() time.Duration {
	return time.Since(s.time)
}