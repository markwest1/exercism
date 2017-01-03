package allergies

const testVersion = 1

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies lists allergens from a given allergy score.
func Allergies(score uint) []string {
	list := make([]string, 0)

	for name, value := range allergens {
		if score&value != 0 {
			list = append(list, name)
		}
	}

	return list
}

// AllergicTo returns true if score indicates an allergy to allergen.
func AllergicTo(score uint, allergen string) bool {
	if v, ok := allergens[allergen]; ok {
		return score&v != 0
	}

	return false
}
