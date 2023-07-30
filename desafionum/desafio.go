//criar codigo que exiba todos os numeros entre 1 e 100 que sejam divisiveis por 3
// utilizar %

package main

import (
	"fmt"
)

func main (){
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Println("Pinpan")
			} else {
				fmt.Println("Pin")
			}
		} else if i % 5 == 0 {
			if i % 3 == 0 {
				fmt.Println("Pinpan")
			} else {
				fmt.Println("Pan")
			}
		} else {
			fmt.Println(i)
		}
	}
}