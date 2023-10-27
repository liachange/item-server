package resources

import (
	"item-server/app/models/category"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type Category struct {
	Model      *category.Category
	ModelSlice []category.Category
	ModelTree  []*category.Category
}

func (p *Category) ShowResource() map[string]any {
	optimus := optimusPkg.NewOptimus()
	parent := p.Model.ParentId
	if parent > 0 {
		parent = optimus.Encode(p.Model.ParentId)
	}
	return map[string]any{
		"id":          optimus.Encode(p.Model.ID),
		"state":       p.Model.State,
		"title":       p.Model.Title,
		"parent":      parent,
		"description": p.Model.Description,
		"icon":        p.Model.IconUrl,
		"sort":        p.Model.Sort,
		"created_at":  helpers.TimeFormat(p.Model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(p.Model.UpdatedAt, "second"),
	}
}
func (p *Category) IndexResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range p.ModelSlice {
		parent := model.ParentId
		if parent > 0 {
			parent = optimus.Encode(model.ParentId)
		}
		s = append(s, map[string]any{
			"id":          optimus.Encode(model.ID),
			"state":       model.State,
			"title":       model.Title,
			"parent":      parent,
			"description": model.Description,
			"icon":        model.IconUrl,
			"sort":        model.Sort,
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}
func (p *Category) InitialResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range p.ModelSlice {
		s = append(s, map[string]any{
			"value": optimus.Encode(model.ID),
			"label": model.Title,
		})
	}
	return s
}

type CategoryTree struct {
	Value    uint64          `json:"value,omitempty"`
	Title    string          `json:"title,omitempty"`
	Children []*CategoryTree `json:"children"`
	ParentId uint64          `json:"parent"`
	Sort     uint64          `json:"sort"`
}

func (p *Category) TreeIterative(parentId uint64) []*CategoryTree {
	optimus := optimusPkg.NewOptimus()
	tree := make(map[uint64]*CategoryTree, 0)
	for _, v := range p.ModelTree {
		id := v.ID
		if id > 0 {
			id = optimus.Optimus.Encode(id)
		}
		parent := v.ParentId
		if parent > 0 {
			parent = optimus.Optimus.Encode(parent)
		}
		if _, ok := tree[id]; ok {
			tree[id] = &CategoryTree{
				Value:    id,
				Title:    v.Title,
				ParentId: parent,
				Sort:     v.Sort,
				Children: tree[id].Children,
			}
		} else {
			tree[id] = &CategoryTree{
				Value:    id,
				Title:    v.Title,
				ParentId: parent,
				Sort:     v.Sort,
				Children: make([]*CategoryTree, 0),
			}
		}
		if _, ok := tree[parent]; ok {
			tree[parent].Children = append(tree[parent].Children, tree[id])
		} else {
			tree[parent] = &CategoryTree{
				Value:    id,
				Title:    v.Title,
				ParentId: parent,
				Sort:     v.Sort,
				Children: []*CategoryTree{tree[id]},
			}
		}
	}
	return tree[parentId].Children
}

//func (p *Category) Convert(parentId uint64) []*CategoryTree {
//	tree := make([]*CategoryTree, 0)
//
//	for _, v := range p.ModelTree {
//		if v.ParentId == parentId {
//			child := &CategoryTree{
//				Value:    v.ID,
//				Title:    v.Title,
//				ParentId: v.ParentId,
//				Sort:     v.Sort,
//			}
//			subList := p.Convert(v.ID)
//			if len(subList) > 0 {
//				child.Children = subList
//			}
//			tree = append(tree, child)
//		}
//	}
//	return tree
//}
