package replaces

import (
	"item-server/app/requests"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/pinyin"
)

type CategoryIndex struct {
	State    uint8    `filter:"state,eq"`
	Title    string   `filter:"title,like"`
	Category []uint64 `filter:"parent_id,in"`
	BetTime  []int64  `filter:"created_at,bet_time"`
}

func (c *CategoryIndex) BrandIndexReplace(req *requests.CategoryFilterRequest) error {

	c.State = req.State
	c.Title = req.Title
	c.BetTime = BetTime(req.BetTime)
	c.Category = IdString(req.Category)
	return nil
}

type CategoryStore struct {
	Abbr     string
	ParentId uint64
}

func (s *CategoryStore) CategoryStoreReplace(req *requests.CategoryRequest) error {
	var parent uint64
	if req.ParentId > 0 {
		parent = optimusPkg.NewOptimus().Decode(req.ParentId)
	}
	s.ParentId = parent
	s.Abbr = pinyin.GetFirstSpell(req.Title)
	return nil
}

type CategorySave struct {
	Abbr        string
	ParentId    uint64
	SelectSlice []string
}

func (s *CategorySave) CategorySaveReplace(req *requests.CategoryRequest) error {
	var parent uint64
	if req.ParentId > 0 {
		parent = optimusPkg.NewOptimus().Decode(req.ParentId)
	}
	s.ParentId = parent
	s.Abbr = pinyin.GetFirstSpell(req.Title)
	s.SelectSlice = []string{
		"state",
		"title",
		"description",
		"icon_url",
		"sort",
		"level",
		"level_tree",
		"abbr",
		"parent_id",
		"updated_at",
	}
	return nil
}
