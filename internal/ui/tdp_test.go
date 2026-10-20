package ui

import (
	"testing"

	"github.com/vulcanshen/terminu-go/tdptest"
)

// TestTDP launches demou as for the first time on many terminal sizes and
// icon widths, and walks what every family app has on each screen and panel:
// the Space menu, the global operation popup, the App Log and the key
// reference; then a terminal under the least size and back (tdptest.Check).
func TestTDP(t *testing.T) {
	tdptest.Check(t, tdptest.App{Name: "demou", New: New})
}
