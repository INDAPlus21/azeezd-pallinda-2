// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answer := make(chan string)

	random_q := make([]string, 5)
	random_q = append(random_q, "the journey of a thousand miles begins with one step") // random quote by Lao Tzu

	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.

	// Question listening goroutine
	go func() {
		for q := range questions {
			if len(q) > 10 && rand.Intn(10) == 0 { // Randomly add questions to the random_q to change the random prophecies
				random_q = append(random_q, q)
			}
			go prophecy(q, answer)
		}
	}()

	// Answer goroutine
	go func() {
		for a := range answer {
			fmt.Printf("%s: %s\n> ", star, a)
		}
	}()

	// Random prophecies goroutine
	go func() {
		for {
			time.Sleep(time.Duration(30+rand.Intn(30)) * time.Second)
			go prophecy(random_q[rand.Intn(len(random_q))], answer) // Ask random question from the random_q stored sentences
		}
	}()

	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	sb := new(strings.Builder)
	question = strings.ToLower(question)

	question_re := regexp.MustCompile(`what|why|how|when|which|where|will|would`)
	time_re := regexp.MustCompile("time|year|month|week|day|hour|minute|second")
	location := question_re.FindStringIndex(question)
	if location != nil && rand.Intn(3) == 0 {
		start := location[1]
		words := strings.Fields(question[start:])
		longest := ""
		for _, w := range words {
			if len(w) > len(longest) {
				longest = w
			}
		}
		switch rand.Intn(4) {
		case 0:
			sb.WriteString("Apollo")
			switch rand.Intn(4) {
			case 0:
				sb.WriteString("'s rays shines upon ")
			case 1:
				sb.WriteString(" warmth thy soul and ")
			case 2:
				sb.WriteString(" hearest what thou speakest regarding ")
			case 3:
				sb.WriteString(" beams his blessing upon ")
			}
			sb.WriteString(longest)
			switch rand.Intn(3) {
			case 0:
				sb.WriteString(". Rejoice!")
			case 1:
				sb.WriteString(". Praises!")
			case 2:
				sb.WriteString(".")
			}
		case 1:
			sb.WriteByte('"')
			sb.WriteString(longest)
			sb.WriteByte('"')
			switch rand.Intn(3) {
			case 0:
				sb.WriteString(" echoes thoughout Olympia.")
			case 1:
				sb.WriteString(" they hear.")
			case 2:
				sb.WriteString(" encoils thy fate.")
			}
		case 2:
			sb.WriteString(longest)
			sb.WriteString(" is ")
			switch rand.Intn(3) {
			case 0:
				sb.WriteString("within thee.")
			case 1:
				sb.WriteString("without thee.")
			case 2:
				sb.WriteString("thee.")
			}
		default:
			if rand.Intn(50) == 0 {
				sb.WriteString("Take counsel from Google.")
			} else {
				sb.WriteString("The divine ears wish to hear thee more.")
			}
		}
	} else if time_re.MatchString(question) || rand.Intn(5) == 0 {
		switch rand.Intn(3) {
		case 0:
			sb.WriteString("Time brings thee the answer.")
		case 1:
			sb.WriteString("Let the sands of time carry thee.")
		case 2:
			sb.WriteString("Time flows...")
		}
	} else {
		words := strings.Fields(question)
		words_amount := len(words)
		longest := ""
		for _, w := range words {
			if len(w) > len(longest) {
				longest = w
			}
		}
		switch rand.Intn(5) {
		case 0:
			sb.WriteString(words[rand.Intn(words_amount)])
			for i := 0; i < rand.Intn(words_amount); i++ {
				sb.WriteString("... ")
				sb.WriteString(words[rand.Intn(words_amount)])
			}
		case 1:
			switch rand.Intn(3) {
			case 0:
				sb.WriteString("Tell us more about ")
				sb.WriteString(longest)
			case 1:
				sb.WriteString(longest)
				sb.WriteString("? We listen for more.")
			case 2:
				sb.WriteString("We hearest thou. What about ")
				sb.WriteString(longest)
				sb.WriteByte('?')
			}

		case 2:
			switch rand.Intn(5) {
			case 0:
				sb.WriteString("Exquisite!")
			case 2:
				sb.WriteString("Marvelous!")
			case 3:
				sb.WriteString("Splendid!")
			case 4:
				sb.WriteString("Extraordinary!")
			}
		case 3:
			switch rand.Intn(4) {
			case 0:
				sb.WriteString("Sun shines upon us all.")
			case 1:
				sb.WriteString("Moonlight guides the lost.")
			case 2:
				sb.WriteString("Blessed be Gaia!")
			case 3:
				sb.WriteString("May thy days be blessed.")
			}
		case 4:
			sb.WriteString("They listen...")
		}
	}

	answer <- sb.String()
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
