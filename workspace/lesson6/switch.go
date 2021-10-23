package main

func main() {
	var grade int = 90
	var level string
	switch grade{
	case 90:
		level = "A"
	case 80:
		level = "B"
	case 70, 60: // case
		level = "C"
	case 50:
		level = "D"
	default:
		level = "D"
	}
	println("grade:", grade, "level:", level)


	var gender string = "male"
	var result string
	switch gender {
	case "male":
		result = "good"
		fallthrough // 配了这个case分支后，因为有了fallthrough，还会执行紧接着的下一个case
	case "female":
		result = "better"
	case "inter":
		result = "nice"
	default:
		result = "exception"
	}
	println("gender:", gender, "result:", result) // the value of result is "better"
}