package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEFAULT_SIZE = -1

type TreeNode struct {
	parent *TreeNode
	childs []*TreeNode
	name   string
	size   int
}

func NewTreeNode(name string, parent *TreeNode, size ...int) *TreeNode {
	s := DEFAULT_SIZE
	if len(size) == 0 {
		s = size[0]
	}
	n := new(TreeNode)
	n.name = name
	n.parent = parent
	n.size = s
	if n.size == DEFAULT_SIZE {
		n.childs = nil
	} else {
		n.childs = []*TreeNode{}
	}
	return n
}

func (t *TreeNode) IsFolder() bool {
	return t.childs != nil
}

func (t *TreeNode) AppendChild(child *TreeNode) {
	if t.IsFolder() {
		child.parent = t
		t.childs = append(t.childs, child)
	} else {
		panic(fmt.Sprintf("Error in child list name parent = %s, child = %s", t.name, child.name))
	}
}

func (t *TreeNode) FindChildBy(name string) *TreeNode {
	for _, val := range t.childs {
		if val.name == name {
			return val
		}
	}

	return nil
}

func ParseInput(pathToFile string) *TreeNode {
	f, err := os.Open(pathToFile)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pathToFile))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	root := NewTreeNode("/", nil)
	var currentNode *TreeNode = nil
	isListing := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd /") {
			currentNode = root
			isListing = false
			continue
		}

		if strings.HasPrefix(line, "$ ls") {
			isListing = true
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			isListing = false
			pathToNavigate := strings.Trim(strings.Replace(line, "$cd ", "", -1), "\n")
			if pathToNavigate == ".." {
				currentNode = currentNode.parent
			} else {
				currentNode = currentNode.FindChildBy(pathToNavigate)
			}
			if currentNode == nil {
				panic(fmt.Sprintf("Parsing broken on line: %s", line))
			}
			continue
		}

		if isListing {
			var node *TreeNode = nil
			if strings.HasPrefix(line, "dir") {
				s := strings.Replace(line, "dir ", "", -1)
				dirName := strings.Trim(s, "/n")
				node = NewTreeNode(dirName, currentNode)
			} else {
				size := -1
				fileName := ""
				for i, val := range strings.Split(line, " ") {
					if i == 0 {
						size, err = strconv.Atoi(val)
						if err != nil {
							panic(fmt.Sprintf("Error in int parsing %s. Line: %s", val, line))
						}
					} else {
						fileName = strings.Trim(val, "\n")
					}
				}
				node = NewTreeNode(fileName, currentNode, size)
			}

			if node != nil {
				currentNode.AppendChild(node)
			} else {
				panic(fmt.Sprintf("Failed to parse line. %s", line))
			}
		}
	}

	return nil
}

func main() {
	isTestFile := false
	flag.BoolVar(&isTestFile, "t", false, "display in uppercase")
	flag.Parse()
	inputFileName := "test.txt"
	if !isTestFile {
		inputFileName = "input.txt"
	}
	fileTree := ParseInput(inputFileName)
	fmt.Printf("Test")
}
