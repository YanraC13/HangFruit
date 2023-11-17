package main

import (
    "fmt"
    "os"
    "strings"
    "math/rand"
    "time"
)

var tableau []string
var maxTentatives = 7
var EtapesPendu = -1

func menu() {
    //Appel de la fonction menu
    time.Sleep(1 * time.Second)
    fmt.Println("\nWelcome =)")
    time.Sleep(1 * time.Second)
    fmt.Println("\nWould you like to play ?\nyes or no")


    var answer string

    fmt.Scanln(&answer)

    if answer == "yes" {
        // Appel de lsword pour obtenir un nouveau mot
        motChoisi = lsword()
        fmt.Println("\nYou're about to start. Insert a letter :")
        boucle()
    } else {
        fmt.Println("\nOkay...Would you like to do the tutorial ?")
        fmt.Println("yes or no")

        fmt.Scanln(&answer)
        if answer == "yes" {
            tuto()
        } else {
            os.Exit(0)
        }
        
    }
}
	
	
var motChoisi string
	
func init() {
    rand.Seed(time.Now().UnixNano())
    motChoisi = lsword()
}
	
func lsword() string {
	mots := []string{
        "apple",
        "banana",
        "cherry",
        "grape",
        "lemon",
        "orange",
        "strawberry",
        "watermelon",
    }
	
	indiceAleatoire := rand.Intn(len(mots))
	return strings.Replace(mots[indiceAleatoire], "\"", "", -1)
	}
	
func tiret() {
	tailledumot := len(motChoisi)
	
	for i := 0; i < tailledumot; i++ {
		if len(tableau) < len(motChoisi) {
			tableau = append(tableau, "_")
		}
	}
	fmt.Println(strings.Join(tableau, ""))
}
	
func comparer(lettreAComparer string) bool {
	lettreCorrespondante := false
	
	for i, lettre := range motChoisi {
		if string(lettre) == lettreAComparer {
			tableau[i] = string(lettre)
			lettreCorrespondante = true
		}
	}
	
	if lettreCorrespondante {
		fmt.Printf("\n%s is part of the word\n", lettreAComparer)
	} else {
			
		fmt.Printf("\n%s is not part of the word.\n", lettreAComparer)
		
	}
	return lettreCorrespondante
}
	
func compare() bool {
		
	var insert string
	fmt.Scanln(&insert)
	return comparer(insert)
}
	
	
	
func boucle() {
        
    MotARemplir := len(motChoisi)
    motGagne := false
    tentativesEffectuees := 0

    for tentativesEffectuees <= maxTentatives && MotARemplir >= 0 {
        AffichePendu()
        tiret()
        
        if compare() {
            MotARemplir--
        } else {
            tentativesEffectuees++
            EtapesPendu++
        }
        
        fmt.Printf("\nRemaining attempt : %d\n", maxTentatives-tentativesEffectuees)
        fmt.Println() 

    if strings.Join(tableau, "") == motChoisi {
        motGagne = true
        break
    }
}

    if motGagne {
        fmt.Println("\nCongratulations! You've guessed the word")
    } else {
        fmt.Println("\nSorry =(, you couldn't guess the word. The word was:", motChoisi)
    }
}
	
func AffichePendu() {
    switch EtapesPendu {
    case 0:
        fmt.Println("        /|\\")
    
    case 1:
        fmt.Println("         |")
        fmt.Println("         |")
        fmt.Println("        /|\\")
		
	case 2:
        fmt.Println("         |  ")
        fmt.Println("         | ")
        fmt.Println("         |")
        fmt.Println("         |")
        fmt.Println("         |")
        fmt.Println("        /|\\")

    case 3:
        fmt.Println("      \\  |  /")
        fmt.Println("       \\ | /")
        fmt.Println("         |")
        fmt.Println("         |")
        fmt.Println("         |")
        fmt.Println("        /|\\")
    
    case 4:
        fmt.Println("      _______     ")
        fmt.Println("   __|       |__")
        fmt.Println("  |             | ")
        fmt.Println("  |   \\  |  /   | ")
        fmt.Println("  |____\\ | /____| ")
        fmt.Println("         |")
        fmt.Println("         |")
        fmt.Println("         |")
        fmt.Println("        /|\\")
	case 5:
        fmt.Println("      _______     ")
        fmt.Println("   __|       |__")
        fmt.Println("  |             | ")
        fmt.Println("  |   \\  |  /   | ")
        fmt.Println("  |____\\ | /____| ")
        fmt.Println("         |     |  ")
        fmt.Println("         |         ")
        fmt.Println("         |")
        fmt.Println("        /|\\")

    case 6:
        fmt.Println("      _______     ")
        fmt.Println("   __|       |__")
        fmt.Println("  |             | ")
        fmt.Println("  |   \\  |  /   | ")
        fmt.Println("  |____\\ | /____| ")
        fmt.Println("         |     |  ")
        fmt.Println("         |     0  ")
        fmt.Println("         |")
        fmt.Println("        /|\\")
    }
}	
