package viewmodel

import (
	"fmt"

	"web-app-go/src/github.com/lss/webapp/model"
)

// Shop struct defines the structure for different types of Shops
type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

// Category struct defines the different categories a shop may have
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

// NewShop returns a Shop that has an underlying number of categories and information about it.
func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:  "Lemonade Stand Supply - Shop",
		Active: "shop",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categorytoVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

func categorytoVM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}
