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
		{Title: "Macaroni ham-kaas"},
		{Title: "Kip in de hoed"},
		{Title: "Griekse salade"},
	}
}
