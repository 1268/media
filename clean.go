package mech

import "strings"

func Clean(path string) string {
   mapping := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return '.'
      }
      return r
   }
   return strings.Map(mapping, path)
}
