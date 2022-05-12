package service

import (
	"github.com/nickczj/web1/cache"
	"github.com/nickczj/web1/database"
	"github.com/nickczj/web1/model"
	"github.com/nickczj/web1/utils"
	"strconv"
)

func GetNetWorth(id int) (finances model.Finances, err error) {
	key := cache.GenerateKey([]string{utils.GetMethodName(), strconv.Itoa(id)}...)
	return cache.GetElse(key, func() (model.Finances, error) {
		return getNetWorth(id)
	})
}

func getNetWorth(id int) (finances model.Finances, err error) {
	err = database.DB.First(&finances, id).Error
	return finances, err
}

func SaveNetWorth(f model.Finances) (model.Finances, error) {
	f.NetWorth = f.Assets - f.Liabilities
	err := database.DB.Create(&f).Error
	return f, err
}

func DeleteNetWorth(id int) {
	database.DB.Delete(&model.Finances{}, id)
}
