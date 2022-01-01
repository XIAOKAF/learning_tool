package dao

import (
	"learning-tool/modles"
	time2 "time"
)

func PublishSign(sign modles.Sign) error {
	sqlStr := "insert into sign(room_id,teacher,publish_time,over_time)values(?,?,?,?)"
	_, err := DB.Exec(sqlStr,sign.RoomId,sign.Teacher,sign.PublishTime,sign.OverTime)
	if err != nil {
		return err
	}
	return nil
}

func SelectSignByRoomId(roomId int) (time2.Time,time2.Time,error) {
	sqlStr := "select publish_time, over_time from sign where room_id = ?"
	rows, err := DB.Query(sqlStr,roomId)
	if err != nil {
		return time2.Now(),time2.Now(),err
	}
	defer rows.Close()

	var publishTime,t time2.Time
	for rows.Next() {
		err = rows.Scan(&publishTime,&t)
		if err != nil {
			return time2.Now(),time2.Now(),err
		}
	}
	return publishTime,t,nil
}

func Sign(t time2.Time,sign modles.Sign) error {
	sqlStr := "update sign set student = ?,signAmount = ? where roomId = ? && publish_time = ?"
	_,err := DB.Exec(sqlStr,sign.Student,sign.SignAmount,sign.RoomId,t)
	if err != nil {
		return err
	}
	return nil
}