package controllers

import (
	"ecommerce/models"
	"github.com/astaxie/beego"
)

// Operations about CoffeeMachines
type CoffeeMachinesController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Coffee Machines
// @Success 200 {object} models.CoffeeMachine
// @router / [get]
func (u *CoffeeMachinesController) GetAll() {
	coffeeMachines := models.GetAllCoffeeMachines()
	u.Data["json"] = coffeeMachines
	u.ServeJSON()
}
