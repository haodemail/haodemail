package models

import (
	"errors"
	"time"
)

// VirtualUser system user for partner application, differnet to postfix user
type VirtualAlias struct {
	ID          uint `gorm:"AUTO_INCREMENT;primary_key" json:"userid"`
	Domain      VirtualDomain
	DomainID    int
	Source      string    `gorm:"type:varchar(100)" json:"source"`
	Destination string    `gorm:"type:varchar(100)" json:"destination"`
	CreateTime  time.Time `sql:"DEFAULT:current_timestamp" json:"createTime"`
	ExpireTime  time.Time `json:"expireTime"`
	LastLogin   time.Time `json:"lastLogin"`
	Active      bool      `json:"active"`
}

// DaoGetVirtualAliasByName get system user object by username
func DaoGetVirtualAliasByName(username string) (user VirtualAlias, err error) {
	var db = DB
	err = db.Where("destination = ?", username).First(&user).Error
	return user, err
}

// DaoGetVirtualUserByID get system user object by user id
func DaoGetVirtualAliasByID(pk uint) (user VirtualAlias, err error) {
	var db = DB
	err = db.First(&user, pk).Error
	return user, err
}

func DaoCreateVirtualAlias(domain VirtualDomain, source string, destination string,
	expireTime time.Time) (user VirtualAlias,
	err error) {
	if !validateUserNameSyntax(source) {
		return user, errors.New("invalid source")
	}
	if !validateUserNameSyntax(destination) {
		return user, errors.New("invalid destination")
	}
	db := DB
	if user, err := DaoGetVirtualAliasByName(destination); err != nil {
		return user, errors.New("alias exists")
	}

	tx := db.Begin()
	user.Domain = domain
	user.Source = source
	user.Destination = destination
	user.Active = true
	user.CreateTime = time.Now()
	user.ExpireTime = expireTime
	if err := tx.Create(&domain).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return
}
