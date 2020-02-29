package storage

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	storage  = "storage"
	filePerm = 0644
)

type Storage interface {
	GetFilePath(name string) string
	Save(name, content string) error
	Get(name string) (string, error)
}

//LocalStorage struct is a struct to keep to dispatch methods that works directly with files
type LocalStorage struct{}

//GetFilePath method: gives the file path with desationation path
func (ls *LocalStorage) GetFilePath(name string) string {
	return storage + "/" + name
}

//Save method: Save files with content string
func (ls *LocalStorage) Save(name, content string) error {
	filePath := ls.GetFilePath(name)
	err := ioutil.WriteFile(filePath, []byte(content), filePerm)
	if err != nil {
		log.Printf("The file %s cannot be saved due: %s", filePath, err.Error())
		return err
	}
	return nil
}

//Get method: return content of the file
func (ls *LocalStorage) Get(name string) (string, error) {
	filePath := ls.GetFilePath(name)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("The file %s cannot be readed due: %s", filePath, err.Error())
		return "", err
	}
	return string(data), nil
}

//Delete method: is to delete a file
func (ls *LocalStorage) Delete(name string) error {
	filePath := ls.GetFilePath(name)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
