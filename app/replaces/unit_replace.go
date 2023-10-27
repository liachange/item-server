package replaces

import (
	"item-server/app/requests"
	"item-server/pkg/helpers"
	"item-server/pkg/pinyin"
	"time"
)

type UnitIndex struct {
	State    uint8   `filter:"state,eq"`
	BetTime  []int64 `filter:"created_at,bet_time"`
	Title    string  `filter:"title,like"`
	Public   uint8   `filter:"is_public,eq"`
	Category []uint64
}

func (u *UnitIndex) UnitIndexReplace(req *requests.UnitFilterRequest) error {
	u.State = req.State
	u.Title = req.Title
	u.BetTime = BetTime(req.BetTime)
	u.Category = IdString(req.Category)
	u.Public = req.Public
	return nil
}

type UnitStore struct {
	Public   uint8
	Category []uint64
	Abbr     string
}

func (u *UnitStore) UnitStoreReplace(req *requests.UnitRequest) error {
	ids, public := IdPublicReplace(req.Category)
	u.Public = public
	u.Category = ids
	u.Abbr = pinyin.GetFirstSpell(req.Title)
	return nil
}

type UnitSave struct {
	IsPublic    uint8
	UpdatedAt   time.Time
	Ids         []uint64
	SelectSlice []string
	Abbr        string
}

func (u *UnitSave) UnitSaveReplace(req *requests.UnitRequest) error {
	ids, public := IdPublicReplace(req.Category)
	u.Ids = ids
	u.IsPublic = public
	u.UpdatedAt = helpers.TimeNow()
	u.Abbr = pinyin.GetFirstSpell(req.Title)
	u.SelectSlice = []string{
		"state",
		"title",
		"description",
		"is_public",
		"abbr",
		"sort",
		"updated_at",
	}
	return nil
}
