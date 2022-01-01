package dao

import (
	"learning-tool/modles"
)

func IsStudentExist(students modles.Students) (modles.Students, error) {

	rows, err := DB.Query("select realname,student_id from list where student_id=?", students.StudentId)
	if err != nil {
		return students, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&students.RealName, &students.StudentId)
		if err != nil {
			return students, err
		}
	}

	return students, nil
}

func IsIdAndNameMatch(students modles.Students) (string, error) {
	rows, err := DB.Query("select realname from list where student_id=?", students.StudentId)
	if err != nil {
		return students.RealName, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&students.RealName)
		if err != nil {
			return students.RealName, err
		}
	}

	return students.RealName, nil
}

func Register(students modles.Students) error {
	sqlStr := "insert into student(realname,password,accessLevel) values(?,?,?)"
	_, err := DB.Exec(sqlStr, students.RealName, students.Password, students.AccessLevel)
	if err != nil {
		return err
	}
	return nil
}

func SelectPasswordByStudentId(students modles.Students) (string, error) {
	rows, err := DB.Query("select password from student where student_id=?", students.StudentId)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err = rows.Scan(&students.Password)
		if err != nil {
			return students.Password, err
		}
	}
	return students.Password, nil
}

func UploadInfo(students modles.Students) error {
	sqlStr := "update student set nickname=?,mobile=?,avatar=? where student_id=?"
	_, err := DB.Exec(sqlStr, students.NickName, students.Mobile, students.Avatar, students.StudentId)
	if err != nil {
		return err
	}
	return nil
}
