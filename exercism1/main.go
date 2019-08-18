// Introduction
//
// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
//
// Bob answers 'Sure.' if you ask him a question.
//
// He answers 'Whoa, chill out!' if you yell at him.
//
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
//
// He says 'Fine. Be that way!' if you address him without actually saying anything.
//
// He answers 'Whatever.' to anything else.
//
// Bob's conversational partner is a purist when it comes to written communication
// and always follows normal rules regarding sentence punctuation in English.


package main

import "fmt"
import "strconv"

func responses(input string) string {
    var punctuation string
    var yell_question string
    b, err := strconv.ParseBool(input)
    if b {
        punctuation = string(input[len(input)-1])
        yell_question = string(input[len(input)-2] + input[len(input)-1])
    } else {
        return "Whatever."
    }

    if yell_question == "!?" || yell_question == "?!" {
        return "Calm down, I know what I'm doing!"
    } else {
        if punctuation == "?" {
            return "Sure."
        } else if punctuation == "!" {
            return "Whoa, chill out!"
        } else if punctuation == "" {
            return "Fine. Be that way!"
        } else {
            return "Whatever."
        }
    }
}

func main(){
    var msgToBob string
    fmt.Scanf("%f", &msgToBob)
    fmt.Println(responses(msgToBob))
}
