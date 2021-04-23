package model

import (
	"errors"
	"go-test/src/global"
	"gorm.io/gorm"
)

func (user *User) GetList(query map[string]string, columns []string, listParams ListParams) (err error) {
	querySql := global.DB
	for k := range columns {
		querySql.Select(columns[k])
	}
	if query["userName"] != "" {
		querySql.Where("name = ?", query["userName"])
	}
	if listParams.Limit != 0 && listParams.Page != 0 {
		querySql.Limit(listParams.Limit).Offset((listParams.Page - 1) * listParams.Limit)
	} else {
		querySql.Limit(listParams.Limit)
	}

	res := querySql.Find(&user)
	if res.Error != nil {
		return res.Error
	}
	return
}

func (user *User) Create() (err error) {
	res := global.DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return
}

func (user *User) Edit(id int) (err error) {
	res := global.DB.Model(&user).Where("id = ?", id).Updates(User{Name: "鲁迅"})
	if res.Error != nil {
		return res.Error
	}
	return
}

func (user *User) Delete(id int) (err error) {
	if id <= 0 {
		return errors.New("ID不能为空")
	}
	res := gorm.DB.Delete(&user, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	return
}
