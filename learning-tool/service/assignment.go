package service

import (
	"learning-tool/dao"
	"learning-tool/modles"
)

func AssignHomework(assignment modles.Assignment) error {
	err := dao.AssignHomework(assignment)
	if err != nil {
		return err
	}
	return nil
}
