package login

import (
	cfg "FaceAnnotation/config"
	manager "FaceAnnotation/service/model/managermodel"
	vars "FaceAnnotation/service/vars"
	security "FaceAnnotation/utils/security"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type LoginParams struct {
	Managername string `json:"username"`
	Password    string `json:"password"`
}

func Login(c *gin.Context) {

	var loginParams LoginParams
	if err := c.BindJSON(&loginParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err.Error()))
		c.JSON(400, gin.H{
			"code":    vars.ErrBindJSON.Code,
			"message": vars.ErrBindJSON.Msg,
		})
		return
	}
	username := loginParams.Managername
	password := loginParams.Password

	managerColl, err := manager.QueryUser(username)
	if err != nil {
		log.Error(fmt.Sprintf("find user error:%s", err.Error()))
		c.JSON(400, gin.H{
			"code":    vars.ErrLoginParams.Code,
			"message": vars.ErrLoginParams.Msg,
		})
		return
	}

	if !security.CheckPasswordHash(password, managerColl.Password) {
		log.Error(fmt.Sprintf("password error.username=%s, password=%s, saved_password=%s", username, password, managerColl.Password))
		c.JSON(400, gin.H{
			"code":    vars.ErrLoginParams.Code,
			"message": vars.ErrLoginParams.Msg,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  managerColl.ManagerId,
		"exp":      time.Now().Add(time.Hour * time.Duration(cfg.Cfg.LoginTokenExpire)).Unix(),
		"username": managerColl.Managername,
	})

	tokenStr, err := token.SignedString([]byte(cfg.Cfg.LoginSecret))
	if err != nil {
		log.Error("gen token failed. err=" + err.Error())
		c.JSON(400, gin.H{
			"message": "gen token failed.",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  0,
		"token": tokenStr,
	})
}
