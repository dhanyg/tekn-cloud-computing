package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

// Fungsi untuk menghubungkan ke database
func connectDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/tekn_cloud")
    if err != nil {
        return nil, err
    }
    return db, nil
}

// Fungsi untuk mengambil data dari tabel 'persons'
func fetchDataFromPersonsTable(db *sql.DB) ([]Person, error) {
    query := "SELECT * FROM persons"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
	fmt.Println("Database connected..")
	fmt.Println("\n")

    var persons []Person
	fmt.Println("Get all persons record: \n")

    for rows.Next() {
        var person Person
        if err := rows.Scan(&person.ID, &person.Name, &person.Address, &person.Email); err != nil {
            return nil, err
        }
        persons = append(persons, person)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return persons, nil
}

// Struct untuk menyimpan data 'persons'
type Person struct {
    ID   int
    Name string
    Address string
	Email string
}

func main() {
    // Menghubungkan ke database
    db, err := connectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Mengambil data dari tabel 'persons'
    persons, err := fetchDataFromPersonsTable(db)
    if err != nil {
        log.Fatal(err)
    }

    // Menampilkan data
    for _, person := range persons {
        fmt.Printf("ID: %d, Name: %s, Address: %s, Email: %s\n", person.ID, person.Name, person.Address, person.Email)
    }
}
