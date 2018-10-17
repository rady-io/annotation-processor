package processor

import (
	"bufio"
	"bytes"
	"regexp"
)

const (
	zeroStr = ""
)

type (
	Processor struct {
		*bufio.Scanner
	}
)

const (
	firstAnnRegex  = `(@[a-zA-Z_][0-9a-zA-Z_]*)\s*(.*)`
	secondAnnRegex = `(@[a-zA-Z_][0-9a-zA-Z_]*)\((.+?)\)\s*(.*)`
)

var (
	firstAnnRe  = regexp.MustCompile(firstAnnRegex)
	secondAnnRe = regexp.MustCompile(secondAnnRegex)
)

func NewProcessor(text string) *Processor {
	return &Processor{bufio.NewScanner(bytes.NewBufferString(text))}
}

func (process *Processor) Scan(fn func(ann, key, value string) (err error)) (err error) {
	for process.Scanner.Scan() {
		text := process.Text()
		if values := secondAnnRe.FindStringSubmatch(text); len(values) == 4 {
			err = fn(values[1], values[2], values[3])
		} else if values := firstAnnRe.FindStringSubmatch(text); len(values) == 3 {
			err = fn(values[1], zeroStr, values[2])
		}

		if err != nil {
			break
		}
	}
	return
}
