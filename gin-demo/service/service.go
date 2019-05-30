package service

import (
	"fmt"
	"github.com/Bowser1704/go-study/gin-demo/model"
	"github.com/Bowser1704/go-study/gin-demo/util"
	"sync"
)

func Listuser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)	//users 是[]*UserModel, count是
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:new(sync.Mutex),
		IdMap:make(map[uint64]*model.UserInfo, len(users)),
	}

	//只放一个chan是为什么呢?
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GetShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:			u.Id,
				Username:	u.Username,
				SayHello:	fmt.Sprint("Hello %s", shortId),
				Password:	u.Password,
				CreatedAt:	u.CreatedAt.Format("2006-01-01 15:15:15"),
				UpdatedAt:	u.UpdatedAt.Format("2006-01-01 15:15:15"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <- finished:
	case err := <- errChan:
		return nil, count, err
	}

	//通过这样达到顺序不乱
	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
