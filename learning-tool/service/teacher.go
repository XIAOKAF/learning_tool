package service

import (
	"learning-tool/dao"
	"learning-tool/modles"
)

func TeacherRegister(teacher modles.Teachers) error {
	err := dao.TeacherRegister(teacher)
	if err != nil {
		return err
	}
	return nil
}

func TeacherLogin(teachers modles.Teachers) (string, error) {
	truePassword, err := dao.SelectPasswordByName(teachers)
	if err != nil {
		return "", err
	}
	return truePassword, nil
}
