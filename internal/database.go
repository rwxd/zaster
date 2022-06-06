package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Account string
type Category string

type DatabaseData struct {
	Transactions map[string]Transaction `json:"transactions"`
	Accounts     []Account              `json:"accounts"`
	Categories   []Category             `json:"budgets"`
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
	data.Categories = []Category{"Food", "Games", "Investment", "Misc", "Inflow"}
	tempTransactions := []Transaction{
		NewTransaction(25.0, time.Now(), "Marten", "Commerzbank", "Food", "wegen Essen", MoneyInflow),
		NewTransaction(36.99, time.Now(), "Peter", "Commerzbank", "Food", "Essen gegangen", MoneyOutflow),
		NewTransaction(29.72, time.Now(), "Versicherung", "Commerzbank", "Misc", "Rückzahlung", MoneyInflow),
		NewTransaction(129.53, time.Now(), "DB", "Commerzbank", "Misc", "9€ Ticket", MoneyOutflow),
		NewTransaction(12.0, time.Now(), "Mc Donalds", "Commerzbank", "Food", "Essen gehen", MoneyOutflow),
		NewTransaction(2502.0, time.Now(), "Firma", "Deutsche Bank", "Inflow", "Gehalt", MoneyInflow),
	}

	data.Transactions = make(map[string]Transaction)
	for _, t := range tempTransactions {
		data.Transactions[t.Id.String()] = t
	}
	db.Data = &data
}

func (db *JSONDatabase) GetTransactionById(id string) (Transaction, error) {
	transaction, found := db.Data.Transactions[id]
	if !found {
		return Transaction{}, errors.New("Transaction does not exist")
	}

	return transaction, nil
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
