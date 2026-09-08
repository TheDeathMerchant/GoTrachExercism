package reversestring

func Reverse(input string) string {
    inputArr := []rune(input)
    outputArr := make([]rune, len(inputArr))
    for i, v:= range inputArr {
        outputArr[len(inputArr)-1-i] = v
    }
    return string(outputArr)
}
