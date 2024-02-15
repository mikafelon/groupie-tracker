package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	F "funct/funct"
)

var data F.PageData

func AlreadyinTheList(st string, list []F.Artists) bool {
	for _, b := range list {
		if b.Name == st {
			return true
		}
	}
	return false
}
func AlreadyinTheList2(st string, list []string) bool {
	for _, c := range list {
		if c == st {
			return true
		}
	}
	return false
}

func Temp(w http.ResponseWriter, r *http.Request) {
	fileSystem := http.Dir("./templates")
	fileServer := http.FileServer(fileSystem)
	_, err := fileSystem.Open(path.Clean(r.URL.Path))
	if os.IsNotExist(err) && r.URL.Path != "/artiste" && r.URL.Path != "/" {
		http.Redirect(w, r, "404.html", http.StatusSeeOther)
		return
	}

	if r.URL.Path == "/index" || r.URL.Path == "/" {

		parseErr := r.ParseForm()
		if parseErr != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		searchText := r.Form.Get("searchText")

		var locations []string
		for _, b := range F.FetchArtists() {
			for j := range b.LocationsRecup {
			if !AlreadyinTheList2(b.LocationsRecup[j], locations) {
					locations = append(locations, b.LocationsRecup[j])
				}
			}
		}

		if len(searchText) <= 0 {
			data = F.PageData{
				ArtistHTML:  F.FetchArtists(),
				ArtistsLIST: F.FetchArtists(),
				Villes:      locations,
			}

			err = tmpl.Execute(w, data)
			log.Println(err)

			return

		}
		var artists []F.Artists
		var villes []string

		for _, b := range F.FetchArtists() {
			if strings.Contains(strings.ToLower(b.Name), strings.ToLower(searchText)) {
				if !AlreadyinTheList(b.Name, artists) {
					artists = append(artists, b)

				}

			}

			for _, v := range b.Members {
				if strings.Contains(strings.ToLower(v), strings.ToLower(searchText)) {
					if !AlreadyinTheList(b.Name, artists) {
						artists = append(artists, b)
					}
				}
			}

			if strings.ToLower(strconv.Itoa(b.Creation)) == strings.ToLower(searchText)  {
				if !AlreadyinTheList(b.Name, artists) {
					artists = append(artists, b)
				}
			}

			if (strings.ToLower(b.FirstAlbum) == strings.ToLower(searchText)) {
				if !AlreadyinTheList(b.Name, artists) {
					artists = append(artists, b)
				}
			}

			for k := range b.RelationsRecup {
				if strings.Contains(strings.ToLower(k), strings.ToLower(searchText)) {
					if !AlreadyinTheList(b.Name, artists) {
						artists = append(artists, b)
					}
				}
			}

			if strings.Contains(strings.ToLower(b.Name), strings.ToLower(searchText)) {
				if !AlreadyinTheList(b.Name, artists) {
					artists = append(artists, b)
				}
			}

			for j := range b.LocationsRecup {
				if !AlreadyinTheList2(b.LocationsRecup[j], villes) {
					villes = append(villes, b.LocationsRecup[j])

				}

			}
		}
		fmt.Println(villes)
		data = F.PageData{
			ArtistHTML:  artists,
			ArtistsLIST: F.FetchArtists(),
			Villes:      villes,
		}

		err = tmpl.Execute(w, data)
		log.Println(err)

		return
	}

	fileServer.ServeHTTP(w, r)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	idform := r.FormValue("artistePage")
	id, _ := strconv.Atoi(idform)
	fmt.Println(idform)

	pageID := data.ArtistHTML[0]
		for _, z := range data.ArtistHTML{
		if id ==  z.ID {
			pageID = z 
		}
	}

	tmpl, err := template.ParseFiles("./templates/artiste.html")
	if err != nil {
		http.Redirect(w, r, "500.html", http.StatusSeeOther)
		fmt.Println("ERREUR")
		return
	}

	tmpl.ExecuteTemplate(w, "artiste.html", pageID)
}
