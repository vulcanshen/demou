package ui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	terminu "github.com/vulcanshen/terminu-go"
)

// moveScreen is [M]OVE, the first chapter: look around and move. [1] SHELVES
// lists the shelves; [2] the books on the one under the cursor, with a tab of
// their authors.
type moveScreen struct {
	d       *demou
	screen  *terminu.Screen
	shelves *terminu.List
	list    *terminu.List
	books   *terminu.Panel
}

const authorsTab = 1

func newMoveScreen(d *demou) *moveScreen {
	m := &moveScreen{d: d}
	// Names (titles, authors) shrink to share the room; short values keep
	// their width. Columns that still do not fit go by Drop, the name never
	// (tdp panel.5).
	m.shelves = &terminu.List{Columns: []terminu.Column{{Title: "Shelf"}, {Title: "Books", Number: true, Drop: 1}},
		OnMove: func(terminu.Row) tea.Cmd { m.fill(); return nil }}
	m.list = &terminu.List{EmptyFact: "No books on this shelf", EmptyHint: "Press [Space] to see what you can do here"}
	m.books = &terminu.Panel{Label: "Books", Tabs: []string{"Authors"}, Body: m.list,
		OnTab: func(int) tea.Cmd { m.fill(); return nil }}
	open := d.openBook(m.list)
	open.Hidden = func() bool { return m.books.Tab == authorsTab }
	m.books.Actions = []terminu.Action{open,
		{Key: "Enter", Label: "Open", Desc: "who they were", Scope: terminu.ItemScope, Hint: "open",
			Hidden: func() bool { return m.books.Tab != authorsTab }, Do: m.openAuthor}}
	shelves := &terminu.Panel{Label: "Shelves", Body: m.shelves, Details: m.books, Actions: []terminu.Action{
		// TODO(tdp DM42): only the focus moves to the details, so it is not
		// logged (tdp note.7); whether an item operation that does only that
		// goes in the App Log, tdp does not say.
		{Key: "Enter", Label: "Books", Desc: "the books on this shelf", Scope: terminu.ItemScope, Hint: "books",
			ToDetails: func() bool { return true }},
	}}
	m.screen = &terminu.Screen{Name: "Move", Key: "M", Desc: "look around and move",
		Panels: []*terminu.Panel{shelves, m.books}, NarrowBelow: 60, Layout: m.layout}
	return m
}

// layout: the shelves narrow on the left, the books on the right.
func (m *moveScreen) layout(w, h int) []terminu.Rect {
	left := 22
	return []terminu.Rect{{X: 0, Y: 0, W: left, H: h}, {X: left, Y: 0, W: w - left, H: h}}
}

func (m *moveScreen) refresh() {
	rows := make([]terminu.Row, len(shelves))
	for i, s := range shelves {
		n := 0
		for _, b := range m.d.books {
			if b.shelf == s {
				n++
			}
		}
		rows[i] = terminu.Row{ID: s, Cells: []string{s, strconv.Itoa(n)}}
	}
	m.shelves.SetRows(rows)
	m.fill()
}

// fill shows the books, or their authors, on the shelf under the cursor.
func (m *moveScreen) fill() {
	shelf, _ := m.shelves.Current()
	var on []*book
	for _, b := range m.d.books {
		if b.shelf == shelf.ID {
			on = append(on, b)
		}
	}
	var rows []terminu.Row
	if m.books.Tab == authorsTab {
		m.list.Columns = []terminu.Column{{Title: "Name", Shrink: true}, {Title: "Born", Number: true, Drop: 2},
			{Title: "Died", Number: true, Drop: 1}}
		seen := map[string]bool{}
		for _, b := range on {
			if seen[b.author] {
				continue
			}
			seen[b.author] = true
			a := authorNamed(b.author)
			rows = append(rows, terminu.Row{ID: a.name, Cells: []string{a.name, strconv.Itoa(a.born), strconv.Itoa(a.died)}})
		}
	} else {
		// A title always shows; the year goes first, then the author (tdp panel.5).
		m.list.Columns = []terminu.Column{{Title: "Title", Shrink: true}, {Title: "Author", Shrink: true, Drop: 2},
			{Title: "Year", Number: true, Drop: 1}}
		for _, b := range on {
			rows = append(rows, terminu.Row{ID: b.id, Cells: []string{b.title, b.author, b.yearText()}})
		}
	}
	m.list.SetRows(rows)
}

// openAuthor is Enter on an author: their years and their books here.
func (m *moveScreen) openAuthor() tea.Cmd {
	r, _ := m.list.Current()
	a := authorNamed(r.ID)
	lines := []string{a.name + ", " + strconv.Itoa(a.born) + "–" + strconv.Itoa(a.died) + "."}
	for _, b := range m.d.books {
		if b.author == a.name {
			lines = append(lines, b.title)
		}
	}
	m.d.a.Log(terminu.Info, "Opened the author "+a.name)
	return m.d.a.Viewer(&terminu.Viewer{Title: "Author", Subject: a.name, Glyph: glyphAuthor, Lines: lines})
}
