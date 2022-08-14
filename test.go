package main

import "fmt"

type ATM struct {
	curMoney []int
	price    []int
}

func Constructor() ATM {
	return ATM{
		curMoney: []int{0, 0, 0, 0, 0},
		price:    []int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for k, v := range banknotesCount {
		this.curMoney[k] += v
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, len(this.curMoney))
	for i := len(this.curMoney) - 1; i > 0; {
		if this.curMoney[i] > 0 {
			if amount == 0 {
				return res
			} else if amount > 0 && amount > this.price[i] {
				amount = amount - this.price[i]
				res[i]++
				this.curMoney[i]--
			} else {
				i--
			}
		} else {
			i--
		}
	}
	return []int{-1}
}

// ["ATM","deposit","withdraw","deposit","withdraw","withdraw"]
// [[],[[0,0,1,2,1]],[600],[[0,1,0,1,1]],[600],[550]]
func main() {
	obj := Constructor()
	fmt.Println()
	obj.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(obj.Withdraw(600))
}
