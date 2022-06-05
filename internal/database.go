package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Account string
type Budget string

type DatabaseData struct {
	Transactions []Transaction `json:"transactions"`
	Accounts     []Account     `json:"accounts"`
	Budgets      []Budget      `json:"budgets"`
}

type JSONDatabase struct {
	filePath string
	Data     *DatabaseData
}

func (db *JSONDatabase) checkPathExists() bool {
	info, err := os.Stat(db.filePath)
	if os.IsNotExist(err) {
		log.Println("No database file found on path: ", db.filePath)
		return false
	} else if info.IsDir() {
		return false
	}

	return true
}

func (db *JSONDatabase) createNewDatabaseFile() {
	file, err := os.Create(db.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("{}"))
	if err != nil {
		log.Fatal(err)
	}
}

func (db *JSONDatabase) loadDatabaseData() {
	log.Println("Load json database data")
	jsonFile, err := os.Open(db.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &db.Data)
}

func (db *JSONDatabase) commitDatabase() {
	log.Println("Commiting data to database")
	jsonFile, err := os.Open(db.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := json.MarshalIndent(db.Data, "", "")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(db.filePath, byteValue, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *JSONDatabase) writeInitialData() {
	data := DatabaseData{}
	data.Accounts = []Account{"Commerzbank", "Deutsche Bank"}
	data.Budgets = []Budget{"Food", "Games", "Investment"}

	transaction1, _ := NewTransaction(25.0, time.Now(), "Marten", "Commerzbank", "", "Essen", MoneyInflow)
	transaction2, _ := NewTransaction(36.99, time.Now(), "Peter", "Commerzbank", "", "", MoneyOutflow)
	transaction3, _ := NewTransaction(29.72, time.Now(), "Versicherung", "Commerzbank", "", "Rückzahlung", MoneyInflow)
	transaction4, _ := NewTransaction(129.53, time.Now(), "DB", "Commerzbank", "", "9€ Ticket", MoneyOutflow)
	transaction5, _ := NewTransaction(12.0, time.Now(), "Mc Donalds", "Commerzbank", "", "Essen gehen", MoneyOutflow)
	transaction6, _ := NewTransaction(2502.0, time.Now(), "Firma", "", "", "Gehalt", MoneyInflow)
	data.Transactions = []Transaction{transaction1, transaction2, transaction3, transaction4, transaction5, transaction6}
	db.Data = &data
}

func NewJSONDatabase(filePath string) JSONDatabase {
	db := JSONDatabase{filePath: filePath}
	if !db.checkPathExists() {
		db.createNewDatabaseFile()
	}
	db.loadDatabaseData()
	db.writeInitialData()
	db.commitDatabase()

	return db
}
