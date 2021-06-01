package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test01_Purchase_Success(t *testing.T) {
	storageName := "TestStorage"
	storage := NewStorageSystemWithName(storageName)
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote1000", 1000, 10, 10})
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote500", 500, 20, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote100", 100, 15, 15})
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote50", 50, 20, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"BankNote20", 20, 30, 30})
	storage.AddBankNoteStorage(BankNoteStorage{"Coin10", 10, 20, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"Coin5", 5, 20, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"Coin1", 1, 20, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"Coin0.25", 0.25, 50, 50})

	change := 78.50
	changeInfo, ok := calculateBanknoteFromChange(&storage, change)
	if !ok {
		t.Fail()
	}

	assert.Equal(t, change, changeInfo.Change)

	sumChangeValue := 0.00
	banknoteCount := make(map[string]int)
	for _, banknote := range changeInfo.BanknoteChange {
		sumChangeValue += (banknote.Value * float64(banknote.Quantity))
		banknoteCount[banknote.Name] = banknote.Quantity
	}

	assert.Equal(t, change, sumChangeValue)

	assert.Equal(t, banknoteCount["BankNote50"], 1)
	assert.Equal(t, banknoteCount["BankNote20"], 1)
	assert.Equal(t, banknoteCount["Coin5"], 1)
	assert.Equal(t, banknoteCount["Coin1"], 3)
	assert.Equal(t, banknoteCount["Coin0.25"], 2)
}

func Test02_PurchaseWithNotEnoughStorage(t *testing.T) {
	storageName := "TestStorage"
	storage := NewStorageSystemWithName(storageName)
	storage.AddBankNoteStorage(BankNoteStorage{"Coin5", 5, 20, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"Coin1", 1, 20, 20})
	storage.AddBankNoteStorage(BankNoteStorage{"Coin0.25", 0.25, 50, 50})

	change := 9999.99

	_, ok := calculateBanknoteFromChange(&storage, change)
	if ok {
		t.Fail()
	}
}
