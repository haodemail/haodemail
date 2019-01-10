package helper

import (
	"errors"
	"strings"
)

// SplitMailbox split mailbox into username and domain parts
func SplitMailbox(mbox string) (username string, domain string, err error) {
	spt := strings.SplitN(strings.ToLower(mbox), "@", 2)
	if len(spt) == 2 {
		username = spt[0]
		domain = spt[1]
	} else {
		err = errors.New("invalid mailbox")
	}
	return
}
