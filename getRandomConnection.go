package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var some_nums []int
	for i := 1; i < 100; i++ {
		some_nums = append(some_nums, i)

	}
	r := getRandNum(some_nums)

	fmt.Printf("%v", r)

}

func getRandNum(nums []int) map[int][]int {

	// use half len(nums) to build connection
	middle_num := len(nums) / 2

	// get new slice with all destation to be selected
	var bSideNums []int
	for i := 0; i < middle_num; i++ {
		bSideNums = append(bSideNums, nums...)
	}

	// return a map who's key = nums.each and value = half nums slice
	// which not include key itself
	var connectMap = make(map[int][]int)

	for _, v := range nums {
		// every round, initial randNums
		var randNums []int
		for i := 0; i < middle_num; i++ {
			// for the sake of dead loop, set most try times
			// when try times end to middle_nums * 2 and not get the num
			// then add a nums to bSideNums to complete it
			for j := 0; j <= middle_num*2; j++ {
				if j == middle_num*2 {
					bSideNums = append(bSideNums, nums...)
				}

				//generate rand with more randomly
				t, err := rand.Int(rand.Reader, big.NewInt(int64(len(bSideNums))))
				if err != nil {
					panic(err)
				}

				// check random choose whether exist in randNums, if yes, continue
				if !checkSliceInclude(randNums, bSideNums[t.Uint64()]) {
					randNums = append(randNums, bSideNums[t.Uint64()])
					bSideNums = append(bSideNums[:t.Uint64()], bSideNums[t.Uint64()+1:]...)
					break
				}
			}
		}

		// key = v and value = randNums
		connectMap[v] = randNums

	}

	return connectMap

}

func checkSliceInclude(x []int, y int) bool {
	if len(x) == 0 {
		return false
	}
	for _, v := range x {
		if y == v {
			return true
		}
	}
	return false
}
