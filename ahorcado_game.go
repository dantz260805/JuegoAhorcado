package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Palabras para jugar
var animales = []string{"perro", "gato", "conejo", "leon", "tigre", "elefante", "jirafa"}
var frutas = []string{"manzana", "pera", "uva", "sandia", "melon", "naranja"}
var paises = []string{"chile", "peru", "mexico", "españa", "china", "japon"}

// Las partes del dibujo
var dibujos = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
========`,
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		// Menu
		fmt.Println("\n=== AHORCADO ===")
		fmt.Println("1. Jugar")
		fmt.Println("2. Salir")
		fmt.Print("Elige: ")
		
		scanner.Scan()
		opcion := scanner.Text()
		
		if opcion == "1" {
			jugar(scanner)
		} else if opcion == "2" {
			fmt.Println("Chao!")
			break
		} else {
			fmt.Println("Esa opcion no sirve")
		}
	}
}

func jugar(scanner *bufio.Scanner) {
	// Elegir categoria
	fmt.Println("\nElige categoria:")
	fmt.Println("1. Animales")
	fmt.Println("2. Frutas")
	fmt.Println("3. Paises")
	fmt.Print("Categoria: ")
	
	scanner.Scan()
	cat := scanner.Text()
	
	var palabras []string
	var nombreCat string
	
	if cat == "1" {
		palabras = animales
		nombreCat = "Animales"
	} else if cat == "2" {
		palabras = frutas
		nombreCat = "Frutas"
	} else {
		palabras = paises
		nombreCat = "Paises"
	}
	
	// Elegir palabra aleatoria
	palabra := palabras[rand.Intn(len(palabras))]
	palabra = strings.ToLower(palabra)
	
	// Variables del juego
	adivinadas := make(map[string]bool)
	errores := 0
	letrasMal := []string{}
	
	// Bucle principal
	for {
		// Mostrar estado
		fmt.Println("\n" + strings.Repeat("=", 40))
		fmt.Println("Categoria:", nombreCat)
		fmt.Println(dibujos[errores])
		
		// Mostrar palabra con guiones
		fmt.Print("Palabra: ")
		for i := 0; i < len(palabra); i++ {
			letra := string(palabra[i])
			if adivinadas[letra] {
				fmt.Print(letra + " ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
		
		// Mostrar letras falladas
		if len(letrasMal) > 0 {
			fmt.Print("Letras mal: ")
			for _, l := range letrasMal {
				fmt.Print(l + " ")
			}
			fmt.Println()
		}
		
		fmt.Println("Errores:", errores, "/6")
		
		// Verificar si gano
		gano := true
		for i := 0; i < len(palabra); i++ {
			letra := string(palabra[i])
			if !adivinadas[letra] {
				gano = false
				break
			}
		}
		
		if gano {
			fmt.Println("\n¡FELICITACIONES! Adivinaste la palabra:", palabra)
			fmt.Println("Presiona Enter para volver al menu")
			scanner.Scan()
			return
		}
		
		// Verificar si perdio
		if errores >= 6 {
			fmt.Println("\nPERDISTE! La palabra era:", palabra)
			fmt.Println("Presiona Enter para volver al menu")
			scanner.Scan()
			return
		}
		
		// Pedir letra
		fmt.Print("\nIngresa una letra: ")
		scanner.Scan()
		letra := strings.ToLower(strings.TrimSpace(scanner.Text()))
		
		// Validar que sea solo una letra
		if len(letra) != 1 {
			fmt.Println("Solo una letra!")
			continue
		}
		
		// Verificar si ya la uso
		if adivinadas[letra] {
			fmt.Println("Ya usaste esa letra!")
			continue
		}
		
		// Verificar si esta en la palabra
		encontrada := false
		for i := 0; i < len(palabra); i++ {
			if string(palabra[i]) == letra {
				encontrada = true
				break
			}
		}
		
		if encontrada {
			adivinadas[letra] = true
			fmt.Println("Bien! Esa letra esta!")
		} else {
			letrasMal = append(letrasMal, letra)
			errores++
			fmt.Println("Mala suerte! Esa letra no esta")
		}
	}
}
