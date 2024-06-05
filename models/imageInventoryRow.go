package models

import (
	"html/template"
)

// ImageInventoryRow represents a row in the image_inventory table.
type ImageInventoryRow struct {
	ImageFname         string
	ImageCamera        string
	ImageDescription   string
	ImageTag           string
	FoapStatus         int
	ShutterstockStatus int
	AlamyStatus        int
	ImageReady         int
	CreatedDt          string
	ImageThumbnail     template.URL
}
