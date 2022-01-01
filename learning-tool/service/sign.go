package service

import (
	"learning-tool/dao"
	"learning-tool/modles"
)

func PublishSign(sign modles.Sign) error {
	err := dao.PublishSign(sign)
	if err != nil {
		return err
	}
	return nil
}

func Sign(sign modles.Sign) error {
	publishTime, _, err := dao.SelectSignByRoomId(sign.RoomId)
	err = dao.Sign(publishTime, sign)
	if err != nil {
		return err
	}
	return nil
}
