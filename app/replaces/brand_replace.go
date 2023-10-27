package replaces

import (
	"item-server/app/requests"
	"item-server/pkg/helpers"
	"time"
)

type BrandIndex struct {
	State    uint8  `filter:"state,eq"`
	Public   uint8  `filter:"is_public,eq"`
	Title    string `filter:"title,like"`
	Category []uint64
	BetTime  []int64 `filter:"created_at,bet_time"`
}

func (b *BrandIndex) BrandIndexReplace(req *requests.BrandFilterRequest) error {
	b.Title = req.Title
	b.State = req.State
	b.Public = req.Public
	b.BetTime = BetTime(req.BetTime)
	b.Category = IdString(req.CategoryId)
	return nil
}

type BrandStore struct {
	IsPublic uint8
	Ids      []uint64
}

func (s *BrandStore) BrandStoreReplace(request *requests.BrandRequest) error {
	ids, isPublic := IdPublicReplace(request.Category)
	s.IsPublic = isPublic
	s.Ids = ids
	return nil
}

type UpdateReplace struct {
	IsPublic    uint8
	UpdatedAt   time.Time
	Ids         []uint64
	SelectSlice []string
}

func (u *UpdateReplace) BrandUpdateReplace(request *requests.BrandRequest) error {
	ids, isPublic := IdPublicReplace(request.Category)
	u.IsPublic = isPublic
	u.UpdatedAt = helpers.TimeNow()
	u.Ids = ids
	u.SelectSlice = []string{
		"state",
		"title",
		"description",
		"icon_url",
		"sort",
		"is_public",
		"updated_at",
	}
	return nil
}
