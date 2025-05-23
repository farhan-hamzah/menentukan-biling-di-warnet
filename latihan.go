package main

import "fmt"

const NMax int = 1000
type tabInt [NMax]int
type tabStr [NMax]string

func main() {
	var n, i, harga, total int
	var nama, ruangan tabStr
	var jam tabInt
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&nama[i], &jam[i], &ruangan[i])
	}
  
	for i = 0; i < n; i++{
		if ruangan[i] == "gaming"{
			harga = 14000 * jam[i]
		}else if ruangan[i] == "biasa"{
			harga = 10000 * jam[i]
		}
		total += harga
		fmt.Println("Total yang harus dibayar", nama[i], "sebesar", harga)
	}
	fmt.Println("Biaya total sebesar", total)
}