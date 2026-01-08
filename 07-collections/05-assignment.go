/*
Find the most occurring "word size" in the given string and print the most occuring "word size" & the number of occurances
Use "strings.Split()" to get the words
Use "len()" to get the length of a word
*/
/*
	ex output format:
	5 letter words occurs the most with 31 occurrances
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// var wordsCountBySize map[int]int = map[int]int{}
	var wordsCountBySize map[int]int = make(map[int]int)

	str := "Commodo Lorem enim dolore culpa minim ipsum ex excepteur in Commodo duis nulla ex laborum irure sunt incididunt Incididunt amet Lorem amet dolor sit consectetur culpa esse quis laborum pariatur laborum fugiat mollit Mollit voluptate aliquip Lorem incididunt mollit pariatur eu enim proident culpa esse laborum voluptate Nostrud aliqua magna ipsum qui duis euNisi pariatur sit do magna Lorem nostrud voluptate occaecat occaecat quis dolore irure Velit aliqua reprehenderit eu duis aliqua excepteur duis non enim Nostrud qui voluptate enim eiusmod dolor proident laboris nostrud commodo laborum aliquip sunt Exercitation do anim do ullamco Ipsum eiusmod aute qui ea consectetur Veniam laborum occaecat mollit pariatur commodo id ullamco dolore ipsum sit dolore elit fugiatUt ad aliqua dolor quis velit reprehenderit Dolore aliquip exercitation do fugiat Pariatur irure aliqua magna quis"

	words := strings.Split(str, " ")
	for _, word := range words {
		wordsCountBySize[len(word)]++
	}
	var maxSize, maxCount int
	for size, count := range wordsCountBySize {
		if count > maxCount {
			maxCount, maxSize = count, size
		}
	}
	fmt.Printf("%d letter words occurs the most with %d occurrences\n", maxSize, maxCount)

}
