package nestapp

import (
	"github.com/inadati/demo.go/nestapp/greet"
	"github.com/inadati/demo.go/nestapp/ps"
)

var Greet = &greetMod{}
var Ps = &psMod{}

type greetMod struct {
    English greet.EnglishMod
    Japanese greet.JapaneseMod
}
type psMod struct {
    PrintEnglishGreet ps.PrintEnglishGreetMod
    PrintJapaneseGreet ps.PrintJapaneseGreetMod
}
