package complete

import (
	"github.com/bytedance/gg/gslice"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var AllEngWords []string

func Init() {
	WordsFilePath := filepath.Join(filepath.Dir(os.Args[0]), "conf", "all_eng_words.txt")
	content, err := os.ReadFile(WordsFilePath)
	if err != nil {
		log.Fatalf("Error reading words file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	AllEngWords = gslice.Map(lines, func(line string) string { return strings.TrimSpace(line) })
	AllEngWords = gslice.Filter(AllEngWords, func(word string) bool { return word != "" })
}
