package main

import (
	"errors"
	"fmt"
	"strings"
)

// Produit représente un article vendu dans la boutique.
type Produit struct {
	ID         int
	Nom        string
	Marque     string
	Prix       float64
	Stock      int
	Categorie  string
	Actif      bool
}

// Catalogue regroupe l'ensemble des produits disponibles.
type Catalogue struct {
	produits []Produit
}

// ErrIDDuplique est renvoyée lorsqu'un produit avec le même ID existe déjà.
var ErrIDDuplique = errors.New("un produit avec cet ID existe déjà dans le catalogue")

// ErrProduitIntrouvable est renvoyée lorsqu'aucun produit ne correspond à l'ID demandé.
var ErrProduitIntrouvable = errors.New("produit introuvable")

// ErrStockInsuffisant est renvoyée lorsque la quantité demandée dépasse le stock disponible.
var ErrStockInsuffisant = errors.New("stock insuffisant pour cette vente")

// AjouterProduit ajoute un produit au catalogue si son ID est unique.
func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, existant := range c.produits {
		if existant.ID == p.ID {
			return fmt.Errorf("%w (ID %d)", ErrIDDuplique, p.ID)
		}
	}
	c.produits = append(c.produits, p)
	return nil
}

// TrouverParID retourne le produit correspondant à l'identifiant donné.
func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.produits {
		if p.ID == id {
			return p, nil
		}
	}
	return Produit{}, fmt.Errorf("%w (ID %d)", ErrProduitIntrouvable, id)
}

// TrouverParCategorie retourne tous les produits d'une catégorie (insensible à la casse).
func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	resultat := make([]Produit, 0)
	for _, p := range c.produits {
		if strings.EqualFold(p.Categorie, cat) {
			resultat = append(resultat, p)
		}
	}
	return resultat
}

// AppliquerReduction diminue le prix des produits d'une catégorie du pourcentage indiqué.
func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	modifies := 0
	facteur := 1.0 - (pct / 100.0)

	for i := range c.produits {
		if strings.EqualFold(c.produits[i].Categorie, categorie) && c.produits[i].Actif {
			c.produits[i].Prix = c.produits[i].Prix * facteur
			modifies++
		}
	}
	return modifies
}

// Vendre retire une quantité du stock d'un produit identifié par son ID.
func (c *Catalogue) Vendre(id int, qte int) error {
	if qte <= 0 {
		return fmt.Errorf("la quantité doit être strictement positive (reçu : %d)", qte)
	}

	for i := range c.produits {
		if c.produits[i].ID != id {
			continue
		}
		if c.produits[i].Stock < qte {
			return fmt.Errorf("%w : disponible %d, demandé %d", ErrStockInsuffisant, c.produits[i].Stock, qte)
		}
		c.produits[i].Stock -= qte
		return nil
	}
	return fmt.Errorf("%w (ID %d)", ErrProduitIntrouvable, id)
}

// Rapport résume le catalogue : nombre de produits uniques et valeur totale du stock.
func (c Catalogue) Rapport() string {
	valeurStock := 0.0
	for _, p := range c.produits {
		valeurStock += p.Prix * float64(p.Stock)
	}
	return fmt.Sprintf(
		"Rapport TechShop\n"+
			"  Produits uniques : %d\n"+
			"  Valeur totale du stock : %.2f €",
		len(c.produits),
		valeurStock,
	)
}

// nouveauCatalogueInitialise crée un catalogue pré-rempli avec 5 produits high-tech.
func nouveauCatalogueInitialise() *Catalogue {
	catalogue := &Catalogue{}

	produitsInit := []Produit{
		{ID: 1, Nom: "iPhone 15", Marque: "Apple", Prix: 969.00, Stock: 25, Categorie: "Smartphone", Actif: true},
		{ID: 2, Nom: "MacBook Pro M3", Marque: "Apple", Prix: 2499.00, Stock: 12, Categorie: "Ordinateur", Actif: true},
		{ID: 3, Nom: "WH-1000XM5", Marque: "Sony", Prix: 399.99, Stock: 40, Categorie: "Audio", Actif: true},
		{ID: 4, Nom: "iPad Air", Marque: "Apple", Prix: 699.00, Stock: 18, Categorie: "Tablette", Actif: true},
		{ID: 5, Nom: "Galaxy S24", Marque: "Samsung", Prix: 899.00, Stock: 30, Categorie: "Smartphone", Actif: true},
	}

	for _, p := range produitsInit {
		if err := catalogue.AjouterProduit(p); err != nil {
			fmt.Println("Erreur initialisation catalogue :", err)
		}
	}
	return catalogue
}

func afficherMenu() {
	fmt.Println("\n========== TechShop CLI ==========")
	fmt.Println("[1] Ajouter un produit")
	fmt.Println("[2] Chercher un produit")
	fmt.Println("[3] Soldes (réduction par catégorie)")
	fmt.Println("[4] Vendre")
	fmt.Println("[5] Rapport")
	fmt.Println("[0] Quitter")
	fmt.Print("Votre choix : ")
}

func lireEntier(prompt string) (int, error) {
	fmt.Print(prompt)
	var valeur int
	_, err := fmt.Scan(&valeur)
	return valeur, err
}

func lireTexte(prompt string) (string, error) {
	fmt.Print(prompt)
	var texte string
	_, err := fmt.Scan(&texte)
	return texte, err
}

func lireFloat(prompt string) (float64, error) {
	fmt.Print(prompt)
	var valeur float64
	_, err := fmt.Scan(&valeur)
	return valeur, err
}

func afficherProduit(p Produit) {
	fmt.Printf(
		"  ID %d | %s (%s) | %.2f € | Stock: %d | Catégorie: %s | Actif: %t\n",
		p.ID, p.Nom, p.Marque, p.Prix, p.Stock, p.Categorie, p.Actif,
	)
}

func menuAjouterProduit(catalogue *Catalogue) {
	fmt.Println("\n--- Ajout d'un produit ---")

	id, err := lireEntier("ID : ")
	if err != nil {
		fmt.Println("Erreur de saisie (ID) :", err)
		return
	}

	nom, err := lireTexte("Nom : ")
	if err != nil {
		fmt.Println("Erreur de saisie (nom) :", err)
		return
	}

	marque, err := lireTexte("Marque : ")
	if err != nil {
		fmt.Println("Erreur de saisie (marque) :", err)
		return
	}

	prix, err := lireFloat("Prix (€) : ")
	if err != nil {
		fmt.Println("Erreur de saisie (prix) :", err)
		return
	}

	stock, err := lireEntier("Stock : ")
	if err != nil {
		fmt.Println("Erreur de saisie (stock) :", err)
		return
	}

	categorie, err := lireTexte("Catégorie : ")
	if err != nil {
		fmt.Println("Erreur de saisie (catégorie) :", err)
		return
	}

	var actifInt int
	fmt.Print("Actif (1 = oui, 0 = non) : ")
	_, err = fmt.Scan(&actifInt)
	if err != nil {
		fmt.Println("Erreur de saisie (actif) :", err)
		return
	}

	produit := Produit{
		ID:        id,
		Nom:       nom,
		Marque:    marque,
		Prix:      prix,
		Stock:     stock,
		Categorie: categorie,
		Actif:     actifInt == 1,
	}

	if err := catalogue.AjouterProduit(produit); err != nil {
		fmt.Println("Impossible d'ajouter le produit :", err)
		return
	}
	fmt.Println("Produit ajouté avec succès.")
}

func menuChercherProduit(catalogue *Catalogue) {
	fmt.Println("\n--- Recherche ---")
	fmt.Println("[A] Par ID")
	fmt.Println("[B] Par catégorie")
	fmt.Print("Mode de recherche : ")

	var mode string
	_, err := fmt.Scan(&mode)
	if err != nil {
		fmt.Println("Erreur de saisie :", err)
		return
	}

	switch strings.ToUpper(mode) {
	case "A":
		id, err := lireEntier("ID recherché : ")
		if err != nil {
			fmt.Println("Erreur de saisie :", err)
			return
		}
		produit, err := catalogue.TrouverParID(id)
		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}
		fmt.Println("Produit trouvé :")
		afficherProduit(produit)

	case "B":
		categorie, err := lireTexte("Catégorie recherchée : ")
		if err != nil {
			fmt.Println("Erreur de saisie :", err)
			return
		}
		resultats := catalogue.TrouverParCategorie(categorie)
		if len(resultats) == 0 {
			fmt.Println("Aucun produit trouvé pour cette catégorie.")
			return
		}
		fmt.Printf("%d produit(s) trouvé(s) :\n", len(resultats))
		for _, p := range resultats {
			afficherProduit(p)
		}

	default:
		fmt.Println("Mode invalide. Utilisez A ou B.")
	}
}

func menuSoldes(catalogue *Catalogue) {
	fmt.Println("\n--- Soldes ---")

	categorie, err := lireTexte("Catégorie : ")
	if err != nil {
		fmt.Println("Erreur de saisie (catégorie) :", err)
		return
	}

	pct, err := lireFloat("Pourcentage de réduction : ")
	if err != nil {
		fmt.Println("Erreur de saisie (pourcentage) :", err)
		return
	}
	if pct < 0 || pct > 100 {
		fmt.Println("Erreur : le pourcentage doit être compris entre 0 et 100.")
		return
	}

	nb := catalogue.AppliquerReduction(categorie, pct)
	fmt.Printf("Réduction de %.0f%% appliquée sur %d produit(s) de la catégorie « %s ».\n", pct, nb, categorie)
}

func menuVendre(catalogue *Catalogue) {
	fmt.Println("\n--- Vente ---")

	id, err := lireEntier("ID du produit : ")
	if err != nil {
		fmt.Println("Erreur de saisie (ID) :", err)
		return
	}

	qte, err := lireEntier("Quantité : ")
	if err != nil {
		fmt.Println("Erreur de saisie (quantité) :", err)
		return
	}

	if err := catalogue.Vendre(id, qte); err != nil {
		fmt.Println("Vente impossible :", err)
		return
	}
	fmt.Println("Vente enregistrée avec succès.")
}

func main() {
	catalogue := nouveauCatalogueInitialise()
	fmt.Println("Bienvenue sur TechShop ! Catalogue initialisé avec 5 produits.")

	for {
		afficherMenu()

		choix, err := lireEntier("")
		if err != nil {
			fmt.Println("Erreur de saisie du menu :", err)
			continue
		}

		switch choix {
		case 1:
			menuAjouterProduit(catalogue)
		case 2:
			menuChercherProduit(catalogue)
		case 3:
			menuSoldes(catalogue)
		case 4:
			menuVendre(catalogue)
		case 5:
			fmt.Println()
			fmt.Println(catalogue.Rapport())
		case 0:
			fmt.Println("Merci d'avoir utilisé TechShop. À bientôt !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez sélectionner une option entre 0 et 5.")
		}
	}
}
