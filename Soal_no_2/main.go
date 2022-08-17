package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int // pemain
	var m int // dadu
	var nama string
	var next int
	ntemp := []string{}
	totalScore := map[string]int{}
	tempScore := []int{}

	fmt.Print("Masukan Jumlah Pemain : ")
	fmt.Scanln(&n)
	fmt.Print("Masukan Jumlah Dadu : ")
	fmt.Scanln(&m)

	fmt.Println("===========================================")

	fmt.Printf("Jumlah Pemain : %d\n", n) //3 pemain
	fmt.Printf("Jumlah Dadu : %d\n", m)   //4 dadu

	fmt.Println("===========================================")

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukan Nama Pemain ke- %d: ", i)
		fmt.Scanln(&nama)
		ntemp = append(ntemp, nama)
	}

	fmt.Println("===========================================")

	for _, v := range ntemp {
		mtemp := []int{}
		totalScore[v] = 0

		fmt.Printf("Giliran %s melampar dadu", v)
		fmt.Scanln(&next)
		fmt.Printf("Hasil %s melempar dadu: \n", v)
		rand.Seed(time.Now().UTC().UnixNano())

		for i := 1; i <= m; i++ {
			result := rand.Intn(7-1) + 1
			mtemp = append(mtemp, result)
			// if result ==1 {

			// }
			if result == 6 {
				value, _ := totalScore[v]
				totalScore[v] = value + 1
			}

		}
		fmt.Println(mtemp)
		fmt.Println("")

	}

	for _, v := range totalScore {
		tempScore = append(tempScore, v)
	}

	win := tempScore[0]
	for k, v := range totalScore {
		for i := 1; i < len(tempScore); i++ {
			if tempScore[i] > win {
				win = tempScore[i]
			}
		}
		if v == win {
			fmt.Printf("Pemenangnya adalah %s dengan total score yang diperoleh: %d \n", k, v)
		}
	}
	fmt.Println("===========================================")
}
