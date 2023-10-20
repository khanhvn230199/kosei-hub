package auto

import (
	"kosei-web/api/models"
)

var users = []models.User{
	models.User{Nickname: "admin", Email: "admin@gmail.com", Password: "123456", Role: "1", Job: "1", Avartar: "https://anhcuoiviet.vn/wp-content/uploads/2022/11/avatar-dep-0.jpg"},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title",
		Content: "Hello World",
	},
}
