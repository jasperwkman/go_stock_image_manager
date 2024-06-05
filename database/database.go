package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connect connects to the MySQL database using the provided configuration values.
// It returns a pointer to the sql.DB object and an error if any.
func Connect(server, port, user, password, dbname string) (*sql.DB, error) {
	fmt.Printf("Connecting to MySQL database %s at %s:%s using username %s and password %s\n", dbname, server, port, user, password)

	// Construct the data source name (DSN) for the MySQL driver
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, server, port, dbname)
	// Open a connection to the database using the DSN
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL database successfully")

	return db, nil
}

// QueryImageInventory queries the image_inventory table and returns the results.
// It returns a pointer to the sql.Rows object and an error if any.
func QueryImageInventory(db *sql.DB) (*sql.Rows, error) {
	// Execute the query on the database
	rows, err := db.Query("SELECT image_fname,image_camera,image_description,image_tag,foap_status,shutterstock_status,alamy_status,image_ready, image_thumbnail,created_dt FROM image_inventory where image_ready = 0 LIMIT 20")
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// SetImageStatus updates the image_status column in the image_inventory table for the specified image.
// It takes a pointer to an sql.DB object, a string variable imageFname, a string variable stock, and an int variable imageStatus as input.
// It returns the number of rows updated and an error if any.
func SetImageStatus(db *sql.DB, imageFname string, stock string, imageStatus int) (int64, error) {
	// Prepare the UPDATE statement
	stmt, err := db.Prepare("UPDATE image_inventory SET " + stock + "_status = ? WHERE image_fname = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the UPDATE statement with the provided imageStatus and imageFname as parameters
	res, err := stmt.Exec(imageStatus, imageFname)
	if err != nil {
		return 0, err
	}

	// Get the number of rows updated
	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

// SetImageDescription updates the image_description column in the image_inventory table for the specified image.
// It takes a pointer to an sql.DB object, a string variable imageFname, and a string variable imageDescription as input.
// It returns the number of rows updated and an error if any.
func SetImageDescription(db *sql.DB, imageFname string, imageDescription string) (int64, error) {
	// Prepare the UPDATE statement
	stmt, err := db.Prepare("UPDATE image_inventory SET image_description = ? WHERE image_fname = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the UPDATE statement with the provided imageDescription and imageFname as parameters
	res, err := stmt.Exec(imageDescription, imageFname)
	if err != nil {
		return 0, err
	}

	// Get the number of rows updated
	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

// QueryImageDetail queries the image_inventory table for a specific image and returns the result.
// It returns a pointer to the sql.Row object and an error if any.
func QueryImageDetail(db *sql.DB, imageFname string) (*sql.Row, error) {
	// Execute the query on the database with the provided image filename as a parameter
	row := db.QueryRow("SELECT image_fname,image_camera,image_description,image_tag,foap_status,shutterstock_status,alamy_status,image_ready, image_thumbnail,created_dt FROM image_inventory WHERE image_fname = ?", imageFname)

	// Check for errors
	if err := row.Err(); err != nil {
		return nil, err
	}

	return row, nil
}

// QueryImageTagsGroup queries the image_tags_group table and returns all rows.
// It returns a slice of ImageTagsGroupRow and an error if any.
func QueryImageTagsGroup(db *sql.DB) (*sql.Rows, error) {
	// Execute the query on the database
	rows, err := db.Query("SELECT tag_group_name,tag_list FROM image_tags_group")
	if err != nil {
		fmt.Println("Query Err:" + err.Error())
		return nil, err
	}

	return rows, nil

}

func SetImageTags(db *sql.DB, imageFname, imageTags string) (int64, error) {
	fmt.Println("SetImageTags: " + imageFname + "  " + imageTags)
	// Prepare the UPDATE statement
	stmt, err := db.Prepare("UPDATE image_inventory SET image_tag = ? WHERE image_fname = ?")

	defer stmt.Close()
	if err != nil {
		return 0, err
	}

	// Execute the UPDATE statement with the provided imageDescription and imageFname as parameters
	res, err := stmt.Exec(imageTags, imageFname)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Get the number of rows updated
	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func SetImageReady(db *sql.DB) (int64, error) {
	fmt.Println("SetImageReady")
	// Prepare the UPDATE statement
	stmt, err := db.Prepare("UPDATE image_inventory SET image_ready = 1 where image_description != '' and image_tag != ''")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the UPDATE statement with the provided imageStatus and imageFname as parameters
	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	// Get the number of rows updated
	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}
