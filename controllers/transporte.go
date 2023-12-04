package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeremiascardoso00/DABD-COLOQUIO/models"
	"gorm.io/gorm"
)

type TransporteController struct {
	Txn *gorm.DB
}

func (tc *TransporteController) GetTransportes(c *gin.Context) {
	var transporte models.Transporte

	query := tc.Txn.Find(&transporte)
	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transporte})
}

func (tc *TransporteController) GetTransporteByID(c *gin.Context) {
	transporteID := c.Param("id")
	var transporte models.Itinerario

	query := tc.Txn.First(&transporte, transporteID)
	if query.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transporte not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transporte})
}

func (tc *TransporteController) DeleteTransporteByID(c *gin.Context) {
	transporteID := c.Param("id")
	var transporte models.Itinerario

	query := tc.Txn.First(&transporte, transporteID)
	if query.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transporte not found"})
		return
	}

	// Eliminar el itinerario
	tc.Txn.Delete(&transporte)

	c.JSON(http.StatusOK, gin.H{"message": "Transporte deleted successfully"})
}
