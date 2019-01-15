package helper

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Password struct {
	Pass            string
	Length          int
	Score           int
	ContainsUpper   bool
	ContainsLower   bool
	ContainsNumber  bool
	ContainsSpecial bool
	DictionaryBased bool
}

// ProcessPassword will parse the password and populate the Password struct attributes.
func (p *Password) ProcessPassword() {
	matchLower := regexp.MustCompile(`[a-z]`)
	matchUpper := regexp.MustCompile(`[A-Z]`)
	matchNumber := regexp.MustCompile(`[0-9]`)
	matchSpecial := regexp.MustCompile(`[\!\@\#\$\%\^\&\*\(\\\)\-_\=\+\,\.\?\/\:\;\{\}\[\]~]`)

	if p.Length < 8 {
		p.Score = 0
		return
	}

	if matchLower.MatchString(p.Pass) {
		p.ContainsLower = true
		p.Score++
	}
	if matchUpper.MatchString(p.Pass) {
		p.ContainsUpper = true
		p.Score++
	}
	if matchNumber.MatchString(p.Pass) {
		p.ContainsNumber = true
		p.Score++
	}
	if matchSpecial.MatchString(p.Pass) {
		p.ContainsSpecial = true
		p.Score++
	}
	if searchDict(p.Pass) {
		p.DictionaryBased = true
		p.Score--
	}
}

// GetScore will provide the score of the password.
func (p *Password) GetScore() int {
	return p.Score
}

// HasUpper indicates whether the password contains an upper case letter.
func (p *Password) HasUpper() bool {
	return p.ContainsUpper
}

// HasLower indicates whether the password contains a lower case letter.
func (p *Password) HasLower() bool {
	return p.ContainsLower
}

// HasNumber indicates whether the password contains a number.
func (p *Password) HasNumber() bool {
	return p.ContainsNumber
}

// HasSpecial indicates whether the password contains a special character.
func (p *Password) HasSpecial() bool {
	return p.ContainsSpecial
}

// InDictionary will return true or false if it's been detected
// that the given password is a dictionary based.
func (p *Password) InDictionary() bool {
	return p.DictionaryBased
}

func searchDict(word string) bool {
	file, err := os.Open("./weakpass.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), word) {
			return true
		}
	}
	return false
}
