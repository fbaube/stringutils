package stringutils

import (
  "fmt"
 "strconv"
 S "strings"
)

type PropSet map[string]string

// YamlMapAsPropSet returns a PropSet, i.e. a map[string]string
func YamlMapAsPropSet(u map[interface{}]interface{}) (ps PropSet) {
  ps = make(map[string]string)
  for k,v := range u {
    ks := S.ToLower(k.(string))
    switch v.(type) {
    case int:
      ps[ks] = strconv.Itoa(v.(int))
    case string:
      ps[ks] = v.(string)
    default:
      fmt.Printf("Bad Prop type: %T \n", v)
    }
  }
  return ps
}
