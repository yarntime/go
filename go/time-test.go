package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	epoch := now.Unix()

	fmt.Println("Now: ", now)
	fmt.Println("Unix Time: ", epoch)

	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	fmt.Println("Year: ", now.Year())
	fmt.Println("Month: ", now.Month())


	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))

	est, _ := time.LoadLocation("EST")
	july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)

	fmt.Println("July 4, 1776 was a ", july4.Format("Monday"))

	if july4.Before(now) {
		fmt.Println("July 4 is before Now ")
	}

	diff := now.Sub(july4)

	days := int(diff.Hours() / 24)
	fmt.Printf("July 4 was about %d days ago \n", days)

	twodaysDiff := time.Hour * 24 * 2
	twodays := now.Add(twodaysDiff)
	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))

	input_form := "1/2/2006 3:04pm"
	user_str := "4/16/2014 11:38am"
	user_date, err := time.Parse(input_form, user_str)
	if err != nil {
		fmt.Println(">>> Error parsing date string")
	}
	fmt.Println("User Date: ", user_date.Format("Jan 2, 2006 @ 3:04pm"))

}