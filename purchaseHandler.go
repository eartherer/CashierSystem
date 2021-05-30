package main

func calculateBanknoteFromChange(change float64) (BankNoteChangeInfo, bool) {
	var quantityCanPick int
	result := BankNoteChangeInfo{Change: change, BanknoteChange: []BankNoteInfo{}}
	remainingChange := change
	bankNoteType := banknoteStorage.getAllBankNoteType()
	for _, typeName := range bankNoteType {
		bankNoteFromStorage, ok := banknoteStorage.BankNoteStorage[typeName]
		if ok {
			quantityCanPick = 0
			value := bankNoteFromStorage.Value
			needBankNoteQuantity := int(remainingChange / value)

			if needBankNoteQuantity > 0 {
				if bankNoteFromStorage.Quantity >= needBankNoteQuantity { // Storage is Enough
					quantityCanPick = needBankNoteQuantity
				} else { // Storage not Enough
					quantityCanPick = bankNoteFromStorage.Quantity
				}

				if quantityCanPick > 0 { //Need to pick some banknote
					//Withdraw from storage
					canWithdraw := banknoteStorage.withdrawBanknoteName(bankNoteFromStorage.Name, quantityCanPick)

					if canWithdraw {
						//Decrease remaining change
						remainingChange = remainingChange - (float64(quantityCanPick) * value)
						//Add banknote to change result
						result.BanknoteChange = append(result.BanknoteChange, BankNoteInfo{
							Name:     bankNoteFromStorage.Name,
							Value:    bankNoteFromStorage.Value,
							Quantity: quantityCanPick,
						})
					}
				}
			}
		}
	}

	if remainingChange > 0 { // Storage cannot make change, not enough banknote
		return result, false
	}
	return result, true
}
