package protein

var codonAminoAcidMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAC": "Tyrosine",
	"UAU": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(codon string) string {
	return codonAminoAcidMap[codon]
}

func FromRNA(rna string) []string {
	ps := make([]string, 0)
	for i := 0; i <= len(rna)-3; i += 3 {
		p := FromCodon(rna[i:i+3])
		if p == "STOP" {
			break
		}
		ps = append(ps, p)
	}
	return ps
}
