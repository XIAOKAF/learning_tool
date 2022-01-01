package modles

import "time"

type Sign struct {
	RoomId      int
	Teacher     string
	Student     []string  //完成签到学生姓名
	SignAmount  int       //签到总人数
	PublishTime time.Time //发布签到时间
	OverTime    time.Time //签到结束时间
}
