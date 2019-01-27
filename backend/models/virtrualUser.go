package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"../helper"
)

// VirtualUser system user for partner application, differnet to postfix user
// uid, gid, home properties will load from global configure
type VirtualUser struct {
	ID           uint          `gorm:"AUTO_INCREMENT;primary_key" json:"userid"`
	UserName     string        `gorm:"type:varchar(50);unique_index" json:"userName"`
	Domain       VirtualDomain `gorm:"ForeignKey:ProfileRefer"`
	DomainID     int
	Password     string    `gorm:"type:varchar(64)" json:"-"`
	NickName     string    `gorm:"type:varchar(50)" json:"nickName"`
	CreateTime   time.Time `sql:"DEFAULT:current_timestamp" json:"createTime"`
	ExpireTime   time.Time `sql:"DEFAULT:current_timestamp" json:"expireTime"`
	MaxQuota     int       `sql:"DEFAULT:-1" json:"maxUserQuota"`
	UserQuota    int       `sql:"DEFAULT:0" json:"userQuota"`
	MaxMailCount int       `sql:"DEFAULT:-1" json:"maxMailCount"`
	MailCount    int       `sql:"DEFAULT:0" json:"mailCount"`
	LastLogin    time.Time `sql:"DEFAULT:current_timestamp" json:"lastLogin"`
	Active       bool      `json:"active"`
	Deleted      bool      `json:"deleted"`
	DeleteTime   time.Time `json:"deleteTime"`
}

type VirtualUserView struct {
	ID           string    `json:"id"`
	Domain       string    `json:"domain"`
	DomainID     string    `json:"domainID"`
	UserName     string    `json:"userName"`
	NickName     string    `json:"nickName"`
	CreateTime   time.Time `json:"createTime"`
	ExpireTime   time.Time `json:"expireTime"`
	LastLogin    time.Time `json:"lastLogin"`
	MaxQuota     int       `json:"maxQuota"`
	UserQuota    int       `json:"userQuota"`
	MaxMailCount int       `json:"maxMailCount"`
	MailCount    int       `json:"mailCount"`
	Active       bool      `json:"active"`
	IsExpired    bool      `json:"expired"`
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

func ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

func DaoCreateVirtualUser(did string, uid string, userName string, password string, nickName string, maxQuota int, maxMailCount int,
	expireTime time.Time) (user VirtualUser, err error) {
	var domainID, userID int
	if did != "" {
		if s, e := helper.DecryptPrimaryKey(did); e == nil {
			domainID, err = strconv.Atoi(s)
			if err != nil {
				return
			}
		} else {
			return user, e
		}
	}
	if uid != "" {
		if s, e := helper.DecryptPrimaryKey(uid); e == nil {
			userID, err = strconv.Atoi(s)
			if err != nil {
				return
			}
		} else {
			return user, e
		}
	}
	domain, err := DaoGetVirtualDomainByID(uint(domainID))
	if err != nil {
		return user, err
	}
	if !validateUserNameSyntax(userName) {
		return user, errors.New("invalid mailbox")
	}
	db := DB
	tx := db.Begin()
	if userID > 0 {
		if user, err = DaoGetVirtualUserByID(uint(userID)); err != nil {
			return user, errors.New("user not exist")
		}
		update := map[string]interface{}{}
		// update only
		if password != "" {
			update["Password"] = HashAndSalt(password)
		}
		update["NickName"] = nickName
		update["MaxQuota"] = maxQuota
		update["MaxMailCount"] = maxMailCount
		if !user.Active {
			update["Active"] = true
			if err = tx.Model(&domain).Where("id=?", domain.ID).Update(map[string]interface{}{"UserCount": domain.UserCount + 1}).Error; err != nil {
				tx.Rollback()
				return
			}
		}
		update["ExpireTime"] = expireTime
		if err := tx.Model(&user).Where("id=?", userID).Update(update).Error; err != nil {
			tx.Rollback()
			return VirtualUser{}, err
		}
	} else {
		if _, err := DaoGetVirtualUserByName(userName); err == nil {
			return user, errors.New("user exist")
		}
		if _, err := DaoGetVirtualAliasByName(userName); err == nil {
			return user, errors.New("alias with that name exists")
		}
		user.DomainID = domain.ID
		user.UserName = userName
		user.Password = HashAndSalt(password)
		user.NickName = nickName
		user.MaxQuota = maxQuota
		user.MaxMailCount = maxMailCount
		user.Active = true
		user.CreateTime = time.Now()
		user.ExpireTime = expireTime
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return VirtualUser{}, err
		}
		domain.UserCount++
		if err = tx.Save(&domain).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

// DaoDeleteVirtualUser delete user
func DaoDeleteVirtualUser(uid string) (user VirtualUser, err error) {
	var userID int
	if s, e := helper.DecryptPrimaryKey(uid); e == nil {
		userID, err = strconv.Atoi(s)
		if err != nil {
			return
		}
	} else {
		return user, e
	}
	user, err = DaoGetVirtualUserByID(uint(userID))
	if err != nil {
		return user, err
	}
	db := DB
	tx := db.Begin()
	user.Deleted = true
	user.DeleteTime = time.Now()
	err = tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		return
	}
	user.Domain.UserCount--
	if err = tx.Save(&user.Domain).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// DaoListVirtualDomain list virtual user by domain name and order given column
func DaoListVirtualUser(did string, name string, orderBy string) (userViews []VirtualUserView, err error) {
	var domainID int
	if s, e := helper.DecryptPrimaryKey(did); e == nil {
		domainID, err = strconv.Atoi(s)
		if err != nil {
			return
		}
	} else {
		return userViews, e
	}
	domain, err := DaoGetVirtualDomainByID(uint(domainID))
	var users = make([]VirtualUser, 0)
	db := DB
	db = db.Model(&VirtualUser{}).Where("domain_id=? AND deleted=0", domain.ID)
	if name != "" {
		db = db.Where("user_name like ?", "%"+name+"%")
	}
	if orderBy == "" {
		orderBy = "create_time desc"
	}
	db = db.Find(&users).Order(orderBy)
	err = db.Error
	if err != nil {
		return
	}
	for _, u := range users {
		var uid string
		var e error
		if uid, e = helper.EncryptPrimaryKey(strconv.Itoa(int(u.ID))); e != nil {
			uid = strconv.Itoa(int(u.ID))
		}
		if did, e = helper.EncryptPrimaryKey(strconv.Itoa(int(domain.ID))); e != nil {
			uid = strconv.Itoa(int(u.ID))
		}
		var isExpired bool
		if u.ExpireTime.IsZero() {
			isExpired = false
		} else {
			if u.ExpireTime.Before(time.Now()) {
				isExpired = true
			}
		}
		userViews = append(userViews, VirtualUserView{
			ID:           uid,
			Domain:       domain.Name,
			DomainID:     did,
			UserName:     u.UserName,
			NickName:     u.NickName,
			CreateTime:   u.CreateTime,
			ExpireTime:   u.ExpireTime,
			LastLogin:    u.LastLogin,
			MaxQuota:     u.MaxQuota,
			UserQuota:    u.UserQuota,
			MaxMailCount: u.MaxMailCount,
			MailCount:    u.MailCount,
			Active:       u.Active,
			IsExpired:    isExpired,
		})
	}
	return
}
