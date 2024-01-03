package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type users struct {
	Name    string `json:"Name"`
	Age     int    `json:"age"`
	IsAdmin bool   `json:"IsAdmin"`
	Tasks   int   `json:"tasks"`
}


var people = []users{

	{Name:"Randy",Age:20,IsAdmin: false,Tasks:0},
	{Name:"Paul",Age:24,IsAdmin: false,Tasks:0},
	{Name:"Jenny",Age:33,IsAdmin: false,Tasks:0},
}


func getUsers(context *gin.Context){
   context.IndentedJSON(http.StatusOK,people)
}
