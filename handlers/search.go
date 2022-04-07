package handlers

import (
	"jabar-nearby-places/services"
	"jabar-nearby-places/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var lat float64 = utils.ParseFloat64(c.Query("latitude"), -999)
	var lon float64 = utils.ParseFloat64(c.Query("longitude"), -999)

	founds, err := services.PlaceService.FindAll(c.Query("category_id"))
	if err != nil {
		c.Error(err)
		return
	}

	res := services.PlaceService.FilterByDistance(*founds, lat, lon)

	c.JSON(http.StatusOK, gin.H{"data": res, "total": len(res)})
}
