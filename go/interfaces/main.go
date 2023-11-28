package main

// English bot
type englishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func printGreeting(eb englishBot) {
	println(eb.getGreeting())
}

// / Spanish bot
type spanishBot struct{}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(sb spanishBot) {
	println(sb.getGreeting())
}

func main() {
	// eb := englishBot.getGreeting(englishBot{})
	// sb := spanishBot.getGreeting(spanishBot{})

	// println(eb)
	// println(sb)
}
