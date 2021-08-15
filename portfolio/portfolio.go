package portfolio

import (
	"time"

	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database"
	"github.com/BoyerDamien/ressources/media"
	"github.com/BoyerDamien/ressources/tag"
)

// Portfolio
//
// swagger:model
type PortFolio struct {
	// Base model
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Gallery
	// required: true
	Gallery []media.Media `json:"gallery" gorm:"many2many:portfolio_medias;" validate:"dive"`

	// Website
	// required: true
	Website string `json:"website" validate:"url,required"`

	// Name
	// required: true
	Name string `json:"name" validate:"required" gorm:"primaryKey"`

	// Description
	// required: true
	Description string `json:"description" validate:"required"`

	// Tags
	// required: true
	Tags []tag.Tag `json:"tags" gorm:"many2many:portfolios_tags;" validate:"required,dive"`
}

func (s *PortFolio) BeforeCreate(tx *database.DB) error {
	s.Gallery = []media.Media{}
	return nil
}

func (s *PortFolio) Retrieve(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	p := new(PortFolio)
	r := db.Where("Name = ?", c.Params("id")).Preload("Tags").Preload("Gallery").First(p)
	*s = *p
	return r, nil
}

func (s *PortFolio) Update(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	if err := db.Model(s).Association("Gallery").Replace(s.Gallery); err != nil {
		return nil, err
	}
	if err := db.Model(s).Association("Tags").Replace(s.Tags); err != nil {
		return nil, err
	}
	if res := db.Model(s).Updates(s); res.Error != nil {
		return res, nil
	}
	p := new(PortFolio)
	r := db.Where("Name = ?", s.Name).Preload("Tags").Preload("Gallery").First(p)
	*s = *p
	return r, nil
}

func (s *PortFolio) Create(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.FirstOrCreate(s, s), nil
}

func (s *PortFolio) Delete(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Model(s).Where("Name = ?", c.Params("id")).Delete(s), nil
}

func (s *PortFolio) DeleteListQuery() gapi.Query {
	return &PortFolioDeleteQuery{}
}

func (s *PortFolio) ListQuery() gapi.Query {
	return &PortFolioListQuery{}
}

type PortFolioListQuery struct {
	ToFind  string `query:"tofind"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=name"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *PortFolioListQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {

	portFolios := new([]PortFolio)
	tmp := db

	if s.Limit > 0 {
		tmp = tmp.Limit(s.Limit)
	}
	if s.Offset > 0 {
		tmp = tmp.Offset(s.Offset)
	}
	if len(s.ToFind) > 0 {
		tmp = tmp.Where("Name LIKE ?", "%"+s.ToFind+"%").Or("Description LIKE ?", "%"+s.ToFind+"%")
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Preload("Tags").Preload("Gallery").Find(portFolios)
	return result, portFolios
}

type PortFolioDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *PortFolioDeleteQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {
	var portFolios []PortFolio

	if result := db.Where("Name IN ?", s.Names).Find(&portFolios); result.Error != nil {
		return result, nil
	}
	return db.Delete(&portFolios, s.Names), nil
}
