package main

import "errors"

type StorageSystem struct {
	Name            string
	BankNoteStorage map[string]BankNoteStorage
	BankNoteTypes   []string
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
	result := StorageSystem{Name: name, BankNoteStorage: make(map[string]BankNoteStorage)}
	return result
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
