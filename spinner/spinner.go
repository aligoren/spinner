package spinner

import (
  "fmt"
  "time"
  "os"
  "os/signal"
  "syscall"
  cs "../charsets"
)

// If user want to stop this program
func SignalHandle() {
  c := make(chan os.Signal, 2)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)
  go func() {
      <-c
      Stop()
  }()
}

func Spinner(key string, iter_time time.Duration, stop_time ...time.Duration) {
  SignalHandle()
  charsets := cs.GetCharset()
  i := 0
  for {
    fmt.Printf("\033[?25l \r %s ", charsets[key][i])
    time.Sleep(iter_time * time.Millisecond)
    i++
    if i == len(charsets[key]) {
      i = 0

      if len(stop_time) > 0 {
        stime := time.Duration(0)
        stime = stop_time[0]
        time.Sleep(time.Second * stime)
        Stop()
      }

    }
  }
}

func Stop() {
  fmt.Printf("\033[?25h\r")
  os.Exit(0)
  fmt.Printf("\033[?25l\r")
}
