package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	uniqueEmail := make(map[string]int)
	for _, email := range emails {
		parts := strings.Split(email, "@")

		leftPart := strings.Replace(parts[0], ".", "", -1)

		var plusIndex int
		// looking for first plus rune
		for i, r := range leftPart {
			// 43 - plus rune
			if r == 43 {
				plusIndex = i
				break
			}
		}

		if plusIndex > 0 {
			leftPart = leftPart[:plusIndex]
		}
		uniqueEmail[leftPart+"|"+parts[1]]++
	}

	return len(uniqueEmail)
}

func main() {
	emails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}

	/*emails := []string{
		"fg.r.u.uzj+o.pw@kziczvh.com",
		"r.cyo.g+d.h+b.ja@tgsg.z.com",
		"fg.r.u.uzj+o.f.d@kziczvh.com",
		"r.cyo.g+ng.r.iq@tgsg.z.com",
		"fg.r.u.uzj+lp.k@kziczvh.com",
		"r.cyo.g+n.h.e+n.g@tgsg.z.com",
		"fg.r.u.uzj+k+p.j@kziczvh.com",
		"fg.r.u.uzj+w.y+b@kziczvh.com",
		"r.cyo.g+x+d.c+f.t@tgsg.z.com",
		"r.cyo.g+x+t.y.l.i@tgsg.z.com",
		"r.cyo.g+brxxi@tgsg.z.com",
		"r.cyo.g+z+dr.k.u@tgsg.z.com",
		"r.cyo.g+d+l.c.n+g@tgsg.z.com",
		"fg.r.u.uzj+vq.o@kziczvh.com",
		"fg.r.u.uzj+uzq@kziczvh.com",
		"fg.r.u.uzj+mvz@kziczvh.com",
		"fg.r.u.uzj+taj@kziczvh.com",
		"fg.r.u.uzj+fek@kziczvh.com",
	}*/

	fmt.Println(numUniqueEmails(emails))
}
