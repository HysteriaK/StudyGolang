package main

import (
	"fmt"
)

var(
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)


func dispatchCoin()int{
	getCoin := 0
	for _,names := range users{
		eachCoin := 0
		for _,v := range names{
			switch v{
			case 'e','E' : getCoin += 1; eachCoin += 1 
			case 'i','I' : getCoin += 2; eachCoin += 2
			case 'o','O' : getCoin += 3; eachCoin += 3
			case 'u','U' : getCoin += 4; eachCoin += 4
			}
		}
		distribution[names] = eachCoin
	}
	coins = coins - getCoin
	return coins
}

func main(){
	left := dispatchCoin()
	fmt.Println("剩下：",left)
	for index,value := range distribution{
		fmt.Println(index)
		fmt.Println(value)
	}
}