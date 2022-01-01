package dao

import "learning-tool/modles"

func TeacherRegister(teachers modles.Teachers) error {
	sqlStr := "insert into teacher(name,password,accessLevel) values(?,?,?)"
	_, err := DB.Exec(sqlStr, teachers.Name, teachers.Password, teachers.AccessLevel)
	if err != nil {
		return err
	}
	return nil
}

func SelectPasswordByName(teachers modles.Teachers) (string, error) {
	rows, err := DB.Query("select password from teacher where name=?", teachers.Name)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		err = rows.Scan(&teachers.Password)
		if err != nil {
			return teachers.Password, err
		}
	}
	return teachers.Password, nil
}
