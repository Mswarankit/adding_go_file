package main

// s = "anagram"
// t = "nagaram"

// func validAnagram(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	s_map := make(map[rune]int)
// 	for _, v := range s {
// 		s_map[v]++
// 	}
// 	for _, v := range t {
// 		if s_map[v] == 0 {
// 			return false
// 		}
// 		s_map[v]--
// 	}
// 	return true
// }

// func NimGame(n int) bool {
//     if n < 3 {
//         return true
//     }
//     return n%4 != 0
// }
// func peakElement(nums []int) int {
// 	if len(nums) == 0 {
// 		return -1
// 	}
// 	if len(nums) == 1 {
// 		return 0
// 	}

// 	if nums[0] > nums[1] {
// 		return 0
// 	}

// 	if nums[len(nums)-1] > nums[len(nums)-2] {
// 		return len(nums) - 1
// 	}
// 	for i := 1; i < len(nums)-1; i++ {
// 		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
// 			return i
// 		}
// 	}
// 	return 0
// }

// using catalan number we can work around with binary search tree

// func uniqueBinarySearchTree(num int) int {
// 	f := make([]int, num+2)
// 	f[0], f[1] = 1, 1
// 	for i := 2; i <= num; i++ {
// 		for j := 0; j < i; j++ {
// 			f[i] += f[i-j-1] * f[j]
// 		}
// 	}
// 	return f[num]

// }

// func main() {
// 	num := []int{1, 2, 1, 3, 5, 6, 4}
// 	fmt.Println(peakElement(num))

// }

// func restoreIpAddresses(s string) (ans []string) {
// 	n := len(s)
// 	t := []string{}
// 	var dfs func(int)
// 	dfs = func(i int) {
// 		if i >= n && len(t) == 4 {
// 			ans = append(ans, strings.Join(t, "."))
// 			return
// 		}
// 		if i >= n || len(t) == 4 {
// 			return
// 		}
// 		x := 0
// 		for j := i; j < i+3 && j < n; j++ {
// 			x = x*10 + int(s[j]-'0')
// 			if x > 255 || (j > i && s[i] == '0') {
// 				break
// 			}
// 			t = append(t, s[i:j+1])
// 			dfs(j + 1)
// 			t = t[:len(t)-1]
// 		}
// 	}
// 	dfs(0)
// 	return
// }

// func main() {
// 	s := "25525511135"
// 	fmt.Println(restoreIpAddresses(s))
// }

// func restoreIpAddresses(s string) []string {
// 	buf := make([]string, 0)
// 	ans := make([]string, 0)
// 	search(s, 3, buf, &ans)
// 	return ans
// }

// const IPV4MaxLen = 3

// func search(s string, dotLeft int, buf []string, ans *[]string) {
// 	if dotLeft == 0 {
// 		*ans = append(*ans, strings.Join(append(buf, s), "."))
// 		return
// 	}

// 	for i := 1; i < len(s); i++ {
// 		leftStr, rightStr := s[:i], s[i:]
// 		if len(rightStr) < dotLeft {
// 			break
// 		}

// 		if !isValid(leftStr) {
// 			break
// 		}
// 		if len(rightStr) > dotLeft*IPV4MaxLen {
// 			continue
// 		}

// 		if dotLeft == 1 && !isValid(rightStr) {
// 			continue
// 		}
// 		bufNew := append(buf, leftStr)
// 		search(rightStr, dotLeft-1, bufNew, ans)
// 	}
// 	return
// }

// func isValid(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}

// 	if len(s) > 1 && s[0] == '0' {
// 		return false
// 	}
// 	v, _ := strconv.Atoi(s)
// 	return v < 256
// }

// func CanFit(new []int, p int) bool {

// 	fitsum := func(new []int) int {
// 		sum := 0
// 		for _, v := range new {
// 			sum += v
// 		}
// 		return sum
// 	}
// 	total := p * 10
// 	fits := fitsum(new)
// 	if fits <= total {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func Encrypt(input string) (result string) {
// 	replaceMap := map[string]string{
// 		"a": "0",
// 		"e": "1",
// 		"i": "2",
// 		"o": "3",
// 		"u": "4",
// 	}

// 	for i := len(input) - 1; i >= 0; i-- {
// 		char := string(input[i])
// 		encryptedChar := replaceMap[string(input[i])]

// 		if encryptedChar != "" {
// 			result += encryptedChar
// 		} else {
// 			result += char
// 		}
// 	}
// 	return result + "aca"
// }

// func Encrypt(str string) string {
// 	var result strings.Builder
// 	result.Grow(len(str) + 3)

// 	for i := len(str) - 1; i >= 0; i-- {
// 		switch str[i] {
// 		case 'a':
// 			result.WriteString("0")
// 		case 'e':
// 			result.WriteString("1")
// 		case 'i':
// 			result.WriteString("2")
// 		case 'o':
// 			result.WriteString("3")
// 		case 'u':
// 			result.WriteString("4")
// 		default:
// 			result.WriteByte(str[i])
// 		}
// 	}
// 	result.WriteString("aca")
// 	return result.String()
// }

//	vowels := func() map[string]int {
//		a_list := []string{"a", "e", "i", "o", "u"}
//		a_map := make(map[string]int)
//		for i, v := range a_list {
//			a_map[v] = i
//		}
//		return a_map
//	}
// func Encrypt(str string) string {
// 	replaceMap := map[rune]string{
// 		'a': "0",
// 		'e': "1",
// 		'i': "2",
// 		'o': "3",
// 		'u': "4",
// 	}
// 	reverse := func(str string) (result string) {
// 		for _, v := range str {
// 			result = string(v) + result
// 		}
// 		return
// 	}
// 	r := reverse(str)

// 	var result strings.Builder

// 	for _, char := range r {
// 		if replacement, ok := replaceMap[char]; ok {
// 			result.WriteString(replacement)
// 		} else {
// 			result.WriteRune(char)
// 		}
// 	}
// 	output := result.String() + "aca"
// 	return output
// }

// Mountains or valleys

// All numbers to the left of the peak are increasing
// All numbers to the right of the  peak are decreasing
// The peak cannot be on the boundary

// mid := func(arr []int32) []int32 {
// 	n := len(arr)
// 	if n == 0 {
// 		return []int32{}
// 	}
// 	if n%2 == 0 {
// 		return []int32{arr[n/2-1], arr[n/2]}
// 	}
// 	return []int32{arr[n/2]}
// }
// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// func pingGoogle(c chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res, _ := http.Get("https://google.com")
// 	c <- res.Status
// }
// func pingDuckDuckGo(c chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res, _ := http.Get("https://duckduckgo.com")
// 	c <- res.Status
// }
// func pingBraveSearch(c chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	res, _ := http.Get("https://search.brave.com")
// 	c <- res.Status
// }

// func main() {
// 	gogChan := make(chan string)
// 	ddgChan := make(chan string)
// 	braveChan := make(chan string)

// 	var wg sync.WaitGroup
// 	wg.Add(3)

// 	go pingGoogle(gogChan, &wg)
// 	go pingDuckDuckGo(ddgChan, &wg)
// 	go pingBraveSearch(braveChan, &wg)

// 	openChannels := 3

// 	go func() {
// 		wg.Wait()
// 		close(gogChan)
// 		close(ddgChan)
// 		close(braveChan)
// 	}()

// 	for openChannels > 0 {
// 		select {
// 		case msg1, ok := <-gogChan:
// 			if !ok {
// 				openChannels--
// 			} else {
// 				fmt.Println("Google responsed:", msg1)
// 			}
// 		case msg2, ok := <-ddgChan:
// 			if !ok {
// 				openChannels--
// 			} else {
// 				fmt.Println("duckduckgo responsed:", msg2)
// 			}
// 		case msg3, ok := <-braveChan:
// 			if !ok {
// 				openChannels--
// 			} else {
// 				fmt.Println("brave responsed:", msg3)
// 			}

// 		}
// 	}

// }

// func receiver(ch <-chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		fmt.Println("Received:", i)
// 	}
// 	wg.Done()
// }

// func sender(ch chan<- int, wg *sync.WaitGroup) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("Sent:", i)
// 		ch <- i
// 	}
// 	close(ch)
// 	wg.Done()
// }

// func main() {
// 	ch := make(chan int)
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go receiver(ch, &wg)
// 	go sender(ch, &wg)
// 	wg.Wait()
// }

// func main() {
// 	ch := make(chan string)
// 	defer close(ch)

// 	go func() {
// 		ch <- "Hello, Gophers!"
// 	}()
// 	fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// }

// Unique Binary Search Trees

// func main() {
// 	fmt.Println(GenerateUUID())
// }

// func GenerateUUID() (string, error) {
// 	uuid := make([]byte, 16)
// 	_, err := rand.Read(uuid)
// 	if err != nil {
// 		return "", err
// 	}

// 	uuid[6] = (uuid[6] & 0x0F) | 0x40
// 	uuid[8] = (uuid[8] & 0x3F) | 0x80

// 	uuidStr := make([]byte, 36)
// 	hex.Encode(uuidStr[0:8], uuid[0:4])
// 	uuidStr[8] = '-'
// 	hex.Encode(uuidStr[9:13], uuid[4:6])
// 	uuidStr[13] = '-'
// 	hex.Encode(uuidStr[14:18], uuid[6:8])
// 	uuidStr[18] = '-'
// 	hex.Encode(uuidStr[19:23], uuid[8:10])
// 	uuidStr[23] = '-'
// 	hex.Encode(uuidStr[24:], uuid[10:])

// 	return string(uuidStr), nil
// }

// func targetSum(arr []int, target *int) []int {
// 	targetMap := make(map[int]int)
// 	for i, r := range arr {
// 		*target = *target - r // Modify the value of target via the pointer
// 		// comp := *target
// 		if j, ok := targetMap[r]; ok {
// 			*target = 1
// 			return []int{j, i}
// 		}
// 		targetMap[*target] = i // Store the current index against the modified target value
// 	}
// 	return nil
// }

//func main() {
// fmt.Println(letterCombination("23"))
// fmt.Println("Happy Coding @Gophers!")
// n := 9
// var target *int = &n
//	var arr []int = []int{10, 9, 2, 5, 3, 7, 101, 18}
// fmt.Println(lIs(arr))

// fmt.Println("Before:", *target)
// fmt.Println(targetSum(arr, target)) // Pass target as a pointer
// fmt.Println("After:", *target)      // Target value should reflect the changes
//}

// func lIs(arr []int) int {
// 	space := []int{}
// 	for _, num := range arr {
// 		bs := sort.SearchInts(space, num)
// 		if bs == len(space) {
// 			space = append(space, num)
// 		} else {
// 			space[bs] = num
// 		}

// 	}
// 	return len(space)
// }

// func binarySearch(arr []int, target int) int {
// 	left, right := 0, len(arr)-1
// 	for left < right {
// 		mid := left + (right-left)/2
// 		if arr[mid] >= target {
// 			right = mid
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return left
// }

// var phoneMap = map[string]string{
// 	"2": "abc",
// 	"3": "def",
// 	"4": "ghi",
// 	"5": "jkl",
// 	"6": "mno",
// 	"7": "pqrs",
// 	"8": "tuv",
// 	"9": "wxyz",
// }

// func letterCombination(digits string) []string {
// 	result := make([]string, 0)
// 	if len(digits) == 0 {
// 		return result
// 	}
// 	generateCombination(0, "", digits, &result)
// 	return result
// }

// func generateCombination(index int, curr string, digits string, result *[]string) {
// 	if index == len(digits) {
// 		*result = append(*result, curr)
// 		return
// 	}
// 	alphabet, _ := phoneMap[string(digits[index])]
// 	fmt.Println(alphabet)
// 	for _, v := range alphabet {
// 		generateCombination(index+1, curr+string(v), digits, result)
// 	}
// }

//	func letterCombination(digits string) []string {
//		ans := []string{}
//		if len(digits) == 0 {
//			return ans
//		}
//		ans = append(ans, "")
//		d := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
//		for _, i := range digits {
//			s := d[i-'2']
//			t := []string{}
//			for _, a := range ans {
//				for _, b := range s {
//					t = append(t, a+string(b))
//				}
//			}
//			ans = t
//		}
//		return ans
//	}
//
// [4, 2, 0, 3, 2, 5]
// func trappedWater(arr []int) int {
// 	if len(arr) <= 2 {
// 		return 0
// 	}
// 	total := 0
// 	for i := 1; i < len(arr)-1; i++ {
// 		leftMax := arr[0]
// 		for j := 0; j < i; j++ {
// 			leftMax = max(leftMax, arr[j])
// 		}
// 		rightMax := arr[len(arr)-1]
// 		for j := i + 1; j < len(arr); j++ {
// 			rightMax = max(rightMax, arr[j])
// 		}
// 		water := min(leftMax, rightMax) - arr[i]
// 		if water > 0 {
// 			total += water
// 		}
// 	}
// 	return total
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
