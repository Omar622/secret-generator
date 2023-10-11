package app

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/Omar622/secret-generator/internal"
)

//go:embed sql/create_report_table.sql
var createReportTableSQL string

//go:embed sql/insert_into_report_table.sql
var insertIntoReportSQL string

//go:embed sql/create_numbers_table.sql
var createNumbersTableSQL string

//go:embed sql/insert_into_numbers_table.sql
var insertIntoNumbersSQL string

// NewWriter is a factory function for writer
func NewWriter(matchingRanges []internal.MatchingRange) Writer {
	return Writer{accessInterface: internal.Access{}, matchingRanges: matchingRanges}
}

// Writer is a writer struct
type Writer struct {
	accessInterface internal.Access
	matchingRanges  []internal.MatchingRange
}

// WriteReport write report in microsoft access file
func (w *Writer) WriteReport() error {
	filePath := fmt.Sprintf("./report-%v.mdb", time.Now().UnixNano())

	if err := w.accessInterface.Connect(filePath); err != nil {
		return fmt.Errorf("could not connect to Microsoft.ACE.OLEDB.12.0: %v", err)
	}
	defer w.accessInterface.Close()

	if err := w.accessInterface.CreateTable(createReportTableSQL); err != nil {
		return fmt.Errorf("could not create table in microsoft access file: %v", err)
	}

	for _, g := range w.matchingRanges {
		err := w.accessInterface.InsertRow(insertIntoReportSQL, []uint{
			g.SettingNumberRange.Start, g.SettingNumberRange.End, g.SecretNumberRange.Start, g.SecretNumberRange.End,
		})

		if err != nil {
			return fmt.Errorf("could not write data in table: %v", err)
		}
	}

	return nil
}

// WriteSecretNumbers write secret numbers in microsoft access file
func (w *Writer) WriteSecretNumbers() error {
	filePath := fmt.Sprintf("./secret-numbers-%v.mdb", time.Now().UnixNano())

	if err := w.accessInterface.Connect(filePath); err != nil {
		return fmt.Errorf("could not connect to Microsoft.ACE.OLEDB.12.0: %v", err)
	}
	defer w.accessInterface.Close()

	if err := w.accessInterface.CreateTable(createNumbersTableSQL); err != nil {
		return fmt.Errorf("could not create table in microsoft access file: %v", err)
	}

	for _, g := range w.matchingRanges {
		for i := uint(0); i < 50; i++ {
			err := w.accessInterface.InsertRow(insertIntoNumbersSQL, []uint{
				g.SettingNumberRange.Start + i, g.SecretNumberRange.Start + i,
			})

			if err != nil {
				return fmt.Errorf("could not write data in table: %v", err)
			}
		}
	}

	return nil
}
