package service

import (
	"learning-tool/dao"
	"learning-tool/modles"
	"time"
)

func CreateClassroom(classroom modles.Classroom) error {
	err := dao.CreateClassroom(classroom)
	if err != nil {
		return err
	}
	return nil
}

func GetStudentName(studentId int) (string, error) {
	name, err := dao.SelectStudentNameByStudentID(studentId)
	if err != nil {
		return "", err
	}
	return name, nil
}

func JoinClassroom(classroom modles.Classroom) (int, error) {
	err := dao.JoinClassroom(classroom)
	if err != nil {
		return 0, err
	}

	//进入教室后查看老师是否有发布签到
	_, t, err := dao.SelectSignByRoomId(classroom.RoomId)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()
	if t.Before(currentTime) {
		return 0, nil
	}
	return 1, nil
}
