package main

import "fmt"

func main() {
	var phoneNumber, phoneNumber1, phoneNumber2, phoneNumber3 string
	fmt.Scanln(&phoneNumber)
	fmt.Scanln(&phoneNumber1)
	fmt.Scanln(&phoneNumber2)
	fmt.Scanln(&phoneNumber3)

	res := checkPhoneNumber(phoneNumber, phoneNumber1, phoneNumber2, phoneNumber3)

	for _, r := range res {
		fmt.Println(r)
	}
}

func clearPhoneOfSymbols(phone string) string {
	for i := 0; i < len(phone); i++ {
		if phone[i] == '(' || phone[i] == ')' || phone[i] == '-' || phone[i] == '+' {
			phone = phone[:i] + phone[i+1:]
			i--
		}
	}

	return phone
}

func transformPhoneNumber(phone string) string {
	if len(phone) == 7 {
		return "495" + phone
	}

	if len(phone) == 11 {
		runes := []rune(phone)
		return string(runes[1:])
	}

	return phone
}

func checkPhoneNumber(phone, phone1, phone2, phone3 string) []string {
	phone = clearPhoneOfSymbols(phone)
	phone = transformPhoneNumber(phone)

	phone1 = clearPhoneOfSymbols(phone1)
	phone1 = transformPhoneNumber(phone1)

	phone2 = clearPhoneOfSymbols(phone2)
	phone2 = transformPhoneNumber(phone2)

	phone3 = clearPhoneOfSymbols(phone3)
	phone3 = transformPhoneNumber(phone3)

	var result []string

	if phone == phone1 {
		result = append(result, "YES")
	} else {
		result = append(result, "NO")
	}

	if phone == phone2 {
		result = append(result, "YES")
	} else {
		result = append(result, "NO")
	}

	if phone == phone3 {
		result = append(result, "YES")
	} else {
		result = append(result, "NO")
	}

	return result
}
