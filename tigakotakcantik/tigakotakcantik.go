package tigakotakcantik

import "fmt"

type kotak int

type tigakotakcantik struct {
	kotak1 kotak
	kotak2 kotak
	kotak3 kotak
}

func NewTigaKotakCantik(kotak1 int, kotak2 int, kotak3 int) tigakotakcantik {
	return tigakotakcantik{
		kotak1: kotak(kotak1),
		kotak2: kotak(kotak2),
		kotak3: kotak(kotak3),
	}
}

func (tk *tigakotakcantik) PerformAdjustment() kotak {
	var permenEaten kotak = 0
	if tk.kotak1 < 1 || tk.kotak2 < 2 || tk.kotak3 < 3 {
		return -1
	}

	kotak32 := tk.kotak3 - tk.kotak2
	if kotak32 <= 0 {
		deduction := kotak32 - 1
		tk.kotak2 = tk.kotak2 + deduction
		permenEaten += deduction * -1
	}

	kotak21 := tk.kotak2 - tk.kotak1
	if kotak21 <= 0 {
		deduction := kotak21 - 1
		tk.kotak1 = tk.kotak1 + deduction
		permenEaten += deduction * -1

	}
	fmt.Printf("%+v kotak result", tk)
	return permenEaten
}
