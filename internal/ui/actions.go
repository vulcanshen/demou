package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	terminu "github.com/vulcanshen/terminu-go"
)

// actionsScreen is [A]CTIONS, the second chapter: find what you can do. One
// panel of every book, with an item operation for each shelf move (disabled
// on the books it does not apply to) and a panel operation that sorts.
type actionsScreen struct {
	d      *demou
	screen *terminu.Screen
	list   *terminu.List
	byYear bool
}

const yearCol = 2

func newActionsScreen(d *demou) *actionsScreen {
	s := &actionsScreen{d: d}
	s.list = &terminu.List{
		Columns:   []terminu.Column{{Title: "Title", Flex: true}, {Title: "Author"}, {Title: "Year", Number: true}, {Title: "Shelf"}},
		EmptyFact: "No books", EmptyHint: "Press [Space] to see what you can do here",
	}
	notOn := func(shelf string) func() bool {
		return func() bool { b := d.current(s.list); return b == nil || b.shelf != shelf }
	}
	books := &terminu.Panel{Label: "Books", Body: s.list, Actions: []terminu.Action{
		d.openBook(s.list),
		{Key: "b", Label: "begin", Desc: "move it to Reading", Scope: terminu.ItemScope, Hint: "begin", Disabled: notOn(toRead),
			Do: func() tea.Cmd { d.moveTo(d.current(s.list), reading); return nil }},
		{Key: "f", Label: "finish", Desc: "move it to Done", Scope: terminu.ItemScope, Hint: "finish", Disabled: notOn(reading),
			Do: func() tea.Cmd { d.moveTo(d.current(s.list), done); return nil }},
		{Key: "s", Label: "sort", Desc: "by year, or by title again", Scope: terminu.PanelScope, Do: s.sort},
	}}
	s.screen = &terminu.Screen{Name: "Actions", Key: "A", Desc: "find what you can do here", Panels: []*terminu.Panel{books}}
	s.list.SortBy(terminu.SortKey{Col: 0})
	return s
}

func (s *actionsScreen) sort() tea.Cmd {
	s.byYear = !s.byYear
	if s.byYear {
		s.list.SortBy(terminu.SortKey{Col: yearCol})
	} else {
		s.list.SortBy(terminu.SortKey{Col: 0})
	}
	return nil
}

func (s *actionsScreen) refresh() {
	rows := make([]terminu.Row, len(s.d.books))
	for i, b := range s.d.books {
		rows[i] = terminu.Row{ID: b.id, Cells: []string{b.title, b.author, b.yearText(), b.shelf}}
	}
	s.list.SetRows(rows)
}
