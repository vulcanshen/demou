// Package ui is demou's screens: one chapter of the terminu family's
// operations on each, over one bookshelf of fake data that is the same every
// time demou starts. Nothing is saved.
package ui

import (
	_ "embed"
	"slices"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	terminu "github.com/vulcanshen/terminu-go"

	"github.com/vulcanshen/demou/internal/version"
)

// icon is demou's family mark (tdp S5), drawn by the splash: a copy of
// docs/icon.svg, which go:embed can't reach from here (a test holds the two
// the same, tdp E6).
//
//go:embed icon.svg
var icon []byte

// The title glyphs of demou's popups (the app's choice).
const (
	glyphBook   = "\uF02D"     // nf-fa-book
	glyphPencil = "\U000F03EB" // nf-md-pencil
	glyphAuthor = "\U000F0004" // nf-md-account
)

// demou is the app and the bookshelf its screens show.
type demou struct {
	a       *terminu.App
	books   []*book
	added   int // books added with the NEW BOOK form, for their IDs
	move    *moveScreen
	actions *actionsScreen
	popups  *popupsScreen
}

// New makes demou.
func New() *terminu.App { return newDemou().a }

func newDemou() *demou {
	d := &demou{a: terminu.New("demou"), books: theBooks()}
	// The least size (tdp L1): a new terminal window is often 80×24, and
	// demou's screens work there; smaller, only its name shows.
	d.a.MinWidth, d.a.MinHeight = 80, 24
	d.a.Quit = &terminu.Ask{Title: "Quit", Question: "Quit demou?", Verb: "quit"}
	d.a.Splash = splash()
	d.move = newMoveScreen(d)
	d.a.AddScreen(d.move.screen)
	d.actions = newActionsScreen(d)
	d.a.AddScreen(d.actions.screen)
	d.popups = newPopupsScreen(d)
	d.a.AddScreen(d.popups.screen)
	d.refresh()
	return d
}

// splash is the family's easter egg, V on a panel (tdp S).
func splash() terminu.Splash {
	s, err := terminu.NewSplash(icon, "demou", version.Display(), "The terminu family's operations, hands-on")
	if err != nil {
		panic(err) // the embedded icon is broken: a mistake in the program
	}
	return s
}

// refresh shows the bookshelf as it is now on every screen.
func (d *demou) refresh() {
	d.move.refresh()
	d.actions.refresh()
	d.popups.refresh()
}

func (d *demou) book(id string) *book {
	i := slices.IndexFunc(d.books, func(b *book) bool { return b.id == id })
	if i < 0 {
		return nil
	}
	return d.books[i]
}

// current is the book under a list's cursor, nil for none.
func (d *demou) current(l *terminu.List) *book {
	r, ok := l.Current()
	if !ok {
		return nil
	}
	return d.book(r.ID)
}

// moveTo, remove and add change the bookshelf, each change going in the App
// Log (tdp F5), which ! opens. An entry starts with its verb, so it starts
// with a capital however a new book's title was typed (tdp note.7).
func (d *demou) moveTo(b *book, shelf string) {
	b.shelf = shelf
	d.refresh()
	d.a.Log(terminu.Success, "Moved "+b.title+" to "+shelf)
}

func (d *demou) remove(b *book) {
	d.books = slices.DeleteFunc(d.books, func(o *book) bool { return o == b })
	d.refresh()
	d.a.Log(terminu.Success, "Took "+b.title+" off the shelf")
}

// add puts a new book on To read and returns its ID.
func (d *demou) add(title, author string) string {
	d.added++
	b := &book{id: "added-" + strconv.Itoa(d.added), title: title, author: author, shelf: toRead}
	d.books = append(d.books, b)
	d.refresh()
	d.a.Log(terminu.Success, "Put "+title+" by "+author+" on "+toRead)
	return b.id
}

// openBook is Enter on a book, the same on every screen (tdp K3): its first
// page, in a viewer with the lib's selection mode (tdp F1). A book added in
// demou has none.
func (d *demou) openBook(l *terminu.List) terminu.Action {
	return terminu.Action{Key: "Enter", Label: "Open", Desc: "its first page", Scope: terminu.ItemScope, Hint: "open",
		Disabled: func() bool { b := d.current(l); return b == nil || b.firstPage == "" },
		Do: func() tea.Cmd {
			b := d.current(l)
			return d.a.Viewer(&terminu.Viewer{Title: "First page", Subject: b.title, Glyph: glyphBook,
				Lines: []string{b.firstPage}})
		}}
}
