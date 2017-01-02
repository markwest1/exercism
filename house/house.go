package house

import "fmt"

var (
	lineEnd = "\n"
	versEnd = "\n\n"

	subject = []string{
		"the horse and the hound and the horn",
		"the farmer sowing his corn",
		"the rooster that crowed in the morn",
		"the priest all shaven and shorn",
		"the man all tattered and torn",
		"the maiden all forlorn",
		"the cow with the crumpled horn",
		"the dog",
		"the cat",
		"the rat",
		"the malt",
		"the house that Jack built.",
	}

	phrase = []string{
		"that belonged to",
		"that kept",
		"that woke",
		"that married",
		"that kissed",
		"that milked",
		"that tossed",
		"that worried",
		"that killed",
		"that ate",
		"that lay in",
	}
)

// Embed embeds a noun phrase as the object of relative clause with a
// transitive verb.
//
// Argument relPhrase is a phrase with a relative clause, minus the object
// of the clause.  That is, relPhrase consists of a subject, a relative
// pronoun, a transitive verb, possibly a preposition, but then no object.
func Embed(relPhrase, nounPhrase string) string {
	return fmt.Sprintf("%s %s", relPhrase, nounPhrase)
}

// Verse generates a verse of a song with relative clauses that have
// a recursive structure.
func Verse(subject string, relPhrases []string, nounPhrase string) string {
	l := len(relPhrases)

	if l == 0 {
		return Embed(subject, nounPhrase)
	}

	return Verse(subject, relPhrases[:l-1], Embed(relPhrases[l-1], nounPhrase))
}

// Song generates the full text of "The House That Jack Built".  Oh yes, you
// could just return a string literal, but humor us; use Verse as a subroutine.
func Song() (song string) {
	vc := len(subject)
	pc := len(phrase)
	song = Embed("This is", subject[vc-1]) + versEnd

	for verse := 1; verse < vc; verse++ {

		song += Embed("This is", subject[vc-verse-1]+lineEnd)
		emptyStringArray := []string{}
		end := lineEnd

		for i := 0; i < verse; i++ {
			if i == verse-1 {
				end = versEnd
				if verse == vc-1 {
					end = ""
				}
			}

			song += Verse(phrase[pc-verse+i], emptyStringArray, subject[vc-verse+i]+end)
		}
	}

	return
}
