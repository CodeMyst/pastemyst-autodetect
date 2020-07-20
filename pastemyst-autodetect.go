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

    langs := [34]string{"C", "Java", "Python", "C++", "C#", "Visual Basic", "JavaScript", "R", "PHP", "Swift", "SQL", "Go", "Perl", "MATLAB", "Ruby", "Rust", "Objective-C", "Dart", "D", "Kotlin", "COBOL", "Lisp", "Julia", "Lua", "Scala", "VBScript", "Haskell", "Scheme", "TypeScript", "Ada", "Prolog", "PostScript", "Erlang", "Fortran"}

    lang, _ := enry.GetLanguageByClassifier(dat, langs[:])

    if lang == "" {
        fmt.Println("unknown");
    } else {
        fmt.Println(lang);
    }
}
