package main

import "fmt"

// Seuils officiels de l'IMC (Indice de Masse Corporelle).
const (
	IMCMaigreur = 18.5
	IMCNormal   = 25.0
	IMCSurpoids = 30.0
)

// Nom de l'utilisateur pour personnaliser l'affichage (bonus).
const Nom = "Alex"

func main() {
	// Données de l'utilisateur (poids en kg, taille en mètres).
	poids := 70.5
	taille := 1.75

	// Calcul de l'IMC : poids / (taille²)
	imc := poids / (taille * taille)

	// Salutation personnalisée
	fmt.Printf("Bonjour %s !\n\n", Nom)

	// Affichage de l'IMC avec exactement 2 décimales
	fmt.Printf("Votre IMC est : %.2f\n", imc)

	// Détermination de la catégorie selon les seuils
	categorie := ""
	if imc < IMCMaigreur {
		categorie = "Maigreur"
	} else if imc >= IMCMaigreur && imc < IMCNormal {
		categorie = "Normal"
	} else if imc >= IMCNormal && imc < IMCSurpoids {
		categorie = "Surpoids"
	} else {
		categorie = "Obésité"
	}

	fmt.Printf("Catégorie : %s\n", categorie)
}
