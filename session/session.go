package session

import (
	"demo_gin/dao"
	"github.com/garyburd/redigo/redis"
	"time"
)

//根据sessionId在redis中取值
func GetSession(sessionId string) (id string,err error) {
	conn := dao.Pool.Get()
	//ping
	err = dao.Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	id, err = redis.String(conn.Do("Get", sessionId))
	if err != nil {
		return
	}
	return
}

//设置redis中session Key--uuid Value--id
func SetSession(sessionId, id string) (err error) {
	conn := dao.Pool.Get()
	//ping
	err = dao.Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return err
	}
	//设置redis session
	_, err = conn.Do("Set", sessionId, id)
	if err != nil {
		return err
	}
	//一分钟过期
	_, err = conn.Do("expire", sessionId, 60)
	if err != nil {
		return err
	}
	return
}
