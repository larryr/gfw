package gfw

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type Section struct {
	Name  string
	lines []string
}

func Process(rd io.Reader, wr io.Writer) error {

	log.Printf("gfw: reading(%v) writing(%v)\n", rd, wr)

	scanner := bufio.NewScanner(rd)
	var sections []Section
	scanner.Scan()
	t := scanner.Text()
	if t != "OPTIONS" {
		return errors.New("Expected OPTIONS")
	}

	s := &Section{"OPTIONS", nil}
	err := s.ParseSection(scanner, "INTERFACES")
	if err != nil {
		return err
	}
	sections = append(sections, *s)
	s = &Section{"INTERFACES", nil}
	err = s.ParseSection(scanner, "ALIASES")
	if err != nil {
		return err
	}
	sections = append(sections, *s)
	s = &Section{"ALIASES", nil}
	err = s.ParseSection(scanner, "FIREWALL")
	if err != nil {
		return err
	}
	sections = append(sections, *s)
	s = &Section{"FIREWALL", nil}
	err = s.ParseSection(scanner, "POLICIES")
	if err != nil {
		return err
	}
	sections = append(sections, *s)
	s = &Section{"POLICIES", nil}
	err = s.ParseSection(scanner, "CUSTOM")
	if err != nil {
		return err
	}
	sections = append(sections, *s)
	s = &Section{"CUSTOM", nil}
	err = s.ParseSection(scanner, "")
	if err != nil {
		return err
	}
	sections = append(sections, *s)

	dumpSectionInfo(sections)
	log.Printf("sections: %v\n", sections)

	return nil
}

func matchSection(t string) bool {
	switch t {
	case "OPTIONS",
		"INTERFACES",
		"ALIASES",
		"FIREWALL",
		"POLICIES",
		"CUSTOM":
		return true
	}
	return false
}

func (sec *Section) ParseSection(scanner *bufio.Scanner, stop string) error {
	//read lines until stop line
	for scanner.Scan() {
		t := scanner.Text()
		if t == stop {
			return nil
		}
		if matchSection(t) {
			return errors.New("unexpected section")
		}
		//skip empty lines
		line := strings.TrimSpace(t)
		if line == "" {
			continue
		}
		sec.lines = append(sec.lines, t)
	}
	return nil
}

func (sec *Section) LineCount() int {
	return len(sec.lines)
}

func dumpSectionInfo(sections []Section) {
	for idx, s := range sections {
		fmt.Printf("[%v] section:%v lines:%v\n", idx, s.Name, s.LineCount())
	}
}
