package main	
import "fmt" 

func main() {

	var koro, waktu, j, m, d, i, n, dKoro, d2Koro int

	fmt.Scan(&n)

	koro = 0 
	for i = 1; i <= n; i++ {

		fmt.Scan(&waktu)
		koro = koro + waktu


	}
	
	d = koro * 60
	dKoro = d / 20
	m = dKoro / 60
	j = m / 60
	d2Koro = d / 120

	fmt.Print(j, m, d2Koro)	

}

// 0 5 3
// 0 3 0
// 0 14 15