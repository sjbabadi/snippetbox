package main

import "com.sjbabadi/snippetbox/pkg/models"

//this type struct will act as a holding structure for data passed to a template
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
