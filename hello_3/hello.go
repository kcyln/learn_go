package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

	//prefix := englishHelloPrefix
	//
	//switch language {
	//case spanish:
	//	prefix = spanishHelloPrefix
	//case french:
	//	prefix = frenchHelloPrefix
	//}
	//
	//return prefix + name

	//if language == spanish {
	//	return spanishHelloPrefix + name
	//}
	//
	//if language == french {
	//	return frenchHelloPrefix + name
	//}

	//return englishHelloPrefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
