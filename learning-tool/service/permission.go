package service

import "learning-tool/dao"

func ReturnPermission(name string) (string, error) {
	permission, err := dao.ReturnPermission(name)
	if err != nil {
		return "", err
	}
	return permission, nil
}
