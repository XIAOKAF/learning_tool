package dao

import (
	"learning-tool/modles"
)

func CreateReportCard(grades modles.Grades) error {
	stmt, err := DB.Prepare("create table " + grades.Subject + "(subject VARCHAR(20) not null, teacher VARCHAR(20) not null, student_id BIGINT(20), student_name VARCHAR(20), scores FLOAT(20), average_score float(20) default 0, primary key(subject));")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func InsertGrades(grades modles.Grades) error {
	sqlStr := "insert into " + grades.Subject + "(subject,teacher,student_id,student_name,scores)values(?,?,?,?,?)"
	_, err := DB.Exec(sqlStr, grades.Subject, grades.Teacher, grades.StudentId, grades.StudentName, grades.Scores)
	if err != nil {
		return err
	}
	return nil
}

func AverageScore(grades modles.Grades) (float32, error) {
	//返回平均分
	rows, err := DB.Query("select AVG(scores) from"+grades.Subject+"where teacher = ?", grades.Teacher)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&grades.AverageScore)
		if err != nil {
			return 0, err
		}
	}
	//将平均分插入到MySQL中
	sqlStr := "insert into " + grades.Subject + "(average_score) values(?)"
	_, err = DB.Exec(sqlStr, grades.AverageScore)
	if err != nil {
		return grades.AverageScore, err
	}
	return grades.AverageScore, nil
}

func HighestScore(grades modles.Grades) (float32, string, error) {
	//找到最高分
	var maxScore float32
	rows, err := DB.Query("select max(scores) from " + grades.Subject)
	if err != nil {
		return 0, "", err
	}
	for rows.Next() {
		err = rows.Scan(&maxScore)
		if err != nil {
			return maxScore, "", err
		}
	}
	//找到取得最高分的同学姓名
	var student string
	rows, err = DB.Query("select student_name from "+grades.Subject+"where scores = ?", maxScore)
	for rows.Next() {
		err = rows.Scan(&student)
		if err != nil {
			return maxScore, "", err
		}
	}
	return maxScore, student, nil
}

func SpecifiedStudent(grades modles.Grades) (modles.Grades, error) {
	rows, err := DB.Query("select scores from "+grades.Subject+" where student_id = ?,student_name=?", grades.StudentId, grades.StudentName)
	if err != nil {
		return grades, err
	}
	for rows.Next() {
		err = rows.Scan(&grades.Scores)
		if err != nil {
			return grades, err
		}
	}
	return grades, nil
}
