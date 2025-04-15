package mylib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StringsFunc() {
	str := "Hello Go"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)
	fmt.Println(str[0])
	num := 18
	str3 := strconv.Itoa(num)
	println(len(str3))

	fruits := "apple, orange, banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(parts)

	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go?"))

	replaced := strings.Replace(str, "Go", "Universe", 2)
	fmt.Println(replaced)

	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))
	fmt.Println(strings.Repeat("foo ", 3))

	fmt.Println(strings.Count("Hello", "l"))
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str5 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")
	builderResult := builder.String()
	fmt.Println(builderResult)

	

}
