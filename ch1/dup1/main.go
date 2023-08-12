// dup1 exibe o texto de toda linha que aparece mais de umas vez na entrada padrão, precedida por sua contagem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		counts[line] = counts[line] + 1
	}
	
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
	}
}
