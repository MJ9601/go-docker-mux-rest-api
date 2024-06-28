package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) * MySQLStorage{
	db, err := sql.Open("mysql", cfg.FormatDSN())
	
	if err != nil {
		 log.Fatal(err)
	}

	err = db.Ping()

	if err != nil{
		log.Fatal(err)
	}
	log.Println("Connected to db ...")

	return &MySQLStorage{db}
}


func (s *MySQLStorage) Init() (*sql.DB, error) {
	// initialize the tables
	// err := s.createUserTables()
	// if err != nil {
	// 	 log.Fatal("User table creation failed")
	// }

	// log.Println("User tables created")

	// err = s.createProjectsTable()
	// if err != nil {
	// 	 log.Fatal("Project table creation failed")
	// }
	// log.Println("Project tables created")

	// err = s.createTaskTables()

	// if err != nil {
	// 	 log.Fatal("Task table creation failed")
	// }

	// log.Println("Task tables created")


	return s.db, nil
}

func (s *MySQLStorage) createProjectsTable() error {

	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS projects (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id)
	)	ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}


func (s *MySQLStorage) createTaskTables() error {
	 _, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		status ENUM("TODO", "IN_PROGRESS", "IN_TEST", "DONE")  NOT NULL DEFAULT "TODO",
		projectID INT UNSIGNED NOT NULL,
		assignedToID INT UNSIGNED NOT NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id),
		FOREIGN KEY (projectID) REFERENCES project(id),
		FOREIGN KEY (assignedToID) REFERENCES users(id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8; 
	 `)

	return err
}

func (s *MySQLStorage) createUserTables() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		email VARCHAR(255) NOT NULL,
		firstName VARCHAR(255) NOT NULL,
		lastName VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id),
		UNIQUE KEY (email)
	)	ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}

