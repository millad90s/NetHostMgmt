package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

//hashmaker: convert bytes into hash
func hashmaker(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash)
}

//Register ...
func Register(router *gin.Engine) {
	router.POST("/registeruser", func(c *gin.Context) {
		db, err := connectdb()
		if err != nil {
			c.String(http.StatusNotFound, "db connectio problem")
			return
		}
		if !db.HasTable(Owner{}) {
			db.CreateTable(Owner{})
		}
		uname := c.PostForm("username")
		name := c.DefaultPostForm("name", uname)
		// password := c.PostForm("password")
		password := hashmaker(c.PostForm("password"))
		email := c.PostForm("email")
		// password := c.DefaultPostForm("password", "anonymous")
		o := Owner{
			Username: uname,
			Name:     name,
			Password: password,
			Email:    email,
		}
		//todo: error handeling
		res := db.Debug().Create(&o)
		if res.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"result": res.Error.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"result": "user added succeesuly :)",
			"userid": o.ID,
		})
	})
}

//Login ...
func Login(router *gin.Engine) {
	router.POST("/login", func(c *gin.Context) {
		db, err := connectdb()
		if err != nil {
			c.String(http.StatusNotFound, "db connectio problem")
			return
		}
		uname := c.PostForm("username")
		password := c.PostForm("password")
		// password := c.DefaultPostForm("password", "anonymous")
		o := Owner{Username: uname}
		db.Where(&o).Find(&o)
		if comparePasswords(o.Password, []byte(password)) {
			c.JSON(200, gin.H{
				"result":  "Okay",
				"message": "you are authorized , welcome " + uname,
			})
			return
		}
		c.JSON(200, gin.H{
			"result":  "NotOkay",
			"message": "you are Not authorized with this credentials",
		})
	})
}

//UserInfo ...
func UserInfo(router *gin.Engine) {
	router.GET("/userinfo/:username", func(c *gin.Context) {
		db, err := connectdb()
		if err != nil {
			c.String(http.StatusNotFound, "db connectio problem")
			return
		}
		uname := c.Param("username")
		o := Owner{Username: uname}
		// db.Debug().Find(&o, Owner{Username: uname})
		db.Where(&o).Find(&o)
		if o.ID == 0 {
			c.String(http.StatusNotFound, "User Not found ")
			return
		}
		json.Marshal(&o)
		fmt.Printf("%v", o)
		c.String(http.StatusOK, "The name of thi suser is %s: ", o.Name)
	})
}

//just for test ... it will be deleted
func test(router *gin.Engine) {

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
}

func ping(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
