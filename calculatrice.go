package main

import (
	"fmt"
	"io"
)

// operer effectue une opération arithmétique sur a et b.
// Retourne une erreur en cas de division par zéro ou d'opération inconnue.
func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+", "-", "*", "/":
		if op == "/" && b == 0 {
			return 0, fmt.Errorf("division par zéro")
		}
		return creerOperation(op)(a, b), nil
	default:
		return 0, fmt.Errorf("opération inconnue : %q", op)
	}
}

// creerOperation retourne une closure pour l'opération demandée.
func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 { return a / b }
	default:
		return func(a, b float64) float64 { return 0 }
	}
}

func main() {
	fmt.Println("Calculatrice — entrez deux nombres et une opération (+, -, *, /).")
	fmt.Println("Tapez 'quit' comme opération pour quitter.")
	fmt.Println()

	for {
		fmt.Print("a b op> ")

		a := 0.0
		b := 0.0
		op := ""

		_, err := fmt.Scan(&a, &b, &op)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Erreur de lecture :", err)
			continue
		}

		if op == "quit" {
			fmt.Println("Au revoir !")
			break
		}

		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Println("Erreur :", err)
			continue
		}

		fmt.Printf("Résultat : %g\n", resultat)
	}
}
