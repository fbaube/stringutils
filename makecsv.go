package stringutils

import S "strings"

func MakeCSV(ss []string) string {
     if ss == nil || len(ss) == 0 {
     	return ""
	}
     if len(ss) == 1 {
     	return ss[0]
	}
     var sb S.Builder
     var s string 
     for _, s = range ss {
     	 sb.WriteString(s + ", ")
	 }
     s = sb.String()
     return s[0:len(s)-2]
}

func MakeQuotedCSV(ss []string) string {
     if ss == nil {
     	return ""
	}
     switch len(ss) {
     case 0: return "\"\""
     case 1: return "\"" + ss[0] + "\""
     }
     var sb S.Builder
     var s string 
     for _, s = range ss {
     	 sb.WriteString("\"" + s + "\", ")
	 }
     s = sb.String()
     return s[0:len(s)-2]
}

