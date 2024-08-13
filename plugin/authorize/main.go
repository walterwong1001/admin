package main

import (
	"github.com/weitien/admin/plugin/trie"
	"log"
	"strings"
)

func main() {
	//pattern := "^\\d+$"
	//str := "12345"
	//matched, err := regexp.MatchString(pattern, str)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//if matched {
	//	fmt.Println("String matches the pattern.")
	//} else {
	//	fmt.Println("String does not match the pattern.")
	//}

	arr := []string{"/user/^\\d+$/GET", "/user/^\\d+$/DELETE", "/account/lock/^\\d+$", "/account/activity/^\\d+$", "/permission/all"}

	trie := trie.New[string]()
	for _, e := range arr {
		slice := strings.Split(e, "/")[1:]
		trie.Add(slice, "")
	}
	path := "user/123"
	nodes := strings.Split(path, "/")
	nodes = append(nodes, "GET")
	n := trie.Find(nodes)
	log.Println(n)
}
