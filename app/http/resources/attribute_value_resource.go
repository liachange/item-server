package resources

import (
	"item-server/app/models/attribute_value"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type AttributeValue struct {
	Model      *attribute_value.AttributeValue
	ModelSlice []*attribute_value.AttributeValue
}

type AttributeValueResource struct {
	ID              uint64 `json:"id"`
	AttributeNameId uint64 `json:"attribute_name"`
	AttrName        string `json:"attr_name"`
	State           uint8  `json:"state"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Sort            uint64 `json:"sort"`
	Abbr            string `json:"abbr"`
	Search          string `json:"search"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (p *AttributeValue) ShowResource() (show AttributeValueResource) {
	optimus := optimusPkg.NewOptimus()
	show.ID = optimus.Encode(p.Model.ID)
	show.State = p.Model.State
	show.Title = p.Model.Title
	show.Description = p.Model.Description
	show.AttributeNameId = optimus.Encode(p.Model.AttributeNameId)
	show.Sort = p.Model.Sort
	show.Abbr = p.Model.Abbr
	show.AttrName = p.Model.AttributeName.Title
	show.Search = p.Model.Search
	show.CreatedAt = helpers.TimeFormat(p.Model.CreatedAt, "second")
	show.UpdatedAt = helpers.TimeFormat(p.Model.UpdatedAt, "second")
	return

}

type AttributeValueIndexResource struct {
	ID              uint64 `json:"id"`
	AttributeNameId uint64 `json:"attribute_name"`
	AttrName        string `json:"attr_name"`
	State           uint8  `json:"state"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Sort            uint64 `json:"sort"`
	Abbr            string `json:"abbr"`
	Search          string `json:"search"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (p *AttributeValue) IndexResource() (index []*AttributeValueIndexResource) {
	optimus := optimusPkg.NewOptimus()
	for _, model := range p.ModelSlice {
		index = append(index, &AttributeValueIndexResource{
			ID:              optimus.Encode(model.ID),
			State:           model.State,
			Title:           model.Title,
			Description:     model.Description,
			AttributeNameId: optimus.Encode(model.AttributeNameId),
			Sort:            model.Sort,
			Abbr:            model.Abbr,
			Search:          model.Search,
			AttrName:        model.AttributeName.Title,
			CreatedAt:       helpers.TimeFormat(model.CreatedAt, "second"),
			UpdatedAt:       helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return
}
