package usermanger

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/harpy-wings/sol-hr/models"
	utilitymanager "github.com/harpy-wings/sol-hr/pkg/utilityManager"
	"github.com/kataras/iris/v12"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type userManager struct {
	db    *gorm.DB
	um    utilitymanager.IUtilityManager
	cache *cache.Cache
	acls  map[string]*models.MAcl
}

var _ IUserManager = &userManager{}

func (u *userManager) setDefaults() {
	if models.DB != nil {
		u.db = models.DB
	}
}

func (u *userManager) init() error {
	return nil
}

func (m *userManager) Login(ctx iris.Context, username string, password string) (u *models.MUser, token string, err error) {
	user := &models.MUser{}
	err = m.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, "", err
	}
	if user.FirstName == "" {
		return nil, "", errors.New("user not found")
	}
	secret := uuid.New().String()
	session := &models.MUserSession{
		UserID:    user.ID,
		Secret:    secret,
		ExpiresAt: time.Now().Add(time.Hour * 12),
		Device:    "N/A",
		IP:        ctx.RemoteAddr(),
		IsSA:      user.IsSA,
	}
	err = m.db.Create(&session).Error
	if err != nil {
		return nil, "", err
	}
	m.cache.Set(secret, session, time.Hour*12)
	m.cache.Set(fmt.Sprintf("user:%d", user.ID), user, time.Hour*12)

	return user, secret, nil
}
func (m *userManager) Acl(token string, fnSignature string, permissionType models.PermissionType) (u *models.MUser, err error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	if fnSignature != "" {
		err = m.initAcl(fnSignature)
		if err != nil {
			return nil, err
		}
	}

	if token == "" {
		return nil, errors.New("token is required")
	}

	Isession, ok := m.cache.Get(token)
	var session *models.MUserSession
	if !ok {
		err = m.db.Where("secret = ?", token).First(&session).Error
		if err != nil {
			return nil, err
		}
	} else {
		session = Isession.(*models.MUserSession)
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("session expired")
	}
	if session.LogoutAt != nil && session.LogoutAt.Before(time.Now()) {
		return nil, errors.New("session logged out")
	}

	user, err := m.GetUser(session.UserID)
	if err != nil {
		return nil, err
	}
	if fnSignature == "" {
		return user, nil
	}
	if user.IsSA {
		return user, nil
	}
	ACL := &models.UserAcl{}
	err = m.db.Where("user_id = ? AND acl_key = ?", user.ID, fnSignature).First(&ACL).Error
	if err != nil {
		return nil, err
	}
	if ACL.Value == 0 {
		return nil, errors.New("user does not have permission")
	}
	if ACL.Value&int64(permissionType) == 0 {
		return nil, errors.New("user does not have permission")
	}
	return user, nil
}

func (m *userManager) GetAcls() ([]*models.MAcl, error) {
	acls := []*models.MAcl{}
	err := m.db.Find(&acls).Error
	if err != nil {
		return nil, err
	}
	if m.acls == nil {
		m.acls = make(map[string]*models.MAcl)
	}
	for _, acl := range acls {
		m.acls[acl.Key] = acl
	}
	return acls, nil
}

func (m *userManager) Logout(token string) error {
	token = strings.Replace(token, "Bearer ", "", 1)
	if token == "" {
		return errors.New("token is required")
	}
	iSession, ok := m.cache.Get(token)
	var session *models.MUserSession
	if !ok {
		err := m.db.Where("secret = ?", token).First(&session).Error
		if err != nil {
			return err
		}
	} else {
		session = iSession.(*models.MUserSession)
	}
	err := m.db.Model(&session).Update("logout_at", time.Now()).Error
	if err != nil {
		return err
	}
	m.cache.Delete(token)
	return nil
}

func (m *userManager) ListUsers(req ListUsersRequest) ([]*models.MUser, error) {
	users := []*models.MUser{}
	err := m.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Model(&models.MUser{})
		if req.BranchID != 0 {
			tx = tx.Where("branch_id = ?", req.BranchID)
		}
		if req.Limit != 0 {
			tx = tx.Limit(int(req.Limit))
		}
		if req.Offset != 0 {
			tx = tx.Offset(int(req.Offset))
		}
		if req.Query != "" {
			tx = tx.Where("first_name LIKE ? OR last_name LIKE ? OR username LIKE ?", "%"+req.Query+"%", "%"+req.Query+"%", "%"+req.Query+"%")
		}
		if req.OrderBy != "" {
			tx = tx.Order(req.OrderBy)
		} else {
			tx = tx.Order("created_at DESC")
		}

		err := tx.Find(&users).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		rank, err := m.um.GetMilitaryRank(user.MilitaryRankID)
		if err != nil {
			return nil, err
		}
		user.MMilitaryRank = rank
	}
	return users, nil
}

func (m *userManager) GetUser(id int64) (*models.MUser, error) {
	user := &models.MUser{}
	err := m.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.Password = ""
	rank, err := m.um.GetMilitaryRank(user.MilitaryRankID)
	if err != nil {
		return nil, err
	}
	user.MMilitaryRank = rank
	return user, nil
}
func (m *userManager) CreateUser(user *models.MUser) error {
	rank, err := m.um.GetMilitaryRank(user.MilitaryRankID)
	if err != nil {
		return err
	}
	user.MMilitaryRank = rank
	userAcls := []*models.UserAcl{}
	for _, acl := range user.ACL {
		userAcls = append(userAcls, &models.UserAcl{
			UserID: user.ID,
			AclKey: acl.AclKey,
			Value:  acl.Value,
			AclID:  acl.AclID,
		})
	}
	err = m.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		if err != nil {
			return err
		}
		err = tx.Create(&userAcls).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *userManager) UpdateUser(user *models.MUser) error {
	editingUser := &models.MUser{ID: user.ID}

	err := m.db.First(&editingUser).Error
	if err != nil {
		return err
	}

	editingUser.UpdatedAt = time.Now()
	editingUser.Username = user.Username
	editingUser.Password = user.Password

	editingUser.Role = user.Role
	editingUser.IsSA = user.IsSA
	editingUser.MilitaryRankID = user.MilitaryRankID
	editingUser.FirstName = user.FirstName
	editingUser.LastName = user.LastName
	rank, err := m.um.GetMilitaryRank(editingUser.MilitaryRankID)
	if err != nil {
		return err
	}
	editingUser.MMilitaryRank = rank
	err = m.db.Where("id = ?", editingUser.ID).Updates(editingUser).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *userManager) UpdateUserPassword(id int64, password string) error {
	err := m.db.Model(&models.MUser{}).Where("id = ?", id).Update("password", password).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *userManager) DeleteUser(id int64) error {
	if id == 0 {
		return errors.New("امکان حذف کاربر اولیه سیستم وجود ندارد")
	}
	err := m.db.Where("id = ?", id).Delete(&models.MUser{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *userManager) GetUserAcls(id int64) ([]*models.MAcl, error) {
	var userAcls []*models.UserAcl
	err := m.db.Where("user_id = ?", id).Find(&userAcls).Error
	if err != nil {
		return nil, err
	}
	acls := []*models.MAcl{}
	for _, acl := range userAcls {
		acls = append(acls, m.acls[acl.AclKey])
	}
	return acls, nil
}

func (m *userManager) UpdateUserAcls(id int64, newAcls []*models.UserAcl) error {
	userAcls := []*models.UserAcl{}
	err := m.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id = ?", id).Delete(&models.UserAcl{}).Error
		if err != nil {
			return err
		}
		for _, acl := range newAcls {
			userAcls = append(userAcls, &models.UserAcl{
				UserID: id,
				AclKey: acl.AclKey,
				AclID:  acl.AclID,
				Value:  acl.Value,
			})
		}
		err = tx.CreateInBatches(&userAcls, 100).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// #########################################################
// ## 				Helper Functions
// #########################################################

func (m *userManager) initAcl(fnSignature string) error {
	if _, ok := m.acls[fnSignature]; ok {
		return nil
	}
	var acl models.MAcl

	err := m.db.Where("key = ?", fnSignature).First(&acl).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if acl.ID != 0 {
		m.acls[fnSignature] = &acl
		return nil
	}
	acl = models.MAcl{
		Key:   fnSignature,
		Title: fnSignature,
	}
	err = m.db.Create(&acl).Error
	if err != nil {
		return err
	}
	m.acls[fnSignature] = &acl
	return nil
}
