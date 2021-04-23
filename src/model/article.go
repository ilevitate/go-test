package model

import "go-test/src/global"

func (query *Article) Create() (err error) {
	res := global.DB.Create(&query)
	if res.Error != nil {
		return err
	}
	return
}
