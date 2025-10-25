package mainp
import "fmt"

func main() {

	var x, n, i, vn, vx int
	var hasil float64

	fmt.Scan(&x, &n)

	for i = 1; i <= n; i++ {

		vn = n + 1 - i
		vx = i * x
		hasil = hasil + float64(vn)/float64(vx)

	}

	fmt.Printf("%.3f", hasil)

}
