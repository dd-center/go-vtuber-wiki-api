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
	defer vc.ServeJSON()
	vtubers, err := models.GetAllVtubers()
	if err != nil {
		vc.Data["json"] = &struct {
			Success bool
			Error   string
		}{false, err.Error()}
		return
	}
	vc.Data["json"] = &struct {
		Success bool
		Vtubers []models.VtuberEntity
	}{true, vtubers}
}

// @router /search [get]
func (vc *VtuberController) SearchVtuber() {
	defer vc.ServeJSON()
	errorTemplate := &struct {
		Success bool
		Error   string
	}{false, ""}
	keyword := vc.Input().Get("keyword")
	if len(keyword) == 0 {
		errorTemplate.Error = "Keyword cannot be null or empty."
		vc.Data["json"] = errorTemplate
		return
	}
	vtuber, err := models.SearchVtuber(keyword)
	if err != nil {
		errorTemplate.Error = err.Error()
		vc.Data["json"] = errorTemplate
		return
	}
	vc.Data["json"] = &struct {
		Success bool
		Vtuber  *models.VtuberEntity
	}{true, vtuber}
}
