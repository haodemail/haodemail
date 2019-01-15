package helper

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
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

// validate email address
func ValidateEmail(email string) (ok bool) {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0," +
		"61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegexp.MatchString(email)
}

// RandomString make random string
func RandomString(l int) string {
	distArray := make([]byte, l)
	rand.Seed(time.Now().Unix())
	samples := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	n := len(samples)
	for i := 0; i < l; i++ {
		x := rand.Intn(n)
		distArray[i] = samples[x]
	}
	return string(distArray)
}

// Escape escape character in sql
func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break
		case '\r':
			escape = 'r'
			break
		case '\\':
			escape = '\\'
			break
		case '\'':
			escape = '\''
			break
		case '"': /* Better safe than sorry */
			escape = '"'
			break
		case '\032': /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}

func Quote(s string) string {
	s = strings.Replace(s, "+", "%2B", -1)
	s = strings.Replace(s, "/", "%2F", -1)
	s = strings.Replace(s, "=", "%3D", -1)
	return s
}

func UnQuote(s string) string {
	s = strings.Replace(s, "%2B", "+", -1)
	s = strings.Replace(s, "%2F", "/", -1)
	s = strings.Replace(s, "%3D", "=", -1)
	s = strings.Replace(s, "%2b", "+", -1)
	s = strings.Replace(s, "%2f", "/", -1)
	s = strings.Replace(s, "%3d", "=", -1)
	return s
}
