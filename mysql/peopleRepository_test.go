package mysql

import (
	"database/sql"
	"errors"
	"feedwell/invitations"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"strings"
	"testing"
)

func TestGetAllUsersShouldReturnAllUsersInDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error opening stub db connection: %s ", err.Error())
	}

	defer db.Close()

	//Creating repo with mock database
	peopleRepository := NewPeopleRepository(gettingMockConnection(db))

	//Setting expectations
	rows := sqlmock.NewRows([]string{"entryID", "name", "email", "registered"}).
		AddRow(1, "pepito", "pepito@email.com", true).
		AddRow(2, "jaimito", "jaimito@email.com", false)

	mock.ExpectQuery("^SELECT (.+) from USERS$").WillReturnRows(rows)

	expectedUsers := 2

	//Acting
	users, errQuery := peopleRepository.GetAllPeople()

	// Asserting
	if errQuery != nil {
		t.Errorf("Error: %s", errQuery)
	} else if len(users) != expectedUsers {
		t.Errorf("Number of returned users is different than expected. Expected: %d / Returned: %d", expectedUsers, len(users))
	}
}

func TestWhenThereisAnErrorPreparingTheQueryAnErrorShouldBeReturned(t *testing.T) {
	//Setup test and mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error opening stub db connection: %s ", err.Error())
	}
	defer db.Close()

	//Creating repo with mock database
	peopleRepository := NewPeopleRepository(gettingMockConnection(db))

	//setting expectations
	expectedError := "an expected error"
	mock.ExpectQuery("^SELECT (.+) from USERS$").WillReturnError(errors.New(expectedError))

	//Acting
	_, errRep := peopleRepository.GetAllPeople()

	//Asserting
	if errRep == nil {
		t.Fatalf("Not error returned")
	}
	if !strings.Contains(errRep.Error(), expectedError) {
		t.Fatalf("Wrong error returned: %s", errRep.Error())
	}
}

func TestWhenRequstingByIdAPersonShouldBeReturned(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error opening stub db connection: %s ", err.Error())
	}
	defer db.Close()

	//Creating repo with mock database
	peopleRepository := NewPeopleRepository(gettingMockConnection(db))

	rows := sqlmock.NewRows([]string{"entryID", "name", "email", "registered"}).
		AddRow(1, "pepito", "pepito@email.com", true)

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)$").WithArgs("1").WillReturnRows(rows)

	//Acting
	person, errRep := peopleRepository.GetPersonByID("1")

	// Asserting
	if errRep != nil {
		t.Errorf("Error: %s", errRep)
	} else if person.PersonBase.Name != "pepito" {
		t.Errorf("Returned person name is not right: %s", person.PersonBase.Name)
	}
}

func TestWhenRequestingByNonExistentIdAnErrorShouldBeReturned(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error opening stub db connection: %s ", err.Error())
	}
	defer db.Close()

	//Creating repo with mock database
	peopleRepository := NewPeopleRepository(gettingMockConnection(db))

	expectedError := "An expected error"
	mock.ExpectQuery("^SELECT (.+) FROM users WHERE (.+)$").WithArgs("1").WillReturnError(errors.New(expectedError))

	//Acting
	_, errRep := peopleRepository.GetPersonByID("1")

	// Asserting
	if errRep == nil {
		t.Fatalf("Not error returned")
	}
	if !strings.Contains(errRep.Error(), expectedError) {
		t.Fatalf("Wrong error returned: %s", errRep.Error())
	}
}

func TestWhenAPersonIsAddedThePersonIsRegisteredInTheDB(t *testing.T) {

	db, mock, peopleRepository := prepareMockedRepository(t)
	defer db.Close()

	//setting expectations
	mock.ExpectExec("INSERT INTO users").WithArgs("Peter", "Peter@cool.com", false).WillReturnResult(sqlmock.NewResult(1, 1))

	//acting
	errRep := peopleRepository.AddPerson(giveMeAPerson())

	//asserting
	errExpec := mock.ExpectationsWereMet()
	if errExpec != nil {
		t.Errorf("Expectations not met: %s", errExpec.Error())
	}

	if errRep != nil {
		t.Errorf("Error: %s", errRep.Error())
	}
}

func gettingMockConnection(db *sql.DB) Connection {
	//Creating mock database
	con := new(MockDbConnection)
	con.mockdb = db
	return con
}

func prepareMockedRepository(t *testing.T) (*sql.DB, sqlmock.Sqlmock, invitations.Repository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error opening stub db connection: %s ", err.Error())
	}
	peopleRepository := NewPeopleRepository(gettingMockConnection(db))
	return db, mock, peopleRepository
}

// MockDbConnection ...
type MockDbConnection struct {
	mockdb *sql.DB
}

//OpenConnection open an mock slq connection
func (db *MockDbConnection) OpenConnection(stringConn string) (*sql.DB, error) {
	return db.mockdb, db.mockdb.Ping()
}

func giveMeAPerson() *invitations.Person {
	person := invitations.Person{Name: "Peter", Email: "Peter@cool.com"}
	return &person
}
