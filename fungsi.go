package main

import (
	"fmt"
)

// Complete the sockMerchant function below.
func sockMerchant(n int, ar [100]int) int {
	var input, idx, count int

	fmt.Scan(&input)
	i := 0
	count = 0
	for i < input {
		fmt.Scan(&n)
		ar[i] = n
		i++
	}
	j := 1
	for j < input {
		nilai := ar[j]
		idx = j
		for idx > 0 && ar[idx-1] > nilai {
			ar[idx] = ar[idx-1]
			idx = idx - 1
		}
		ar[idx] = nilai
		j++
	}
	j = 0
	if j < input && ar[j] == ar[j+1] {
		count = count + 1
		j = +2
	}

	return count
}

func main() {

	var N, hasil int
	var array [100]int

	hasil = sockMerchant(N, array)
	fmt.Println(hasil)
	/*
	   func readLine(reader *bufio.Reader) string {
	       str, _, err := reader.ReadLine()
	       if err == io.EOF {
	           return ""
	       }

	       return strings.TrimRight(string(str), "\r\n")
	   }

	   func checkError(err error) {
	       if err != nil {
	           panic(err)
	       }
	   }
	*/
}
