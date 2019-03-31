package controllers

import (
	"fmt"
	. "go_transfer/models"
	. "go_transfer/services"
	. "go_transfer/utilities"
	"gopkg.in/unrolled/render.v1"
	"net/http"
	"strconv"
)

type Action func(rw http.ResponseWriter, r *http.Request) error

type AppController struct{}

func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

type MyController struct {
	AppController
	*Users
	*render.Render
}

func (c *MyController) GetUser(rw http.ResponseWriter, r *http.Request) error {
	name, _ := r.URL.Query()["name"]
	var user User
	for _, item := range c.Users.Users {
		if item.Name == name[0] {
			user = item
			break
		}
	}
	c.JSON(rw, 200, map[string]float64{"balance": user.Balance})
	return nil
}

func (c *MyController) Transfer(rw http.ResponseWriter, r *http.Request) error {
	GenerateHTML(rw, c.Users.Users, "layout", "index")
	return nil
}

func (c *MyController) SendFounds(rw http.ResponseWriter, r *http.Request) error {
	nameSender := r.FormValue("fromUser")
	nameTarget := r.FormValue("toUser")
	availCreditStr := r.FormValue("fromUserAvail")
	amountValStr := r.FormValue("amountVal")
	availCredit, _ := strconv.ParseFloat(availCreditStr, 64)
	amountVal, _ := strconv.ParseFloat(amountValStr, 64)
	ok, _ := ProcessTransfer(c.Users, nameSender, availCredit, nameTarget, amountVal)
	if ok {
		var balance = fmt.Sprintf("%.2f", availCredit-amountVal)
		var res = map[string]string{
			"sender":  nameSender,
			"balance": balance,
			"target":  nameTarget,
			"amount":  amountValStr,
		}
		GenerateHTML(rw, res, "layout", "success")
		var msg = "Mr(s) " + nameSender + " our have transfered $" + amountValStr
		msg += " for Mr(s). " + nameTarget + ".\n"
		msg += "Your Account Balance now is $" + balance + "."
		Info(msg)
	} else {
		var messages []string
		if nameSender == "default" || nameTarget == "default" {
			messages = append(messages, "Need select sender and target users for transfer founds!!!")
		}
		if availCredit == float64(0) || amountVal > availCredit {
			messages = append(messages, "Insufficient funds in your account for transfer!!!")
			Info("Mr(s). " + nameSender + "you have $" + availCreditStr + " what is insufficient funds for transfer $" + amountValStr + "!!!")
		}
		if amountValStr == "0.00" || amountValStr == "" {
			messages = append(messages, "Need one amount value for transfer founds!!!")
		}
		GenerateHTML(rw, messages, "layout", "error")
	}
	return nil
}
