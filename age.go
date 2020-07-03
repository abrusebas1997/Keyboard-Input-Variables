package main

import (
  "fmt"
  "os"
  "bufio"
  "time"
  "strconv"
)

func main() {
  t := time.Now()
  month := t.Month()
  day := t.Day()
  y := t.Year()
  fmt.Println("current mm/dd/yyyy:", y, int(month), day)

  your_month := bufio.NewReader(os.Stdin)
  fmt.Print("Enter you month of birth:")

  your_month_, _ := your_month.ReadString('\n')
  fmt.Print("month: " + your_month_)

  your_day := bufio.NewReader(os.Stdin)
  fmt.Print("Enter you day of birth:")

  your_day_, _ := your_day.ReadString('\n')
  fmt.Print("day: " + your_day_)

  your_age := bufio.NewReader(os.Stdin)
  fmt.Print("Enter you age:")

  your_age_, _ := your_age.ReadString('\n')
  // fmt.Print("age: " + your_age_)
  yearInt, _ := strconv.Atoi(your_age_)
  current := int(y) - yearInt
  fmt.Print(current)
}
