package service

import (
	"database/sql"
	"learning-tool/dao"
	"learning-tool/modles"
)

// IsStudentExist 学号是否存在
func IsStudentExist(students modles.Students) (bool, error) {
	_, err := dao.IsStudentExist(students)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil //学号不存在
		}
		return false, err //出错
	}
	return true, nil //学号存在且无错
}

// IsIdAndNameMatch 学号与输入的真实姓名是否匹配
func IsIdAndNameMatch(students modles.Students) (string, error) {
	name, err := dao.IsIdAndNameMatch(students)
	if err != nil {
		return name, err
	}
	return name, nil
}

// Register 注册
func Register(student modles.Students) error {
	err := dao.Register(student)
	if err != nil {
		return err
	}
	return nil
}

// IsPasswordCorrect 查询正确密码
func IsPasswordCorrect(student modles.Students) (string, error) {
	truePassword, err := dao.SelectPasswordByStudentId(student)
	if err != nil {
		return "", err
	}
	return truePassword, nil
}

// UpdateInfo 登录成功后可进行个人信息完善
func UpdateInfo(student modles.Students) error {
	err := dao.UploadInfo(student)
	if err != nil {
		return err
	}
	return nil
}
