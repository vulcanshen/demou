package ui

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	terminu "github.com/vulcanshen/terminu-go"
)

// popupsScreen is [P]OPUPS, the third chapter: popups over popups. Remove
// asks first; New book opens a small form whose Author is a select over it.
type popupsScreen struct {
	d      *demou
	screen *terminu.Screen
	list   *terminu.List
}

func newPopupsScreen(d *demou) *popupsScreen {
	p := &popupsScreen{d: d}
	p.list = &terminu.List{
		Columns:   []terminu.Column{{Title: "Title"}, {Title: "Author", Drop: 2}, {Title: "Shelf", Drop: 1}},
		EmptyFact: "No books on the shelf", EmptyHint: "Press [n] to add a book",
	}
	books := &terminu.Panel{Label: "Books", Body: p.list, Actions: []terminu.Action{
		d.openBook(p.list),
		// The question names the book under the cursor. Run from the Space
		// menu, Esc on it goes back to the menu; once the book is gone the
		// menu closes, there being nothing left for it (tdp T1).
		{Key: "r", Label: "remove", Desc: "take it off the shelf", Scope: terminu.ItemScope, Hint: "remove",
			AskFor: func() *terminu.Ask {
				return &terminu.Ask{Title: "Remove", Verb: "remove",
					Question: "Remove " + d.current(p.list).title + " from the shelf?"}
			},
			Do: func() tea.Cmd { d.remove(d.current(p.list)); return nil }},
		{Key: "n", Label: "new book", Desc: "put another book on the shelf", Scope: terminu.PanelScope, Hint: "new", Do: p.newBook},
	}}
	p.screen = &terminu.Screen{Name: "Popups", Key: "P", Desc: "popups over popups", Panels: []*terminu.Panel{books}}
	return p
}

func (p *popupsScreen) refresh() {
	rows := make([]terminu.Row, len(p.d.books))
	for i, b := range p.d.books {
		rows[i] = terminu.Row{ID: b.id, Cells: []string{b.title, b.author, b.shelf}}
	}
	p.list.SetRows(rows)
}

// newBook is the NEW BOOK form: a title typed in a text popup, an author
// chosen in a select. Add puts the book on To read, the cursor on it (tdp
// K3). Run from the Space menu, the form opens over the menu and Esc goes
// back to it (tdp T1); Add closes the menu with the form, the menu being
// about the book that was under the cursor before.
func (p *popupsScreen) newBook() tea.Cmd {
	title, author := "", -1
	names := make([]string, len(authors))
	for i, a := range authors {
		names[i] = a.name
	}
	return p.d.a.Form(&terminu.Form{Title: "New book", Name: "New book", Glyph: glyphBook, Button: "Add",
		Fields: []*terminu.FormField{
			{Label: "Title", Value: func() string { return title }, Empty: "not set",
				Edit: func() terminu.Input {
					return terminu.TextInput{Title: "Title", Glyph: glyphPencil, Prompt: "The title of the book", Old: title,
						// Enter on an empty title says so rather than close (tdp K3).
						Check: func(s string) error {
							if s == "" {
								return errors.New("Title is required")
							}
							return nil
						},
						Done: func(s string) tea.Cmd { title = s; return nil }}
				},
				Check: func() string {
					if title == "" {
						return "Title is required"
					}
					return ""
				}},
			{Label: "Author", Value: func() string {
				if author < 0 {
					return ""
				}
				return names[author]
			}, Empty: "not set",
				Edit: func() terminu.Input {
					return terminu.Select{Title: "Author", Glyph: glyphAuthor, Options: names, Selected: author,
						Done: func(i int) tea.Cmd { author = i; return nil }}
				},
				Check: func() string {
					if author < 0 {
						return "Author is required"
					}
					return ""
				}},
		},
		Submit: func() (tea.Cmd, error) {
			p.list.Select(p.d.add(title, names[author]))
			return p.d.a.CloseAll(), nil
		},
	})
}
