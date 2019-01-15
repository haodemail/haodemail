package main

import (
	"./models"
	"testing"
	"time"
)

func TestCreateVirtualDomain(t *testing.T) {
	_, err := models.DaoCreateVirtualDomain("163.com", "1234AbcdR", 100, 1000000000, 100000, time.Now().Add(time.Hour*time.Duration(24*365*10)))
	if err == nil {
		t.Log("create domain 163.com ok")
	} else {
		t.Fatal("create domain failed", err)
	}
}

func TestCreateVirtualUser(t *testing.T) {
	if domain, err := models.DaoGetVirtualDomainByName("haodemail.com"); err != nil {
		t.Fatal("domain not exist")
	} else {
		if _, err := models.DaoCreateVirtualUser(domain, "admin", "postfix123", "Administrator",
			time.Now().Add(time.Hour*time.Duration(24*365*10))); err != nil {
			t.Fatal("create admin failed", err)
		}
	}
}
