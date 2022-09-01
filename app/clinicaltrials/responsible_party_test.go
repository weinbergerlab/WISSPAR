package clinicaltrials

import (
	"context"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/stretchr/testify/assert"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"log"
	"os"
	"testing"
	"time"
)

func TestResultsContext_GetResponsibleParties(t *testing.T) {
	os.Remove("test.db") // I delete the file to avoid duplicated records.
	//// SQLite is a file based database.
	//
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("test.db") // Create SQLite file
	assert.Nil(t, err)
	file.Close()
	log.Println("test.db created")

	ctx := context.Background()

	//driver := "postgres"
	//dataSource := "host=localhost port=18034 user=local dbname=local password=password sslmode=disable"

	driver := "sqlite3"
	dataSource := "file:./test.db?cache=shared&mode=memory&_fk=true"

	db, err := models.Open(driver, dataSource)
	assert.Nil(t, err)
	assert.Nil(t, db.Schema.Create(ctx))

	_, err = db.ClinicalTrial.Create().
		SetStudyType("").
		SetStudyName("").
		SetSponsor("").
		SetPrimaryPurpose("").
		SetPhase("").
		SetOfficialTitle("").
		SetMasking("").
		SetInterventionModel("").
		SetAllocation("").
		SetActualStudyStartDate(time.Now()).
		SetActualStudyCompletionDate(time.Now()).
		SetActualPrimaryCompletionDate(time.Now()).
		SetActualEnrollment("").
		SetStudyID("").
		SetResponsibleParty("Test 1").Save(ctx)
	assert.Nil(t, err)
	_, err = db.ClinicalTrial.Create().
		SetStudyType("").
		SetStudyName("").
		SetSponsor("").
		SetPrimaryPurpose("").
		SetPhase("").
		SetOfficialTitle("").
		SetMasking("").
		SetInterventionModel("").
		SetAllocation("").
		SetActualStudyStartDate(time.Now()).
		SetActualStudyCompletionDate(time.Now()).
		SetActualPrimaryCompletionDate(time.Now()).
		SetActualEnrollment("").
		SetStudyID("").
		SetResponsibleParty("Test 2").Save(ctx)
	assert.Nil(t, err)
	_, err = db.ClinicalTrial.Create().
		SetStudyType("").
		SetStudyName("").
		SetSponsor("").
		SetPrimaryPurpose("").
		SetPhase("").
		SetOfficialTitle("").
		SetMasking("").
		SetInterventionModel("").
		SetAllocation("").
		SetActualStudyStartDate(time.Now()).
		SetActualStudyCompletionDate(time.Now()).
		SetActualPrimaryCompletionDate(time.Now()).
		SetActualEnrollment("").
		SetStudyID("").
		SetResponsibleParty("Test 3").Save(ctx)
	assert.Nil(t, err)
	_, err = db.ClinicalTrial.Create().
		SetStudyType("").
		SetStudyName("").
		SetSponsor("").
		SetPrimaryPurpose("").
		SetPhase("").
		SetOfficialTitle("").
		SetMasking("").
		SetInterventionModel("").
		SetAllocation("").
		SetActualStudyStartDate(time.Now()).
		SetActualStudyCompletionDate(time.Now()).
		SetActualPrimaryCompletionDate(time.Now()).
		SetActualEnrollment("").
		SetStudyID("").
		SetResponsibleParty("Test 4").Save(ctx)
	assert.Nil(t, err)
	_, err = db.ClinicalTrial.Create().
		SetStudyType("").
		SetStudyName("").
		SetSponsor("").
		SetPrimaryPurpose("").
		SetPhase("").
		SetOfficialTitle("").
		SetMasking("").
		SetInterventionModel("").
		SetAllocation("").
		SetActualStudyStartDate(time.Now()).
		SetActualStudyCompletionDate(time.Now()).
		SetActualPrimaryCompletionDate(time.Now()).
		SetActualEnrollment("").
		SetStudyID("").
		SetResponsibleParty("Test 3").Save(ctx)
	assert.Nil(t, err)


	rc := NewResultsContext(ctx, db)

	sut, err := rc.GetResponsibleParties()
	assert.Nil(t, err)

	correctTestResult := []string{"Test 1", "Test 2", "Test 3", "Test 4"}

	assert.Equal(t, correctTestResult, sut)
}
