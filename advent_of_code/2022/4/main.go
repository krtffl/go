package src

import (
	"fmt"
	"krtffl/aoc/helpers"
	"log"
	"strconv"
	"strings"
)

func Solve_Day_4() {
    data, error := helpers.ReadInput("day4.txt")
    if error != nil {
        log.Fatal(error)
    }

    count := 0

    for i:=0; i < len(data); i++ {
        sections := strings.Split(data[i], ",") 

        left_section := make([]int, 2)
        right_section := make([]int, 2)

        left_section[0], error = strconv.Atoi(strings.Split(sections[0], "-")[0])
        if error != nil {
            log.Fatal(error)
        }

        left_section[1], error = strconv.Atoi(strings.Split(sections[0], "-")[1])
        if error != nil {
            log.Fatal(error)
        }

        right_section[0], error = strconv.Atoi(strings.Split(sections[1], "-")[0])
        if error != nil {
            log.Fatal(error)
        }

        right_section[1], error = strconv.Atoi(strings.Split(sections[1], "-")[1])
        if error != nil {
            log.Fatal(error)
        }

        if sections_overlap(left_section, right_section) {
            count++
        }


    }

    fmt.Println(count)
}

func sections_overlap(left, right []int) bool {
    for i:= 0; i < len(left); i++ {
        if left[i] >= right[0] && left[i] <= right[1] {
            return true
        }
        
        if right[i] >= left[0] && right[i] <= left[1] {
            return true
        }
        
    }
    return false
}

func left_contains(left, section []int) bool {
    if left[0] <= section[0] && left[1] >= section[1] {
        return true
    }
    return false
}

func left_is_contained(left, section []int) bool {
    if left[0] >= section[0] && left[1] <= section[1] {
        return true
    }
    return false
}
