package main

import (
	"fmt"
	"strings"
)

var testMap = make(map[string]string)

type Trie struct {
	pattern  string
	part     string
	children map[string]*Trie
	isEnd    bool
}

func new() *Trie {
	t := &Trie{isEnd: false, pattern: "", part: ""}
	t.children = make(map[string]*Trie)
	return t
}

func (t *Trie) searchPrefix(prefix []string) *Trie {
	root, ok := t, true
	for i := 0; i < len(prefix); i++ {
		_, ok = root.children[prefix[i]]
		if !ok {
			flag := false
			for child := range root.children {
				if strings.HasPrefix(child, ":") {
					flag = true
					testMap[child[1:]] = prefix[i]
					root = root.children[child]
				} else if strings.HasPrefix(child, "*") {
					flag = true
					testMap[child[1:]] = strings.Join(prefix[i:], "")
					root = root.children[child]
				}
			}
			if !flag {
				return nil
			}
		} else {
			root = root.children[prefix[i]]
		}
	}
	return root
}

func (t *Trie) insert(parts []string) {
	root, ok := t, true
	for _, part := range parts {
		_, ok = root.children[part]
		if !ok {
			root.children[part] = &Trie{
				pattern:  root.pattern + "/" + part,
				part:     part,
				children: make(map[string]*Trie),
			}
		}
		root = root.children[part]
	}
	root.isEnd = true
}

func ParseURLs(path string) (parts []string) {
	vs := strings.Split(path, "/")
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return
}

func main() {
	parts := ParseURLs("/hello/:name/*any")
	t := new()
	t.insert(parts)
	// test := ParseURLs("/hello/test/")
	test := ParseURLs("/hello/test/newbee")
	tnode := t.searchPrefix(test)

	// 获得动态匹配的数据映射
	// 获得匹配得到的路由，对应相应的handleFunc
	fmt.Printf("map_data : %v, pattern : %s", testMap, tnode.pattern)
}
