package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    
	units := map[string]int{}
    
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, err := units[unit]
	if !err {
        return false
    } 
    _, err = bill[item]
    if !err {
        bill[item] = units[unit]
        return true
    } else  {
        bill[item] += units[unit]
        return true
    }
    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, err := bill[item]; !err {
        return false
    } else  if _, err = units[unit]; !err{
        return false
    } else if bill[item] - units[unit] < 0 {
        return false
    } else if bill[item] - units[unit] == 0 {
        delete(bill, item)
        return true
    } else {
        bill[item] -= units[unit]
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, err := bill[item]; !err {
        return 0, false
    } else {
        return bill[item], true
    }
}
