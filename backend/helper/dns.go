package helper

import (
	"errors"
	"net"
	"time"
)

func QueryDomainMX(domain string) (ipList []net.IP, err error) {
	var done = make(chan bool, 1)
	var timeout = make(chan bool, 1)

	go func() {
		time.Sleep(time.Second * time.Duration(1))
		timeout <- true
	}()

	go func() {
		mxs, err := net.LookupMX(domain)
		if err != nil {
			return
		}
		for _, mx := range mxs {
			if ips, err := net.LookupIP(mx.Host); err == nil {
				ipList = append(ipList, ips...)
			}
		}
		done <- true
	}()

	select {
	case <-done:

	case <-timeout:
		err = errors.New("query mx record timeout")
	}
	return
}
