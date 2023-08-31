package karten

import (. "Spielkarten/zufallszahlen")

type Karte interface {
	//Vor.: -
	//Erg.: 
	//New() erzeugt eine zufällige Instance einer Spielkarte (52er-Kartendeck) von Karo, Herz, Pik, Kreuz und den Werten von 2,3,4,5,6,7,8,9,10,B,D,K,A
	//daher ist es möglich, wenn man einen Ziehstapel machen möchte, einfach eine neu Instanz der Spielkarte zu erzeugen und so ein imprinzip endlosen Ziehstapel zu haben.
	//Daher sind Kartendoppelungen natürlich möglich. Wer keine Kartendoppelungen möchte, oder die Spielkarten eines 32er-Kartendeck für sein Spiel braucht, findet weiter
	//unten zwei Funktionen (Deck52 und Deck32) welche eine sortierte Instanz dieser Decks erzeugen. 
	
	//Vor.: -
	//Erg.: Der Wert einer Karte (2,3,4,5,6,7,8,9,10,B,D,K,A) ist zurückgegeben.
	GibWert () string
	
	//Vor.: -
	//Erg.: Die Farbe (Suit - Karo, Herz, Pik, Kreuz) der Karte ist zurückgegeben. 
	GibSuit () string
	
	//Vor.: -
	//Erg.: Die gespeicherte Größe (Size) ist zurückgegeben. - Sinnvolles Ergebnis gibt es erst, nach dem die Karte gezeichnet wurde - s. Draw(x,y,size uint16)
	GibSize () (size uint16)

	//Vor.: -
	//Erg.: Die gespeicherten Koordinaten in x und y sind zurückgegeben. - Sinnvolle Ergebnise gibt es erst, nach dem die Karte gezeichnet wurde - s. Draw(x,y,size uint16)
	GibKoordinaten () (x,y uint16)
	
	//Vor.: -
	//Erg.: Die Farbe in der Wert der Karte gezeichnet wird ist in r, g, b zurückgegeben.
	GibWertFarbe () (r,g,b uint8)
	
	//Vor.: -
	//Erg.: Die Farbe in der die Farbe (Suit - Karo, Herz, Pik, Kreuz) der Karte gezeichnet wird ist in r, g, b zurückgegeben.
	GibSuitFarbe () (r,g,b uint8)
	
	//Vor.: -
	//Erg.: Die Farbe in der die Highlightumrandung der Karte gezeichnet wird ist in r, g, b zurückgegeben.
	GibHighlightFarbe() (r,g,b uint8)
	
	//Vor.: -
	//Eff.: Die Karte besitzt jetzt den gegeben Wert (2,3,4,5,6,7,8,9,10,B,D,K,A) und die entsprechende Farbe (Suit - Karo,Herz,Pik,Kreuz).
	//		Diese Funktion setzt auch die Default Farben für den Wert & den Suit der Karte. Möchte man andere Farben haben, muss man sie anschließend.
	//		wieder ändern. Die Default-Farben sind: 255,0,0 für Wert & Suit bei Karo & Herz; 0,0,0 für Wert & Suit bei Pik & Kreuz.
	//		Alle anderen Farben (Highlighting & Deckfarbe (Farbe der Rückseite) werden nicht verändert.
	//Erg.: True ist zurück gegeben, genau dann wenn beide Werte angewendet wurden. Andernfalls ist ein false zurückgegeben.
	//Bei Schreibfehler in einem der beiden übergebenen Strings kommt es zu keiner Änderung. Bitte der Notierung unter Eff. folgen.  
	SetzeKarte (wert, suit string) (verändert bool)
	
	//Vor.: -
	//Eff.: Das Highlighting wird auf den übergeben boolschen Wert gesetzt. true - Highlighting: an; false - Highlighting: aus
	SetzeHighlight(highlight bool)
	
	//Vor.: -
	//Erg.: Die Karte ist als string zurückgegeben.
	//		Ist die Karte aufgedeckt in der Formatierung: Wert (2 Stellen,rechtsbünding) Leerzeichen Farbe (suit) (5 Stellen, linksbündig) Bsp.: "10 Kreuz" & " A Pik  "
	//		Ist die Karte verdeckt als 8 Sternchen. "********"
	String () string
	
	//Vor.: -
	//Eff.: Die Karte wird umgedreht. War die Karte verdeckt (Rückseite zu sehen) wird sie jetzt aufgedeckt (Wert & Farbe sind zu sehen) dargestellt (Draw/String-Funktion).
	//		War die Karte aufgedeckt wird sie jetzt verdeckt dargestellt.
	Umdrehen ()
	
	//Vor.: -
	//Eff.: Die Karte ist aufgedeckt. War die Karte verdeckt ist sie jetzt aufgedeckt (Wert & Farbe sind zu sehen), war sie aufgedeckt bleibt sie aufgedeckt.
	Aufdecken ()
	
	//Vor.: -
	//Eff.: Die Karte ist zugedeckt. War die Karte aufgedeckt ist sie jetzt zugedeckt (nur die Rückseite ist zusehen, Wert & Farbe (Suit) nicht), war sie zugedeckt bleibt sie zugedeckt.
	Zudecken ()
	
	//Vor.: - ,aber eine sinnvolle Verwendung dieser Funkion ist nur möglich, wenn die Karte vorher gezeichnet wurde, weil dadurch die Koordinaten (x & y) & die Größe (size)
	// 		in der Karte gespeichert werden. Noch nicht gezeichnete Karten besitzen die Werte x,y,size = 0,0,0. Weswegen es hier immer nur der Punkt (0|0) zu einem Treffer führt.
	//Erg.: True ist geliefert, genau dann wenn der Punkt zur Karte gehört. (Highlighting gehört nicht zur Karte.) Andernfalls ist false geliefert.
	//		Diese Funktion ermöglicht es die Karten per Maus zu steuern und zu bewegen.
	GehörtPunktzurKarte (xp,yp uint16) bool
	
	//Vor.: Ein gfx-Fenster ist geöffnet und die Koordinaten x & y befinden sich innerhalb des Fensters. Wenn die Karte nicht vollständig ins Fenster passt,
	//		kommt es zu Teildarstellungen. 
	//Eff.: Die Spielkarte ist an der Position gegeben durch Koordinaten x & y (~ linke obere Ecke - 1/10size) und in der Größe Breite: 2xsize+2/10*size Höhe: 3xsize+2/10*size.
	//		im gfx-Fenster dargestellt.
	//		Die gegebenen Koordinaten (x,y) und Größe (size) werden in der Karte gespeichert. Noch nicht gezeichnete Karten besitzen die Werte x,y,size = 0,0,0.
	//		Je nach dem ob die Karte aufgedeckt ist (Wert & Farbe sind zu sehen) oder ob die Karte verdeckt ist (nur die Rückseite der Karte ist zu sehen) wird die Karte gezeichnet.
	Draw (x,y,size uint16)
}


//Vor.: -
//Erg.: Ein slice mit 52 Karten ist geliefert in der sortierung 2,3,4,5,6,7,8,9,10,B,D,K,A in Karo,Herz,Pik,Kreuz. 
func Deck52 () []Karte {
var deck []Karte																//Eine Variable deck wird erzeugt vom Typ slice von Karten = leerer Slice []
	werte:=[13]string{"2","3","4","5","6","7","8","9","10","B","D","K","A"}		//Eine Variable werte wird erzeugt vom Typ Liste der Länge 13 von string und beladen mit "2","3","4","5","6","7","8","9","10","B","D","K","A"
 	suits:=[4]string{"Karo","Herz","Pik","Kreuz"}								//Eine Variable suits wird erzeugt vom Typ Liste der Länge 4 von string und beladen mit "Karo","Herz","Pik","Kreuz"
	for i:=0;i<len(suits);i++{													//äußere Schleife, geht durch die Liste der suits durch
		for j:=0;j<len(werte);j++{												//innere Schleife, geht durch die Liste der werte durch
			var k Karte															//Eine Variable k vom Typ Karte wird initialisiert, also nur ein ins leere zeigender Zeiger wird generiert
			k = New()															//Jetzt wird der Speicher auf den der Zeiger k zeigt mit Inhalten belegt (zufällig)
			k.SetzeKarte(werte[j],suits[i])										//Die Karte wird mit den gewünschten Werten gegeben durch die Listen werte und suits überschrieben
			deck = append(deck,k)												//Die Karte wird hinten an das Deck angefügt.
		}
	}
	return deck																	//Das fertig zusammen gestellte deck wird zurückgegeben
}

//Vor.: -
//Erg.: Ein slice mit 32 Karten ist geliefert in der sortierung 7,8,9,10,B,D,K,A in Karo,Herz,Pik,Kreuz.
func Deck32 () []Karte {														//Aufbau wie bei Deck52 nur jetzt mit 32 Karten
	var deck []Karte
	werte:=[8]string{"7","8","9","10","B","D","K","A"}
	suits:=[4]string{"Karo","Herz","Pik","Kreuz"}
	for i:=0;i<len(suits);i++{
		for j:=0;j<len(werte);j++{
			var k Karte
			k = New()
			k.SetzeKarte(werte[j],suits[i])
			deck = append(deck,k)
		}
	}
	return deck
}

//Vor.: -
//Erg.: Ein Slice mit 104 Karten ist geliefert - 52er Deck mit allen Karten doppelt für Memory - Slice auswählen umkopieren und erst anschließend mischen.
func MemoryDeck () []Karte {
	var deck []Karte
	werte:=[13]string{"A","K","D","B","10","2","3","4","5","6","7","8","9"}		//Eine Variable werte wird erzeugt vom Typ Liste der Länge 13 von string und beladen mit "2","3","4","5","6","7","8","9","10","B","D","K","A"
 	suits:=[4]string{"Karo","Herz","Pik","Kreuz"}								//Eine Variable suits wird erzeugt vom Typ Liste der Länge 4 von string und beladen mit "Karo","Herz","Pik","Kreuz"
	for i:=0;i<len(werte);i++{													//äußere Schleife, geht durch die Liste der suits durch
		for j:=0;j<len(suits);j++{												//innere Schleife, geht durch die Liste der werte durch
			for l:=0;l<2;l++{													//Für Memory benötigt man Karten doppelt
				var k Karte															//Eine Variable k vom Typ Karte wird initialisiert, also nur ein ins leere zeigender Zeiger wird generiert
				k = New()															//Jetzt wird der Speicher auf den der Zeiger k zeigt mit Inhalten belegt (zufällig)
				k.SetzeKarte(werte[i],suits[j])										//Die Karte wird mit den gewünschten Werten gegeben durch die Listen werte und suits überschrieben
				deck = append(deck,k)												//Die Karte wird hinten an das Deck angefügt.
			}
		}
	}
	return deck																	//Das fertig zusammen gestellte deck wird zurückgegeben
}

func IstGleich (k1,k2 Karte) bool {
	return k1.GibWert()==k2.GibWert() && k1.GibSuit()==k2.GibSuit()
}

//Vor.: -
//Erg.: Ein neues gemischtes Deck mit den vorher enthaltenen Karten ist zurückgegeben
func Mischen (deck []Karte) []Karte { 							//Die Funktion ist recht einfach, ziehe zufällig eine Karte aus dem bisherigen Deck und füge sie zufällig vorne oder hinten an des neue gemischte Deck an. Wiederhole, mit allen Karten des Decks.
	Randomisieren()												//Erstelle einen zufälligen Keim aus der Systemzeit für die Zufallszahl
	var ndeck []Karte											//generiere ein neues Deck, in das die Karten zufällig hinein gesteckt werden
	var vorne int64												//erzeuge eine Variable vorne (wenn 0, dann wird die Karte hinen eingefügt, wenn 1 dann wird die Karte vorne eingefügt) - eigentlich unnötig 
	for 0<len(deck) {											//Schleife, wird solange ausgeführt, wie noch Karten im deck sind
		var k Karte												//Variable erzeugen, die die Karte aufnehmen kann
		var zahl int64											//Variable, die den zufällig gewählten Index aufnehmen kann
		zahl = Zufallszahl(0,int64(len(deck)-1))				//Hier wird ein zufälliger Index erzeugt von 0-len(deck)-1, also über alle Indices 
		k = deck[int(zahl)]										//dann wird die Karte aus dem Deck an die Variable k übergeben
		deck = append(deck[:int(zahl)],deck[int(zahl)+1:]...)	//dann wird die Karte aus dem Deck entfernt - es muss eine Typumwandlung von int64 in int erfolgen
																//											- der Befehl deck [:int(zahl)], wählt alle zahlen von 0 bis zum index aus, lässt aber genau diesen aus
																//											- deck [in(zahl)+1:] nimmt alle Indices des decks, ab zahl+1 mit
																//											- durch den Befehl append(teildeck1,teildeck2...) werden die beiden slices aneinander gehangen. Dies wird durch die drei Punkte (...) erreicht
																//Der Effekt ist, dass man genau die Karte aus dem deck schmeißt, die man zufällig vorher gezogen hat
		vorne = Zufallszahl(0,1)								//Hier wird zufällig gewählt, ob man die Karte vorne (1) oder hinten (0) an das bisherige neue deck anfügt
		if vorne==0 {											//Hinten - anfügen mit append
			ndeck = append(ndeck,k)
		} else {
			var nndeck []Karte									//Vorne - da es kein eigenständigen Befehl für vorne gibt, muss erst ein weiteres neues nndeck geschaffen werden
			nndeck = append(nndeck,k)							//in diese leere neue nndeck wird die Karte gelegt
			ndeck = append(nndeck,ndeck...)						//Anschließend wird das restliche gemischte ndeck hintendran gepackt und gesagt, jetzt sei doch bitte das ndeck
		}
	}
	return ndeck												//am Ende gibt man das gemischte ndeck zurück.
}

//Vor.: Das verwendete ist durch die Draw-Funktion so dargestellt, dass kleinere Indices immer unter größeren Indices gezeichnet sind.
//Erg.: True ist geliefert, genau dann wenn der übergeben Punkt zu min. einer Karte im deck gehört (treffer). False ist geliefert, wenn der übergebene Punkt zu keiner Karte gehört.
//		index ist = 0 wenn treffer = false ist. Daher ist es nur sinnvoll den index auszuwerten, wenn treffer = true ist. Wenn treffer = true, dann liefert index den Index der 
//		obersten Karte im Deck, oder anders ausgedrückt, es wird der index zurückgegeben, der Karte, auf dem der Mauszeiger direkt liegt, nicht aller Karten, sollten sie sich in dem 
//		übergebenen Punkt überlappen. So ist es leichter, die oberste Karte aus einem Kartenstapel mit der Maus auszuwählen. 
func GetTopSelected (deck []Karte, x,y uint16) (treffer bool, index int) { 	//zeigt an, ob eine Karte durch den Punkt getroffen wurde, und gibt den index der obersten Karte aus.
	for i:=0;i<len(deck);i++ {												//geht das deck durch (von untern nach oben! - s. Vorraussetzung)
		if deck[i].GehörtPunktzurKarte (x,y) {								//wenn der Punkt zu einer Karte im Deck gehört, dann
			index = i														//merke dir den Index
			if !treffer {													//setze treffer auf true, wenn es noch false ist
				treffer = true
			}	
		}
	}																		//Da nur der index gemerkt wird wenn ein Treffer gelandet wurde und der index immer mit dem letzten treffer
																			//überschrieben wird, bekommt man am Ende die oberste Karte, zu der der Punkt gehört. 
	return																	//Werte werden zurückgegeben.
}
