package db

import "gorm.io/gorm"

func CreateOrm(model interface{}) *gorm.DB {
	//TODO 개선의 여지 있음
	orm := GetDataBase()
	err := orm.AutoMigrate(model)
	if err != nil {
		return nil
	}
	return orm
}
