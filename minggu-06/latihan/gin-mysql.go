package main

import (
    "database/sql"
    // "fmt"
    "log"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

// Struct untuk menyimpan data 'persons'
type Person struct {
    ID      int     `json:"id"`
    Name    string  `json:"name"`
    Address string  `json:"address"`
    Email   string  `json:"email"`
}

func main() {
    // Membuat router Gin
    router := gin.Default()

    // Menghubungkan ke database
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/tekn_cloud")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Mengambil data dari tabel 'persons'
    router.GET("/persons", func(c *gin.Context) {
        var persons []Person
        query := "SELECT * FROM persons"
        rows, err := db.Query(query)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        for rows.Next() {
            var person Person
            if err := rows.Scan(&person.ID, &person.Name, &person.Address, &person.Email); err != nil {
                c.JSON(500, gin.H{"error": err.Error()})
                return
            }
            persons = append(persons, person)
        }

        c.JSON(200, persons)
    })

    // Menambahkan data ke dalam tabel 'persons'
    router.POST("/persons", func(c *gin.Context) {
        var newPerson Person
        if err := c.BindJSON(&newPerson); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        query := "INSERT INTO persons (name, address, email) VALUES (?, ?, ?)"
        _, err := db.Exec(query, newPerson.Name, newPerson.Address, newPerson.Email)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }

        c.JSON(201, newPerson)
    })

    // Menjalankan server Gin pada port 8080
    router.Run(":8080")
}
