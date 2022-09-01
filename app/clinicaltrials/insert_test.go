package clinicaltrials

import (
	"context"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"

	//_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	//os.Remove("test.db") // I delete the file to avoid duplicated records.
	//// SQLite is a file based database.
	//
	//log.Println("Creating sqlite-database.db...")
	//file, err := os.Create("test.db") // Create SQLite file
	//assert.Nil(t, err)
	//file.Close()
	//log.Println("test.db created")

	ctx := context.Background()

	//testID := "NCT03197376"

	//testID = "NCT03547167"

	//testIDs := []string{
		//"NCT03950856",
		//"NCT03950622",
		//"NCT03835975",
		//"NCT03828617",
		//"NCT03760146",
		//"NCT03547167",
		//"NCT03480802",
		//"NCT03480763",
		//"NCT03313037", // This one presents an issue we need to look into (invalid array length)
		//"NCT03197376",

		//"NCT02787863",
		//"NCT02717494",
		//"NCT02573181",
		//"NCT02547649",
		//"NCT02308540",
		//"NCT02225587",
		//"NCT02097472",
		//"NCT02037984",
		//"NCT01646398",
		//"NCT01641133",

		//"NCT01545375",
		//"NCT01513551",
		//"NCT01215188",
		//"NCT01215175",
		//"NCT00962780",
		//"NCT00836641",
		//"NCT00689351",
		//"NCT00688870",
		//"NCT00680914",
		//"NCT00676091",

		//"NCT00574548",
		//"NCT00546572",
		//"NCT00508742",
		//"NCT00475033",
		//"NCT00474539",
		//"NCT00463437",
		//"NCT00457977",
		//"NCT00452790",
		//"NCT00444457",
		//"NCT00427895",

		//"NCT00384059",
		//"NCT00373958",
		//"NCT00370396",
		//"NCT00368966",
		//"NCT00366899",
		//"NCT00366678",
		//"NCT00366340",
		//"NCT00205803", // Issue parsing date at insert
	//}

	driver := "postgres"
	dataSource := "host=localhost port=18034 user=local dbname=local password=password sslmode=disable"

	//driver = "sqlite3"
	//dataSource = "file:./test.db?cache=shared&mode=memory&_fk=true"

	db, err := models.Open(driver, dataSource)
	assert.Nil(t, err)
	assert.Nil(t, db.Schema.Create(ctx))

	//rc := NewResultsContext(ctx, db)

	//for _, testIdInsert := range testIDs {
	//	assert.Nil(t, rc.Insert(testIdInsert, false))
	//}

	_, err = GetTrial("NCT03197376")
	assert.Nil(t, err)

	//assert.Nil(t, rc.InsertAll(study, GetOutcomeOverviewList(study)), []Metadata{})

	//assert.Nil(t, rc.Insert(testID, false))
}
