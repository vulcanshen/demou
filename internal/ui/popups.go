package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	terminu "github.com/vulcanshen/terminu-go"
)

// popupsScreen is [P]OPUPS, the third chapter: popups over popups. Remove
// asks first; New book opens a small form whose Author is a select over it.
type popupsScreen struct {
	d      *demou
	screen *terminu.Screen
	list   *terminu.List
	// removeAsk names the book under the cursor; it is read when Remove runs.
	removeAsk *terminu.Ask
}

func newPopupsScreen(d *demou) *popupsScreen {
	p := &popupsScreen{d: d, removeAsk: &terminu.Ask{Title: "Remove", Verb: "remove"}}
	p.list = &terminu.List{
		Columns:   []terminu.Column{{Title: "Title", Flex: true}, {Title: "Author"}, {Title: "Shelf"}},
		EmptyFact: "No books on the shelf", EmptyHint: "Press [n] to add a book",
		OnMove: func(r terminu.Row) tea.Cmd {
			p.removeAsk.Question = "Remove " + d.book(r.ID).title + " from the shelf?"
			return nil
		},
	}
	books := &terminu.Panel{Label: "Books", Body: p.list, Actions: []terminu.Action{
		d.openBook(p.list),
		{Key: "r", Label: "remove", Desc: "take it off the shelf", Scope: terminu.ItemScope, Hint: "remove", Ask: p.removeAsk,
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
// chosen in a select. Add puts the book on To read, the cursor on it.
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
			return nil, nil
		},
	})
}
