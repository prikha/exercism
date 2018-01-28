package house

const testVersion = 1

var Firstliners = []string{"This is the house that Jack built.",
"This is the malt",
"This is the rat",
"This is the cat",
"This is the dog",
"This is the cow with the crumpled horn",
"This is the maiden all forlorn",
"This is the man all tattered and torn",
"This is the priest all shaven and shorn",
"This is the rooster that crowed in the morn",
"This is the farmer sowing his corn",
"This is the horse and the hound and the horn"}

var Lastliners = []string{"that belonged to the farmer sowing his corn",
"that kept the rooster that crowed in the morn",
"that woke the priest all shaven and shorn",
"that married the man all tattered and torn",
"that kissed the maiden all forlorn",
"that milked the cow with the crumpled horn",
"that tossed the dog",
"that worried the cat",
"that killed the rat",
"that ate the malt",
"that lay in the house that Jack built."}

const VerseCount int = 12


func Song() (out string) {
	for i := 1; i <= VerseCount; i++ {
		out += Verse(i)
		if i != VerseCount {
			out += "\n\n"
		}
	}
	return

}

func Verse(num int) (out string) {
	idx := num - 1
	out += Firstliners[idx]
	for i := (VerseCount - num); i < VerseCount; i++ {
		if i > len(Lastliners) - 1 {
			continue
		}
		out += "\n"
		out += Lastliners[i]
	}
	return
}