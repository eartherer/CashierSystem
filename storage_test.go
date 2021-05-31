package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test01_NewStorage(t *testing.T) {

	storageName := "TestStorage"
	storage := NewStorageSystemWithName(storageName)

	assert.Equal(t, storageName, storage.Name)

	assert.Equal(t, 0, len(storage.getAllBankNoteType()))

	assert.Equal(t, 0, len(storage.BankNoteStorage))
}

func Test02_AddNewBankNoteToStorage(t *testing.T) {

	storageName := "TestStorage"
	storage := NewStorageSystemWithName(storageName)

	name := "BankNote1000"
	value := 1000.00
	quantity := 10
	maxQuantity := 10
	storage.AddBankNoteStorage(BankNoteStorage{name, value, quantity, maxQuantity})
	assert.Equal(t, 1, len(storage.getAllBankNoteType()))
	assert.Equal(t, 1, len(storage.BankNoteStorage))
	banknoteInfo := storage.BankNoteStorage[name]
	assert.Equal(t, value, banknoteInfo.Value)
	assert.Equal(t, quantity, banknoteInfo.Quantity)
	assert.Equal(t, maxQuantity, banknoteInfo.MaxQuantity)

	name = "BankNote222"
	value = 222.22
	quantity = 100
	maxQuantity = 100
	storage.AddBankNoteStorage(BankNoteStorage{name, value, quantity, maxQuantity})
	assert.Equal(t, 2, len(storage.getAllBankNoteType()))
	assert.Equal(t, 2, len(storage.BankNoteStorage))
	banknoteInfo = storage.BankNoteStorage[name]
	assert.Equal(t, value, banknoteInfo.Value)
	assert.Equal(t, quantity, banknoteInfo.Quantity)
	assert.Equal(t, maxQuantity, banknoteInfo.MaxQuantity)

	name = "BankNote0.55"
	value = 0.55
	quantity = 105
	maxQuantity = 105
	storage.AddBankNoteStorage(BankNoteStorage{name, value, quantity, maxQuantity})
	assert.Equal(t, 3, len(storage.getAllBankNoteType()))
	assert.Equal(t, 3, len(storage.BankNoteStorage))
	banknoteInfo = storage.BankNoteStorage[name]
	assert.Equal(t, value, banknoteInfo.Value)
	assert.Equal(t, quantity, banknoteInfo.Quantity)
	assert.Equal(t, maxQuantity, banknoteInfo.MaxQuantity)
}

func Test03_RefillBanknoteToStorage(t *testing.T) {

	storageName := "TestStorage"
	storage := NewStorageSystemWithName(storageName)
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote1000", 1000, 5, 10})
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote500", 500, 6, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote100", 100, 5, 15})
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote50", 50, 2, 20})

	name := "BankNote1000"
	refillQuantity := 3
	originalQuantity := -1
	banknote, ok := storage.BankNoteStorage[name]
	if !ok {
		t.Fail()
	}
	originalQuantity = banknote.Quantity
	storage.refillBankNoteStorage(name, refillQuantity)
	banknote, ok = storage.BankNoteStorage[name]
	if !ok {
		t.Fail()
	}
	assert.Equal(t, name, banknote.Name)
	assert.Equal(t, originalQuantity+refillQuantity, banknote.Quantity)

}

func Test04_WithdrawBanknoteFromStorage(t *testing.T) {

	storageName := "TestStorage"
	storage := NewStorageSystemWithName(storageName)
	storage.initBankNoteStorageWithDefault()

	types := storage.getAllBankNoteType()

	for withdrawQuantity, banknoteType := range types {
		banknote, ok := storage.BankNoteStorage[banknoteType]
		if !ok {
			t.Fail()
		}
		originalQuantity := banknote.Quantity
		expectQuantity := originalQuantity - withdrawQuantity

		storage.withdrawBanknoteName(banknoteType, withdrawQuantity)

		banknote, ok = storage.BankNoteStorage[banknoteType]
		if !ok {
			t.Fail()
		}
		assert.Equal(t, banknoteType, banknote.Name)
		assert.Equal(t, expectQuantity, banknote.Quantity)
	}
}
