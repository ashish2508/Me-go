//error

/*
type error interfrace {
Error() string 
}
*/

package main

import ("main")

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error)  {
	cost, err := sendSMS(msgToCustomer); 
	if err != nil {
		return 0.0, err
}
	cost2, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = 0.0002
	if len(message) > maxTextLen {
		return 0.0,  fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}


