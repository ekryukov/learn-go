package main

import (
        "fmt"
)

func e101() {
        fmt.Println("1.01 Вывести на экран число π с точностью до сотых.")
        pi := 3.14
        fmt.Println("The first 2 digits of Pi: ", pi)
}

func e102() {
        fmt.Println("1.02 Вывести на экран число e (основание натурального логарифма) с точностью до десятых.")
        e := 1.2
        fmt.Println("E = ", e)
}

func e103() {
        fmt.Println("1.03 Составить программу вывода на экран числа, вводимого с клавиатуры. Выводимому числу должно предшествовать сообщение Вы ввели число.")
        var i int
        fmt.Scanf("%d", &i)
        fmt.Println("Вы ввели число: ", i)
}

func e104() {
        fmt.Println("1.04 Составить программу вывода на экран числа, вводимого с клавиатуры.")
        fmt.Println("После выводимого числа должно следовать сообщение  - вот какое число Вы ввели.")
        var i int
        fmt.Scan(&i)
        fmt.Println(i, " - вот какое число вы ввели")
}

func e105() {
        fmt.Println("1.05 Вывести на одной строке числа 1, 13 и 49 с одним пробелом между ними.")
        fmt.Println(1, 13, 49)
}

func e106() {
        fmt.Println("1.06 Вывести на одной строке числа 7, 15 и 100 с двумя пробелами между ними.")
        var sep string = ""
        fmt.Println(7, sep, 15, sep, 100)

}

func e107() {
        fmt.Println("1.07 Составить программу вывода на экран в одну строку трех любых чисел с двумя пробелами между ними.")
        var (
                a   = 3
                b   = 5
                c   = 56
                sep = ""
        )
        fmt.Println(a, sep, b, sep, c)

}

func e108() {
        fmt.Println("1.08 Составить программу вывода на экран в одну строку четырех любых чисел с одним пробелом между ними.")
        var (
                a = 3
                b = 4
                c = 76
                d = 34
        )
        fmt.Println(a, b, c, d)
}

func main() {
        e101()
        e102()
        e103()
        e104()
        e105()
        e106()
        e107()
        e108()
}
