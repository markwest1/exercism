package protein

const testVersion = 1

// Proteins maps proteins to codons
var Proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromRNA translates an RNA sequence into a polypeptide protein sequence.
func FromRNA(rna string) (chain []string) {
	if len(rna) < 3 || len(rna)%3 != 0 {
		return
	}

	for i := 3; i <= len(rna); i += 3 {
		if protein, ok := Proteins[rna[i-3:i]]; ok {
			if protein == "STOP" || protein == "" {
				return
			}

			chain = append(chain, protein)
		}
	}

	return
}

// FromCodon translates a codon to a protein.
func FromCodon(codon string) string {
	if protein, ok := Proteins[codon]; ok {
		return protein
	}

	return ""
}
