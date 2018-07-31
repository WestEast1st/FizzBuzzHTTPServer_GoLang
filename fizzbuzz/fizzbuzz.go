package fizzbuzz

import "strconv"

var response string

func fizzBuzz() func(int) string {
    ret := ""
    return func(i int) string {
        switch {
            case i % 15   ==  0 :
                ret = "fizz buzz"
            case i % 3    ==  0 :
                ret = "fizz"
            case i % 5    ==  0 :
                ret = "buzz"
            default :
                ret =  strconv.Itoa(i)
        }
        return ret
    }
}

func Handler(inputNum int) string{
    fb := fizzBuzz()
    return fb(inputNum)
}
