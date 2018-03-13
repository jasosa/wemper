package mysql

import (
	"database/sql"
	"feedwell/invitations"
	"fmt"
	"strconv"
	//we will use it
	_ "github.com/go-sql-driver/mysql"
)

//Connection ...
type Connection interface {
	OpenConnection(stringConn string) (*sql.DB, error)
}

//PeopleRepository mysql repository for people entities
type PeopleRepository struct {
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
func (pr PeopleRepository) GetAllPeople() ([]invitations.AppUser, error) {

	db, err := openConnection(pr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * from USERS")
	if err != nil {
		return nil, fmt.Errorf("executing query failed: %s", err.Error())
	}

	users := make([]invitations.AppUser, 0)

	for rows.Next() {
		var entryID int
		var name string
		var email string
		var registered bool
		var admin bool
		err = rows.Scan(&entryID, &name, &email, &registered, &admin)

		var appUser invitations.AppUser
		if admin {
			appUser = invitations.NewAdminUser(strconv.Itoa(entryID), name, email)
		} else if registered {
			appUser = invitations.NewRegisteredUser(strconv.Itoa(entryID), name, email)
		} else {
			appUser = invitations.NewInvitedUser(strconv.Itoa(entryID), name, email)
		}

		users = append(users, appUser)
	}

	return users, nil
}

//GetPersonByID gets the person with the specified id from the db source
func (pr PeopleRepository) GetPersonByID(id string) (invitations.AppUser, error) {
	db, err := openConnection(pr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entryID int
	var name string
	var email string
	var registered bool
	var admin bool
	errScan := db.QueryRow("SELECT entryId, name, email, registered, admin FROM users WHERE entryID = ?", id).Scan(&entryID, &name, &email, &registered, &admin)

	if errScan != nil {
		return nil, fmt.Errorf("Executing query failed: %s", errScan.Error())
	}

	var appUser invitations.AppUser
	if admin {
		appUser = invitations.NewAdminUser(strconv.Itoa(entryID), name, email)
	} else if registered {
		appUser = invitations.NewRegisteredUser(strconv.Itoa(entryID), name, email)
	} else {
		appUser = invitations.NewInvitedUser(strconv.Itoa(entryID), name, email)
	}

	return appUser, nil
}

//AddPerson adds a new person to the db source
func (pr *PeopleRepository) AddPerson(p invitations.AppUser) error {
	db, err := openConnection(*pr)
	if err != nil {
		return err
	}
	defer db.Close()

	result, errExec := db.Exec("INSERT INTO users (name, email,registered, admin) VALUES (?, ?, ?, ?)",
		p.GetPersonInfo().Name,
		p.GetPersonInfo().Email,
		p.GetPersonInfo().Registered,
		p.GetPersonInfo().Admin)

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
