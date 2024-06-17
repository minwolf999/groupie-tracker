package ApiAndDataRequest

import "strings"

// Make a string to lower case and remove all the accent of a string
func RemoveAccent(s string) string {
	s = strings.ToLower(s)

	if strings.HasPrefix(s, "cluj-") {
		s = strings.ReplaceAll(s, "cluj-", "")
	} else if strings.HasSuffix(s, " city") {
		s = strings.ReplaceAll(s, " city", "")
	} else if strings.HasPrefix(s, "asia/") {
		s = strings.ReplaceAll(s, "asia/", "")
	} else if strings.HasPrefix(s, "borriana / ") {
		s = strings.ReplaceAll(s, "borriana / ", "")
	} else if s == "netherlands antilles" {
		s = "netherlands"
	}

	accentMap := map[rune]string{
		'à': "a", 'á': "a", 'â': "a", 'ä': "a", 'ã': "a", 'å': "a",
		'ç': "c",
		'è': "e", 'é': "e", 'ê': "e", 'ë': "e",
		'ì': "i", 'í': "i", 'î': "i", 'ï': "i",
		'ñ': "n",
		'ò': "o", 'ó': "o", 'ô': "o", 'ö': "o", 'õ': "o", 'ø': "o",
		'ù': "u", 'ú': "u", 'û': "u", 'ü': "u",
		'ý': "y", 'ÿ': "y",
		'ß': "ss",
		'-': "", '.': "", ' ': "",
	}

	s2 := ""
	for _, i := range s {
		nRune := accentMap[i]

		if nRune != "" || i == '.' || i == '-' || i == ' ' {
			s2 += nRune
		} else {
			s2 += string(i)
		}
	}

	return s2
}
