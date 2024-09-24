package hackerrank

import (
	"strings"
)

func gridSearch(G []string, P []string) string {
	// Write your code here
	// fmt.Println(G)
	// fmt.Println(P)
	// found := make(map[string]bool)
	idx := -1
	R := -1
	for i := len(G) - 1; i > 0; i-- {
		idx = strings.Index(G[i], P[len(P)-1])
		if idx > 0 {
			R = i
			break
		}
	}
	if R == -1 || idx == -1 {
		return "NO"
	}

	j := len(P) - 1
	for i := 0; i < len(P); i++ {
		if strings.Index(G[R-i], P[j]) != idx {
			return "NO"
		}
		if i == len(P)-1 && strings.Index(G[R-i], P[j]) == idx {
			return "YES"
		}
		j--
	}

	// all := strings.Join(G, " ")
	// lg := len(G[0])
	// lp := len(P[0])
	// idx := strings.Index(all, P[len(P)-1])
	// if idx < 0 {
	// 	return "NO"
	// }

	// j := 1
	// for i := len(P) - 2; i >= 0; i-- {
	// 	if idx-j-(lg*j) < 0 {
	// 		return "NO"
	// 	}
	// 	if P[i] != all[(idx-j-(lg*j)):(idx-j-(lg*j))+lp] {
	// 		return "NO"
	// 	}
	// 	j++
	// }

	// for i:= 0;i<len(P);i++ {
	//     found[P[i]] = strings.Contains(all, P[i])
	// }
	// for i := range found {
	//     fmt.Println(found[i])
	//     if !found[i] {
	//         return "NO"
	//     }
	// }
	return "YES"
}
