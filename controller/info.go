package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"klsb/common"
	"klsb/model"
	"klsb/response"
	"net/http"
)

type data struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Timer     float64  `json:"timer"`
	Price     string   `json:"price"`
	Sumpeople string   `json:"sumpeople"`
	Tags      []string `json:"tags"`
	Url       string   `json:"url"`
	Islimit   string   `json:"islimit"`
}

func Get(c *gin.Context) {
	// 获取db
	db := common.GetDB()
	var infos []model.Info
	db.Order("is_limit desc, max").Find(&infos)
	var ret []data
	for _, info := range infos {
		jbInfo := db2json(info)
		ret = append(ret, jbInfo)
	}
	response.Success(c, gin.H{"data": ret, "errno": 0, "errmsg": "OK"})
}

func Getjb(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		response.Response(c, http.StatusNotFound, gin.H{})
		return
	}
	db := common.GetDB()
	var infos []model.Info
	db.Where("name = ?", name).Find(&infos)

	var ret []data
	for _, info := range infos {
		jbInfo := db2json(info)
		ret = append(ret, jbInfo)
	}
	response.Success(c, gin.H{"data": ret, "errno": 0, "errmsg": "OK"})

}

func db2json(info model.Info) data {
	db := common.GetDB()
	var jbInfo = data{
		Id:    int(info.ID),
		Name:  info.Name,
		Timer: float64(info.Timer),
		Price: fmt.Sprintf("%g/人", info.Price),
		Url:   info.Url,
	}
	if info.IsLimit {
		jbInfo.Islimit = "限定"
	}
	if info.Min == info.Max {
		jbInfo.Sumpeople = fmt.Sprintf("%d人", info.Min)
	} else {
		jbInfo.Sumpeople = fmt.Sprintf("%d人-%d人", info.Min, info.Max)
	}
	if info.IsInd {
		jbInfo.Sumpeople = fmt.Sprintf("%s(不分男女)", jbInfo.Sumpeople)
	} else {
		jbInfo.Sumpeople = fmt.Sprintf("%s(%d男%d女)", jbInfo.Sumpeople, info.Boy, info.Girl)
	}
	if info.Npc != 0 {
		jbInfo.Sumpeople = fmt.Sprintf("%s%dnpc", jbInfo.Sumpeople, info.Npc)
	}
	var tags []model.Tag
	db.Model(&info).Association("Tags").Find(&tags)
	for _, tag := range tags {
		jbInfo.Tags = append(jbInfo.Tags, tag.Name)
	}
	return jbInfo
}
