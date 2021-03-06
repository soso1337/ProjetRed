package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (p *Personnage) spiritTurn() {
	fmt.Println("1 >> Attaquer")
	fmt.Println("2 >> Inventaire")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.spiritAttack()
	case "2":
		ClearLog()
		p.accesInventory()
		p.useInventory(true)
	default:
		p.spiritTurn()
	}
}

func (p Personnage) spiritAttack() {
	ClearLog()
	fmt.Println(BIWhite + ">> Combat d'entraînement <<" + Reset)
	fmt.Println("1 >> Tempête verte")
	fmt.Println("2 >> Boule de feu")
	fmt.Println("3 >> Retour")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		p.useStorm()
	case "2":
		p.useFireball("Esprit de la forêt")
	case "3":
		p.spiritTurn()
	default:
		p.spiritAttack()
	}
}

func (p Personnage) useStorm() {
	rand.Seed(time.Now().UnixNano())
	stormDmg := rand.Intn(11)
	if p.skill["Tempête verte"] >= 1 {
		fmt.Println(p.nom, "utilise", BIGreen+"Tempête verte"+Reset, "et inflige", stormDmg, "dégats à", M1.name)
		M1.HP -= stormDmg
		time.Sleep(400 * time.Millisecond)
		if M1.HP > 0 {
			stormDmg = rand.Intn(11)
			fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
			time.Sleep(2 * time.Second)
			fmt.Println("La", BIGreen+"Tempête verte"+Reset, "surgit à nouveau et fait", stormDmg, "dégats à", M1.name)
			M1.HP -= stormDmg
			if M1.HP > 0 {
				fmt.Println(M1.name, "a maintenant", M1.HP, "HP/", M1.maxHP, "HP")
			}
		}

	} else {
		fmt.Println(">> Impossible <<")
		p.spiritAttack()
	}
}
