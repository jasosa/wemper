package mysql

import (
	"database/sql"
	"feedwell/invitations"
	"fmt"
	"strconv"
	//we will use it
	_ "github.com/go-sql-driver/mysql"
)

//PeopleRepository mysql repository for people entities
type PeopleRepository struct {
	//TODO: use encrypted credentials instead user and password
	connection   Connection
	user         string
	password     string
	databasename string
}

//NewPeopleRepository creates a new instance of mysql people repository
func NewPeopleRepository(connection Connection) invitations.Repository {
	pr := new(PeopleRepository)
	pr.user = "wempathy"
	pr.password = "wempathy2018"
	pr.databasename = "wempathy"
	pr.connection = connection
	return pr
}

// GetAllPeople gets all people from db source
func (pr PeopleRepository) GetAllPeople() ([]invitations.User, error) {

	db, err := openConnection(pr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	/* 	queryStmt, err := db.Prepare("SELECT * from USERS")
	   	if err != nil {
	   		return nil, fmt.Errorf("preparing query failed: %s", err.Error())
	   	}
	   	defer queryStmt.Close() */

	rows, err := db.Query("SELECT * from USERS")
	if err != nil {
		return nil, fmt.Errorf("executing query failed: %s", err.Error())
	}

	users := make([]invitations.User, 0)

	for rows.Next() {
		var entryID int
		var name string
		var email string
		var registered bool
		err = rows.Scan(&entryID, &name, &email, &registered)
		person := invitations.Person{
			ID:    strconv.Itoa(entryID),
			Email: email,
			Name:  name,
		}
		users = append(users, invitations.User{PersonBase: person, Registered: registered})
	}

	return users, nil
}

//GetPersonByID gets the person with the specified id from the db source
func (pr PeopleRepository) GetPersonByID(id string) (*invitations.User, error) {
	db, err := openConnection(pr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entryID int
	var name string
	var email string
	var registered bool
	errScan := db.QueryRow("SELECT entryId, name, email, registered FROM users WHERE entryID = ?", id).Scan(&entryID, &name, &email, &registered)
	if errScan != nil {
		return nil, fmt.Errorf("Executing query failed: %s", errScan.Error())
	}

	person := invitations.Person{
		ID:    strconv.Itoa(entryID),
		Email: email,
		Name:  name,
	}

	user := invitations.User{
		PersonBase: person,
		Registered: registered,
	}

	return &user, nil
}

//AddPerson adds a new person to the db source
func (pr *PeopleRepository) AddPerson(p *invitations.Person) error {
	db, err := openConnection(*pr)
	if err != nil {
		return err
	}
	defer db.Close()

	result, errExec := db.Exec("INSERT INTO users (name, email,registered) VALUES (?, ?, ?)", p.Name, p.Email, false)

	if errExec != nil {
		return fmt.Errorf("Inserting person failed: %s", errExec.Error())
	}

	rowsAffected, errorRowsAff := result.RowsAffected()

	if errorRowsAff != nil {
		return fmt.Errorf("Inserting person failed: %s", errorRowsAff.Error())
	}

	if rowsAffected != 1 {
		return fmt.Errorf("Inserting person failed: Rows affected (%d)", rowsAffected)
	}

	_, errorRLI := result.LastInsertId()

	if errorRLI != nil {
		return fmt.Errorf("Inserting person failed: %s", errorRLI.Error())
	}

	return nil
}

func openConnection(pr PeopleRepository) (*sql.DB, error) {
	stringCon := fmt.Sprintf("%s:%s@/%s", pr.user, pr.password, pr.databasename)
	db, err := pr.connection.OpenConnection(stringCon)
	if err != nil {
		return nil, fmt.Errorf("mySQLConnection failed: %s", err.Error())
	}
	return db, nil
}
