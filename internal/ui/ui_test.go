package ui

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	terminu "github.com/vulcanshen/terminu-go"
)

// states are the keys from the start to every screen and every kind of
// popup demou opens.
var states = []struct {
	name string
	keys []string
}{
	{"MOVE", nil},
	{"MOVE, [2] focused", []string{"Tab"}},
	{"MOVE, the AUTHORS tab", []string{"Tab", "l"}},
	{"Space menu", []string{"Space"}},
	{"global operation popup over the Space menu", []string{"Space", "G", "Enter"}},
	{"key reference", []string{"?"}},
	{"first page", []string{"Tab", "Enter"}},
	{"an author", []string{"Tab", "l", "Enter"}},
	{"ACTIONS", []string{"A"}},
	{"ACTIONS, Space menu", []string{"A", "Space"}},
	{"POPUPS", []string{"P"}},
	{"remove confirm", []string{"P", "r"}},
	{"remove confirm over the Space menu", []string{"P", "Space", "j", "Enter"}},
	{"NEW BOOK form", []string{"P", "n"}},
	{"title text popup over the form", []string{"P", "n", "Enter"}},
	{"author select over the form", []string{"P", "n", "j", "Enter"}},
	{"discard confirm over the form", []string{"P", "n", "Enter", "x", "Enter", "Esc"}},
	{"form errors", []string{"P", "n", "Tab", "Enter"}},
	{"quit confirm", []string{"q"}},
	{"splash", []string{"V"}},
}

// iconWidths are the widths of an icon and a powerline symbol the screen
// tests run with (tdp E5).
var iconWidths = [][2]int{{1, 1}, {2, 1}, {2, 2}}

// eachIconWidth runs check in parallel for every case, once for each of
// iconWidths. Only the widths change between the rounds, while nothing runs.
func eachIconWidth[C any](t *testing.T, cases []C, name func(C) string, check func(*testing.T, C)) {
	defer terminu.SetIconWidth(1, 1)
	for _, iw := range iconWidths {
		terminu.SetIconWidth(iw[0], iw[1])
		t.Run(fmt.Sprintf("icon %d, powerline %d", iw[0], iw[1]), func(t *testing.T) {
			for _, c := range cases {
				t.Run(name(c), func(t *testing.T) {
					t.Parallel()
					check(t, c)
				})
			}
		})
	}
}

type sized struct {
	w, h  int
	name  string
	keys  []string
	small [2]int // the size to shrink to once it is open; 0 for none
}

// TestL4EveryLineIsTheTerminalsWidth is tdp L4: on every screen and under
// every popup, at many sizes.
func TestL4EveryLineIsTheTerminalsWidth(t *testing.T) {
	var cases []sized
	for _, sz := range [][2]int{{80, 40}, {120, 40}, {80, 24}, {40, 12}, {24, 8}} {
		for _, st := range states {
			cases = append(cases, sized{w: sz[0], h: sz[1], name: st.name, keys: st.keys})
		}
	}
	eachIconWidth(t, cases, func(c sized) string { return fmt.Sprintf("%dx%d %s", c.w, c.h, c.name) }, checkSized)
}

// TestE5PopupBiggerThanTheScreen is tdp E5: a popup keeps the size it opened
// with (Rules F7), so shrinking the terminal leaves it wider than the
// screen, taller, or both; every line still fits.
func TestE5PopupBiggerThanTheScreen(t *testing.T) {
	var cases []sized
	for _, st := range states[3:] { // the popups
		for _, small := range [][2]int{{30, 40}, {120, 10}, {30, 10}} {
			cases = append(cases, sized{w: 120, h: 40, name: st.name, keys: st.keys, small: small})
		}
	}
	eachIconWidth(t, cases, func(c sized) string { return fmt.Sprintf("%s, then %dx%d", c.name, c.small[0], c.small[1]) }, checkSized)
}

func checkSized(t *testing.T, c sized) {
	hs := newHarness(t, c.w, c.h)
	hs.keys(c.keys...)
	if c.small[0] > 0 {
		hs.send(tea.WindowSizeMsg{Width: c.small[0], Height: c.small[1]})
	}
	hs.fullWidth(c.name)
}

// The same books, in the same order, on the same shelves, every time demou
// starts: what is recorded once looks the same again.
func TestSameEveryStart(t *testing.T) {
	for _, st := range states {
		t.Run(st.name, func(t *testing.T) {
			t.Parallel()
			a, b := newHarness(t, 80, 40), newHarness(t, 80, 40)
			a.keys(st.keys...)
			b.keys(st.keys...)
			if a.text() != b.text() {
				t.Fatalf("it differs between two starts:\n%s\n\n%s", a.text(), b.text())
			}
		})
	}
}

func TestMoveShelvesShowTheirBooks(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.contains("To read      7", "Reading      3", "Done        10", "Pride and Prejudice", "Dracula")
	hs.lacks("Frankenstein")
	hs.keys("j") // Reading
	hs.contains("Frankenstein", "Jane Eyre", "The Picture of Dorian Gray")
	hs.lacks("Pride and Prejudice")
	hs.keys("j", "j") // past Done, back to To read (tdp K12)
	hs.contains("Pride and Prejudice")
	hs.keys("G") // Done
	hs.contains("Emma", "The Secret Garden")
}

func TestMoveTabsAndPanels(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("2", "l")
	hs.contains("NAME", "BORN", "DIED", "Jane Austen", "1775", "1817")
	hs.lacks("Pride and Prejudice")
	hs.keys("l") // past the last tab, back to BOOKS
	hs.contains("Pride and Prejudice")
	hs.keys("1", "Enter") // a shelf's Enter sends the focus to its books (tdp K3)
	hs.keys("Enter")      // so this Enter opens the first book
	hs.contains("FIRST PAGE — Pride and Prejudice", "It is a truth universally acknowledged")
}

func TestMoveOpenAnAuthor(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("Tab", "l", "Enter")
	hs.contains("AUTHOR — Jane Austen", "Jane Austen, 1775–1817.", "Pride and Prejudice", "Emma")
}

// A first page folds between words, never inside one (tdp components,
// "Wrapping"; TODO(lib DM2)).
func TestFirstPageFoldsBetweenWords(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("Tab", "Enter")
	hs.contains("a single man in", "possession of a good fortune")
}

func TestActionsBeginAndFinish(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("A") // sorted by title: A Tale of Two Cities, To read
	two := hs.d.book("two-cities")
	hs.keys("f") // not Reading: Finish is disabled (tdp M6)
	if two.shelf != toRead {
		t.Fatalf("Finish moved a book on To read to %s", two.shelf)
	}
	hs.keys("b")
	if two.shelf != reading {
		t.Fatalf("Begin left it on %s", two.shelf)
	}
	hs.keys("b") // now Begin is the disabled one
	if two.shelf != reading {
		t.Fatalf("Begin moved a book on Reading to %s", two.shelf)
	}
	hs.keys("f")
	if two.shelf != done {
		t.Fatalf("Finish left it on %s", two.shelf)
	}
	hs.keys("M") // the shelves count it where it is now
	hs.contains("To read      6", "Done        11")
}

func TestActionsSpaceMenu(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("A", "Space")
	hs.contains("ITEM OPERATION", "[Enter] Open", "[b]egin", "[f]inish", "PANEL OPERATION", "[s]ort", "Global operation")
	hs.keys("G", "Enter")
	hs.contains("GLOBAL OPERATION", "[M]ove", "[A]ctions", "[P]opups", "Quit")
	hs.keys("Esc") // one layer at a time (tdp K4)
	hs.contains("ITEM OPERATION")
	hs.keys("Esc")
	hs.lacks("ITEM OPERATION")
}

func TestActionsSort(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("A")
	first := func() string {
		hs.t.Helper()
		r, _ := hs.d.actions.list.Current()
		return r.ID
	}
	if got := first(); got != "two-cities" {
		t.Fatalf("by title, the cursor is on %s first", got)
	}
	hs.keys("s", "gg")
	if got := first(); got != "pride" {
		t.Fatalf("by year, the first book is %s", got)
	}
	hs.keys("s", "gg")
	if got := first(); got != "two-cities" {
		t.Fatalf("by title again, the first book is %s", got)
	}
}

func TestPopupsRemoveAsksFirst(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("P", "r")
	hs.contains("REMOVE", "Remove Pride and Prejudice from the shelf?", "Enter:remove Esc:cancel")
	hs.keys("Esc")
	if hs.d.book("pride") == nil {
		t.Fatal("Esc removed the book")
	}
	hs.keys("j", "r") // the question names the book under the cursor
	hs.contains("Remove Emma from the shelf?")
	hs.keys("Enter")
	if hs.d.book("emma") != nil {
		t.Fatal("Enter did not remove the book")
	}
	hs.lacks("Emma ")
}

// From the Space menu, the confirm goes over the menu and Esc goes back to
// it (tdp F4).
func TestPopupsRemoveFromTheMenu(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("P", "Space", "j", "Enter")
	hs.contains("Remove Pride and Prejudice from the shelf?")
	hs.keys("Esc")
	hs.contains("ITEM OPERATION", "[r]emove")
	hs.keys("Esc")
	hs.lacks("ITEM OPERATION")
}

func TestPopupsNewBook(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("P", "n")
	hs.contains("NEW BOOK", "Title", "Author", "not set", "Add")
	hs.keys("Enter", "Persuasion", "Enter")
	hs.contains("Persuasion")
	hs.keys("j", "Enter") // the select goes over the form
	hs.contains("AUTHOR", "Bram Stoker", "Enter:choose")
	hs.keys("Esc") // closes the select only
	hs.contains("NEW BOOK")
	hs.keys("Enter", "Tab", "Austen", "Enter", "Enter")
	hs.contains("Jane Austen")
	hs.keys("Tab", "Enter") // Add
	hs.lacks("NEW BOOK")
	b := hs.d.book("added-1")
	if b == nil || b.title != "Persuasion" || b.author != "Jane Austen" || b.shelf != toRead {
		t.Fatalf("added %+v", b)
	}
	if r, _ := hs.d.popups.list.Current(); r.ID != "added-1" {
		t.Fatalf("the cursor is on %s, not the new book", r.ID)
	}
	hs.keys("M")
	hs.contains("To read      8")
}

func TestPopupsNewBookNeedsTitleAndAuthor(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("P", "n", "Tab", "Enter")
	hs.contains("NEW BOOK", "Title is required")
	if len(hs.d.books) != 20 {
		t.Fatal("an empty form added a book")
	}
}

func TestPopupsChangedFormAsksBeforeEsc(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("P", "n", "Esc") // not changed: Esc closes it
	hs.lacks("NEW BOOK")
	hs.keys("n", "Enter", "x", "Enter", "Esc")
	hs.contains("Discard changes to New book?")
	hs.keys("Esc") // back to the form, the value still there
	hs.contains("NEW BOOK", "Title   x")
	hs.keys("Esc", "Enter")
	hs.lacks("NEW BOOK")
	if len(hs.d.books) != 20 {
		t.Fatal("discarding added a book")
	}
}

func TestQuitAsksFirst(t *testing.T) {
	t.Parallel()
	hs := newHarness(t, 80, 40)
	hs.keys("q")
	hs.contains("QUIT", "Quit demou?", "Enter:quit Esc:cancel")
	hs.keys("Esc")
	hs.lacks("Quit demou?")
	if hs.quit {
		t.Fatal("Esc quit")
	}
	hs.keys("q", "Enter")
	if !hs.quit {
		t.Fatal("Enter did not quit")
	}
}

// Every book's Enter is the same on every screen: its first page (tdp K3).
func TestEnterOpensTheFirstPageEverywhere(t *testing.T) {
	t.Parallel()
	for _, keys := range [][]string{{"Tab", "Enter"}, {"A", "Enter"}, {"P", "Enter"}} {
		hs := newHarness(t, 80, 40)
		hs.keys(keys...)
		hs.contains("FIRST PAGE —")
	}
}

// TestSplashIsDocsIcon is tdp E6: the splash draws docs/icon.svg, demou's
// mark as S5 has it, the U frame and then D, E, M, O.
func TestSplashIsDocsIcon(t *testing.T) {
	t.Parallel()
	want, err := os.ReadFile("../../docs/icon.svg")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(icon, want) {
		t.Fatal("internal/ui/icon.svg is not docs/icon.svg: copy docs/icon.svg over it")
	}
	if _, err := terminu.NewSplash(icon, "demou", "dev", "intro"); err != nil {
		t.Fatalf("the lib's splash does not take docs/icon.svg: %v", err)
	}
	var comments []string
	for _, l := range strings.Split(string(icon), "\n") {
		if c, ok := strings.CutPrefix(strings.TrimSpace(l), "<!-- "); ok {
			comments = append(comments, strings.TrimSuffix(c, " -->"))
		}
	}
	if !slices.Equal(comments, []string{"U", "D", "E", "M", "O"}) {
		t.Fatalf("the letters are %q, want U D E M O", comments)
	}
	s := splash()
	var lines []string
	for f := 0; ; f++ {
		var done bool
		if lines, done = s(f, 120, 40); done {
			break
		}
		if f > 1000 {
			t.Fatal("the reveal never ends")
		}
	}
	text := terminu.StripStyles(strings.Join(lines, "\n"))
	for _, want := range []string{"demou", "dev", "The terminu family's operations, hands-on", "Press any key to close"} {
		if !strings.Contains(text, want) {
			t.Fatalf("the splash does not show %q:\n%s", want, text)
		}
	}
}
