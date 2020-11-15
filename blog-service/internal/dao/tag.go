package dao

import (
	"github.com/go-microservices/blog-service/internal/model"
	"github.com/go-microservices/blog-service/pkg/app"
)

func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name: name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{} {
		"state": state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffSet := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) GetTagListByIds(ids []uint32, state uint8) ([]*model.Tag, error) {
	tag := model.Tag{State: state}
	return tag.ListByIds(d.engine, ids)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}