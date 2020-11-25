package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type GlobalConfig struct {
	Name string
	Version string
	Id int
	Aim string

}

func main() {
	s:= "hello"
	fmt.Println(string(s[0]))
}

func Openfile(path string) {

	f, _ := os.OpenFile(filepath.Clean(path), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	info, _ := f.Stat()
	fmt.Println("mod is", uint32(info.Mode()))
}

func LoadCOnfig() *GlobalConfig {
	//file, _ := os.Create("test.json")
	//write := io.Writer(file)
	config := &GlobalConfig{
	}
	//bytes, _ := json.Marshal(&config)
	//write.Write(bytes)

	bytes,_ := ioutil.ReadFile("test.json")

	json.Unmarshal(bytes, &config)
	fmt.Println(config)

	return  config
}

func moveZeroes(nums []int)  {
	count := 0
	for i := range nums{
		if nums[i]==0 {
			count++
		} else {
			if count!=0 {
				nums[i-count], nums[i] = nums[i], nums[i-count]
			}
		}
	}
}


func maxProduct(nums []int) int {
	fmt.Println(nums)
	max := nums[0]
	sum := nums[0]
	for index := 1; index < len(nums); index ++ {

		//sum > sum * curnum
		fmt.Println("num is ", nums[1])
		if  sum != 0 {
			if sum > sum * nums[index] {
				if max < sum {
					max = sum
				}
				sum = nums[index]
				fmt.Println("sum1 is ", sum)
			} else {
				sum = sum * nums[index]
				fmt.Println("sum2 is ", sum)
			}
		} else {
			if max < 0 {
				max = 0
			}
			fmt.Println("sum3 is ", sum)
			sum = nums[index]

		}

	}

	if max < sum {
		max = sum
	}

	return  max
}

