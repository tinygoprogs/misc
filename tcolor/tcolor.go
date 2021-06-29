/*
Tiny terminal color-book.

Colors get by default enabled by init().
This does not happed if either the environment variable 'NOCOLOR' is set or the
stdout-device is not a terminal.
*/
package tcolor

import (
  "os"
  "syscall"
)

const (
  pipe = 12
  term = 20
)

var (
  White = ""
  Black = ""
  Red   = ""
  Green = ""
  Yellow= ""
  Blue  = ""
  Bold  = ""
)

func init () {
  if _, exist := os.LookupEnv("NOCOLOR"); exist { return }
  st, _ := os.Stdout.Stat()
  devinfo := st.Sys().(*syscall.Stat_t)
	//println("stdout.stat devinfo.Dev= ", devinfo.Dev)
  if devinfo != nil && devinfo.Dev != term { return }
  Enable()
}

func Enable() {
  White = "\x1b[0;6m"
  Black = "\x1b[0;30m"
  Red   = "\x1b[0;31m"
  Green = "\x1b[0;32m"
  Yellow= "\x1b[0;33m"
  Blue  = "\x1b[0;34m"
  Bold  = "\x1b[0;1m"
}

func Disable() {
  White = ""
  Black = ""
  Red   = ""
  Green = ""
  Yellow= ""
  Blue  = ""
  Bold  = ""
}

