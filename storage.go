package main

import "errors"

type StorageSystem struct {
	Name            string                     `json:"StorageName"`
	BankNoteStorage map[string]BankNoteStorage `json:"BankNoteInfo"`
	BankNoteTypes   []string                   `json:"AllBankNote"`
}

type BankNoteStorage struct {
	Name        string
	Value       float64
	Quantity    int
	MaxQuantity int
}

func (s *StorageSystem) getAllBankNoteType() []string {
	return s.BankNoteTypes
}
func (s *StorageSystem) withdrawBanknoteName(name string, quantity int) bool {
	banknote, ok := s.BankNoteStorage[name]
	if ok {
		if banknote.Quantity >= quantity {
			banknote.Quantity -= quantity
			s.BankNoteStorage[name] = banknote
			return true
		}
	}
	return false
}
func (s *StorageSystem) AddBankNoteStorage(newBank BankNoteStorage) {
	s.BankNoteStorage[newBank.Name] = newBank
	s.BankNoteTypes = append(s.BankNoteTypes, newBank.Name)
}
func NewStorageSystemWithName(name string) StorageSystem {
	banknoteStorage := StorageSystem{Name: name, BankNoteStorage: make(map[string]BankNoteStorage)}
	return banknoteStorage
}

func (s *StorageSystem) refillBankNoteStorage(name string, quantity int) error {
	if storage, ok := s.BankNoteStorage[name]; ok {
		if (storage.Quantity + quantity) > storage.MaxQuantity {
			return errors.New("exceed MaxQuantity")
		} else {
			storage.Quantity += quantity
			s.BankNoteStorage[name] = storage
			return nil
		}
	}
	return errors.New(name + " not found")
}
func (banknoteStorage *StorageSystem) initBankNoteStorageWithDefault() error {
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"BankNote1000", 1000, 10, 10})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"BankNote500", 500, 20, 20})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"BankNote100", 100, 15, 15})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"BankNote50", 50, 20, 20})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"BankNote20", 20, 30, 30})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"Coin10", 10, 20, 20})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"Coin5", 5, 20, 20})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"Coin1", 1, 20, 20})
	banknoteStorage.AddBankNoteStorage(BankNoteStorage{"Coin0.25", 0.25, 50, 50})
	return nil
}
