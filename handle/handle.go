package handle

import (
	"fmt"
	help "groupie-tracker/help"
	"net/http"
	"strconv"
	"strings"
)

var allocation help.AllLocationIndex
var alldate help.AllDatesIndex

func HandleHomeRequest(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"/",
		"/location/",
		"/date/",
		"/relation/",
		"/locations/",
		"/locations-artist/",
		"/dates/",
		"/dates-artist/",
	}

	if !help.IsMatch(r.URL.Path, paths) {

		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}
	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err := help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	context := help.Context{
		Artists: artists,
	}

	help.RenderTemplate(w, context, "templates/index.html")
}

func HandleLocationRequest(w http.ResponseWriter, r *http.Request) {

	// Récupérer l'ID de la location depuis l'URL
	id := r.URL.Path[len("/location/"):]
	intid, _ := strconv.Atoi(id)
	if intid <= 0 || intid > 52 {
		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}

	// Fetch location data
	locationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id)
	var locationData help.LocationData
	err := help.FetchDataFromAPI(locationURL, &locationData)
	if err != nil {
		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}

	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err = help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}

	// Process data and render template
	n := locationData.ID - 1

	artist := artists[n]

	context := help.Context{
		Artists:      artists,
		LocationData: locationData,
		Image:        artist.Image,
		Name:         artist.Name,
	}

	help.RenderTemplate(w, context, "templates/location.html")
}

func HandleDatesRequest(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la location depuis l'URL
	id := r.URL.Path[len("/date/"):]

	intid, _ := strconv.Atoi(id)

	if intid <= 0 || intid > 52 {
		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}

	// Fetch dates data
	datesURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", id)
	var datesData help.DateData
	err := help.FetchDataFromAPI(datesURL, &datesData)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err = help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Process data and render template
	n := datesData.ID - 1

	artist := artists[n]
	trimeDateStruct := help.DateData{
		ID:    datesData.ID,
		Dates: help.TrimStart(datesData.Dates),
	}

	context := help.Context{
		DateData: trimeDateStruct,
		Image:    artist.Image,
		Name:     artist.Name,
	}

	help.RenderTemplate(w, context, "templates/date.html")
}

func HandleRelationRequest(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de la location depuis l'URL
	id := r.URL.Path[len("/relation/"):]
	intid, _ := strconv.Atoi(id)

	if intid <= 0 || intid > 52 {
		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}
	// Fetch relations data
	relationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)
	var relationData help.RelationsData
	err := help.FetchDataFromAPI(relationURL, &relationData)

	var locationDates []help.LocationDate

	for location, dates := range relationData.DatesLocations {
		for _, date := range dates {
			locationDate := help.LocationDate{
				Location: location,
				Date:     date,
			}
			locationDates = append(locationDates, locationDate)
		}
	}

	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err = help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Process data and render template
	n := relationData.ID - 1
	artist := artists[n]

	context := help.Context{
		Artists:       artists,
		RelationData:  relationData,
		LocationDates: locationDates,
		Image:         artist.Image,
		Name:          artist.Name,
	}

	help.RenderTemplate(w, context, "templates/relation.html")
}

func HandleAllLocationRequest(w http.ResponseWriter, r *http.Request) {

	paths := []string{
		"/",
		"/location/",
		"/date/",
		"/relation/",
		"/locations/",
		"/locations-artist/",
		"/dates/",
		"/dates-artist/",
	}

	if !help.IsMatch(r.URL.Path, paths) {

		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}
	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err := help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Fetch allocation data
	err = help.FetchDataFromAPI("https://groupietrackers.herokuapp.com/api/locations", &allocation)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Process data and render template
	uniqueLocations := make(map[string][]int)
	for _, entry := range allocation.AllLocationIndex {
		for _, location := range entry.Locations {
			location = strings.ToLower(location)
			uniqueLocations[location] = append(uniqueLocations[location], entry.ID)
		}
	}

	var locationList help.LocationList
	for location, artistIDs := range uniqueLocations {
		var tab []help.Artists
		locationList.ListOfLocation = append(locationList.ListOfLocation, location)
		for _, id := range artistIDs {
			tab = append(tab, artists[id-1])
		}
		locationList.ListOfArtist = append(locationList.ListOfArtist, tab)
	}

	context := help.Context{
		LocationList: locationList,
	}

	help.RenderTemplate(w, context, "templates/locations.html")
}

func HandleArtistLocationRequest(w http.ResponseWriter, r *http.Request) {

	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err := help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Fetch allocation data
	err = help.FetchDataFromAPI("https://groupietrackers.herokuapp.com/api/locations", &allocation)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Process data and render template
	uniqueLocations := make(map[string][]int)
	for _, entry := range allocation.AllLocationIndex {
		for _, location := range entry.Locations {
			location = strings.ToLower(location)
			uniqueLocations[location] = append(uniqueLocations[location], entry.ID)
		}
	}

	var locationList help.LocationList
	for location, artistIDs := range uniqueLocations {
		var tab []help.Artists
		locationList.ListOfLocation = append(locationList.ListOfLocation, location)
		for _, id := range artistIDs {
			tab = append(tab, artists[id-1])
		}
		locationList.ListOfArtist = append(locationList.ListOfArtist, tab)
	}

	location := r.URL.Path[len("/locations-artist/"):]
	index := help.FindStringIndex(locationList.ListOfLocation, location)
	if index < 0 {
		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}
	context := help.Context{
		Artists: locationList.ListOfArtist[index],
	}
	help.RenderTemplate(w, context, "templates/artist.html")
}

func HandleAllDatesRequest(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"/",
		"/location/",
		"/date/",
		"/relation/",
		"/locations/",
		"/locations-artist/",
		"/dates/",
		"/dates-artist/",
	}

	if !help.IsMatch(r.URL.Path, paths) {

		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return

	}
	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err := help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Fetch alldate data
	err = help.FetchDataFromAPI("https://groupietrackers.herokuapp.com/api/dates", &alldate)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Process data and render template
	uniqueDates := make(map[string][]int)
	for _, entry := range alldate.AllDatesIndex {
		for _, date := range entry.Dates {
			uniqueDates[date] = append(uniqueDates[date], entry.ID)
		}
	}

	var dateList help.DateList
	for date, artistIDs := range uniqueDates {
		var tab []help.Artists
		dateList.ListOfDate = append(dateList.ListOfDate, date)
		for _, id := range artistIDs {
			tab = append(tab, artists[id-1])
		}
		dateList.ListOfArtist = append(dateList.ListOfArtist, tab)
	}

	trimDateStruct := help.DateList{
		ListOfDate:   help.TrimStart(dateList.ListOfDate),
		ListOfArtist: dateList.ListOfArtist,
	}

	context := help.Context{
		DateList: trimDateStruct,
	}

	help.RenderTemplate(w, context, "templates/dates.html")
}

func HandleArtistDateRequest(w http.ResponseWriter, r *http.Request) {

	// Fetch artists data
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []help.Artists
	err := help.FetchDataFromAPI(artistsURL, &artists)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Fetch alldate data
	err = help.FetchDataFromAPI("https://groupietrackers.herokuapp.com/api/dates", &alldate)
	if err != nil {
		w.WriteHeader(500)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "500",
			ErrorMessage: "Erreur lors de la récupération des données de l'API ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(500)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}

	// Process data and render template
	uniqueDates := make(map[string][]int)
	for _, entry := range alldate.AllDatesIndex {
		for _, date := range entry.Dates {
			uniqueDates[date] = append(uniqueDates[date], entry.ID)
		}
	}

	var dateList help.DateList
	for date, artistIDs := range uniqueDates {
		var tab []help.Artists
		dateList.ListOfDate = append(dateList.ListOfDate, date)
		for _, id := range artistIDs {
			tab = append(tab, artists[id-1])
		}
		dateList.ListOfArtist = append(dateList.ListOfArtist, tab)
	}

	date := r.URL.Path[len("/dates-artist/"):]
	index := help.FindStringIndex(help.TrimStart(dateList.ListOfDate), date)
	if index < 0 {
		w.WriteHeader(404)
		ErrMessage := help.ErrorStruct{
			ErrorName:    "404",
			ErrorMessage: "The page you are looking for cannot be found. It might have been moved or doesn't exist. ",
		}
		context := help.Context{
			ErrorStruct: ErrMessage,
		}
		w.WriteHeader(404)
		help.RenderTemplate(w, context, "templates/error.html")
		return
	}
	context := help.Context{
		Artists: dateList.ListOfArtist[index],
	}
	help.RenderTemplate(w, context, "templates/artist.html")
}
