package main

import (
  spinner "./spinner"
)

func main() {
  /*
    params:
      character name
      delay
      stop time (optional), if stop time is blank, spinner won't be stop
  */
  spinner.Spinner("big_circle", 150)
}
