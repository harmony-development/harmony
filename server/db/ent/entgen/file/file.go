// Code generated by entc, DO NOT EDIT.

package file

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldContenttype holds the string denoting the contenttype field in the database.
	FieldContenttype = "contenttype"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// Table holds the table name of the file in the database.
	Table = "files"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldContenttype,
	FieldSize,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
