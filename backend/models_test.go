package main

import (
	"./helper"
	"./models"
	"fmt"
	"testing"
	"time"
)

func TestCreateVirtualDomain(t *testing.T) {
	_, err := models.DaoCreateVirtualDomain("", "163.com", "1234AbcdR", 100, 1000000000, 100000, time.Now().Add(time.Hour*time.Duration(24*365*10)))
	if err == nil {
		t.Log("create domain 163.com ok")
	} else {
		t.Fatal("create domain failed", err)
	}
}

func TestCreateVirtualUser(t *testing.T) {
	did, _ := helper.EncryptPrimaryKey("1")
	if _, err := models.DaoCreateVirtualUser(did, "", "admin", "postfix123", "Administrator", -1, -1, time.Now().Add(time.Hour*time.Duration(24*365*10))); err != nil {
		t.Fatal("create admin failed", err)
	}
}

func TestDecrypt(t *testing.T) {
	//s := helper.UnQuote("a5tBRmz7u7VR%2BXOoW%2FRgpg%3D%3D")
	fmt.Println(helper.DecryptPrimaryKey("Bb4M0IeJRp0DZRwDS7ktcx+JZfmjJQoC1SEhEIp85UA=="))
}

func TestPassword(t *testing.T) {
	fmt.Println(models.HashAndSalt("postfix123"))
}
