package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type SakerOversikt struct {
	ResponsDatoTid string `xml:"respons_dato_tid"`
	Versjon        string `xml:"versjon"`
	SakerListe     []Sak  `xml:"saker_liste>sak"`
}

type Sak struct {
	ResponsDatoTid    string         `xml:"respons_dato_tid"`
	Versjon           string         `xml:"versjon"`
	BehandletSesjonId string         `xml:"behandlet_sesjon_id"`
	Dokumentgruppe    string         `xml:"dokumentgruppe"`
	EmneListe         []Emne         `xml:"emne_liste>emne"`
	Henvisning        string         `xml:"henvisning"`
	Id                int            `xml:"id"`
	InnstillingId     int            `xml:"innstilling_id"`
	InnstillingKode   string         `xml:"innstilling_kode"`
	Komitee           Komitee        `xml:"komitee"`
	Korttittel        string         `xml:"korttittel"`
	SakFremmetId      int            `xml:"sak_fremmet_id"`
	SaksordførerListe []Representant `xml:"saksordfoerer_liste>representant"`
	SistOppdatertDato string         `xml:"sist_oppdatert_dato"`
	Status            string         `xml:"status"`
	Tittel            string         `xml:"tittel"`
	Type              string         `xml:"type"`
}

type Emne struct {
	ResponsDatoTid string `xml:"respons_dato_tid"`
	Versjon        string `xml:"versjon"`
	ErHovedemne    bool   `xml:"er_hovedemne"`
	HovedemneId    int    `xml:"hovedemne_id"`
	Id             int    `xml:"id"`
	Navn           string `xml:"navn"`
}

type Komitee struct {
	ResponsDatoTid string `xml:"respons_dato_tid"`
	Versjon        string `xml:"versjon"`
	Id             string `xml:"id"`
	Navn           string `xml:"navn"`
}

type Representant struct {
	ResponsDatoTid   string  `xml:"respons_dato_tid"`
	Versjon          string  `xml:"versjon"`
	Doedsdato        *string `xml:"doedsdato"`
	Etternavn        string  `xml:"etternavn"`
	Foedselsdato     string  `xml:"foedselsdato"`
	Fornavn          string  `xml:"fornavn"`
	Id               string  `xml:"id"`
	Kjoenn           string  `xml:"kjoenn"`
	Fylke            Fylke   `xml:"fylke"`
	Parti            Parti   `xml:"parti"`
	VaraRepresentant bool    `xml:"vara_representant"`
}

type Fylke struct {
	ResponsDatoTid string `xml:"respons_dato_tid"`
	Versjon        string `xml:"versjon"`
	HistoriskFylke bool   `xml:"historisk_fylke"`
	Id             string `xml:"id"`
	Navn           string `xml:"navn"`
}

type Parti struct {
	ResponsDatoTid    string `xml:"respons_dato_tid"`
	Versjon           string `xml:"versjon"`
	Id                string `xml:"id"`
	Navn              string `xml:"navn"`
	RepresentertParti bool   `xml:"representert_parti"`
}

func main() {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://data.stortinget.no/eksport/saker", nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var oversikt SakerOversikt
	err = xml.Unmarshal(b, &oversikt)
	if err != nil {
		panic(err)
	}

	// todo: tabwriter ?
	for _, sak := range oversikt.SakerListe {
		fmt.Println(sak.Id, sak.Tittel)
	}

}
