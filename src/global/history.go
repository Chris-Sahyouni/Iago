package global

import (
	"bufio"
	"os"
	"strings"
	"io"
)

var MAX_RECORDS int = 100
var HISTORY_FILE_PATH string = ".iago_history"

type History struct {
	records []string
}

func (h *History) Init() {

	records := make([]string, 0)
	historyFile, err := os.Open(HISTORY_FILE_PATH)

	if err == nil {
		defer historyFile.Close()
		historyReader := bufio.NewReader(historyFile)
		for range MAX_RECORDS {
			line, err := historyReader.ReadString('\n');
			if err != nil {
				if err == io.EOF {
					records = append(records, line)
				}
				break
			}
			records = append(records, line)
		}
	}
	h.records = records
}

func (h *History) Add(line string) {

	if strings.TrimSpace(line) == "quit" {
		return
	}

	if strings.TrimSpace(line) == h.records[len(h.records) - 1] {
		return
	}

	h.records = append(h.records, line)

}

func (h *History) Flush() {
	cutoff := max(len(h.records) - MAX_RECORDS, 0)
	keep := h.records[cutoff:]
	historyFile, err := os.OpenFile(HISTORY_FILE_PATH, os.O_WRONLY, 0644)
	if err == nil {
		defer historyFile.Close()
		historyWriter := bufio.NewWriter(historyFile)
		keepConcat := strings.Join(keep, "")
		historyWriter.WriteString(keepConcat)
		historyWriter.Flush()
	}
}