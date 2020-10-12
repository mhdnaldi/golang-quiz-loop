package main

import (
	"fmt"
	"log"
)

func cetak_angka(num int) {
	if num % 2 == 0 {
		log.Println("PARAMETER HARUS BILANGAN GANJIL :)")
	} else {
	
		for i:= 0 ; i < num ; i++ {
			for j:= 0; j < num; j++ {
				if j < 1 || j == num-1 || i == num/2 {
					fmt.Print("* ") 
				} else {
					fmt.Print("= ")
				}
			}
			fmt.Println("\n")
		}
	}


}

func main() {
	cetak_angka(5)
}

