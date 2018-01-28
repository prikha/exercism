package twelve

import "fmt"

const testVersion = 1

const verseCount = 12

var verseItems = []string{"a Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds", "Gold Rings", "Geese-a-Laying", "Swans-a-Swimming", "Maids-a-Milking", "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming"}

var verseDays = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth" }
var verseCounts = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve" }

func Song()(out string) {
  out = ""
  for i := 1; i <= verseCount; i++ {
		out += Verse(i)
    out += "\n"
	}
  return
}

func Verse(day int) string {
  out := fmt.Sprintf("On the %s day of Christmas my true love gave to me", verseDays[day-1])
  buf := ""
  if day == 1 {
    return out + fmt.Sprintf(", %s", verseItems[0]) + "."
  }

  for i:= 0; i < day; i++ {
    if day != 1 && i == 0 {
      buf = fmt.Sprintf(", and %s", verseItems[i]) + buf
    } else {
      buf = fmt.Sprintf(", %s %s", verseCounts[(i)] ,verseItems[i]) + buf
    }
  }
  return out + buf + "."
}
