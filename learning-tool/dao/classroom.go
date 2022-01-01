package dao

import (
	"fmt"
	"learning-tool/modles"
	"strings"
)

func CreateClassroom(classroom modles.Classroom) error {
	sqlStr := "insert into classroom(classroom_name,capacity,teacher_name) values(?,?,?)"
	_, err := DB.Exec(sqlStr, classroom.RoomName, classroom.Capacity, classroom.Teacher)
	if err != nil {
		return err
	}
	return nil
}

func SelectStudentNameByStudentID(studentId int) (string, error) {
	var name string

	rows, err := DB.Query("select realname from student where student_id = ?", studentId)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
	}

	return name, nil
}

func JoinClassroom(classroom modles.Classroom) error {
	students := fmt.Sprintf(strings.Join(classroom.Members, ","))
	sqlStr := "update classroom set students = ? where classroom_id = ?"
	_, err := DB.Exec(sqlStr, students, classroom.RoomId)
	if err != nil {
		return err
	}
	return nil
}
