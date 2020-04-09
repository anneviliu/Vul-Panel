package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/wuhan005/govalid"
)

type RegUser struct {
	gorm.Model
	Username   string `gorm:"type:varchar(100);unique_index" json:"username" valid:"required;username" label:"用户名"`
	Password   string `json:"password" valid:"required;min:0;max:32" label:"密码"`
	Email      string `gorm:"type:varchar(100);unique_index" json:"email" valid:"required;email" label:"邮箱"`
	InviteCode string `gorm:"-" json:"inviteCode" valid:"required;" label:"邀请码"`
}

func (s *Service) register(c *gin.Context) {
	var UserData RegUser
	err := c.BindJSON(&UserData)
	if err != nil {
		c.JSON(400, gin.H{"errcode": 400, "msg": "error"})
		return
	}
	v := govalid.New(UserData)
	if !v.Check() {
		for _, err := range v.Errors {
			c.JSON(200, gin.H{"errcode": 400, "msg": err.Message})
			return
		}
	}

	if UserData.InviteCode != s.Conf.Admin.InviteCode {
		c.JSON(200, gin.H{"errcode": 400, "msg": "邀请码错误"})
		return
	}

	a := s.Mysql.Where("username = ? or email=?", UserData.Username, UserData.Email).Find(&UserData).RowsAffected
	if a != 0 {
		c.JSON(200, gin.H{"errcode": 400, "msg": "用户名或邮箱已存在"})
		return
	}

	userdata := RegUser{
		Model:    gorm.Model{},
		Username: UserData.Username,
		Password: UserData.Password,
		Email:    UserData.Email,
	}
	s.Mysql.Create(&userdata)
	c.JSON(200, gin.H{"errcode": 0, "msg": "注册成功"})
}

func (s *Service) login(c *gin.Context) {
	var loginData RegUser
	err := c.BindJSON(&loginData)
	if err != nil {
		c.JSON(400, gin.H{"errcode": 400, "msg": "error"})
		return
	}

	a := s.Mysql.Where("email = ? AND password=?", loginData.Email, loginData.Password).Find(&loginData).RowsAffected
	if a > 0 {
		session := sessions.Default(c)
		session.Set("mail", loginData.Email)
		session.Save()
		c.JSON(200, gin.H{"errcode": 0, "msg": "登录成功"})
	} else {
		c.JSON(200, gin.H{"errcode": 0, "msg": "邮箱或密码错误或不合法"})
		return
	}
}

func (s *Service) getSession(c *gin.Context) bool {
	session := sessions.Default(c)
	mail := session.Get("mail")
	fmt.Println("mail:", mail)
	if mail != nil {
		return true
	} else {
		return false
	}
}

func (s *Service) getNameByEmail(c *gin.Context) string {
	var loginData RegUser
	session := sessions.Default(c)
	mail := session.Get("mail")
	err := s.Mysql.Where("email = ?", mail).Find(&loginData).Error
	if err != nil {
		fmt.Println(err)
	}
	return loginData.Username
}

func (s *Service) logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("mail")
	session.Save()
	session.Clear()
	fmt.Println("delete session...", session.Get("mail"))
	c.Redirect(302, "/")
}
