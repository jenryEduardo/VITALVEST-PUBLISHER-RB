package controllers

import (
	"net/http"
	"publisher/application"
	"publisher/domain"
	"publisher/infraestructure/adapters"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {


	var data domain.DatosSensor

	if err:= c.ShouldBindJSON(&data);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se encontro los mismos datos o son nulos"})
	}

	repo := adapters.NewConn()
	use_Case:=application.NewGetData(repo)

	if err:=use_Case.Execute(data);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo enviar los datos"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"ok":"se publico el objeto a la cola correctamente"})

}