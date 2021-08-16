package offer

import (
	"time"

	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database"
	"github.com/BoyerDamien/ressources/tag"
)

// Offer
//
// swagger:model
type Offer struct {
	// Base model
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Nom de l'offre
	// required: true
	Name string `json:"name" validate:"required" gorm:"primaryKey"`

	// Description de l'offre
	// required: true
	Description string `json:"description" validate:"required"`

	// Tags liés l'offre
	// required: true
	Tags []tag.Tag `json:"tags" gorm:"many2many:offer_tags;constraint:OnUpdate:CASCADE;References:Name" validate:"required,dive"`
}

func (s *Offer) BeforeCreate(tx *database.DB) error {
	if err := gapi.Validate(s); err != nil {
		return err
	}
	return nil
}

func (s *Offer) BeforeUpdate(tx *database.DB) error {
	return s.BeforeCreate(tx)
}

func (s *Offer) Retrieve(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Where("Name = ?", c.Params("id")).Preload("Tags").First(s), nil
}

func (s *Offer) Update(c *gapi.Ctx, db *database.DB) (*database.DB, error) {

	if err := db.Model(s).Association("Tags").Replace(s.Tags); err != nil {
		return nil, err
	}

	res := db.Model(s).Updates(s)
	if res.Error != nil {
		return res, nil
	}
	return db.Where("Name = ?", s.Name).Preload("Tags").First(s), nil
}

func (s *Offer) Create(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.FirstOrCreate(s, Offer{Name: s.Name}), nil
}

func (s *Offer) Delete(c *gapi.Ctx, db *database.DB) (*database.DB, error) {
	return db.Where("Name = ?", c.Params("id")).Delete(s), nil
}

func (s *Offer) DeleteListQuery() gapi.Query {
	return &OfferDeleteQuery{}
}

func (s *Offer) ListQuery() gapi.Query {
	return &OfferListQuery{}
}

type OfferListQuery struct {
	ToFind  string `query:"tofind" validate:"omitempty"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=name"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *OfferListQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {

	offers := new([]Offer)
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
	result := tmp.Preload("Tags").Find(offers)
	return result, offers
}

type OfferDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *OfferDeleteQuery) Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{}) {
	var offers []Offer

	if result := db.Where("Name IN ?", s.Names).Find(&offers); result.Error != nil {
		return result, nil
	}
	return db.Delete(&offers, s.Names), nil
}
