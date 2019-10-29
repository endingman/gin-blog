package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
	DeletedOn  int    `json:"deleted_on"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

/**
 *这属于gorm的Callbacks，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm将停止未来操作并回滚所有更改。

gorm所支持的回调方法：

创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind
*/

/**
使用钩子函数代替
*/

//func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}

func ExistTagByID(id int) (bool, error) {
	var tag Tag

	err := db.Select("id").Where("id = ?", id).First(&tag).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func GetTag(id int) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &tag, nil
}

func CleanAllTags() bool {
	//硬删除要使用 Unscoped()，这是 GORM 的约定
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error; err != nil {
		return false
	}

	return true
}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
