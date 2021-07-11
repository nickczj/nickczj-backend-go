package service

import (
	"github.com/nickczj/web1/database"
	"github.com/nickczj/web1/model"
)

func GetNetWorth(id int) (finances model.Finances, err error) {
	err = database.DB.First(&finances, id).Error
	return
}

func SaveNetWorth(f model.Finances) (model.Finances, error) {
	f.NetWorth = f.Assets - f.Liabilities
	err := database.DB.Create(&f).Error
	return f, err
}

func DeleteNetWorth(id int) {
	database.DB.Delete(&model.Finances{}, id)
}