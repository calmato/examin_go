package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	jwt-tokenを検証する
	tokenからlogin_idとroleを取得する
*/
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginID, role, err := Parse(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":      http.StatusText(http.StatusUnauthorized),
				"description": "認証エラーです．",
			})
			c.Abort()
			return
		}

		c.Set("LoginID", loginID)
		c.Set("Role", role)
		c.Next()
	}
}

/*
	studentsの権限確認をする
	前提条件としては，上のAuthの後に呼び出すこと
	router.goでmiddlewareを使用するrouter.Groupを定義する
	ex.)
	studnets := router.Group("/studnets")
	students.Use(middleware.StudentsAuth())
*/
func StudentsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleCheck(c, 0)
	}
}

/*
	teachersの権限確認をする
	前提条件としては，上のAuthの後に呼び出すこと
	router.goでmiddlewareを使用するrouter.Groupを定義する
	ex.)
	teachers := router.Group("/teachers")
	teachers.Use(middleware.TeachersAuth())
*/
func TeachersAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleCheck(c, 1)
	}
}

func roleCheck(c *gin.Context, r int) {
	role := c.MustGet("Role")
	if role != r {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":      http.StatusText(http.StatusForbidden),
			"description": "認証エラーです．",
		})
		c.Abort()
		return
	}
	c.Next()
}
