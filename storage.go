package main

type StorageSystem struct {
	Name            string
	BankNoteStorage map[string]BankNoteStorage
}

type BankNoteStorage struct {
	Name        string
	Value       float64
	Quantity    int
	MaxQuantity int
}

func (s *StorageSystem) getAllBankNoteType() []string {
	return []string{"BankNote1000", "BankNote500", "BankNote100", "BankNote50", "BankNote20", "Coin10", "Coin5", "Coin1", "Coin0.25"}
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
func NewStorageSystemWithName(name string) StorageSystem {
	result := StorageSystem{Name: name, BankNoteStorage: make(map[string]BankNoteStorage)}
	result.BankNoteStorage["BankNote1000"] = BankNoteStorage{"BankNote1000", 1000, 10, 10}
	result.BankNoteStorage["BankNote500"] = BankNoteStorage{"BankNote500", 500, 20, 20}
	result.BankNoteStorage["BankNote100"] = BankNoteStorage{"BankNote100", 100, 15, 15}
	result.BankNoteStorage["BankNote50"] = BankNoteStorage{"BankNote50", 50, 20, 20}
	result.BankNoteStorage["BankNote20"] = BankNoteStorage{"Coin20", 20, 30, 30}
	result.BankNoteStorage["Coin10"] = BankNoteStorage{"Coin10", 10, 20, 20}
	result.BankNoteStorage["Coin5"] = BankNoteStorage{"Coin5", 5, 20, 20}
	result.BankNoteStorage["Coin1"] = BankNoteStorage{"Coin1", 1, 20, 20}
	result.BankNoteStorage["Coin0.25"] = BankNoteStorage{"Coin0.25", 0.25, 50, 50}
	return result
}
