package main

import (
	"crypto/md5"
	"fmt"
	pageable "github.com/BillSJC/gorm-pageable"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func (s *Service) genFilename(data VulInfo) {
	if data.Detail.Host != "" {
		hostSlice := strings.Split(data.Detail.Host, ".")
		md5s := []byte(data.Detail.Url + data.Detail.Payload + time.Now().String())
		has := md5.Sum(md5s)
		ext := fmt.Sprintf("%x", has)
		s.Conf.TempFileName = hostSlice[len(hostSlice)-2] + "-" + ext
	}
	//else if data.Detail.Url != "" {
	//	hostSlice := strings.Split(data.Detail.Host, ".")
	//	md5s := []byte(data.Detail.Url + data.Detail.Payload + time.Now().String())
	//	has := md5.Sum(md5s)
	//	ext := fmt.Sprintf("%x", has)
	//	s.Conf.TempFileName = hostSlice[len(hostSlice)-2] + "-" + ext
	//}

}

// 返回漏洞条目总数
func (s *Service) getTotalItems(c *gin.Context) {
	var total int
	s.Mysql.Table("vuls").Count(&total)
	//fmt.Println(total)
	c.JSON(200, gin.H{"code": 200, "msg": total})
}

// 分页
func (s *Service) getListByPage(t string, pageNum int, pageSize int, c *gin.Context) {
	resultSet := make([]*Vul, 0, 30)
	handler := s.Mysql.Model(&Vul{}).Order("created_at desc")
	//pageNumint, _ := strconv.Atoi(pageNum)
	//fmt.Println(pageNum)
	r, err := pageable.PageQuery(pageNum, pageSize, handler, &resultSet)

	if err != nil {
		panic(err)
	}
	if t == "totalPages" {
		c.JSON(200, r.PageCount)
		return
	}
	type RecentList struct {
		ID        uint
		Host      string
		CreatedAt string
		VulUrl    string
		Url       string
		Title     string
		Times     string
		Read      bool
		Status    string
	}
	var res []RecentList
	for _, v := range resultSet {
		time := v.CreatedAt.Format("2006-01-02 15:04:05")
		res = append(res, RecentList{
			ID:        v.ID,
			Host:      v.Host,
			VulUrl:    v.Url,
			CreatedAt: time,
			Url:       "/vulinfo/" + v.TempFilename,
			Title:     v.VulClass,
			Read:      v.Read,
			Status:    v.Status,
		})
	}
	c.JSON(200, res)
}
