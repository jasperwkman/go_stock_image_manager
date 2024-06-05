package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"go_stock_image_manager/database"
	"go_stock_image_manager/models"

	"github.com/gin-gonic/gin"
)

// IndexHandler handles GET requests for the / route.
// It retrieves the image inventory from the database and displays it on the page.
func IndexHandler(c *gin.Context, db *sql.DB) {
	fmt.Println("Received GET request for /")

	// Query the database for the image inventory
	rows, err := database.QueryImageInventory(db)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	imageInventoryRows := []models.ImageInventoryRow{}

	// Iterate over the rows and scan them into the imageInventoryRows slice
	for rows.Next() {
		var row models.ImageInventoryRow
		var imageThumbnail []byte

		err = rows.Scan(&row.ImageFname, &row.ImageCamera, &row.ImageDescription, &row.ImageTag, &row.FoapStatus, &row.ShutterstockStatus, &row.AlamyStatus, &row.ImageReady, &imageThumbnail, &row.CreatedDt)
		if err != nil {
			panic(err)
		}

		// Convert the image thumbnail to a base64 encoded string and set it in the row struct
		row.ImageThumbnail = template.URL(fmt.Sprintf("data:image/jpeg;base64,%s", imageThumbnail))

		imageInventoryRows = append(imageInventoryRows, row)
	}

	fmt.Printf("Rendering index.tmpl with %d rows of data\n", len(imageInventoryRows))

	// Render the HTML template with the image inventory
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"imageInventoryRows": imageInventoryRows,
	})
}
