package main

func SumAll(slices ...[]int) []int {
    result := make([]int, len(slices))

    for i, slice := range slices {
        result[i] = Sum(slice)
    }

    return result
}

func Sum(numbers []int) int{
    var sum int

    for _, number := range numbers {
        sum += number
    }
    return sum
}


func SumAllTails(slices ...[]int) []int {
    var tails []int

    for _, slice := range slices {
        if len(slice) == 0 {
            tails = append(tails, 0)
        } else {
            tails = append(tails, Sum(slice[1:]))
        }
    }

    return tails
}
