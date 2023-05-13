package main

type Recipe struct {
	Title       string
	URL         string
	Ingredients string
	Steps       string
}

func (r Recipe) String() string {
	return r.Title
}

func GetRecipes() []Recipe {
	return []Recipe{
		Recipe{Title: "Maceroni ham-kaas"},
		Recipe{Title: "Kip in de hoed"},
		Recipe{Title: "Griekse salade"},
	}
}
