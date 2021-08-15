package tag

import (
	"time"

	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database"
)

// Tag
//
// swagger:model
type Tag struct {
	// Base model
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Nom du Tag
	// required: true
	// example: #python
	Name string `json:"name" validate:"required,min=3,max=255" gorm:"primaryKey"`
}

// swagger:operation GET /tag/{id} Tag RetrieveTag
//
// Retourne des informations détaillées sur un tag
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: nom du tag
//   required: true
//   type: string
// responses:
//   '200':
//     description: Retourne un tag
//     schema:
//         "$ref": "#/definitions/Tag"
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
func (s *Tag) Retrieve(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	tag := new(Tag)
	res := db.Where("Name = ?", c.Params("id")).First(tag)
	*s = *tag
	return res, nil
}

// swagger:operation POST /tag Tag CreateTag
//
// Créé un nouveau tag
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// parameters:
// - name: tag
//   in: body
//   description: Données du tag
//   schema:
//       "$ref": "#/definitions/Tag"
// responses:
//   '200':
//     description: Retourne le tag
//     schema:
//         "$ref": "#/definitions/Tag"
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
func (s *Tag) Create(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.FirstOrCreate(s, s), nil
}

// swagger:operation DELETE /tag/{id} Tag DeleteTag
//
// Supprime un tag existant
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: nom tu tag
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
func (s *Tag) Delete(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Where("Name = ?", c.Params("id")).Delete(s), nil
}

// swagger:operation DELETE /tags Tag DeleteTagList
//
// Supprime une liste de tags
//
// ---
// produces:
// - application/json
// parameters:
// - name: names
//   in: query
//   description: Liste de noms
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
func (s *Tag) DeleteListQuery() gapi.Query {
	return &TagDeleteQuery{}
}

// swagger:operation GET /tags Tag TagList
//
// Retourne des informations détaillées sur une liste de tags
// ---
// produces:
// - application/json
// parameters:
// - name: toFind
//   in: query
//   description: Permet le filtre par nom. Retourne les tags pour lesquels le nom contient la chaîne de caractères à rechercher
//   required: false
//   type: string
// - name: orderBy
//   description: Permet de trier les résultats par champs
//   pattern: " name | created_at | updated_at"
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
//     description: Retourne une liste de tags
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Tag"
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
func (s *Tag) ListQuery() gapi.Query {
	return &TagListQuery{}
}

type TagListQuery struct {
	ToFind  string `query:"tofind"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=name"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *TagListQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {

	tags := new([]Tag)
	tmp := db

	if s.Limit > 0 {
		tmp = tmp.Limit(s.Limit)
	}
	if s.Offset > 0 {
		tmp = tmp.Offset(s.Offset)
	}
	if len(s.ToFind) > 0 {
		tmp = tmp.Where("Name LIKE ?", "%"+s.ToFind+"%")
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	return tmp.Find(tags), tags
}

type TagDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *TagDeleteQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {
	var tags []Tag

	if result := db.Where("Name IN ?", s.Names).Find(&tags); result.Error != nil {
		return result, nil
	}
	return db.Delete(&tags, s.Names), nil
}
