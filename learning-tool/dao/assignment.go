package dao

import "learning-tool/modles"

func AssignHomework(assignment modles.Assignment) error {
	sqlStr := "insert into assignment(classroom_id,teacher_name,homework) values(?,?,?)"
	_, err := DB.Exec(sqlStr, assignment.RoomId, assignment.Teacher, assignment.Homework)
	if err != nil {
		return err
	}
	return nil
}
