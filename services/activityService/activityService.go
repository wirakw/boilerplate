package activityService

import (
	"boilerplate/database"
	"boilerplate/entities"
	"boilerplate/utils"
	"errors"
)

func GetAll() (*[]entities.Activity, error) {
	var err error
	var datas *[]entities.Activity
	tx := database.Conn.Find(&datas)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if *datas == nil {
		err = errors.New("datas not found")
	}

	return datas, err
}

func Get(id int) (*entities.Activity, error) {
	var err error
	var resp *entities.Activity
	tx := database.Conn.Find(&resp, id)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if resp.ID == 0 {
		err = errors.New("datas not found")
	}
	return resp, err
}

func Create(body *entities.Activity) (*entities.Activity, error) {
	var err error
	tx := database.Conn.Create(&body)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return body, err
}

func Update(id int, body *entities.Activity) (*entities.Activity, error) {
	var err error
	var resp *entities.Activity
	if err := database.Conn.First(&resp, id).Error; err != nil {
		return nil, err
	}
	resp.Title = body.Title
	tx := database.Conn.Save(&resp)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return resp, err
}

func Delete(id int) error {
	var err error
	var resp *entities.Activity
	if err := database.Conn.First(&resp, id).Error; err != nil {
		return err
	}
	tx := database.Conn.Delete(&resp, id)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}
	return err
}
