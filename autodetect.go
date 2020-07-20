package main

import "fmt"
import "github.com/go-enry/go-enry/v2"
import "io/ioutil"
import "os"

func main() {
    if len(os.Args) != 2 {
        panic("missing file path argument")
    }

    dat, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        panic(err)
    }

    langs := [45]string{"C", "Java", "Python", "C++", "C#", "Visual Basic", "JavaScript", "R", "PHP", "Swift", "SQL", "Go", "Assembly", "Perl", "MATLAB", "Ruby", "Scratch", "Rust", "Objective-C", "Groovy", "Dart", "D", "Kotlin", "COBOL", "ABAP", "Delphi", "Logo", "Powershell", "Lisp", "Julia", "Lua", "Scala", "VBScript", "Haskell", "Scheme", "TypeScript", "Ada", "Prolog", "PostScript", "Erlang", "RPG", "Apex", "C shell", "Fortran"}

    lang, safe := enry.GetLanguageByClassifier(dat, langs[:])

    fmt.Println(lang, safe)
}
