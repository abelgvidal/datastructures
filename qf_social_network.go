//Social network connectivity

//Given a social network containing N members and a log file containing M timestamps at which times pairs of members formed friendships, design an algorithm to determine the earliest time at which all members are connected. Asumme that the log file is sorted by timestamp and that friendship is an equivalence relation. The running time of your algorithm should b "m log n" and use extra space proportional to n. ---> logarithmic with respect to n and linear with respect to m


package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"datastructures/unionfind"
)

func main() {
	// Open the log file for reading
	file, err := os.Open("qf_social_network.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a bufio scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ", ")
		if len(parts) != 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		// Extract the three variables
		timestamp := parts[0]
		fmt.Println("line:", line, parts[1])
		integer1, err := strconv.Atoi(parts[1])
		if err != nil { panic(err) }
		integer2, err := strconv.Atoi(parts[2])
		if err != nil { panic(err) }

		// Process the line or print it
		fmt.Println(timestamp, integer1, integer2)
		unionfind.QUWCUnion(integer1, integer2)

	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
