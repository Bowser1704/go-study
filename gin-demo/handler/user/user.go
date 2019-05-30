package user

import "github.com/Bowser1704/go-study/gin-demo/model"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username 	string `json:"username"`
	Offest		string `json:"offest"`
	Limit 		string `json:"limit"`
}

type ListResponse struct {
	TotalCount	uint64				`json:"total_count"`
	UserList 	[]*model.UserInfo	`json:"user_list"`
}


