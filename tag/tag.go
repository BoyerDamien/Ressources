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

func (s *Tag) Retrieve(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Where("Name = ?", c.Params("id")).First(s), nil
}

func (s *Tag) Create(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.FirstOrCreate(s, s), nil
}

func (s *Tag) Delete(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Where("Name = ?", c.Params("id")).Delete(s), nil
}

func (s *Tag) DeleteListQuery() gapi.Query {
	return &TagDeleteQuery{}
}

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
