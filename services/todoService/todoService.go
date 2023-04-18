package todoService

import (
	"boilerplate/database"
	"boilerplate/entities"
	"boilerplate/utils"
	"errors"
)

func GetAll(activity_group_id int) (*[]entities.Todo, error) {
	var err error
	var datas *[]entities.Todo
	tx := database.Conn
	if activity_group_id > 0 {
		tx.Where("activity_group_id = ? ", activity_group_id)
	}
	tx.Order("id").Find(&datas)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if *datas == nil {
		err = errors.New("datas not found")
	}

	return datas, err
}

func Get(id int) (*entities.Todo, error) {
	var err error
	var resp *entities.Todo
	tx := database.Conn.Find(&resp, id)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if resp.ID == 0 {
		err = errors.New("datas not found")
	}
	return resp, err
}

func Create(body *entities.Todo) (*entities.Todo, error) {
	var err error
	tx := database.Conn.Create(&body)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return body, err
}

func Update(id int, body map[string]interface{}) (*entities.Todo, error) {
	var err error
	var resp *entities.Todo
	if err := database.Conn.First(&resp, id).Error; err != nil {
		return nil, err
	}
	resp.Title = body["title"].(string)
	tx := database.Conn.Save(&resp)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}

	return resp, err
}

func Delete(id int) error {
	var err error
	var resp *entities.Todo
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
