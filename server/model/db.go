package model

import (
	"myblog/global"
)

type User struct {
	global.GMODEL
	Name      string `json:"name" gorm:"comment:用户名"`
	UserType  int    `json:"userType" gorm:"comment:用户类型"`
	PhoneNum  string `json:"phoneNum" gorm:"comment:手机号"`
	ImageUrl  string `json:"imageUrl" gorm:"comment:用户封面"`
	Email     string `json:"email" gorm:"comment:邮箱"`
	Introduce string `json:"introduce" gorm:"comment:简介"`
	Avatar    string `json:"avatar" gorm:"comment:头像"`
	Password  string `json:"password" gorm:"comment:密码"`
}

type Category struct {
	global.GMODEL
	Name string `json:"name" gorm:"comment:文章类型名称"`
}

type Article struct {
	global.GMODEL
	Title       string   `json:"title" gorm:"comment:标题"`
	KeyWord     string   `json:"keyWord" gorm:"comment:关键字"`
	Author      User     `json:"author" gorm:"foreignKey:ID;references:ID;comment:作者"`
	Description string   `json:"description" gorm:"comment:描述"`
	Content     string   `json:"content" gorm:"comment:内容"`
	Numbers     int64    `json:"numbers" gorm:"comment:字数"`
	Image       string   `json:"image" gorm:"comment:封面"`
	ArticleType int      `json:"articleType" gorm:"comment:类型"`
	State       int      `json:"state" gorm:"comment:状态"`
	Origin      int      `json:"origin" gorm:"comment:来源"`
	Tags        string   `json:"tags" gorm:"comment:标签"`
	Category    Category `json:"category" gorm:"foreignKey:ID;references:ID;comment:文章类型"`
	LikeUser    User     `json:"likeUser" gorm:"foreignKey:ID;references:ID;comment:点赞的用户"`
	ViewsNum    int64    `json:"viewsNum" gorm:"comment:阅读数量"`
	LikesNum    int64    `json:"likesNum" gorm:"comment:点赞数量"`
	CommentsNum int64    `json:"commentsNum" gorm:"comment:评论数量"`
}

type Comments struct {
	global.GMODEL
	ArticleId Article `json:"articleId" gorm:"foreignKey:ID;comment:关联文章ID"`
	Content   string  `json:"content" gorm:"comment:内容"`
	IsTop     bool    `json:"isTop" gorm:"comment:是否置顶"`
	LikesNum  int64   `json:"likesNum" gorm:"comment:点赞数"`
	UserId    User    `json:"userId" gorm:"ID;foreignKey:ID"`
	State     int     `json:"state" gorm:"comment:状态"`
}
