package main

import (
	"flag"
	"fmt"
)

func main() {

	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur, fr...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language code
type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         //Hebrew
	"ur": "ہیلو",              // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
	"pt": "Olá mundo",         // Portuguese
}

// greet says hello to the world in hte specified language
func greet(l language) string {

	greeting, ok := phrasebook[l]
	if !ok {
		fmt.Sprintf("Unsuported language: %q\n", l)
	}
	return greeting

}
