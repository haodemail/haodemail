package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"../helper"
)

// VirtualUser system user for partner application, differnet to postfix user
// uid, gid, home properties will load from global configure
type VirtualUser struct {
	ID         uint          `gorm:"AUTO_INCREMENT;primary_key" json:"userid"`
	UserName   string        `gorm:"type:varchar(50);unique_index" json:"userName"`
	Domain     VirtualDomain `gorm:"ForeignKey:ProfileRefer"`
	DomainID   int
	Password   string    `gorm:"type:varchar(64)" json:"-"`
	NickName   string    `gorm:"type:varchar(50)" json:"nickName"`
	CreateTime time.Time `sql:"DEFAULT:current_timestamp" json:"createTime"`
	ExpireTime time.Time `sql:"DEFAULT:current_timestamp" json:"expireTime"`
	LastLogin  time.Time `sql:"DEFAULT:current_timestamp" json:"lastLogin"`
	Active     bool      `json:"active"`
}

// DaoGetVirtualAliasByName get system user object by username
func DaoGetVirtualUserByName(mailbox string) (user VirtualUser, err error) {
	username, domain, err := helper.SplitMailbox(mailbox)
	if err != nil {
		return
	}
	vdomain, err := DaoGetVirtualDomainByName(domain)
	if err != nil {
		return
	}
	var db = DB
	err = db.Where("domain_id=? AND user_name = ?", vdomain.ID, username).First(&user).Error
	return user, err
}

// DaoGetVirtualUserByID get system user object by user id
func DaoGetVirtualUserByID(pk uint) (user VirtualUser, err error) {
	var db = DB
	err = db.First(&user, pk).Error
	return user, err
}

// HashAndSalt make password with salt
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func DaoCreateVirtualUser(domain VirtualDomain, userName string, password string, nickName string,
	expireTime time.Time) (user VirtualUser, err error) {
	if !validateUserNameSyntax(userName) {
		return user, errors.New("invalid mailbox")
	}
	db := DB
	if _, err := DaoGetVirtualUserByName(userName); err == nil {
		return user, errors.New("user exists")
	}
	if _, err := DaoGetVirtualAliasByName(userName); err == nil {
		return user, errors.New("alias with that name exists")
	}
	tx := db.Begin()
	fmt.Println("domain is", domain)
	user.DomainID = domain.ID
	user.UserName = userName
	user.Password = HashAndSalt(password)
	user.NickName = nickName
	user.Active = true
	user.CreateTime = time.Now()
	user.ExpireTime = expireTime
	if err := tx.Create(&user).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	tx.Commit()
	return
}
