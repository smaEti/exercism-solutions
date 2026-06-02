package raindrops
import "strconv"

func Convert(number int) string {
    string := ""

    if number%3 == 0 {
        string = string + "Pling"
    }
    
    if number%5 == 0 {
        string = string + "Plang"
    }
    
    if number%7 == 0 {
        string = string + "Plong"
    }

    if string == "" { 
        return strconv.Itoa(number)
    }
    return string
}
