package utils

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

const (
	RecipeSeparator    = "=>"
	ComponentSeparator = ","
	ComponentFormat    = "%d %s"

	FuelComponent = "FUEL"
	OreComponent  = "ORE"

	ErrInvalidRecipe         = Error("invalid recipe")
	ErrTooFewInputComponents = Error("too few input components")
	ErrComponentNotFound     = Error("component not found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type Recipe struct {
	Input  []Component
	Output Component
}

type Component struct {
	Quantity int
	Name     string
}

func Parse(reader io.Reader) ([]Recipe, error) {
	scanner := bufio.NewScanner(reader)
	recipes := make([]Recipe, 0)
	for scanner.Scan() {
		row := scanner.Text()
		split := strings.Split(row, RecipeSeparator)
		if len(split) != 2 {
			return nil, ErrInvalidRecipe
		}
		input, output := split[0], split[1]
		inputs := strings.Split(input, ComponentSeparator)
		if len(inputs) < 1 {
			return nil, ErrTooFewInputComponents
		}

		recipe := Recipe{}
		for _, i := range inputs {
			c, err := parseComponent(i)
			if err != nil {
				return nil, err
			}
			recipe.Input = append(recipe.Input, *c)
		}

		o, err := parseComponent(output)
		if err != nil {
			return nil, err
		}
		recipe.Output = *o
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func parseComponent(s string) (*Component, error) {
	var q int
	var n string
	_, err := fmt.Sscanf(strings.TrimSpace(s), ComponentFormat, &q, &n)
	if err != nil {
		return nil, err
	}
	return &Component{
		Quantity: q,
		Name:     n,
	}, nil
}

func FuelToOre(recipes []Recipe) (int, error) {
	return minOre(recipes, 1, FuelComponent, map[string]int{})
}

func minOre(recipes []Recipe, quantity int, name string, left map[string]int) (int, error) {
	// if nothing is wanted, then nothing is required
	if quantity == 0 {
		return 0, nil
	}

	// find the recipe producing what we want
	recipe, err := find(recipes, name)
	if err != nil {
		return 0, err
	}

	// check how many times this recipe need to be executed to produce at least 'quantity' of required component
	times := int(math.Ceil(float64(quantity) / float64(recipe.Output.Quantity)))

	qt := 0
	// for each required components
	for _, i := range recipe.Input {
		if i.Name == OreComponent {
			// if required component is Ore then simply add it since it is like a leaf in a graph
			qt += times * i.Quantity
		} else {
			// otherwise, use recursion

			// required components could be reduced using leftover
			l := int(math.Min(float64(times*i.Quantity), float64(left[i.Name])))
			left[i.Name] -= l
			target := times*i.Quantity - l

			// recursive call
			q, err := minOre(recipes, target, i.Name, left)
			if err != nil {
				return 0, err
			}
			// add required Ores to total
			qt += q
		}
	}

	// compute leftover
	left[name] += times*recipe.Output.Quantity - quantity
	return qt, nil
}

func find(recipes []Recipe, name string) (*Recipe, error) {
	for _, r := range recipes {
		if r.Output.Name == name {
			return &r, nil
		}
	}
	return nil, ErrComponentNotFound
}
