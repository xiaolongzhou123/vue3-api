package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type RS struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}

func main() {
	r := gin.Default()
	arr := make([]Data, 0)
	a := Data{
		Id:   1,
		Name: "aa",
	}
	b := Data{
		Id:   2,
		Name: "bb",
	}
	arr = append(arr, a, b)

	rs := &RS{
		Code:    200,
		Message: "ok",
		Data:    arr,
	}

	r.GET("/nav", func(c *gin.Context) {
		c.JSON(http.StatusOK, rs)
	})
	r.Run(":10001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
