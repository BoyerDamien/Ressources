package user

import (
	"fmt"
	"time"

	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User
//
// swagger:model
type User struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Nom de l'utilisateur
	// required: true
	FirstName string `json:"firstName"`

	// Prénom de l'utilisateur
	// required: true
	LastName string `json:"lastName"`

	// Mot de passe de l'utilisateur
	// required: true
	Password string `json:"password" validate:"required"`

	// Age de l'utilisateur
	// required: false
	Age uint8 `json:"age" validate:"gte=0,lte=130"`

	// Email de l'utilisateur
	// required: true
	Email string `json:"email" validate:"required,email" gorm:"primaryKey"`

	// Address de l'utilisateur
	// require: false
	Address string `json:"address"`

	// Numéros de téléphone de l'utilisateur
	// require: false
	Phone string `json:"phone"`

	// Role de l'utilisateur
	// pattern: " customer | admin"
	// required: true
	Role string `json:"role" validate:"required,eq=admin|eq=customer|eq=user"`
}

func (s *User) BeforeCreate(tx *database.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	s.Password = string(bytes)
	return
}

func (s *User) AfterCreate(tx *database.DB) (err error) {
	s.Password = ""
	return
}

func (s *User) AfterFind(tx *database.DB) (err error) {
	s.Password = ""
	return
}

func (s *User) AfterUpdate(tx *database.DB) (err error) {
	s.Password = ""
	return
}

func (s *User) Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Email = ?", c.Params("id")).First(s), nil
}

func (s *User) Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	res := db.Model(s).Omit("Email", "Role", "Password").Updates(s)
	if res.Error != nil {
		return res, nil
	}
	return db.Where("Email = ?", s.Email).First(s), nil
}

func (s *User) Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	if len(s.Password) == 0 {
		return db, fmt.Errorf("no password")
	}
	return db.FirstOrCreate(s, s), nil
}

func (s *User) Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Email = ?", c.Params("id")).Delete(s), nil
}

func (s *User) DeleteListQuery() gapi.Query {
	return &UserDeleteQuery{}
}

func (s *User) ListQuery() gapi.Query {
	return &UserListQuery{}
}

type UserListQuery struct {
	ToFind  string `query:"tofind"`
	Role    string `query:"role" validate:"omitempty,eq=admin|eq=customer|eq=user"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=first_name|eq=last_name|eq=age|eq=address|eq=email"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *UserListQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {

	users := new([]User)
	tmp := db.Model(&User{})

	if s.Limit > 0 {
		tmp = tmp.Limit(s.Limit)
	}
	if s.Offset > 0 {
		tmp = tmp.Offset(s.Offset)
	}
	if len(s.ToFind) > 0 {
		tmp = tmp.Where("Email LIKE ?", "%"+s.ToFind+"%").Or("first_name LIKE ?", "%"+s.ToFind+"%").Or("last_name LIKE ?", "%"+s.ToFind+"%")
	}
	if len(s.Role) > 0 {
		tmp = tmp.Where("Role = ?", s.Role)
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Find(users)
	return result, users
}

type UserDeleteQuery struct {
	Emails []string `query:"emails"`
}

func (s *UserDeleteQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {
	var users []User

	if result := db.Where("Email IN ?", s.Emails).Find(&users); result.Error != nil {
		return result, nil
	}
	return db.Delete(&users, s.Emails), nil
}
