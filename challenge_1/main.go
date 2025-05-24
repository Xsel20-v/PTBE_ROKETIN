package main

import "fmt"

func main() {
	var hour, min, sec int
	fmt.Print("Insert the time on earth (hh:mm:ss) : ")
	fmt.Scanf("%d:%d:%d", &hour, &min, &sec)
	roketinTime := convertTime(hour, min, sec)
	fmt.Printf("on earth %d:%d:%d, on planet Roketin Planet %s\n", hour, min, sec, roketinTime)
}

func convertTime(hour, min, sec int) string {

	hour = hour*60*60 // change hour to second
	min = min*60 // change min to second

	totalSec := hour+min+sec

	//convert earth's time to roketin's time
	//1 hour = 100 min, 1 min = 100 sec
	hour = totalSec/10000 
	min = (totalSec%10000)/100
	sec = ((totalSec%10000)%100)
	
	//make the integer in format of 2decimals if below 10
	rocketHour := fmt.Sprintf("%02d", hour) 
	rocketMin := fmt.Sprintf("%02d", min)
	rocketSec := fmt.Sprintf("%02d", sec)

	return fmt.Sprintf("%s:%s:%s", rocketHour, rocketMin, rocketSec)
}