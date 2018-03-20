package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	//needed for initialization
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasosa/wemper/invitations"
	"strconv"
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

	var query = "SELECT * from USERS"
	rows, err := db.Query(query)
	if err != nil {
		return nil, &SQLQueryError{Query: query, BaseError: err}
	}

	users := make([]invitations.AppUser, 0)
	for rows.Next() {
		var entryID int
		var name, email string
		var registered, admin bool
		err = rows.Scan(&entryID, &name, &email, &registered, &admin)
		if err != nil {
			return nil, &SQLQueryError{Query: query, BaseError: err}
		}
		users = append(users, pr.createUser(entryID, name, email, registered, admin))
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
	var query = "SELECT entryId, name, email, registered, admin FROM users WHERE entryID = ?"
	errScan := db.QueryRow(query, id).Scan(&entryID, &name, &email, &registered, &admin)

	if errScan != nil {
		return nil, &SQLQueryError{Query: query, BaseError: errScan}
	}

	appUser := pr.createUser(entryID, name, email, registered, admin)
	return appUser, nil
}

//AddPerson adds a new person to the db source
func (pr *PeopleRepository) AddPerson(p invitations.AppUser) error {
	db, err := openConnection(*pr)
	if err != nil {
		return err
	}
	defer db.Close()

	var query = "INSERT INTO users (name, email,registered, admin) VALUES (?, ?, ?, ?)"
	result, errExec := db.Exec(query,
		p.GetPersonInfo().Name,
		p.GetPersonInfo().Email,
		p.GetPersonInfo().Registered,
		p.GetPersonInfo().Admin)

	if errExec != nil {
		return &SQLQueryError{Query: query, BaseError: errExec}
	}

	rowsAffected, errorRowsAff := result.RowsAffected()
	if errorRowsAff != nil {
		return &SQLQueryError{Query: query, BaseError: errorRowsAff}
	}

	if rowsAffected != 1 {
		return &SQLQueryError{Query: query, BaseError: errors.New("1 row affected was expected, but %s rows were affected")}
	}

	_, errorRLI := result.LastInsertId()
	if errorRLI != nil {
		return &SQLQueryError{Query: query, BaseError: errorRLI}
	}

	return nil
}

func (pr PeopleRepository) createUser(entryID int, name, email string, registered, admin bool) invitations.AppUser {
	var appUser invitations.AppUser
	if admin {
		appUser = invitations.NewAdminUser(strconv.Itoa(entryID), name, email)
	} else if registered {
		appUser = invitations.NewRegisteredUser(strconv.Itoa(entryID), name, email)
	} else {
		appUser = invitations.NewInvitedUser(strconv.Itoa(entryID), name, email)
	}
	return appUser
}

func openConnection(pr PeopleRepository) (*sql.DB, error) {
	stringCon := fmt.Sprintf("%s:%s@/%s", pr.user, pr.password, pr.databasename)
	db, err := pr.connection.OpenConnection(stringCon)
	if err != nil {
		return nil, &SQLOpeningDBError{BaseError: err}
	}
	return db, nil
}
