package main

func main() {

}

/*
	func Season(month int) string {
		switch month {  // switch on the basis of value of the months(1-12)
			case 12,1,2:  return "Winter"  // Jan, Feb and Dec have winter
			case 3,4,5:	  return "Spring"  // March, Apr and May have spring
			case 6,7,8:   return "Summer"  // June, July and Aug have summer
			case 9,10,11: return "Autumn"  // Sept, Oct and Nov have autumn

			default: return "Season unknown"         //value outside [1,12], then season is unkown
		}
	}
*/
func Season(month int) string {
	var sson string
	switch {
	case month == 1 || month == 2 || month == 12:
		sson = "Winter"
	case month >= 3 && month <= 5:
		sson = "Spring"
	case month >= 6 && month <= 8:
		sson = "Summer"
	case month >= 9 && month <= 11:
		sson = "Autumn"
	default:
		sson = "Season unknown"

	}
	return sson
}
