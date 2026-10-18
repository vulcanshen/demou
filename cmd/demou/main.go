// demou is the terminu family's app for learning the operations.
package main

import (
	"fmt"
	"os"

	terminu "github.com/vulcanshen/terminu-go"

	"github.com/vulcanshen/demou/internal/version"
)

const usage = `demou: learn the terminu family's operations hands-on

Usage:
  demou            start
  demou version    print the version (-v, --version)
  demou help       print this help (-h, --help)
`

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version", "-v", "--version":
			fmt.Println(version.Display())
			return
		case "help", "-h", "--help":
			fmt.Print(usage)
			return
		default:
			fmt.Fprintf(os.Stderr, "demou: unknown command %q\n\n%s", os.Args[1], usage)
			os.Exit(2)
		}
	}
	app := terminu.New("demou")
	// A placeholder so the app runs; the real screens replace it.
	start := &terminu.List{Columns: []terminu.Column{{Title: "Start", Flex: true}}, EmptyFact: "Nothing here yet"}
	app.AddScreen(&terminu.Screen{Name: "Start", Panels: []*terminu.Panel{{Label: "Start", Body: start}}})
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "demou:", err)
		os.Exit(1)
	}
}
