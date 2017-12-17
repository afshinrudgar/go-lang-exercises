package tournament

import (
  "strings"
  "bufio"
  "sort"
  "fmt"
  "io"
  "errors"
)

type Team struct {
  Name string
  MP, W, D, L, P int
}

func (t *Team) win() {
  t.MP++
  t.W++
  t.P += 3
}

func (t *Team) loss() {
  t.MP++
  t.L++
}

func (t *Team) draw() {
  t.MP++
  t.D++
  t.P++
}

type TeamSorter []*Team
func (a TeamSorter) Len() int { return len(a) }
func (a TeamSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a TeamSorter) Less(i, j int) bool {
  a1, a2 := a[i], a[j]
  if a1.P > a2.P {
    return true
  }
  return a1.P == a2.P && a1.Name < a2.Name
}

func Tally(in io.Reader, out io.Writer) error {

  scanner := bufio.NewScanner(in)
  table := make(map[string]*Team)

  var p1, p2, status string
  for scanner.Scan() {
    game := scanner.Text()
    if game == "" || game[0] == '#' {
      continue
    }
    objs := strings.Split(game, ";")
    if len(objs) != 3 {
      return errors.New("Invalid number of parameters")
    }
    p1, p2, status = objs[0], objs[1], objs[2]

    t1, ok := table[p1]
    if !ok {
      table[p1] = &Team{Name: p1}
      t1 = table[p1]
    }

    t2, ok := table[p2]
    if !ok {
      table[p2] = &Team{Name: p2}
      t2 = table[p2]
    }

    if status == "win" {
      t1.win()
      t2.loss()
    } else if status == "draw" {
      t1.draw()
      t2.draw()
    } else if status == "loss" {
      t1.loss()
      t2.win()
    } else {
      return errors.New("Invalid status")
    }

  }

  list := make([]*Team, 0, len(table))
  for _, t := range table {
    list = append(list, t)
  }

  sort.Sort(TeamSorter(list))
  fmt.Fprintf(out, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P");
  for _, t := range list {
    fmt.Fprintf(out, "%-31s|%3d |%3d |%3d |%3d |%3d\n", t.Name, t.MP, t.W, t.D, t.L, t.P);
  }
  return nil
}
