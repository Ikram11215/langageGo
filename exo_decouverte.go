package main

import "fmt"

// Partie 1 : j'ai utilisé iota pour numéroter automatiquement 4 mentions (0 → 3 sans les écrire à la main).
const (
	Insuffisant = iota
	Passable
	Bien
	Excellent
)

var nomsMentions = []string{"Insuffisant", "Passable", "Bien", "Excellent"}

func main() {
	// Partie 1 : affichage des valeurs générées par iota pour vérifier l'incrémentation.
	fmt.Printf("Insuffisant = %d\n", Insuffisant)
	fmt.Printf("Passable    = %d\n", Passable)
	fmt.Printf("Bien        = %d\n", Bien)
	fmt.Printf("Excellent   = %d\n", Excellent)
	fmt.Println()

	// Partie 2 : slice littéral + append + len + boucle for classique (Go n'a pas de while).
	notes := []int{14, 7, 18}
	notes = append(notes, 16, 9)
	fmt.Printf("Nombre de notes : %d\n", len(notes))
	for i := 0; i < len(notes); i++ {
		fmt.Printf("Note[%d] = %d\n", i, notes[i])
	}
	fmt.Println()

	// Partie 3 : switch avec fallthrough sur note < 8 pour voir deux messages d'affilée (ex. note 7).
	for i := 0; i < len(notes); i++ {
		note := notes[i]
		fmt.Printf("Note %d : ", note)
		switch {
		case note < 8:
			fmt.Printf("%s ", nomsMentions[Insuffisant])
			fallthrough
		case note < 12:
			fmt.Printf("%s ", nomsMentions[Passable])
		case note < 16:
			fmt.Printf("%s ", nomsMentions[Bien])
		default:
			fmt.Printf("%s ", nomsMentions[Excellent])
		}
		fmt.Println()
	}
	fmt.Println()

	// Partie 4 : for range pour récupérer index et valeur du slice en une seule boucle.
	for i, note := range notes {
		fmt.Printf("Index %d → note %d\n", i, note)
	}
}
