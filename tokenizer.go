package main

import "fmt"

func tokenize(contents string) {
    for _, char := range contents {
        switch char {
        case LEFT_PAREN:
            fmt.Printf("LEFT_PAREN ( null\n")
            break
        case RIGHT_PAREN:
            fmt.Printf("RIGHT_PAREN ) null\n")
            break
        case LEFT_BRACE:
            fmt.Printf("LEFT_BRACE { null\n")
            break
        case RIGHT_BRACE:
            fmt.Printf("RIGHT_BRACE } null\n")
            break
        case STAR:
            fmt.Printf("STAR * null\n")
            break
        case DOT:
            fmt.Printf("DOT . null\n")
            break
        case COMMA:
            fmt.Printf("COMMA , null\n")
            break
        case PLUS:
            fmt.Printf("PLUS + null\n")
            break
        case MINUS:
            fmt.Printf("MINUS - null\n")
            break
        case SEMICOLON:
            fmt.Printf("SEMICOLON ; null\n")
            break
        }
    }
}
