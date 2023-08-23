package cleaner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Cleaner struct {
	actions []Action
}

func New() *Cleaner {
	cleaner := &Cleaner{}
	cleaner.actions = []Action{}

	return cleaner
}

func (c *Cleaner) LoadFile(filename string) error {
	b, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		return err
	}

	return c.Load(string(b))
}

func (c *Cleaner) Load(content string) error {

	scanner := bufio.NewScanner(strings.NewReader(content))
	keep := scanner.Scan()
	// First line is a header
	if !keep {
		return fmt.Errorf("file is empty")
	}

	keep = scanner.Scan()
	var err error = nil

	for keep && err == nil {
		keep, err = c.readGroup(scanner)
	}

	return err
}

func (c *Cleaner) readGroup(scanner *bufio.Scanner) (bool, error) {
	rxGroup := regexp.MustCompile(`Duplicate (\d+) Size (\d+) Files (\d+) MD5 (\w+)`)

	rxFile := regexp.MustCompile(`^\[([\w]?)\]\s+'(.+)'`)

	// scanner.Scan()
	groupLine := scanner.Text()

	matches := rxGroup.FindStringSubmatch(groupLine)
	if len(matches) == 0 {
		return false, fmt.Errorf("expected group line but found '%s'", groupLine)
	}

	scan := true
	groupId, _ := strconv.Atoi(matches[1])
	actions := []Action{}

	// Read files
	for {
		scan = scanner.Scan()
		if !scan {
			break
		}
		line := scanner.Text()

		fileMatches := rxFile.FindStringSubmatch(line)
		if len(fileMatches) == 0 {
			break
		}
		// Add file to scanner
		actionType, err := StringToAction(fileMatches[1])
		if err != nil {
			return scan, err
		}
		actions = append(actions, Action{Group: groupId, Type: actionType, Filename: fileMatches[2]})
	}

	c.actions = append(c.actions, actions...)
	return scan, nil
}

func (c *Cleaner) Actions() []Action {
	return c.actions
}
