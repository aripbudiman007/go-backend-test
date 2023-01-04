package main

import "fmt"

func main() {
	input := 100000

	current := input
	pecahan := map[string]int{}

	for current != 0 {
		fmt.Println(current, current != 0)

		ratusan := current / 100000
		current = current % 100000

		limaPuluhRibuan := current / 50000
		current = current % 50000

		duaPuluhRibuan := current / 20000
		current = current % 20000

		sepuluhRibuan := current / 10000
		current = current % 10000

		limaRibuan := current / 5000
		current = current % 5000

		duaRibuan := current / 2000
		current = current % 2000

		seribu := current / 1000
		current = current % 1000

		limaRatus := current / 500
		current = current % 500

		duaRatus := current / 200
		current = current % 200

		seratus := current / 100
		current = current % 100

		pecahan["Rp. 100.000"] = ratusan
		pecahan["RP. 50.000"] = limaPuluhRibuan
		pecahan["Rp. 20.000"] = duaPuluhRibuan
		pecahan["Rp. 10.000"] = sepuluhRibuan
		pecahan["Rp. 5.000"] = limaRibuan
		pecahan["Rp. 2.000"] = duaRibuan
		pecahan["Rp. 1.000"] = seribu
		pecahan["Rp. 500"] = limaRatus
		pecahan["Rp. 200"] = duaRatus
		pecahan["Rp. 100"] = seratus
	}

	fmt.Println(pecahan)
}
