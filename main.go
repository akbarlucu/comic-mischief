package main

import "fmt"

func main() {
	var (
		writer    = "Tracey Hatchet"
		artist    = "Jewel Tampson"
		title     = "Mr. GoToSleep"
		publisher = "DizzyBooks Publishing Inc."
	)

	var (
		year       = 1997
		pageNumber = 14
		grade      = 6.5
	)

	fmt.Println(title, "written by", writer, "drawn by", artist)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0

	// Assigns
	fmt.Println()
	fmt.Println(title, writer, artist, publisher, year, pageNumber, grade)
}
