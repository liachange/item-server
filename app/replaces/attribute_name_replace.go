package replaces

import (
	"item-server/app/requests"
	"item-server/pkg/helpers"
	"item-server/pkg/pinyin"
	"time"
)

type AttributeNameIndex struct {
	State    uint8   `filter:"state,eq"`
	BetTime  []int64 `filter:"created_at,bet_time"`
	Genre    uint8   `filter:"genre,eq"`
	Title    string  `filter:"title,like"`
	Public   uint8   `filter:"is_public,eq"`
	Category []uint64
}

func (a *AttributeNameIndex) AttributeNameReplace(req *requests.AttributeNameFilterRequest) error {
	a.State = req.State
	a.Title = req.Title
	a.Genre = req.Genre
	a.BetTime = BetTime(req.BetTime)
	a.Category = IdString(req.Category)
	a.Public = req.Public
	return nil
}

type AttributeNameStore struct {
	Public   uint8
	Category []uint64
	Abbr     string
}

func (s *AttributeNameStore) AttributeNameStoreReplace(req *requests.AttributeNameRequest) error {
	ids, public := IdPublicReplace(req.Category)
	s.Public = public
	s.Category = ids
	s.Abbr = pinyin.GetFirstSpell(req.Title)
	return nil
}

type AttributeNameSave struct {
	IsPublic    uint8
	UpdatedAt   time.Time
	Ids         []uint64
	SelectSlice []string
	Abbr        string
}

func (a *AttributeNameSave) AttributeNameSaveReplace(req *requests.AttributeNameRequest) error {
	ids, public := IdPublicReplace(req.Category)
	a.Ids = ids
	a.IsPublic = public
	a.UpdatedAt = helpers.TimeNow()
	a.Abbr = pinyin.GetFirstSpell(req.Title)
	a.SelectSlice = []string{
		"state",
		"title",
		"genre",
		"description",
		"is_public",
		"abbr",
		"sort",
		"search",
		"updated_at",
	}

	return nil
}
