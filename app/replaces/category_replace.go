package replaces

import "item-server/app/requests"

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
