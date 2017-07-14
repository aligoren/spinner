# Go Spinner

Terminal Spinner for Golang

Inspired by [askn](https://github.com/askn/spinner)

# Usage

```go
import (
  spinner "./spinner"
)

func main() {
  char_name := "big_circle"
  delay := 150
  stop_time := 2

  spinner.Spinner(char_name, delay, stop_time)
}
```
