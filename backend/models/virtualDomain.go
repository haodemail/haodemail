package models

import (
	"../helper"
	"errors"
	"regexp"
	"strconv"
	"time"
)

// VirtualDomain master dovecot domain table
type VirtualDomain struct {
	ID           int       `gorm:"AUTO_INCREMENT;primary_key" json:"domainID"`
	Name         string    `gorm:"type:varchar(120);unique_index" json:"domainName"`
	CreateTime   time.Time `sql:"DEFAULT:current_timestamp" json:"createTime"`
	ExpireTime   time.Time `json:"expireTime"`
	MaxUserCount int       `sql:"DEFAULT:0" json:"maxUserCount"`
	UserCount    int       `sql:"DEFAULT:0" json:"userCount"`
	MaxUserQuota int       `sql:"DEFAULT:0" json:"maxUserQuota"`
	UserQuota    int       `sql:"DEFAULT:0" json:"userQuota"`
	MaxMailCount int       `sql:"DEFAULT:0" json:"maxMailCount"`
	MailCount    int       `sql:"DEFAULT:0" json:"mailCount"`
	Active       bool      `json:"active"`
}

type VirtualDomainView struct {
	ID           string    `json:"id"`
	Domain       string    `json:"domain"`
	CreateTime   time.Time `json:"createTime"`
	ExpireTime   time.Time `json:"expireTime"`
	MaxUserCount int       `json:"maxUserCount"`
	UserCount    int       `json:"userCount"`
	MaxUserQuota int       `json:"maxUserQuota"`
	UserQuota    int       `json:"userQuota"`
	MaxMailCount int       `json:"maxMailCount"`
	MailCount    int       `json:"mailCount"`
	Active       bool      `json:"active"`
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
func DaoCreateVirtualDomain(domainName string, password string, maxUserCount int, maxUserQuota int, maxMailCount int, expireTime time.Time) (domain VirtualDomain, err error) {
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
	domain.MaxUserCount = maxUserCount
	domain.UserQuota = 0
	domain.MaxUserQuota = maxUserQuota
	domain.MailCount = 0
	domain.MaxMailCount = maxMailCount
	if err := tx.Create(&domain).Error; err != nil {
		tx.Rollback()
		return domain, err
	}
	// create postmaster in this domain, identify with password applied.
	user := VirtualUser{
		DomainID:   domain.ID,
		UserName:   "postmaster",
		Password:   HashAndSalt(password),
		NickName:   "postmaster",
		Active:     true,
		CreateTime: time.Now(),
		ExpireTime: expireTime,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return domain, err
	}
	tx.Commit()
	return
}

// DaoListVirtualDomain list virtual domain by domain name and order given column
func DaoListVirtualDomain(name string, orderBy string) (domainViews []VirtualDomainView, err error) {
	domains := []VirtualDomain{}
	db := DB
	db = db.Model(&VirtualDomain{})
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	if orderBy == "" {
		orderBy = "create_time desc"
	}
	db = db.Find(&domains).Order(orderBy)
	err = db.Error
	if err != nil {
		return
	}
	for _, d := range domains {
		var sid string
		var e error
		if sid, e = helper.EncryptPrimaryKey(strconv.Itoa(d.ID)); e != nil {
			sid = strconv.Itoa(d.ID)
		}
		domainViews = append(domainViews, VirtualDomainView{
			ID:           sid,
			Domain:       d.Name,
			CreateTime:   d.CreateTime,
			ExpireTime:   d.ExpireTime,
			MaxUserCount: d.MaxUserCount,
			UserCount:    d.UserQuota,
			MaxUserQuota: d.MaxUserQuota,
			UserQuota:    d.UserQuota,
			MaxMailCount: d.MaxMailCount,
			MailCount:    d.MailCount,
			Active:       d.Active,
		})
	}
	return
}
