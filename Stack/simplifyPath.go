package stack

import (
	"strings"
)

// LC #71
func SimplifyPath(path string) string {
	stack := []string{}

	for _, str := range strings.Split(path, "/") {
		if str != "" && str != "." && str != ".." {
			stack = append(stack, str)
		} else if str == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return "/" + strings.Join(stack, "/")
}

func SimplifyPath2(path string) string {
	stack := []string{}
	curr := ""
	for _, char := range path + "/" {
		if char != '/' {
			curr += string(char)
		} else if char == '/' {
			if curr == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else if curr != "" && curr != "." {
				stack = append(stack, curr)
				curr = ""
			}
		}
	}

	return "/" + strings.Join(stack, "/")
}
