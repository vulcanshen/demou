package ui

import (
	"slices"
	"strconv"
)

// The shelves, in their order on [1] SHELVES.
const (
	toRead  = "To read"
	reading = "Reading"
	done    = "Done"
)

var shelves = []string{toRead, reading, done}

// book is one book on the bookshelf.
type book struct {
	id, title, author string
	year              int // 0 for a book added in demou
	shelf             string
	firstPage         string // "" for a book added in demou
}

func (b *book) yearText() string {
	if b.year == 0 {
		return ""
	}
	return strconv.Itoa(b.year)
}

type author struct {
	name       string
	born, died int
}

// authors are everyone with a book on the shelf, in the order of the
// NEW BOOK form's select.
var authors = []author{
	{"Bram Stoker", 1847, 1912},
	{"Charles Dickens", 1812, 1870},
	{"Charlotte Brontë", 1816, 1855},
	{"Emily Brontë", 1818, 1848},
	{"Frances Hodgson Burnett", 1849, 1924},
	{"H. G. Wells", 1866, 1946},
	{"Herman Melville", 1819, 1891},
	{"Jack London", 1876, 1916},
	{"Jane Austen", 1775, 1817},
	{"Jules Verne", 1828, 1905},
	{"L. Frank Baum", 1856, 1919},
	{"L. M. Montgomery", 1874, 1942},
	{"Lewis Carroll", 1832, 1898},
	{"Louisa May Alcott", 1832, 1888},
	{"Mark Twain", 1835, 1910},
	{"Mary Shelley", 1797, 1851},
	{"Oscar Wilde", 1854, 1900},
	{"Robert Louis Stevenson", 1850, 1894},
}

func authorNamed(name string) author {
	i := slices.IndexFunc(authors, func(a author) bool { return a.name == name })
	return authors[i]
}

// theBooks is the bookshelf demou starts with, the same every time: public
// domain classics, in the order they came out.
func theBooks() []*book {
	return []*book{
		{"pride", "Pride and Prejudice", "Jane Austen", 1813, toRead,
			"It is a truth universally acknowledged, that a single man in possession of a good fortune, must be in want of a wife."},
		{"emma", "Emma", "Jane Austen", 1815, done,
			"Emma Woodhouse, handsome, clever, and rich, with a comfortable home and happy disposition, seemed to unite some of the best blessings of existence; and had lived nearly twenty-one years in the world with very little to distress or vex her."},
		{"frankenstein", "Frankenstein", "Mary Shelley", 1818, reading,
			"You will rejoice to hear that no disaster has accompanied the commencement of an enterprise which you have regarded with such evil forebodings. I arrived here yesterday, and my first task is to assure my dear sister of my welfare and increasing confidence in the success of my undertaking."},
		{"jane-eyre", "Jane Eyre", "Charlotte Brontë", 1847, reading,
			"There was no possibility of taking a walk that day. We had been wandering, indeed, in the leafless shrubbery an hour in the morning; but since dinner (Mrs. Reed, when there was no company, dined early) the cold winter wind had brought with it clouds so sombre, and a rain so penetrating, that further out-door exercise was now out of the question."},
		{"wuthering", "Wuthering Heights", "Emily Brontë", 1847, done,
			"1801.—I have just returned from a visit to my landlord—the solitary neighbour that I shall be troubled with. This is certainly a beautiful country!"},
		{"moby-dick", "Moby-Dick", "Herman Melville", 1851, toRead,
			"Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world."},
		{"two-cities", "A Tale of Two Cities", "Charles Dickens", 1859, toRead,
			"It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of Darkness, it was the spring of hope, it was the winter of despair…"},
		{"expectations", "Great Expectations", "Charles Dickens", 1861, toRead,
			"My father's family name being Pirrip, and my Christian name Philip, my infant tongue could make of both names nothing longer or more explicit than Pip. So, I called myself Pip, and came to be called Pip."},
		{"alice", "Alice's Adventures in Wonderland", "Lewis Carroll", 1865, done,
			"Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, \"and what is the use of a book,\" thought Alice \"without pictures or conversations?\""},
		{"little-women", "Little Women", "Louisa May Alcott", 1868, done,
			"\"Christmas won't be Christmas without any presents,\" grumbled Jo, lying on the rug."},
		{"eighty-days", "Around the World in Eighty Days", "Jules Verne", 1872, toRead,
			"Mr. Phileas Fogg lived, in 1872, at No. 7, Saville Row, Burlington Gardens, the house in which Sheridan died in 1814."},
		{"tom-sawyer", "The Adventures of Tom Sawyer", "Mark Twain", 1876, done,
			"\"TOM!\" No answer. \"TOM!\" No answer. \"What's gone with that boy, I wonder? You TOM!\" No answer."},
		{"treasure", "Treasure Island", "Robert Louis Stevenson", 1883, toRead,
			"Squire Trelawney, Dr. Livesey, and the rest of these gentlemen having asked me to write down the whole particulars about Treasure Island, from the beginning to the end, keeping nothing back but the bearings of the island, and that only because there is still treasure not yet lifted, I take up my pen in the year of grace 17— and go back to the time when my father kept the Admiral Benbow inn and the brown old seaman with the sabre cut first took up his lodging under our roof."},
		{"dorian-gray", "The Picture of Dorian Gray", "Oscar Wilde", 1890, reading,
			"The studio was filled with the rich odour of roses, and when the light summer wind stirred amidst the trees of the garden, there came through the open door the heavy scent of the lilac, or the more delicate perfume of the pink-flowering thorn."},
		{"time-machine", "The Time Machine", "H. G. Wells", 1895, done,
			"The Time Traveller (for so it will be convenient to speak of him) was expounding a recondite matter to us. His grey eyes shone and twinkled, and his usually pale face was flushed and animated."},
		{"dracula", "Dracula", "Bram Stoker", 1897, toRead,
			"3 May. Bistritz.—Left Munich at 8:35 P. M., on 1st May, arriving at Vienna early next morning; should have arrived at 6:46, but train was an hour late."},
		{"oz", "The Wonderful Wizard of Oz", "L. Frank Baum", 1900, done,
			"Dorothy lived in the midst of the great Kansas prairies, with Uncle Henry, who was a farmer, and Aunt Em, who was the farmer's wife."},
		{"call-wild", "The Call of the Wild", "Jack London", 1903, done,
			"Buck did not read the newspapers, or he would have known that trouble was brewing, not alone for himself, but for every tide-water dog, strong of muscle and with warm, long hair, from Puget Sound to San Diego."},
		{"green-gables", "Anne of Green Gables", "L. M. Montgomery", 1908, done,
			"Mrs. Rachel Lynde lived just where the Avonlea main road dipped down into a little hollow, fringed with alders and ladies' eardrops and traversed by a brook that had its source away back in the woods of the old Cuthbert place…"},
		{"secret-garden", "The Secret Garden", "Frances Hodgson Burnett", 1911, done,
			"When Mary Lennox was sent to Misselthwaite Manor to live with her uncle everybody said she was the most disagreeable-looking child ever seen. It was true, too."},
	}
}
