package protein

const testVersion = 1

// Proteins maps proteins to codons
var Proteins = map[string][]string{
	"Methionine":    []string{"AUG"},
	"Phenylalanine": []string{"UUU", "UUC"},
	"Leucine":       []string{"UUA", "UUG"},
	"Serine":        []string{"UCU", "UCC", "UCA", "UCG"},
	"Tyrosine":      []string{"UAU", "UAC"},
	"Cysteine":      []string{"UGU", "UGC"},
	"Tryptophan":    []string{"UGG"},
	"STOP":          []string{"UAA", "UAG", "UGA"},
}

// FromRNA translates an RNA sequence into a polypeptide protein sequence.
func FromRNA(rna string) (chain []string) {
	if len(rna) < 3 || len(rna)%3 != 0 {
		return
	}

	for i := 3; i <= len(rna); i += 3 {
		protein := FromCodon(rna[i-3 : i])
		if protein == "STOP" || protein == "" {
			return
		}
		chain = append(chain, protein)
	}

	return
}

// FromCodon translates a codon to a protein.
func FromCodon(codon string) string {
	for protein, codons := range Proteins {
		for _, c := range codons {
			if c == codon {
				return protein
			}
		}
	}

	return ""
}
