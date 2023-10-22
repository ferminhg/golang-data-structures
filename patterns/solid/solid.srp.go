package solid

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) bool {
	if index < 0 || index >= len(j.entries) {
		return false
	}

	j.entries = append(j.entries[:index], j.entries[index+1:]...)
	return true
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// separation of concerns
// God Object -> anti-pattern

func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	//
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	//
}

// Alternative approach Single Responsibility Principle
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(
		strings.Join(
			j.entries,
			LineSeparator,
		)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(
		strings.Join(
			j.entries,
			p.lineSeparator,
		)), 0644)
}

func solidSRP() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Printf("%s\n", j.String())
	//
	SaveToFile(&j, "journal.txt")
	//
	p := Persistence{"\n"}
	p.SaveToFile(&j, "journal.txt")
}
