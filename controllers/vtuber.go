package controllers

import (
	"github.com/astaxie/beego"
	"go-vtuber-wiki-api/models"
)

type VtuberController struct {
	beego.Controller
}

// @router /list [get]
func (vc *VtuberController) GetAllVtubers() {
	vtubers, err := models.GetAllVtubers()
	if err != nil {
		vc.Data["json"] = &struct {
			Success bool
			Error   string
		}{false, err.Error()}
		vc.ServeJSON()
		return
	}
	vc.Data["json"] = &struct {
		Success bool
		Vtubers []models.VtuberEntity
	}{true, vtubers}
	vc.ServeJSON()
}

// @router /search [get]
func (vc *VtuberController) SearchVtuber() {
	errorTemplate := &struct {
		Success bool
		Error   string
	}{false, ""}
	keyword := vc.Input().Get("keyword")
	if len(keyword) == 0 {
		errorTemplate.Error = "Keyword cannot be null or empty."
		vc.Data["json"] = errorTemplate
		vc.ServeJSON()
		return
	}
	vtuber, err := models.SearchVtuber(keyword)
	if err != nil {
		errorTemplate.Error = err.Error()
		vc.Data["json"] = errorTemplate
		vc.ServeJSON()
		return
	}
	vc.Data["json"] = &struct {
		Success bool
		Vtuber  *models.VtuberEntity
	}{true, vtuber}
	vc.ServeJSON()
}
