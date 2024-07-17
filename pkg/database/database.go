package database

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type DB struct {
	neo4j.Driver
}

// GetDriver initializes a new Neo4j driver.
func GetDriver(uri, username, password string) (neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func NewDatabase(uri, username, password string) *DB {
	driver, err := GetDriver(uri, username, password)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		Driver: driver,
	}
}

// GetSession creates a new Neo4j session.
func (driver *DB) GetSession() (neo4j.Session, error) {
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite, // Adjust access mode as needed
		DatabaseName: "neo4j",               // Specify the database name
	})
	return session, nil
}

func CloseDriver(driver neo4j.Driver) {
	driver.Close()
}

// CloseSession closes the Neo4j session.
func CloseSession(session neo4j.Session) {
	if err := session.Close(); err != nil {
		log.Printf("Error closing Neo4j session: %v", err)
	}
}

// ValidateDriver checks the connectivity of the Neo4j driver.
func ValidateDriver(driver neo4j.Driver) error {
	return driver.VerifyConnectivity()
}
