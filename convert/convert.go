package convert

import (
	"regexp"
	"strings"
)

// func ToCamelCase(s string) string {
//	// decide _ or -
//	var underscores, dashes bool
//	underscores, _ = regexp.MatchString("_", s)
//	dashes, _ = regexp.MatchString("-", s)

//	if underscores && dashes {
//		fmt.Printf("Cannot mix underscores and dashes\n")
//	} else {
//		fmt.Printf("string: %s. underscores: %t dashes: %t\n", s, underscores, dashes)
//	}
//	// split  by separator
//		var words []string
//	if underscores {
//		words = strings.Split(s, "_")
//	} else {
//		words = strings.Split(s, "-")
//	}

//	// capitalize all (except first if it is not capitalized
//	for i, w := range words[1:] {
//		words[i+1] = strings.Title(w)
//	}

//	fmt.Println(words)



//	// join words
//	result := strings.Join(words, "")
//	return  result
// }

func ToCamelCase(s string) string {
  words := regexp.MustCompile("-|_").Split(s, -1)

  for i, w := range words[1:] {
	words[i+1] = strings.Title(w)
  }

  return strings.Join(words, "")
}
