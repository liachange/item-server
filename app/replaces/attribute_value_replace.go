package replaces

import "item-server/app/requests"

type AttributeValueIndex struct {
	State         uint8    `filter:"state,eq"`
	Title         string   `filter:"title,like"`
	AttributeName []uint64 `filter:"attribute_name_id,in"`
}

func (a *AttributeValueIndex) AttrValueReplace(rep *requests.AttributeValueFilterRequest) error {
	a.State = rep.State
	a.Title = rep.Title
	a.AttributeName = IdString(rep.AttributeName)
	return nil
}
