package main

import (
	loginendpoint "FaceAnnotation/service/api/login"
	userendpoint "FaceAnnotation/service/api/user"
	middleware "FaceAnnotation/service/middleware"
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/itsjamie/gin-cors"
)

var (
	port  = flag.Int64("p", 80, "port")
	debug = flag.Bool("d", false, "debug model")
)

func ready() {

	flag.Parse()
	if *port == 8060 {
		gin.SetMode(gin.ReleaseMode)
		log.Info(fmt.Sprintf("ReleaseMode"))
	}
}

func main() {
	ready()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "POST",
		RequestHeaders:  "Origin, Authorization, Content-Type, Access-Control-Allow-Headers, LoginToken",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	r.Static("origin_images", "./origin_images")

	authorized := r.Group("/user")
	authorized.Use(middleware.AuthToken)
	{
		authorized.POST("login", loginendpoint.Login)
	}

	if *debug == false {
		r.POST("add_user", userendpoint.AddUser)
	}

	p := fmt.Sprintf(":%d", *port)
	log.Info("listen port " + p)
	r.Run(p)
}
