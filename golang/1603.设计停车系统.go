/*
 * @lc app=leetcode.cn id=1603 lang=golang
 *
 * [1603] 设计停车系统
 */
package main

import "fmt"

// @lc code=start
type ParkingSystem struct {
	small  int
	medium int
	big    int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{small: small, medium: medium, big: big}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big--
			return true
		} else {
			return false
		}
	case 2:
		if this.medium > 0 {
			this.medium--
			return true
		} else {
			return false
		}

	case 3:
		if this.small > 0 {
			this.small--
			return true
		} else {
			return false
		}

	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end
func main() {
	obj := Constructor(1, 1, 0)
	fmt.Println(obj.AddCar(1))
}
