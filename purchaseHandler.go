package main

func calculateSolutionFromChange(storage *StorageSystem, change float64) (BankNoteChangeInfo, bool) {
	return calculateSubSolutionFromChange(storage, 1000.00+1, change)
}

type ChangeProblem struct {
	Change            float64
	SubProblemOfValue float64
}

func calculateSubSolutionFromChange(storage *StorageSystem, topBanknoteValue float64, change float64) (BankNoteChangeInfo, bool) {

	result := BankNoteChangeInfo{Change: change, BanknoteChange: []BankNoteInfo{}}

	remainingChange := change
	bankNoteType := storage.getAllBankNoteType()
	for _, typeName := range bankNoteType {
		bankNoteFromStorage, ok := storage.BankNoteStorage[typeName]
		if !ok {
			continue
		}

		//Find Banknote which value less than remaining
		if (bankNoteFromStorage.Value <= remainingChange) && (bankNoteFromStorage.Value < topBanknoteValue) {
			maximunNumberToPick := int(remainingChange / bankNoteFromStorage.Value)
			remainingInStorage := bankNoteFromStorage.Quantity
			if remainingInStorage < maximunNumberToPick {
				maximunNumberToPick = remainingInStorage
			}

			if maximunNumberToPick == 0 { // Cannot pick for this banknote type
				continue
			}

			for numberToPick := maximunNumberToPick; numberToPick > 0; numberToPick-- {
				needMoreChange := remainingChange - (float64(numberToPick) * bankNoteFromStorage.Value)

				result := BankNoteChangeInfo{Change: change, BanknoteChange: []BankNoteInfo{}}
				needToAddCurrentPick := true
				if needMoreChange > 0 {

					//This is lowest value banknote so no need to find solution of "needMoreChange"
					if bankNoteFromStorage.Value == 0.25 {
						return result, false
					}

					subSolution, work := calculateSubSolutionFromChange(storage, bankNoteFromStorage.Value, needMoreChange)
					if !work {
						if bankNoteFromStorage.Value != 50 {
							// If current banknote is 100 which can divide by next banknote is 50 so no need to decrease pick amount
							// but 50 cannot divide by 20 so incase change 110 After try to pick 50x2 which cannot solve problem
							// It must try to pick 50x1 + 20x3
							break
						}
						continue
					}
					for _, banknote := range subSolution.BanknoteChange {
						if banknote.Name == typeName {
							banknote.Quantity += numberToPick
							needToAddCurrentPick = false
						}
						result.BanknoteChange = append(result.BanknoteChange, banknote)
					}
				}
				if needToAddCurrentPick {
					result.BanknoteChange = append(result.BanknoteChange, BankNoteInfo{
						Name:     bankNoteFromStorage.Name,
						Value:    bankNoteFromStorage.Value,
						Quantity: numberToPick,
					})
				}
				return result, true
			}
		}
	}

	//Try all banknote type and all number to pick but not found solution
	return result, false
}
func calculateBanknoteFromChange(storage *StorageSystem, change float64) (BankNoteChangeInfo, bool) {

	solution, canChange := calculateSolutionFromChange(storage, change)
	if !canChange {
		return BankNoteChangeInfo{Change: change}, false
	}

	withdrawFromStorage(storage, solution)
	return solution, true
}

func withdrawFromStorage(storage *StorageSystem, solution BankNoteChangeInfo) {
	for _, banknote := range solution.BanknoteChange {
		storage.withdrawBanknoteName(banknote.Name, banknote.Quantity)
	}
}
