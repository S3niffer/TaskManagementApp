package store

import (
	"encoding/json"
	"os"
)

type Store struct {
	User UserStore
}

type DataBase struct {
	Users         []user `json:"users"`
	ModifiedTimes int    `json:"modified_times"`
	FileName      string `json:"-"`
}

func New() (Store, error) {
	db := &DataBase{
		FileName: "Database.txt",
	}

	if err := db.createFile(); err != nil {
		return Store{}, err
	}

	err := db.LoadFromFile()
	if err != nil {
		return Store{}, err
	}

	return Store{
		User: UserStore{db},
	}, nil
}

func (db DataBase) createFile() error {
	// if utilities.IsFileExist(db.FileName) {
	// 	return nil
	// }

	// file, err := os.Create(db.FileName)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()
	// return nil

	file, err := os.OpenFile(db.FileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func (db *DataBase) SaveToFile() error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	if err = os.WriteFile(db.FileName, data, 0644); err != nil {
		return err
	}

	return nil
}

func (db *DataBase) LoadFromFile() error {
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	if err = json.Unmarshal(data, db); err != nil {
		return err
	}

	return nil
}
