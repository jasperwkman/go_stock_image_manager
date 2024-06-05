package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go_stock_image_manager/database"
	"go_stock_image_manager/models"

	"github.com/gin-gonic/gin"
)

// ImageDetailHandler handles GET requests for the /imageDetail route.
// It retrieves the image details from the database and displays them on the page.
func ImageDetailHandler(c *gin.Context, db *sql.DB) {

	// Get the image filename from the query string
	imageFname := c.Query("imageFname")

	// Query the database for the image details
	rows, err := database.QueryImageDetail(db, imageFname)
	if err != nil {
		panic(err)
	}

	// Create variables to hold the image details and thumbnail
	var imageDetail models.ImageInventoryRow
	var imageThumbnail []byte

	// Scan the returned row into the imageDetail and imageThumbnail variables
	err = rows.Scan(&imageDetail.ImageFname, &imageDetail.ImageCamera, &imageDetail.ImageDescription, &imageDetail.ImageTag, &imageDetail.FoapStatus, &imageDetail.ShutterstockStatus, &imageDetail.AlamyStatus, &imageDetail.ImageReady, &imageThumbnail, &imageDetail.CreatedDt)
	if err != nil {
		panic(err)
	}

	// Initialize the ImageThumbnail field to an empty string
	imageDetail.ImageThumbnail = ""

	// If the image thumbnail is not nil, convert it to a base64 encoded string and store it in the ImageThumbnail field
	if imageThumbnail != nil {
		imageDetail.ImageThumbnail = template.URL(fmt.Sprintf("data:image/jpeg;base64,%s", imageThumbnail))
	}

	// Query the database for the image tags group
	rows2, err := database.QueryImageTagsGroup(db)
	if err != nil {
		fmt.Println("QueryImageTagsGroup Error: " + err.Error())
		panic(err)
	}
	defer rows2.Close()

	// Create variables to hold the image details and thumbnail
	imageTagsGroupRows := []models.ImageTagsGroupRow{}

	// Iterate over the rows and scan them into the imageTagsGroupRows slice
	for rows2.Next() {
		var currentRow models.ImageTagsGroupRow
		err = rows2.Scan(&currentRow.TagGroupName, &currentRow.TagList)
		if err != nil {
			panic(err)
		}

		imageTagsGroupRows = append(imageTagsGroupRows, currentRow)
	}
	// Render the imageDetail.tmpl template with the provided data
	c.HTML(http.StatusOK, "imageDetail.tmpl", gin.H{
		"ImageFname":         imageDetail.ImageFname,
		"ImageCamera":        imageDetail.ImageCamera,
		"ImageDescription":   imageDetail.ImageDescription,
		"ImageTag":           imageDetail.ImageTag,
		"FoapStatus":         imageDetail.FoapStatus,
		"ShutterstockStatus": imageDetail.ShutterstockStatus,
		"AlamyStatus":        imageDetail.AlamyStatus,
		"CreatedDt":          imageDetail.CreatedDt,
		"ImageThumbnail":     template.URL(imageDetail.ImageThumbnail),
		"ImageTagsGroupRows": imageTagsGroupRows,
	})

}

// imageDetail.go
// SetImageStatusHandler handles POST requests for the /setImageStatus route.
// It receives JSON data and calls the SetImageStatus function to update the image status in the database.
func SetImageStatusHandler(c *gin.Context, db *sql.DB) {
	// Parse the JSON request body
	var req struct {
		ImageFname  string `json:"imageFname"`
		Stock       string `json:"stock"`
		StatusValue int    `json:"statusValue"`
	}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"rowsUpdated": -1, "error": err.Error()})
		return
	}

	// Call the SetImageStatus function with the provided data
	rowsUpdated, err := database.SetImageStatus(db, req.ImageFname, req.Stock, req.StatusValue)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"rowsUpdated": -1, "error": err.Error()})
		return
	}

	// Return a success response with the number of rows updated
	c.JSON(http.StatusOK, gin.H{"rowsUpdated": rowsUpdated, "error": "no"})
}

// SetImageDesctiptionHandler updates the image_description column in the image_inventory table for the specified image.
// It takes a pointer to an sql.DB object, a string variable imageFname, and a string variable imageDescription as input.
// It returns the number of rows updated and an error if any.
func SetImageDesctiptionHandler(c *gin.Context, db *sql.DB) {

	var req struct {
		ImageFname       string `json:"imageFname"`
		ImageDescription string `json:"imageDescription"`
	}
	if err := c.BindJSON(&req); err != nil {
		log.Println(err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"rowsUpdated": -1, "error": err.Error()})
		return
	}

	// Call the SetImageStatus function with the provided data
	rowsUpdated, err := database.SetImageDescription(db, req.ImageFname, req.ImageDescription)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"rowsUpdated": -1, "error": err.Error()})
		return
	}

	// Return a success response with the number of rows updated
	c.JSON(http.StatusOK, gin.H{"rowsUpdated": rowsUpdated, "error": "no"})
}

func SetImageTagsHandler(c *gin.Context, db *sql.DB) {

	// Parse request body
	var req struct {
		ImageFname string `json:"imageFname"`
		ImageTags  string `json:"imageTags"`
	}

	if err := c.BindJSON(&req); err != nil {
		log.Println(err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"rowsUpdated": -1, "error": err.Error()})
		return
	}

	// Call the SetImageStatus function with the provided data
	rowsUpdated, err := database.SetImageTags(db, req.ImageFname, req.ImageTags)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"rowsUpdated": -1, "error": err.Error()})
		return
	}

	// Return a success response with the number of rows updated
	c.JSON(http.StatusOK, gin.H{"rowsUpdated": rowsUpdated, "error": "no"})
}

func SetImageReadyHandler(c *gin.Context, db *sql.DB) {

	// Call the SetImageStatus function with the provided data
	rowsUpdated, err := database.SetImageReady(db)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"rowsUpdated": -1, "error": err.Error()})
		return
	}

	// Return a success response with the number of rows updated
	c.JSON(http.StatusOK, gin.H{"rowsUpdated": rowsUpdated, "error": "no"})
}
