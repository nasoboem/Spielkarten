package main

import (
		. "Spielkarten/gfx" //der Punkt sorgt dafür, dass bei Aufrufen von Funktionen aus dem gfx-Paket nicht immer gfx.Funktionsname geschrieben werden muss
		"Spielkarten/karten"
		)


func main () {
	var x,y,size uint16 											//Initialisierung der Variablen für Koordinaten(x,y) und Größe der Spielkarten
	x,y,size = 50,50,52 											//Belegung der Koordinate des Starpunktes an dem die erste Karte gezeichnet werden soll + Größe 
	var deck []karten.Karte											//Initialisierung der Variable "deck", die gleich ein 52er Kartendeck entgegen nimmt
	deck = karten.Deck52()											//Es wird ein sortirtes Kartendeck mit 52 Karten generiert und an die Variable "deck" übergeben
	Fenster(1200,800)												//Ein gfx-Fenster mit der Breite 1200 & Höhe 800 wird geöffnet
	Stiftfarbe(47,186,51)											//Die Stiftfarbe wird auf grün gesetzt
	Vollrechteck(0,0,1200,800)										//Zeichne ein grünes Rechteck in der größe des Fensters
	//MausLesen1 () (taste uint8, status int8, mausX, mausY uint16)	//Überreste: Wenn man auf Karten klicken können möchte braucht man var taste & status +
	//var taste uint8												//Funktionsrumpf aus dem gfx-Paket, da ich mir die reihenfolge der Variablen nicht merken kann
	//var status int8
	var mx,my uint16												//Erzeugt zwei variablen, die die Koordinaten des Mauszeigers aufnehmen
	deck = karten.Mischen(deck)										//Das Deck wird gemischt
	
	UpdateAus()														//Zeichne im Hintergrund nicht sichtbar
	Stiftfarbe(255,255,255)											//Setze die Stiftfarbe auf weiß
	Cls()															//Mach das ganze Fenster weiß
	Stiftfarbe(47,186,51)											//Setze die Stiftfarbe auf grün
	Vollrechteck(0,0,1200,800)										//Zeichne ein grünes Rechteck in der größe des Fensters
	for i:=0;i<len(deck);i++{										//Zähle durch die Karten des Decks durch, und mache folgendes
		deck[i].Draw(x+uint16(i)%13*size,y+uint16(i)/13*size*10/3,size)	//deck[i] wählt die aktuelle Karte aus über den Index;
																		//Draw zeichnet die Karte
																		//x+uint16(i)%13*size - sorgt dafür, dass die Karten in x Richtung mit jedem i um 50px (size - ca. halbe Kartenbreite) verschoben werden.
																		//das passiert aber immer nur dann, wenn die Zahl mit Rest durch 13 teilbar ist, ist sie durch 13 glatt teilbar wird wieder bei 0 angefangen, das macht der Modolo-Operator %13
																		//Effekt: man bekommt Zeilen mit 13 Karten darin
																		//y+uint16(i)/13*size*10/3 - sorgt dafür, dass die Karten um 3 x 50px (volle Karten höhe) in y Richtung verschoben werden, aber nur dann, wenn die i sich durch 13 Teilen lässt (i:0-12 = 0; 13-25 = 1; 26-38 = 2; 39-52 = 3).
																		//Effekt: alle 13 Karten wird eine neue Zeile begonnen
																		//Gesamteffekt ist, dass 4 zeilen a 13 Karten auf dem Bildschirm gezeichnet werden in der Größe 50
	}
	UpdateAn()					
	
	for {													//Endlosschleife Anfang
		_,_,mx,my = MausLesen1()									//Koordinate des Mauszeigers werden ausgelesen und an die Variablen mx & my übergeben; _,_,: verwirft die Werte für taste & status, die nicht gebraucht werden
		treffer,index:=karten.GetTopSelected(deck,mx,my)			//Die Variablen treffer (ob der Mauszeiger auf einer Karte liegt) & index (index der obersten getroffenen Karte im Deck)
		if treffer {												//Wenn der Mauszeiger auf einer Karte liegt, dann...
			deck[index].Aufdecken()									//decke die oberste Karte auf der der Mauszeiger liegt auf,
			for i:=0;i<len(deck);i++{								//alle anderen Karten
				if i!=index{										//bis auf die Karte die du aufgedeckt hast,
					deck[i].Zudecken()								//deckst du zu.
				}
			}
		} else {													//Wenn der Mauszeiger nicht auf einer Karte liegt, dann ...
			for i:=0;i<len(deck);i++{								//decke alle Karte zu.
					deck[i].Zudecken()
			}
		}
		
	
	UpdateAus()														//Zeichne im Hintergrund nicht sichtbar
	Stiftfarbe(255,255,255)											//Setze die Stiftfarbe auf weiß
	Cls()															//Mach das ganze Fenster weiß
	Stiftfarbe(47,186,51)											//Setze die Stiftfarbe auf grün
	Vollrechteck(0,0,1200,800)										//Zeichne ein grünes Rechteck in der größe des Fensters
	for i:=0;i<len(deck);i++{										//Zähle durch die Karten des Decks durch, und mache folgendes
		deck[i].Draw(x+uint16(i)%13*size,y+uint16(i)/13*size*10/3,size)	//deck[i] wählt die aktuelle Karte aus über den Index;
																		//Draw zeichnet die Karte
																		//x+uint16(i)%13*size - sorgt dafür, dass die Karten in x Richtung mit jedem i um 50px (size - ca. halbe Kartenbreite) verschoben werden.
																		//das passiert aber immer nur dann, wenn die Zahl mit Rest durch 13 teilbar ist, ist sie durch 13 glatt teilbar wird wieder bei 0 angefangen, das macht der Modolo-Operator %13
																		//Effekt: man bekommt Zeilen mit 13 Karten darin
																		//y+uint16(i)/13*size*10/3 - sorgt dafür, dass die Karten um 3 x 50px (volle Karten höhe) in y Richtung verschoben werden, aber nur dann, wenn die i sich durch 13 Teilen lässt (i:0-12 = 0; 13-25 = 1; 26-38 = 2; 39-52 = 3).
																		//Effekt: alle 13 Karten wird eine neue Zeile begonnen
																		//Gesamteffekt ist, dass 4 zeilen a 13 Karten auf dem Bildschirm gezeichnet werden in der Größe 50
	}
	UpdateAn()														//Alles was im Hintergrund gezeichnet wurde wird auf einmal sichtbar - verhindert Flackern in den Bildern
	}													//Endlosschleife Ende
}
	
