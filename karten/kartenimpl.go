package karten

import (
		. "Spielkarten/gfx"
		. "Spielkarten/zufallszahlen"
		"strconv")
		
type data struct {
	wert uint8 			//Zahlenwert wird gespeichert 2,3,4,5,6,7,8,9,10,11 = B, 12 = D, 13 = K, 1 = A
	suit uint8			//0=Karo,1=Herz,2=Pik,3=Kreuz
	x,y,size uint16		//Position & Größe
	sr,sg,sb uint8		//Farbe für den Suit
	lr,lg,lb uint8		//Farbe für die Buchstaben (Letter)30,144,255
	hr,hg,hb uint8		//Farbe fürs Highlighting
	dr,dg,db uint8		//Farbe fürs Deck - Farbe der verdeckten Karte
	aufgedeckt bool 	//True wenn die Seite mit den Werten sichtbar ist
	highlighting bool	//True wenn Highlighting an ist
}

func New () *data {							//Erzeugt einen neue instanz des Datentyps Karte
	var k *data								//Erzeugt einen Zeiger auf den Datentyp Karte (oder data)
	k = new(data)							//Erzeugt den notwendigen Speicher der im struct spezifiziert wurde
	Randomisieren()							//Erzeugt einen neuen zufälligen Keim aus der Systemzeit in Millisekunden für die Berechnung der folgenden Zufallszahlen
	(*k).wert = uint8(Zufallszahl(1,13))	//Wählt zufällig eine Karte aus - Wert (A,2,3,4,5,6,7,8,9,10,B,D,K)
	(*k).suit = uint8(Zufallszahl(0,3))		//								- Suit (Karo,Herz,Pik,Kreuz)
	if (*k).suit<2 {
		(*k).sr,(*k).sg,(*k).sb = 255,0,0 	//Setz die Default-Farbe für die Suits 	- rot für Karo & Herz
		(*k).lr,(*k).lg,(*k).lb = 255,0,0
	} else {
		(*k).sr,(*k).sg,(*k).sb = 0,0,0		// 										- schwarz für Pik & Kreuz 
		(*k).lr,(*k).lg,(*k).lb = 0,0,0
	}
	(*k).dr,(*k).dg,(*k).db = 30,144,255 	//Setzt die Default-Farbe für das Deck 	- Farbe der Rückseite blau
	(*k).hr,(*k).hg,(*k).hb = 212,8,8		//Setzt die Default-Farbe fürs Highlighting - Farbe fürs Highlighting rot, etwas anders als die Suit-Farbe
	
//Die hier aufgelisteten Werte werden von Go automatisch gesetzt. Sie werden hier nur zum Verständnis aufgelistet.
	//(*k).aufgdeckt = false 		- bedeutet, alle Karten sind erstmal verdeckt, also nur die Rückseite zu sehen 
	//(*k).highlighting = false		- bedeutet, alle Karten haben erstmal kein Highlighting
	//(*k).x = 0					- Diese Werte machen erst Sinn, wenn sie in einem gfx-Fenster dargestellt wurden, daher werden sie auch von der Draw-Funktion gesetzt
	//(*k).y = 0					  Daher macht es keinen Sinn die Funktion "GehörtPunktzurKarte" aufzurufen, wenn die Karte nicht vorher durch die Draw-Funktion gezeichnet wurde.
	//(*k).size = 0
	
	return k
}

func (k *data) GibWert () string { 	// Wandelt den unter wert gespeicherten Zahlenwert in einen String um und gibt diesen zurück
	switch (*k).wert {				//fängt die Buchstaben ab (B,D,K,A)
		case 11:
			return "B"
		case 12:
			return "D"
		case 13:
			return "K"
		case 1:
			return "A"
		default:					//Wandelt die Zahlen von 2-10 in einen String um.
			var sWert string
			sWert = strconv.Itoa(int((*k).wert))
			return sWert
	}
}

func (k *data) GibSuit () string {	// Wandelt den unter suit gespeicherten Zahlenwert in einen String um und gibt diesen zurück.
	var erg string
	switch (*k).suit {
		case 0:
			erg = "Karo"
		case 1:
			erg = "Herz"
		case 2:
			erg = "Pik"
		case 3:
			erg = "Kreuz"
	}
	return erg
}

func (k *data) GibWertFarbe () (r,g,b uint8) {	//Gibt die Farbe in der der Wert in der Draw-Funktion dargestellt wird in r,g,b zurück.
	return (*k).lr,(*k).lg,(*k).lb
}

func (k *data) GibSuitFarbe () (r,g,b uint8) {	//Gibt die Farbe in der der Suit in der Draw-Funktion dargestellt wird in r,g,b zurück.
	return (*k).sr,(*k).sg,(*k).sb
}

func (k *data) GibHighlightFarbe() (r,g,b uint8) {	//Gibt die Farbe des Highlighting, das in der Draw-Funktion dargestellt wird in r,g,b zurück.
	return (*k).hr,(*k).hg,(*k).hb
}

func (k *data) GibSize () (size uint16) { //Gibt die gespeicherte Größe zurück
	return (*k).size
}

func (k *data) GibKoordinaten () (x,y uint16) { //Gibt die gespeicherten Koordinaten der Karte zurück
	return (*k).x,(*k).y
}

func (k *data) SetzeHighlight (highlight bool) {	//Setzt das Highlighting auf den übergebenen Wert (true: Highlighting an; false: Highlighting aus
	(*k).highlighting = highlight
}

func (k *data) SetzeKarte (wert, suit string) (verändert bool) {	//Verändert die Karte
	var vwert,vsuit bool
	switch suit {
		case "Karo":
			(*k).suit = 0
			(*k).lr,(*k).lg,(*k).lb = 255,0,0
			(*k).sr,(*k).sg,(*k).sb = 255,0,0
			vsuit = true
		case "Herz":
			(*k).suit = 1
			(*k).lr,(*k).lg,(*k).lb = 255,0,0
			(*k).sr,(*k).sg,(*k).sb = 255,0,0
			vsuit = true
		case "Pik":
			(*k).suit = 2
			(*k).lr,(*k).lg,(*k).lb = 0,0,0
			(*k).sr,(*k).sg,(*k).sb = 0,0,0
			vsuit = true
		case "Kreuz":
			(*k).suit = 3
			(*k).lr,(*k).lg,(*k).lb = 0,0,0
			(*k).sr,(*k).sg,(*k).sb = 0,0,0
			vsuit = true
	}
	switch wert {
		case "A":
			(*k).wert = 1
			vwert = true
		case "K":
			(*k).wert = 13
			vwert = true
		case "D":
			(*k).wert = 12
			vwert = true
		case "B":
			(*k).wert = 11
			vwert = true
		default:
			erg,err:=strconv.Atoi(wert)
			if err == nil && erg>=2 && erg<=10 {
				(*k).wert = uint8(erg)
				vwert = true
			}
	}
	return vwert && vsuit
}

func (k *data) Umdrehen () {
	if (*k).aufgedeckt {
		(*k).aufgedeckt = false
	} else {
		(*k).aufgedeckt = true
	}
}

func (k *data) Aufdecken () {
	(*k).aufgedeckt = true
}

func (k *data) Zudecken () {
	(*k).aufgedeckt = false
}

func (k *data) String () string {
	var erg string
	if (*k).aufgedeckt {
		if (*k).wert != 10 {
			erg = erg + " "
		}
		erg = erg + k.GibWert()
		erg = erg + " "
		erg = erg + k.GibSuit()
		if (*k).suit!=3 {
			switch (*k).suit {
				case 2:
					erg = erg + "  "
				default:
					erg = erg + " "
			}
		}
	} else {
		erg = "********"
	}
	return erg
}

func kreuz (x,y,size uint16) {
	Vollrechteck(x+size/3,y,size/3,size)
	Vollrechteck(x,y+size/3,size,size/3)
}
	
func karo (x,y,size uint16) {
	Volldreieck(x+size/2,y,x+size/12,y+size/2,x+size*9/12,y+size/2)
	Volldreieck(x+size/2,y+size,x+size*3/12,y+size/2,x+size*11/12,y+size/2)
}

func herz (x,y,size uint16) {
	Volldreieck(x+size/3,y,x+size/12,y+size/3,x+size/2,y+size/3)
	Volldreieck(x+size*2/3,y,x+size*11/12,y+size/3,x+size/2,y+size/3)
	Volldreieck(x+size/12,y+size/3,x+size/2,y+size,x+size/2,y+size/3)
	Volldreieck(x+size*11/12,y+size/3,x+size/2,y+size,x+size/2,y+size/3)
}

func pik (x,y,size uint16) {
	Volldreieck(x+size/6,y,x+size/6,y+size*7/12,x+size*5/6,y+size/2-size/12)
	Volldreieck(x+size/6,y+size*5/12,x+size/6,y+size,x+size*5/6,y+size-size/12)
}

func (k *data) GehörtPunktzurKarte (xp,yp uint16) bool {
	var xk,yk,size uint16
	xk = (*k).x
	yk = (*k).y
	size = (*k).size
	//Vollrechteck(x-1,y-1,2*size+2,3*size+2)
	if xp>=xk-1 && xp<=xk-1+2*size+2 && yp>=yk-1 && yp<=yk-1+3*size+2 {
		return true
	} 
	//Vollrechteck(x-size/10-1,y-1,size/10+2,3*size+2)
	if xp>=xk-size/10-1 && xp<=xk-size/10-1+size/10+2 && yp>=yk-1 && yp<=yk-1+3*size+2{
		return true
	}
	//Vollrechteck(x-1,y-size/10-1,2*size+2,size/10+2)
	if xp>=xk-1 && xp<=xk-1+2*size+2 && yp>=yk-size/10-1 && yp<=yk-size/10-1+size/10+2 {
		return true
	}
	//Vollrechteck(x+size*2-1,y,size/10+1+2,3*size+2)
	if xp>=xk+size*2-1 && xp<=xk+size*2-1+size/10+1+2 && yp>=yk && yp<=yk+3*size+2 {
		return true
	}
	//Vollrechteck(x-1,y+3*size-1,size*2+2,size/10+1+2)
	if xp>=xk-1 && xp<=xk-1+size*2+2 && yp>=yk+3*size-1 && yp<=yk+3*size-1+size/10+1+2 {
		return true
	}
	//Vollkreis(x-1,y-1,size/10)
	if (int(xp)-int(xk-1))*(int(xp)-int(xk-1))+(int(yp)-int(yk-1))*(int(yp)-int(yk-1)) <= int(size/10)*int(size/10) {
		return true
	}
	//Vollkreis(x-1,y+3*size+1,size/10)
	if (int(xp)-int(xk-1))*(int(xp)-int(xk-1))+(int(yp)-int(yk+3*size+1))*(int(yp)-int(yk+3*size+1)) <= int(size/10)*int(size/10) {
		return true
	}
	//Vollkreis(x+size*2-1+2,y-1,size/10)
	if (int(xp)-int(xk+size*2-1+2))*(int(xp)-int(xk+size*2-1+2))+(int(yp)-int(yk-1))*(int(yp)-int(yk-1)) <= int(size/10)*int(size/10) {
		return true
	}
	//Vollkreis(x+size*2-1+2,y+3*size+1,size/10)
	if (int(xp)-int(xk+size*2-1+2))*(int(xp)-int(xk+size*2-1+2))+(int(yp)-int(yk+3*size+1))*(int(yp)-int(yk+3*size+1)) <= int(size/10)*int(size/10) {
		return true
	}
	return false
}

func (k *data) Draw (x,y,size uint16) {
	(*k).x = x
	(*k).y = y
	(*k).size = size
	if (*k).highlighting {
		var hr,hg,hb uint8
		hr,hg,hb = k.GibHighlightFarbe()
		Stiftfarbe(hr,hg,hb)
		Vollrechteck(x-3,y-3,2*size+6,3*size+6)
		Vollrechteck(x-size/10-3,y-3,size/10+6,3*size+6)
		Vollrechteck(x-3,y-size/10-3,2*size+6,size/10+6)
		Vollrechteck(x+size*2-3,y,size/10+1+6,3*size+6)
		Vollrechteck(x-3,y+3*size-3,size*2+6,size/10+1+6)
		Vollkreis(x-3,y-3,size/10)
		Vollkreis(x-3,y+3*size+3,size/10)
		Vollkreis(x+size*2-3+6,y-3,size/10)
		Vollkreis(x+size*2-3+6,y+3*size+3,size/10)
	}
	if (*k).aufgedeckt {
	
		//Körper
		Stiftfarbe(0,0,0)
		Vollrechteck(x-1,y-1,2*size+2,3*size+2)
		Vollrechteck(x-size/10-1,y-1,size/10+2,3*size+2)
		Vollrechteck(x-1,y-size/10-1,2*size+2,size/10+2)
		Vollrechteck(x+size*2-1,y,size/10+1+2,3*size+2)
		Vollrechteck(x-1,y+3*size-1,size*2+2,size/10+1+2)
		Vollkreis(x-1,y-1,size/10)
		Vollkreis(x-1,y+3*size+1,size/10)
		Vollkreis(x+size*2-1+2,y-1,size/10)
		Vollkreis(x+size*2-1+2,y+3*size+1,size/10)
		
		
		Stiftfarbe(255,255,255)
		Vollrechteck(x,y,2*size,3*size)
		Vollrechteck(x-size/10,y,size/10,3*size)
		Vollrechteck(x,y-size/10,2*size,size/10)
		Vollrechteck(x+size*2,y,size/10+1,3*size)
		Vollrechteck(x,y+3*size,size*2,size/10+1)
		Vollkreis(x,y,size/10)
		Vollkreis(x,y+3*size,size/10)
		Vollkreis(x+size*2,y,size/10)
		Vollkreis(x+size*2,y+3*size,size/10)
		
		//Wert
		var br,bg,bb uint8
		br,bg,bb = k.GibWertFarbe()
		Stiftfarbe(br,bg,bb)
		SetzeFont ("./gfx/fonts/Lato-Black.ttf", int(size/3))
		w:=k.GibWert()
		
		if (*k).wert==10 {						//Zweistellige Zahlen (10)	
			SchreibeFont (x,y,w)
			SchreibeFont (x+2*size-size/3-size*7/100,y+3*size-size/3-size*9/100,w)
			SetzeFont ("./gfx/fonts/Lato-Black.ttf", int(size))
			SchreibeFont (x+size/2-size*11/100,y+size/2+size*36/100,w)	
		} else if (*k).wert>10 || (*k).wert==1{ //Buchstaben
			SchreibeFont (x+size*9/100,y,w)
			SchreibeFont (x+2*size-size/3-size*1/100,y+3*size-size/3-size*9/100,w)
			SetzeFont ("./gfx/fonts/Lato-Black.ttf", int(size))
			SchreibeFont (x+size/2+size*14/100,y+size/2+size*36/100,w)
		} else {								//Einstellige Zahlen
			SchreibeFont (x+size*12/100,y,w)
			SchreibeFont (x+2*size-size/3+size*2/100,y+3*size-size/3-size*9/100,w)
			SetzeFont ("./gfx/fonts/Lato-Black.ttf", int(size))
			SchreibeFont (x+size/2+size*25/100,y+size/2+size*36/100,w)
		}
		
		//Suit
		var sr,sg,sb uint8
		sr,sg,sb = k.GibSuitFarbe()
		Stiftfarbe(sr,sg,sb)
		switch k.GibSuit() {
			case "Karo":
				karo(x+size*9/100,y+size/3+size*2/100,size/4)
				karo(x-size*17/100-size/3+2*size+size/6,y-size*7/12+3*size-size*3/100,size/4)
				karo(x+size-size*24/100,y+size/3+size*20/100,size/2)
				karo(x+size-size*24/100,y+size*5/3+size*29/100,size/2)
			case "Herz":
				herz(x+size*9/100,y+size/3+size*2/100,size/4)
				herz(x-size*17/100-size/3+2*size+size/6,y-size*7/12+3*size-size*3/100,size/4)
				herz(x+size-size*24/100,y+size/3+size*20/100,size/2)
				herz(x+size-size*24/100,y+size*5/3+size*29/100,size/2)
			case "Pik":
				pik(x+size*9/100,y+size/3+size*2/100,size/4)
				pik(x-size*17/100-size/3+2*size+size/6,y-size*7/12+3*size-size*3/100,size/4)
				pik(x+size-size*24/100,y+size/3+size*20/100,size/2)
				pik(x+size-size*24/100,y+size*5/3+size*29/100,size/2)
			case "Kreuz":
				kreuz(x+size*9/100,y+size/3+size*2/100,size/4)
				kreuz(x-size*17/100-size/3+2*size+size/6,y-size*7/12+3*size-size*3/100,size/4)
				kreuz(x+size-size*24/100,y+size/3+size*20/100,size/2)
				kreuz(x+size-size*24/100,y+size*5/3+size*29/100,size/2)
		}
	} else {
		//Körper
		//Rand
		Stiftfarbe(0,0,0)
		Vollrechteck(x-1,y-1,2*size+2,3*size+2)
		Vollrechteck(x-size/10-1,y-1,size/10+2,3*size+2)
		Vollrechteck(x-1,y-size/10-1,2*size+2,size/10+2)
		Vollrechteck(x+size*2-1,y,size/10+1+2,3*size+2)
		Vollrechteck(x-1,y+3*size-1,size*2+2,size/10+1+2)
		Vollkreis(x-1,y-1,size/10)
		Vollkreis(x-1,y+3*size+1,size/10)
		Vollkreis(x+size*2-1+2,y-1,size/10)
		Vollkreis(x+size*2-1+2,y+3*size+1,size/10)
		
		//Fläche
		Stiftfarbe(255,255,255)
		Vollrechteck(x,y,2*size,3*size)
		Vollrechteck(x-size/10,y,size/10,3*size)
		Vollrechteck(x,y-size/10,2*size,size/10)
		Vollrechteck(x+size*2,y,size/10+1,3*size)
		Vollrechteck(x,y+3*size,size*2,size/10+1)
		Vollkreis(x,y,size/10)
		Vollkreis(x,y+3*size,size/10)
		Vollkreis(x+size*2,y,size/10)
		Vollkreis(x+size*2,y+3*size,size/10)
		
		//Muster
		var dr,dg,db uint8
		dr,dg,db = (*k).dr,(*k).dg,(*k).db
		Stiftfarbe(dr,dg,db)
		Vollrechteck(x-size/10,y+size*2/12+size*1/100,size*2+2*size/10+1,size/3)
		Vollrechteck(x-size/10,y+size*9/12+size*1/100,size*2+2*size/10+1,size/3)
		Vollrechteck(x-size/10,y+size*16/12+size*1/100,size*2+2*size/10+1,size/3)
		Vollrechteck(x-size/10,y+size*23/12+size*1/100,size*2+2*size/10+1,size/3)
		Vollrechteck(x-size/10,y+size*30/12+size*1/100,size*2+2*size/10+1,size/3)
	}
}
