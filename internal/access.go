package internal

import (
	"database/sql"
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	_ "github.com/mattn/go-adodb"
)

// Access access interface struct
type Access struct {
	db *sql.DB
}

func (a *Access) createMdb(dbFilePath string) error {
	unk, err := oleutil.CreateObject("ADOX.Catalog")
	if err != nil {
		return err
	}
	defer unk.Release()

	cat, err := unk.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer cat.Release()

	provider := "Microsoft.ACE.OLEDB.12.0"
	r, err := oleutil.CallMethod(cat, "Create", "Provider="+provider+";Data Source="+dbFilePath+";")
	if err != nil {
		return err
	}

	r.Clear()
	return nil
}

// Connect connect to access file via file path
func (a *Access) Connect(dbFilePath string) error {
	if err := a.createMdb(dbFilePath); err != nil {
		return err
	}

	connectionString := fmt.Sprintf("Provider=Microsoft.ACE.OLEDB.12.0;Data Source=%v;Persist Security Info=False;", dbFilePath)

	var err error
	a.db, err = sql.Open("adodb", connectionString)

	return err
}

// Close close the connection
func (a *Access) Close() {
	a.db.Close()
}

// CreateTable execute create table query in database
func (a *Access) CreateTable(sqlStr string) error {
	_, err := a.db.Exec(sqlStr)
	return err
}

// InsertRows execute insert row query in database
func (a *Access) InsertRow(sqlStr string, values []uint) error {
	valueInterfaces := make([]interface{}, len(values))
	for i, v := range values {
		valueInterfaces[i] = v
	}

	_, err := a.db.Exec(sqlStr, valueInterfaces...)
	return err
}
