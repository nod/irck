package irck

import (
  "testing"
  "time"
)


func TestParseISO(t *testing.T) {
  want := time.Date(2011, 11, 21, 9, 45, 51, 0, time.UTC)
  got := ParseISOTime("2011-11-21T09:45:51Z")
  if got != want {
    t.Errorf("ParseISOTime() = %q, want %q", got, want)
  }
}

