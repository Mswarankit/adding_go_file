package main

// Convert ages to days

// func convertAges(n int) int {
// 	if n < 0 {
// 		return 0
// 	}
// 	var age int = 365
// 	return age * n

// }

// Finding Nemo
// func findingNemo(str string) string {
// 	strw := strings.Fields(str)
// 	for i, val := range strw {
// 		if val == "Nemo" {
// 			return fmt.Sprintf("I found Nemo at %d", i)
// 		}
// 	}
// 	return "I can't find Nemo :("
// }

// Barbeque Skewers
// func barbequeSkewers(arr []string) []int {

// }

// const (
// 	Keyword      = "Nemo"
// 	NotFoundText = "I can't find Nemo :("
// 	FoundText    = "I found Nemo at %v!\n"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Type the Nemo sentenece:")
// 	input, err := reader.ReadString('\n')
// 	fmt.Println("........Looking for Nemo............")
// 	if err != nil {
// 		fmt.Printf("Something went wrong: %v\n", err)
// 	} else {
// 		nemoPosition := findNemo(input)
// 		if nemoPosition > 0 {
// 			fmt.Printf(FoundText, nemoPosition)
// 		} else {
// 			fmt.Printf(NotFoundText)
// 		}
// 	}
// 	// fmt.Println(convertAges(-1))
// 	// fmt.Println(findingNemo("I am finding NeMo !"))
// }

// func findNemo(input string) int {
// 	for i, word := range strings.Fields(input) {
// 		if word == Keyword {
// 			return i + 1
// 		}
// 	}
// 	return 0
// }

// finding barbeque

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Printf("Enter the bbq configuration: \n")
// 	input, _ := reader.ReadString('\n')
// 	result, err := parseJson(input)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	fmt.Print(result)
// }

// const (
// 	vegan    = "o"
// 	nonVegan = "x"
// )

// func parseJson(input string) (result string, error error) {
// 	veganC, NonVeganC := 0, 0
// 	var SkewArr []string
// 	err := json.Unmarshal([]byte(input), &SkewArr)
// 	if err != nil {
// 		return result, fmt.Errorf("Something went wrong parsing the input")
// 	}
// 	for _, s := range SkewArr {
// 		if strings.Contains(s, vegan) {
// 			veganC++
// 		} else {
// 			if strings.Contains(s, nonVegan) {
// 				NonVeganC++
// 			}
// 		}

// 	}
// 	result = fmt.Sprintf("[%v,%v]\n", veganC, NonVeganC)
// 	return result, nil

// }

// fmt.Println(prevPrime(50))
// socks := SocksPairs("ACACBCA")
// fmt.Println(socks)
// progress := []int{10, 11, 12, 9, 10, 11, 12, 13, 15, 18}
// fmt.Println(progressDays(progress))
// arr1 := []int{2, 4, 0, 0}
// arr2 := []int{1, 3}

// var count = 0

// func progressDays(arr []int) int {
// 	var first = arr[0]
// 	for i := 1; i < len(arr); i++ {
// 		if first < arr[i] {
// 			count++
// 		}
// 		first = arr[i]
// 	}
// 	return count
// }

// func progressDays(arr []int) int {
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i-1] < arr[i] {
// 			count++
// 		}
// 	}
// 	return count
// }

// func SocksPairs(str string) int {
// 	if len(str) < 0 {
// 		return 0
// 	}
// 	var count = 0
// 	socks := make(map[string]int)

// 	for _, r := range str {
// 		socks[string(r)]++
// 		if socks[string(r)]%2 == 0 {
// 			count++
// 		}
// 	}
// 	fmt.Println(socks)
// 	return count
// }

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func prevPrime(n int) int {
// 	prime := isPrime(n)
// 	if prime == true {
// 		return n
// 	}
// 	return prevPrime(n - 1)
// }

// func main() {
// 	array_1 := []int{1, 2, 3, 4, 5, 0, 0, 0}
// 	array_2 := []int{2, 5, 6}

// 	fmt.Println(MergeSortedArray(array_1, array_2))
// }

// func MergeSortedArray(arr1, arr2 []int) []int {
// 	i, j := len(arr1)-len(arr2)-1, len(arr2)-1
// 	k := i + j + 1
// 	for j >= 0 {
// 		if i >= 0 && arr1[i] >= arr2[j] {
// 			arr1[k] = arr1[i]
// 			i--
// 		} else {
// 			arr1[k] = arr2[j]
// 			j--
// 		}
// 		k--
// 	}
// 	return arr1
// }
