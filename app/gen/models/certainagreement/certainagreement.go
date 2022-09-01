// Code generated (@generated) by entc, DO NOT EDIT.

package certainagreement

const (
	// Label holds the string label denoting the certainagreement type in the database.
	Label = "certain_agreement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAgreementPiSponsorEmployee holds the string denoting the agreement_pi_sponsor_employee field in the database.
	FieldAgreementPiSponsorEmployee = "agreement_pi_sponsor_employee"
	// FieldAgreementRestrictionType holds the string denoting the agreement_restriction_type field in the database.
	FieldAgreementRestrictionType = "agreement_restriction_type"
	// FieldAgreementRestrictiveAgreement holds the string denoting the agreement_restrictive_agreement field in the database.
	FieldAgreementRestrictiveAgreement = "agreement_restrictive_agreement"
	// FieldAgreementOtherDetails holds the string denoting the agreement_other_details field in the database.
	FieldAgreementOtherDetails = "agreement_other_details"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the certainagreement in the database.
	Table = "certain_agreement"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "certain_agreement"
	// ParentInverseTable is the table name for the MoreInfoModule entity.
	// It exists in this package in order to avoid circular dependency with the "moreinfomodule" package.
	ParentInverseTable = "more_info_module"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "certain_agreement_parent"
)

// Columns holds all SQL columns for certainagreement fields.
var Columns = []string{
	FieldID,
	FieldAgreementPiSponsorEmployee,
	FieldAgreementRestrictionType,
	FieldAgreementRestrictiveAgreement,
	FieldAgreementOtherDetails,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "certain_agreement"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"certain_agreement_parent",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
