package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	existingValue, exists := units[unit]

	if exists == false {
		return false
	}

	bill[item] += existingValue
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exist := bill[item]

	if exist == false {
		return false
	}

	_, exist = units[unit]

	if exist == false {
		return false
	}

	if bill[item] < units[unit] {
		return false
	}

	bill[item] -= units[unit]

	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, exists := bill[item]

	if exists == false {
		return 0, false
	}

	return bill[item], true
}
