package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Mahasiswa struct {
		Nim  int
		Nama string
	}
	router := gin.Default()

	// GET a mahasiswa detail
	router.GET("/mahasiswa/:nim", func(c *gin.Context) {
		var (
			mahasiswa Mahasiswa
			result    gin.H
		)
		nim := c.Param("nim")
		row := db.QueryRow("select nim, name from mahasiswa where nim = ?;", nim)
		err = row.Scan(&mahasiswa.Nim, &mahasiswa.Nama)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": mahasiswa,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all mahasiswa
	router.GET("/mahasiswas", func(c *gin.Context) {
		var (
			mahasiswa  Mahasiswa
			mahasiswas []Mahasiswa
		)
		rows, err := db.Query("select nim, nama from mahasiswa;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&mahasiswa.Nim, &mahasiswa.Nama)
			mahasiswas = append(mahasiswas, mahasiswa)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": mahasiswa,
			"count":  len(mahasiswas),
		})
	})

	// POST new pmahasiswa details
	router.POST("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		nim := c.PostForm("nim")
		nama := c.PostForm("nama")
		stmt, err := db.Prepare("insert into mahasiswa (nim, nama) values(?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nim, nama)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(nim)
		buffer.WriteString(" ")
		buffer.WriteString(nama)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", name),
		})
	})

	// PUT - update a mahasiswa details
	router.PUT("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		nim := c.Query("nim")
		nama := c.PostForm("nama")
		stmt, err := db.Prepare("update mahasiswa set nim= ?, nama= ? where nim= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nama, nim)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(nim)
		buffer.WriteString(" ")
		buffer.WriteString(nama)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name),
		})
	})

	// Delete resources
	router.DELETE("/mahasiswa", func(c *gin.Context) {
		nim := c.Query("nim")
		stmt, err := db.Prepare("delete from mahasiswa where nim= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nim)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted mahasiswa: %s", nim),
		})
	})
	router.Run(":3000")
}
