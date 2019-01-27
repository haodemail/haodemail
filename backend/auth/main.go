package main

import (
	"../models"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func GetMechList() (mechList [][]byte) {
	mechList = append(mechList, []byte("PLAIN\tplaintext"))
	mechList = append(mechList, []byte("LOGIN\tplaintext"))
	return
}

func HandShake(cuid int) (challenge []byte) {
	challenge = append(challenge, []byte("VERSION\t1\t0\n")...)
	for _, m := range GetMechList() {
		challenge = append(challenge, []byte("MECH\t")...)
		challenge = append(challenge, m...)
		challenge = append(challenge, []byte("\n")...)
	}
	pid := os.Getpid()
	challenge = append(challenge, []byte(fmt.Sprintf("SPID\t%d\n", pid))...)
	challenge = append(challenge, []byte(fmt.Sprintf("CUID\t%d\n", cuid))...)
	challenge = append(challenge, []byte("DONE\n")...)
	return
}

func authenticate(username, password string) bool {
	if strings.Index(username, "@") < 0 {
		username = username + "@" + models.Config.MailServer.PrimaryDomain
	}
	username = strings.ToLower(username)
	user, err := models.DaoGetVirtualUserByName(username)
	if err == nil {
		return models.ComparePassword(user.Password, password)
	}
	return false
}

type authSession struct {
	cid      string
	userSent bool
	pwdSent  bool
	mech     string
	username string
	password string
	lasttime time.Time
}

var authMap = make(map[string]*authSession)

// postfix connect to sasl-server, and keep alive, so do not close the connection
func connHandler(c net.Conn) {
	defer c.Close()
	if c == nil {
		return
	}
	buf := make([]byte, 8192)
	var handSent bool
	//log.Println("accept connection from", c.RemoteAddr())
	if !handSent {
		if _, err := c.Write(HandShake(rand.Int())); err == nil {
			//log.Println("hand shake ok")
			handSent = true
		} else {
			log.Println("hand shake err", err)
			return
		}
	}

	for {
		var err error
		output := make([]byte, 0)
		cnt, ioErr := c.Read(buf)
		if ioErr == io.EOF {
			break
		}
		if ioErr != nil || cnt == 0 {
			err = errors.New("error occur" + err.Error())
		}
		line := string(buf[0:cnt])
		//fmt.Println("Got", line)
		if strings.HasPrefix(line, "AUTH") {
			l := strings.Split(line, "\t")
			if len(l) < 2 {
				err = errors.New("bad command format")
				break
			}
			cid := l[1]
			_, ok := authMap[cid]
			if !ok {
				authMap[cid] = &authSession{
					cid: cid,
				}
			}
			a := authMap[cid]
			a.lasttime = time.Now()
			a.mech = l[2]
			if a.mech == "login" {
				if !a.userSent {
					dst := base64.StdEncoding.EncodeToString([]byte("Username:"))
					output = []byte("CONT" + "\t" + cid + "\t" + dst + "\n")
					if _, err = c.Write(output); err == nil {
						a.userSent = true
					}
				}
			} else if strings.ToLower(a.mech) == "plain" {
				var resp string
				for _, e := range l[2:] {
					f := strings.SplitN(e, "=", 2)
					if len(f) == 2 && strings.ToLower(f[0]) == "resp" {
						resp = f[1]
					}
				}
				if resp == "" {
					err = errors.New("resp not found")
				} else {
					if deResp, derr := base64.StdEncoding.DecodeString(resp); derr == nil {
						if x := strings.SplitN(string(deResp), "\000", 3); len(x) == 3 {
							a.username = x[1]
							a.password = x[2]
						}
					}
				}
				if len(a.username) > 0 && len(a.password) > 0 {
					if authenticate(string(a.username), string(a.password)) {
						c.Write([]byte(fmt.Sprintf("OK\t%s\tuser=%s\n", cid, a.username)))
					} else {
						err = errors.New("authenticate failed")
					}
				}
				if err != nil {
					output = []byte("FAIL" + "\t" + string(cid) + "\n")
					c.Write(output)
				}
			}
		} else if strings.HasPrefix(line, "CONT") {
			l := strings.Split(line, "\t")
			if len(l) < 2 {
				err = errors.New("bad command format")
				break
			}
			cid := l[1]
			_, ok := authMap[cid]
			if !ok {
				authMap[cid] = &authSession{
					cid: cid,
				}
			}
			a := authMap[cid]
			a.lasttime = time.Now()
			arg := strings.Trim(l[2], "\n")
			if a.cid == l[1] && a.mech == "login" {
				if a.userSent == true && a.pwdSent == false {
					a.username = arg
					dst := base64.StdEncoding.EncodeToString([]byte("Password:"))
					output = []byte("CONT" + "\t" + cid + "\t" + dst + "\n")
					if _, err = c.Write(output); err == nil {
						a.pwdSent = true
					}
				} else if a.userSent == true && a.pwdSent == true {
					//check username and password
					a.password = arg
					deUsername, _ := base64.StdEncoding.DecodeString(a.username)
					dePassword, _ := base64.StdEncoding.DecodeString(a.password)
					if len(deUsername) > 0 && len(dePassword) > 0 {
						if authenticate(string(deUsername), string(dePassword)) {
							c.Write([]byte(fmt.Sprintf("OK\t%s\tuser=%s\n", cid, deUsername)))
						} else {
							err = errors.New("authenticate failed")
						}
					} else {
						err = errors.New("decode username or password failed")
					}
				}
			}
			if err != nil {
				output = []byte("FAIL" + "\t" + string(cid) + "\n")
				c.Write(output)
			}
		}
	}
	fmt.Printf("close auth from %v. \n", c.RemoteAddr())
}

func main() {
	server, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Panic("Fail to start server", err)
	}
	log.Println("Server Started ...")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Printf("Fail to process connect, %s\n", err)
			break
		}
		go connHandler(conn)
	}
}
