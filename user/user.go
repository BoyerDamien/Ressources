package user

import (
	"fmt"
	"time"

	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database"
	"golang.org/x/crypto/bcrypt"
)

// User
//
// swagger:model
type User struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Nom de l'utilisateur
	// required: true
	FirstName string `json:"first_name"`

	// Prénom de l'utilisateur
	// required: true
	LastName string `json:"last_name"`

	// Mot de passe de l'utilisateur
	// required: true
	Password string `json:"password" validate:"required"`

	// Age de l'utilisateur
	// required: false
	// min: 0
	// max: 130
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
	// pattern: " customer | admin | user"
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

// swagger:operation GET /user/{id} User RetrieveUser
//
// Retourne des informations détaillées sur un utilisateur
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: l'email de l'utilisateur
//   required: true
//   type: string
// responses:
//   '200':
//     description: Retourne un utilisateur
//     schema:
//         "$ref": "#/definitions/User"
//   '404':
//     description: StatusNotFound
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '400':
//     description: StatusBadRequest
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '500':
//     description: StatusInternalServerError
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   default:
//     description: Erreur
//     schema:
//       "$ref": "#/definitions/ErrResponse"
func (s *User) Retrieve(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	u := new(User)
	r := db.Where("Email = ?", c.Params("id")).First(u)
	*s = *u
	return r, nil
}

// swagger:operation PUT /user User Update
//
// Modifie un utilisateur existant
//
// ---
// produces:
// - application/json
// consume:
// - application/json
// parameters:
// - name: user
//   in: body
//   description: Données de l'utilisateur
//   schema:
//       "$ref": "#/definitions/User"
// responses:
//   '200':
//     description: Retourne l'utilisateur modifié
//     schema:
//         "$ref": "#/definitions/User"
//   '400':
//     description: StatusBadRequest
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '500':
//     description: StatusInternalServerError
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   default:
//     description: Erreur
//     schema:
//       "$ref": "#/definitions/ErrResponse"
func (s *User) Update(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	res := db.Model(s).Omit("Email", "Role", "Password").Updates(s)
	if res.Error != nil {
		return res, nil
	}
	return db.Where("Email = ?", s.Email).First(s), nil
}

// swagger:operation POST /user User CreateUser
//
// Créé un nouvel utilisateur
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// parameters:
// - name: user
//   in: body
//   description: Données de l'utilisateur
//   schema:
//       "$ref": "#/definitions/User"
// responses:
//   '200':
//     description: Retourne l'utilisateur créé
//     schema:
//         "$ref": "#/definitions/User"
//   '500':
//     description: StatusInternalServerError
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '400':
//     description: StatusBadRequest
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   default:
//     description: Erreur
//     schema:
//       "$ref": "#/definitions/ErrResponse"
func (s *User) Create(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	if s.Password == "" {
		return db, fmt.Errorf("no password")
	}
	return db.FirstOrCreate(s, s), nil
}

// swagger:operation DELETE /user/{id} User DeleteUser
//
// Supprime un utilisateur existant
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: email de l'utilisateur
//   required: true
//   type: string
// responses:
//   '200':
//     description: Valide la suppression
//   '202':
//     description: StatusAccepted
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '500':
//     description: StatusInternalServerError
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   default:
//     description: Erreur
//     schema:
//       "$ref": "#/definitions/ErrResponse"
func (s *User) Delete(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Where("Email = ?", c.Params("id")).Delete(s), nil
}

// swagger:operation DELETE /users User DeleteUserList
//
// Supprime une liste d'utilisateurs
//
// ---
// produces:
// - application/json
// parameters:
// - name: emails
//   in: query
//   description: Liste de mails
//   required: true
//   type: array
//   items:
//       type: string
// responses:
//   '200':
//     description: Valide la suppression
//   '400':
//     description: StatusBadRequest
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '202':
//     description: StatusAccepted
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '500':
//     description: StatusInternalServerError
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   default:
//     description: Erreur
//     schema:
//       "$ref": "#/definitions/ErrResponse"
func (s *User) DeleteListQuery() gapi.Query {
	return &UserDeleteQuery{}
}

// swagger:operation GET /users User UserList
//
// Retourne des informations détaillées sur une liste d'utilisateurs
// ---
// produces:
// - application/json
// parameters:
// - name: role
//   in: query
//   description: Permet le filtre par role
//   required: false
//   type: string
//   pattern: " user | admin | customer"
// - name: orderBy
//   description: Permet de trier les résultats par champs
//   pattern: " first_name | last_name | age | adress | email | created_at | updated_at"
//   type: string
//   in: query
//   required: false
// - name: limit
//   description: Limite le nombre de résultats au nombre passé en paramètre
//   type: number
//   in: query
// - name: offset
//   description: Filtre les résultats a partir de l'index passé en paramètre
//   type: number
//   in: query
// responses:
//   '200':
//     description: Retourne une liste d'utilisateurs
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/User"
//   '404':
//     description: StatusNotFound
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '400':
//     description: StatusBadRequest
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   '500':
//     description: StatusInternalServerError
//     schema:
//       "$ref": "#/definitions/ErrResponse"
//   default:
//     description: Erreur
//     schema:
//       "$ref": "#/definitions/ErrResponse"
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
