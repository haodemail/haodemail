package models

import (
	"errors"
	"regexp"
	"time"
)

// VirtualDomain master dovecot domain table
type VirtualDomain struct {
	ID         int       `gorm:"AUTO_INCREMENT;primary_key" json:"domainID"`
	Name       string    `gorm:"type:varchar(120);unique_index" json:"domainName"`
	CreateTime time.Time `sql:"DEFAULT:current_timestamp" json:"createTime"`
	ExpireTime time.Time `json:"expireTime"`
	UserCount  int       `sql:"DEFAULT:0" json:"userCount"`
	Active     bool      `json:"active"`
}

func validateDomainSyntax(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]{0,1}\.([a-zA-Z]{1,6}|[a-zA-Z0-9-]{1,
30}\.[a-zA-Z]{2,3})$`)
	return re.MatchString(s)
}

func validateUserNameSyntax(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{1,100}$`)
	return re.MatchString(s)
}

// DaoGetVirtualDomainByName get virtual domain object by domain name
func DaoGetVirtualDomainByName(domainName string) (domain VirtualDomain, err error) {
	var db = DB
	domain = VirtualDomain{}
	err = db.Where("name = ?", domainName).First(&domain).Error
	return domain, err
}

// DaoGetVirtualDomainByID get virtual domain object by id
func DaoGetVirtualDomainByID(pk uint) (domain VirtualDomain, err error) {
	var db = DB
	err = db.First(&domain, pk).Error
	return domain, err
}

// CreateVirtualDomain new domain
func DaoCreateVirtualDomain(domainName string, expireTime time.Time) (domain VirtualDomain, err error) {
	if !validateDomainSyntax(domainName) {
		return domain, errors.New("invalid domain")
	}
	domainRe := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]{0,1}\.([a-zA-Z]{1,6}|[a-zA-Z0-9-]{1,30}\.[a-zA-Z]{2,3})$`)
	if !domainRe.MatchString(domainName) {
		return domain, errors.New("invalid domain name")
	}
	db := DB
	var count int
	db.Model(&domain).Where("name=?", domainName).Count(&count)
	if count > 0 {
		return domain, errors.New("domain exists")
	}
	tx := db.Begin()
	domain.Name = domainName
	domain.Active = true
	domain.CreateTime = time.Now()
	domain.ExpireTime = expireTime
	domain.UserCount = 0
	if err := tx.Create(&domain).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return
}
