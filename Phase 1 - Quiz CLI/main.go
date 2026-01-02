package main

import (
	"fmt"
	"strings"
	"time"
)

type Question struct {
	enonce  string
	choix   []string
	reponse string
}

var questionsFaciles []Question
var questionsMoyennes []Question
var questionsDifficiles []Question

func main() {
	ClearScreen()
	var menuchoice float64
	var reponsequestion string
	sleep := time.Duration(2) * time.Second
	var videur string
	var compteur int
	questionsFaciles = []Question{
		{
			enonce:  "Quel est le nom du studio qui a créé la console PlayStation ?",
			choix:   []string{"A - Nintendo", "B - Sega", "C - Sony", "D - Microsoft"},
			reponse: "C",
		},
		{
			enonce:  "Dans quel jeu doit-on survivre face à une zone qui rétrécit sur une île ?",
			choix:   []string{"A - FIFA", "B - Fortnite", "C - Tetris", "D - Sims"},
			reponse: "B",
		},
		{
			enonce:  "Comment s'appelle le personnage principal de la saga The Legend of Zelda ?",
			choix:   []string{"A - Zelda", "B - Ganondorf", "C - Peach", "D - Link"},
			reponse: "D",
		},
		{
			enonce:  "Quelle est la couleur du fantôme le plus connu dans Pac-Man ?",
			choix:   []string{"A - Rouge", "B - Noir", "C - Vert", "D - Marron"},
			reponse: "A",
		},
		{
			enonce:  "Quel bloc doit-on casser pour obtenir des ressources dans Minecraft ?",
			choix:   []string{"A - Du plastique", "B - Du bois", "C - Du carton", "D - Du tissu"},
			reponse: "B",
		},
		{
			enonce:  "Dans Pokémon, quel est le numéro du célèbre Pikachu dans le Pokédex ?",
			choix:   []string{"A - 1", "B - 150", "C - 25", "D - 42"},
			reponse: "C",
		},
		{
			enonce:  "Quel jeu consiste à empiler des briques de formes différentes pour faire des lignes ?",
			choix:   []string{"A - Tetris", "B - Pong", "C - Snake", "D - Doom"},
			reponse: "A",
		},
		{
			enonce:  "Quel est le nom de la ville dans laquelle se déroule GTA V ?",
			choix:   []string{"A - Liberty City", "B - Los Santos", "C - Vice City", "D - San Fierro"},
			reponse: "B",
		},
		{
			enonce:  "Quel animal est Sonic, le personnage bleu ultra rapide ?",
			choix:   []string{"A - Un chat", "B - Un chien", "C - Un hérisson", "D - Un raton laveur"},
			reponse: "C",
		},
		{
			enonce:  "Sur quelle console est sorti le tout premier jeu Super Mario Bros ?",
			choix:   []string{"A - GameCube", "B - NES", "C - Wii", "D - Nintendo Switch"},
			reponse: "B",
		},
	}

	questionsMoyennes = []Question{
		{
			enonce:  "Quel studio est à l'origine de la série des 'Dark Souls' ?",
			choix:   []string{"A - FromSoftware", "B - Ubisoft", "C - Bethesda", "D - Capcom"},
			reponse: "A",
		},
		{
			enonce:  "En quelle année est sorti le premier jeu Minecraft (version Alpha) ?",
			choix:   []string{"A - 2005", "B - 2009", "C - 2011", "D - 2013"},
			reponse: "B",
		},
		{
			enonce:  "Quel est le nom de l'intelligence artificielle dans la saga 'Halo' ?",
			choix:   []string{"A - Alexa", "B - Siri", "C - Cortana", "D - Glados"},
			reponse: "C",
		},
		{
			enonce:  "Dans 'The Witcher 3', quel est le nom de la fille adoptive de Geralt ?",
			choix:   []string{"A - Yennefer", "B - Triss", "C - Ciri", "D - Keira"},
			reponse: "C",
		},
		{
			enonce:  "Quel jeu a popularisé le genre 'Battle Royale' avant Fortnite ?",
			choix:   []string{"A - H1Z1", "B - PUBG", "C - DayZ", "D - Apex Legends"},
			reponse: "B",
		},
		{
			enonce:  "Quel est le nom du monde dans lequel se déroule 'League of Legends' ?",
			choix:   []string{"A - Azeroth", "B - Runeterra", "C - Sanctuaire", "D - Hyrule"},
			reponse: "B",
		},
		{
			enonce:  "Dans 'Metal Gear Solid', quel est le vrai nom de Solid Snake ?",
			choix:   []string{"A - John", "B - Jack", "C - David", "D - Liquid"},
			reponse: "C",
		},
		{
			enonce:  "Quelle entreprise a créé la console Dreamcast ?",
			choix:   []string{"A - Nintendo", "B - Sony", "C - Sega", "D - Atari"},
			reponse: "C",
		},
		{
			enonce:  "Quel jeu met en scène un groupe de survivants contre des infectés appelés 'Claqueurs' ?",
			choix:   []string{"A - Resident Evil", "B - Dying Light", "C - The Last of Us", "D - Days Gone"},
			reponse: "C",
		},
		{
			enonce:  "Quel est le nom du moteur de jeu développé par Epic Games ?",
			choix:   []string{"A - Unity", "B - Frostbite", "C - CryEngine", "D - Unreal Engine"},
			reponse: "D",
		},
	}

	questionsDifficiles = []Question{
		{
			enonce:  "Quel est le premier jeu vidéo de l'histoire à avoir été commercialisé sur une borne d'arcade en 1971 ?",
			choix:   []string{"A - Pong", "B - Computer Space", "C - Space Invaders", "D - Pac-Man"},
			reponse: "B",
		},
		{
			enonce:  "Qui est le créateur original de la série 'Metal Gear' et de 'Death Stranding' ?",
			choix:   []string{"A - Shinji Mikami", "B - Hidetaka Miyazaki", "C - Hideo Kojima", "D - Shigeru Miyamoto"},
			reponse: "C",
		},
		{
			enonce:  "Dans 'The Elder Scrolls V: Skyrim', comment appelle-t-on la langue des dragons ?",
			choix:   []string{"A - Le Dovahzul", "B - Le Draconique", "C - Le Thalmor", "D - Le Daedrique"},
			reponse: "A",
		},
		{
			enonce:  "Quel était le nom de code interne de la console Nintendo GameCube durant son développement ?",
			choix:   []string{"A - Project Reality", "B - Atlantis", "C - Dolphin", "D - Revolution"},
			reponse: "C",
		},
		{
			enonce:  "Dans quel jeu incarne-t-on un personnage nommé 'Sora' qui voyage avec Donald et Dingo ?",
			choix:   []string{"A - Final Fantasy", "B - Kingdom Hearts", "C - Dragon Quest", "D - NieR"},
			reponse: "B",
		},
		{
			enonce:  "Quel jeu détient le record du budget de développement le plus élevé de l'histoire (hors marketing) ?",
			choix:   []string{"A - GTA V", "B - Red Dead Redemption 2", "C - Star Citizen", "D - Cyberpunk 2077"},
			reponse: "C",
		},
		{
			enonce:  "En quelle année la console Super Nintendo (SNES) est-elle sortie en Europe ?",
			choix:   []string{"A - 1990", "B - 1991", "C - 1992", "D - 1993"},
			reponse: "C",
		},
		{
			enonce:  "Quel compositeur est célèbre pour avoir écrit les musiques de la saga 'Final Fantasy' ?",
			choix:   []string{"A - Koji Kondo", "B - Nobuo Uematsu", "C - Akira Yamaoka", "D - Yoko Shimomura"},
			reponse: "B",
		},
		{
			enonce:  "Comment s'appelle le premier boss que l'on affronte dans le jeu 'Dark Souls' ?",
			choix:   []string{"A - Démon de l'Asile", "B - Gargouilles de la Cloche", "C - Queelag", "D - Ornstein"},
			reponse: "A",
		},
		{
			enonce:  "Dans la série 'Resident Evil', quel virus est à l'origine de la catastrophe de Raccoon City ?",
			choix:   []string{"A - Virus G", "B - Virus T", "C - Virus Las Plagas", "D - Virus C"},
			reponse: "B",
		},
	}

	for {
		ClearScreen()
		fmt.Println("Bonjour quel Quiz voulez vous faire aujourd'hui ?")
		fmt.Println("1 - Quiz Facile")
		fmt.Println("2 - Quiz Moyen")
		fmt.Println("3 - Quiz Difficile")
		fmt.Println("4 - Quitter")
		_, err := fmt.Scan(&menuchoice)
		if err != nil {
			fmt.Println("❌ Choix impossible, réessayez.")
			fmt.Scanln(&videur)
			time.Sleep(sleep)
			continue

		}
		switch menuchoice {
		case 1:
		questionsfaciles:
			for {

				ClearScreen()

				for i := 0; i < len(questionsFaciles); i++ {
					fmt.Println(questionsFaciles[i].enonce)
					fmt.Println(questionsFaciles[i].choix[0])
					fmt.Println(questionsFaciles[i].choix[1])
					fmt.Println(questionsFaciles[i].choix[2])
					fmt.Println(questionsFaciles[i].choix[3])

					_, errreponse := fmt.Scan(&reponsequestion)
					if reponsequestion == "A" || reponsequestion == "B" || reponsequestion == "C" || reponsequestion == "D" {
						reponsequestion = strings.ToLower(reponsequestion)
					}
					if errreponse != nil || len(reponsequestion) > 1 || reponsequestion != "a" && reponsequestion != "b" && reponsequestion != "c" && reponsequestion != "d" {
						fmt.Println("❌ Choix impossible, réessayez.")
						fmt.Scanln(&videur)
						time.Sleep(sleep)
						continue
					}

					if strings.ToUpper(reponsequestion) == questionsFaciles[i].reponse {
						compteur++
						continue
					}
				}
				if compteur <= 4 {
					fmt.Println("Vous êtes nul revoyez vos basiques, vous avez eu un score de", compteur)
				} else if compteur <= 7 {
					fmt.Println("Vous êtes Mid, vous avez eu un score de", compteur)
				} else if compteur <= 9 {
					fmt.Println("Vous êtes fort,bien joué à vous, vous avez eu un score de", compteur)
				} else {
					fmt.Println("Vous êtes le GOAT, bien joué à vous, vous avez eu un score de", compteur)
				}
				fmt.Println("Que voulez vous faire maintenant ?")
				fmt.Println("1 - Continuez les quizz !")
				fmt.Println("0 - Quittez")

				_, err2 := fmt.Scan(&menuchoice)
				if err2 != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				switch menuchoice {
				case 1:
					compteur = 0
					break questionsfaciles
				case 0:
					return
				}
			}
		case 2:
		questionsMoyennes:
			for {

				ClearScreen()

				for i := 0; i < len(questionsMoyennes); i++ {
					fmt.Println(questionsMoyennes[i].enonce)
					fmt.Println(questionsMoyennes[i].choix[0])
					fmt.Println(questionsMoyennes[i].choix[1])
					fmt.Println(questionsMoyennes[i].choix[2])
					fmt.Println(questionsMoyennes[i].choix[3])

					_, errreponse := fmt.Scan(&reponsequestion)
					if reponsequestion == "A" || reponsequestion == "B" || reponsequestion == "C" || reponsequestion == "D" {
						reponsequestion = strings.ToLower(reponsequestion)
					}
					if errreponse != nil || len(reponsequestion) > 1 || reponsequestion != "a" && reponsequestion != "b" && reponsequestion != "c" && reponsequestion != "d" {

						fmt.Println("❌ Choix impossible, réessayez.")
						fmt.Scanln(&videur)
						time.Sleep(sleep)
						continue
					}
					if strings.ToUpper(reponsequestion) == questionsMoyennes[i].reponse {
						compteur++
						continue
					}
				}
				if compteur <= 4 {
					fmt.Println("Vous êtes nul revoyez vos basiques, vous avez eu un score de", compteur)
				} else if compteur <= 7 {
					fmt.Println("Vous êtes Mid, vous avez eu un score de", compteur)
				} else if compteur <= 9 {
					fmt.Println("Vous êtes fort,bien joué à vous, vous avez eu un score de", compteur)
				} else {
					fmt.Println("Vous êtes le GOAT, bien joué à vous, vous avez eu un score de", compteur)
				}
				fmt.Println("Que voulez vous faire maintenant ?")
				fmt.Println("1 - Continuez les quizz !")
				fmt.Println("0 - Quittez")

				_, err2 := fmt.Scan(&menuchoice)
				if err2 != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				switch menuchoice {
				case 1:
					compteur = 0
					break questionsMoyennes
				case 0:
					return
				}
			}
		case 3:
		questionsDifficiles:
			for {

				ClearScreen()

				for i := 0; i < len(questionsDifficiles); i++ {
					fmt.Println(questionsDifficiles[i].enonce)
					fmt.Println(questionsDifficiles[i].choix[0])
					fmt.Println(questionsDifficiles[i].choix[1])
					fmt.Println(questionsDifficiles[i].choix[2])
					fmt.Println(questionsDifficiles[i].choix[3])

					_, errreponse := fmt.Scan(&reponsequestion)
					if reponsequestion == "A" || reponsequestion == "B" || reponsequestion == "C" || reponsequestion == "D" {
						reponsequestion = strings.ToLower(reponsequestion)
					}
					if errreponse != nil || len(reponsequestion) > 1 || reponsequestion != "a" && reponsequestion != "b" && reponsequestion != "c" && reponsequestion != "d" {
						fmt.Println("❌ Choix impossible, réessayez.")
						fmt.Scanln(&videur)
						time.Sleep(sleep)
						continue
					}
					if strings.ToUpper(reponsequestion) == questionsDifficiles[i].reponse {
						compteur++
						continue
					}
				}
				if compteur <= 4 {
					fmt.Println("Vous êtes nul revoyez vos basiques, vous avez eu un score de", compteur)
				} else if compteur <= 7 {
					fmt.Println("Vous êtes Mid, vous avez eu un score de", compteur)
				} else if compteur <= 9 {
					fmt.Println("Vous êtes fort,bien joué à vous, vous avez eu un score de", compteur)
				} else {
					fmt.Println("Vous êtes le GOAT, bien joué à vous, vous avez eu un score de", compteur)
				}
				fmt.Println("Bravo vous avez eu un score de", compteur)
				fmt.Println("Que voulez vous faire maintenant ?")
				fmt.Println("1 - Continuez les quizz !")
				fmt.Println("0 - Quittez")

				_, err2 := fmt.Scan(&menuchoice)
				if err2 != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				switch menuchoice {
				case 1:
					compteur = 0
					break questionsDifficiles
				case 0:
					return
				}
			}
		case 4:
			return
		default:
			fmt.Println("❌ Choix impossible, réessayez.")
			fmt.Scanln(&videur)
			time.Sleep(sleep)
		}
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
