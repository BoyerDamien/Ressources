package media

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database"
)

// Media
//
// swagger:model
type Media struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Path string `json:"-"`

	// Nom du média
	// required: true
	// example: image.png
	Name string `json:"name" gorm:"primaryKey"`

	// Taille du fichier en ko
	// required: true
	// min: 1
	// example: 140
	Size uint `json:"file_size"`

	// Mime type du fichier
	// required: true
	// example: png
	Type string `json:"type"`

	// Url du fichier
	// required: false
	// example: /chemin/vers/le/fichier.png
	Url string `json:"url"`

	// Status du l'image
	// required: true
	// pattern: " protected | open"
	// example: open
	Status string `json:"status"`
}

func (s *Media) AfterDelete(tx *database.DB) (err error) {
	os.Remove(s.Path)
	return
}

func (s *Media) BeforeFind(tx *database.DB) (err error) {
	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory")
	}
	return
}

// swagger:operation GET /media/{id} Media RetrieveMedia
//
// Retourne des informations détaillées sur un média
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: le nom du média
//   required: true
//   type: string
// responses:
//   '200':
//     description: Retourne un média
//     schema:
//         "$ref": "#/definitions/Media"
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
func (s *Media) Retrieve(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	m := new(Media)
	r := db.Where("name = ?", c.Params("id")).First(m)
	*s = *m
	return r, nil
}

// swagger:operation PUT /media Media UpdateMedia
//
// Modifie un média existant
//
// ---
// produces:
// - application/json
// consume:
// - application/json
// parameters:
// - name: media
//   in: body
//   description: Données du média
//   schema:
//       "$ref": "#/definitions/Media"
// responses:
//   '200':
//     description: Retourne le média modifié
//     schema:
//         "$ref": "#/definitions/Media"
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
func (s *Media) Update(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	if s.Status == "open" || s.Status == "protected" {
		if res := db.Model(s).Where("Name = ?", s.Name).Update("status", s.Status); res.Error != nil {
			return db, nil
		}
		m := new(Media)
		r := db.Where("Name = ?", s.Name).First(m)
		s = m
		return r, nil
	}
	return nil, fmt.Errorf("wrong status")
}

// swagger:operation POST /media Media CreateMedia
//
// Créé un nouveau média
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// parameters:
// - name: media
//   in: formData
//   description: Contenu du média
//   type: file
// responses:
//   '200':
//     description: Retourne le média créé
//     schema:
//         "$ref": "#/definitions/Media"
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
func (s *Media) Create(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	media, err := c.FormFile("media")
	if err != nil {
		return nil, err
	}

	s.Type = media.Header.Get("Content-Type")
	s.Name = media.Filename
	s.Size = uint(media.Size)
	s.Path = path.Join(os.Getenv("MEDIA_PATH"), media.Filename)
	s.Url = path.Join(os.Getenv("STATIC_FILES_MEDIA_URL"), media.Filename)
	s.Status = "protected"

	result := db.FirstOrCreate(s, s)
	if result.Error != nil {
		return result, nil
	}

	if err := c.SaveFile(media, s.Path); err != nil {
		return db.Where("Name = ?", s.Name).Delete(s), err
	}
	return result, nil
}

// swagger:operation DELETE /media/{id} Media DeleteMedia
//
// Supprime un média existant
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: nom du média
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
//
func (s *Media) Delete(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Where("Name = ?", c.Params("id")).Delete(s), nil
}

// swagger:operation DELETE /medias Media DeleteMediaList
//
// Supprime une liste de médias
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
func (s *Media) DeleteListQuery() gapi.Query {
	return &MediaDeleteQuery{}
}

// swagger:operation GET /medias Media MediaList
//
// Retourne des informations détaillées sur une liste de médias
// ---
// produces:
// - application/json
// parameters:
// - name: status
//   in: query
//   description: Permet le filtre par status
//   required: false
//   type: string
//   pattern: " protected | open"
// - name: type
//   in: query
//   description: Permet le filtre par mime type
//   required: false
//   type: string
// - name: orderBy
//   description: Permet de trier les résultats par champs
//   pattern: " name | created_at | updated_at | size"
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
//     description: Retourne une liste de médias
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
func (s *Media) ListQuery() gapi.Query {
	return &MediaListQuery{}
}

type MediaListQuery struct {
	ToFind  string `query:"tofind" validate:"omitempty"`
	Type    string `query:"type" validate:"omitempty"`
	Status  string `query:"status" validate:"omitempty,eq=protected|eq=open"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=name|eq=size"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *MediaListQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {

	medias := new([]Media)
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
	if len(s.Status) > 0 {
		tmp = tmp.Where("Status = ?", s.Status)
	}
	if len(s.Type) > 0 {
		tmp = tmp.Where("Type = ?", s.Type)
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Find(medias)
	return result, medias
}

type MediaDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *MediaDeleteQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {
	var medias []Media

	if result := db.Where("Name IN ?", s.Names).Find(&medias); result.Error != nil {
		return result, nil
	}
	return db.Delete(&medias, s.Names), nil
}
