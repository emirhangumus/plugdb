package dblook

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"emirhangumus.com/plugdb/main/contants"
)

type Database struct {
	Name        string
	Description string
}

type Table struct {
	Name string
}

type DBLook struct {
	Database Database
	Tables   []Table
}

func GetDBLook(db string) (DBLook, error) {
	// available attributes
	dbAvailableAttributes := []string{"database", "tables"}
	// open the file
	file, err := os.Open(contants.DBPath + db + "/entry.plugd")
	if err != nil {
		fmt.Println(err)
	}

	// read the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// read the file
	fileContent := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContent)
	if err != nil {
		fmt.Println(err)
	}

	// generate the database look
	dbLook := DBLook{}
	dbLook.Database.Name = db
	dbLook.Tables = []Table{}

	// parse [...] to get the tables
	heading := string("")
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] == '[' {
			for j := i + 1; j < len(fileContent); j++ {
				if fileContent[j] == ']' {
					heading = string(fileContent[i+1 : j])
					if !slices.Contains(dbAvailableAttributes, heading) {
						return DBLook{}, fmt.Errorf("unknown database file attribute %s", heading)
					}
					i = j + 1
					break
				}
			}
		}
		content := string("")
		if heading == "database" {
			for j := i + 1; j < len(fileContent); j++ {
				if fileContent[j] == '\n' {
					content = string(fileContent[i+1 : j])
					_, err := regexp.Compile("^\t.*")
					if err != nil {
						return DBLook{}, fmt.Errorf("invalid database file format (missing tab character)")
					}
					// remove the tab
					content = content[1:]
					fmt.Println(content)
					// find the first index of "="
					equalIndex := strings.Index(content, "=")
					if equalIndex == -1 {
						return DBLook{}, fmt.Errorf("invalid database file format (missing equal sign)")
					}
					// split by the equal sign
					key := content[:equalIndex]
					value := content[equalIndex+1:]

					if key == "name" {
						dbLook.Database.Name = value
					}

					if key == "description" {
						dbLook.Database.Description = value
					}

					i = j
					if fileContent[j+1] == '[' {
						break
					}
				}
			}
		}
	}

	return dbLook, nil
}

func WriteDBLook(look DBLook, db string) error {
	// convert DBLook to string
	dbEntry := "[database]\n"
	dbEntry += "\tname=" + look.Database.Name + "\n"
	dbEntry += "\tdescription=" + look.Database.Description + "\n"
	dbEntry += "[tables]\n"
	for _, table := range look.Tables {
		dbEntry += "\t" + table.Name + "\n"
	}

	// open the file
	file, err := os.OpenFile(contants.DBPath+db+"/entry.plugd", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// write the file
	_, err = file.WriteAt([]byte(dbEntry), 0)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
