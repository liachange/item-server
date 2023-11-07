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
type CategoryResource struct {
	ID          uint64 `json:"id"`
	State       uint8  `json:"state"`
	Title       string `json:"title"`
	ParentId    uint64 `json:"parent"`
	Icon        string `json:"icon"`
	Sort        uint64 `json:"sort"`
	Abbr        string `json:"abbr"`
	Description string `json:"desc"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (p *Category) ShowResource() (show CategoryResource) {
	optimus := optimusPkg.NewOptimus()
	parent := p.Model.ParentId
	if parent > 0 {
		parent = optimus.Encode(p.Model.ParentId)
	}
	show.ID = optimus.Encode(p.Model.ID)
	show.State = p.Model.State
	show.Title = p.Model.Title
	show.ParentId = parent
	show.Description = p.Model.Description
	show.Icon = p.Model.IconUrl
	show.Sort = p.Model.Sort
	show.Abbr = p.Model.Abbr
	show.CreatedAt = helpers.TimeFormat(p.Model.CreatedAt, "second")
	show.UpdatedAt = helpers.TimeFormat(p.Model.UpdatedAt, "second")
	return
}
func (p *Category) IndexResource() (index []*CategoryResource) {
	optimus := optimusPkg.NewOptimus()
	for _, model := range p.ModelSlice {
		parent := model.ParentId
		if parent > 0 {
			parent = optimus.Encode(model.ParentId)
		}
		index = append(index, &CategoryResource{
			ID:          optimus.Encode(model.ID),
			State:       model.State,
			Title:       model.Title,
			ParentId:    parent,
			Description: model.Description,
			Icon:        model.IconUrl,
			Sort:        model.Sort,
			Abbr:        model.Abbr,
			CreatedAt:   helpers.TimeFormat(model.CreatedAt, "second"),
			UpdatedAt:   helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return
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
	if len(p.ModelTree) == 0 {
		return make([]*CategoryTree, 0)
	}
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

type CategoryHas struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

func CategoryHasResource(r []*category.Many) (c []*CategoryHas) {
	optimus := optimusPkg.NewOptimus()
	for _, v := range r {
		c = append(c, &CategoryHas{
			ID:    optimus.Encode(v.ID),
			Title: v.Title,
		})
	}
	return
}

type CategorySelect struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Level uint8  `json:"level"`
	Abbr  string `json:"abbr"`
}

func (p *Category) CategorySelectResource() (sel []*CategorySelect) {
	optimus := optimusPkg.NewOptimus()
	for _, model := range p.ModelTree {
		sel = append(sel, &CategorySelect{
			ID:    optimus.Encode(model.ID),
			Title: model.Title,
			Abbr:  model.Abbr,
			Level: model.Level,
		})
	}
	return
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
