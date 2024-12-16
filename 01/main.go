package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "slices"
)

func read_input() ([]string, error) {
    file, err := os.Open("input.txt")
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(file)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("Could not read file: %w", err)
    }

    return lines, nil
}

func parse_input() ([]int, []int, error) {
    lines, err := read_input()
    if err != nil {
        return nil, nil, err
    }

    var left_list []int
    var right_list []int
    for _, line := range lines {
        data := strings.Split(line, " ")

        left_num, err := strconv.Atoi(data[0])
        if err != nil { return nil, nil, err }
        right_num, err := strconv.Atoi(data[3])
        if err != nil { return nil, nil, err }

        left_list = append(left_list, left_num)
        right_list = append(right_list, right_num)
    }

    return left_list, right_list, nil
}

func calculate_frequency_map(list []int) map[int]int {
    frequency_map := make(map[int]int)
    
    for _, n := range list {
        frequency_map[n]++
    }
    
    return frequency_map
}

func main() {
    left_list, right_list, err := parse_input()
    if err != nil {
        panic(err)
    }

    slices.Sort(left_list)
    slices.Sort(right_list)

    total_distance := 0
    
    for i := 0; i < len(left_list); i++ {
        distance := left_list[i] - right_list[i]
        if distance < 0 { distance = -distance }

        total_distance += distance
    }

    fmt.Println("Total distance:", total_distance)

    frequency_map := calculate_frequency_map(right_list)
    similarity := 0

    for _, n := range left_list {
        similarity += n * frequency_map[n]
    }

    fmt.Println("Similarity:", similarity)
}

