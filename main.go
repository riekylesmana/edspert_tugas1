package main

import "fmt"
import libGG1 "github.com/riekylesmana/edspert_tugas1_lib"
import libGG2 "github.com/riekylesmana/edspert_tugas1_lib/v2"
 

func main(){
	result := libGG1.CekGanjilGenap(1)
	result2 := libGG2.CekGanjilGenap(1,2,3,4,5)
	fmt.Println(result)
	fmt.Println(result2)
}
