package main

import (
	"container/list"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Test Part 1:", partOne("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 1:", partOne("input.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test Part 2:", partTwo("test.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Test 2 Part 2:", partTwo("test2.txt"), "time spent", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", partTwo("input.txt"), "time spent", time.Since(start))
}

type Pos struct {
	x        int
	y        int
	parent   string
	cost     int
	estimate int
	dir      int
	dirCount int
	id       string
}

var dirMap = [4]Dir{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type Dir struct {
	x int
	y int
}

func makePos(world *[][]int, x int, y int, parent *Pos, dir int) Pos {
	cost := -1
	parentId := ""
	dirCount := 1
	if parent == nil {
		cost = 0
		dirCount = 0
	} else if x >= 0 && y >= 0 && x < len((*world)[0]) && y < len(*world) {
		cost = parent.cost + (*world)[y][x]
		parentId = parent.id
		if parent.dir == dir {
			dirCount += parent.dirCount
		}
	}
	return Pos{x, y, parentId, cost, heuristic(world, x, y), dir, dirCount, generateID(x, y, dir, dirCount)}
}

var checkList = list.New()
var checked = make(map[string]*list.Element)

func generateID(x int, y int, dir int, dirCount int) string {
	return fmt.Sprintf("%03dx%03d/%d/%02d", x, y, dir, dirCount)
}

func heuristic(world *[][]int, x int, y int) int {
	return (len(*world) - y) + (len((*world)[0]) - x)
}

func takeFirst() Pos {
	el := checkList.Front()
	if el == nil {
		fmt.Println("WTF?!")
	}
	val := checkList.Remove(el)
	return val.(Pos)
}

func insertToPriorityPosition(pos Pos) {
	if pos.cost < 0 {
		return
	}
	if val, ok := checked[pos.id]; ok {
		if val.Value.(Pos).cost <= pos.cost {
			return
		}

		checkList.Remove(val)
	}
	for e := checkList.Front(); e != nil; e = e.Next() {
		value := e.Value.(Pos)
		if value.cost+value.estimate > pos.cost+pos.estimate {
			el := checkList.InsertBefore(pos, e)
			checked[pos.id] = el
			return
		}
	}
	el := checkList.PushBack(pos)
	checked[pos.id] = el
}

func createPosForDir(world *[][]int, parent *Pos, dir int) Pos {
	x := parent.x + dirMap[dir].x
	y := parent.y + dirMap[dir].y
	return makePos(world, x, y, parent, dir)
}

func drawWorld(filename string, world *[][]int, winner Pos) {
	multiplier := 3
	im := image.NewRGBA(image.Rectangle{Max: image.Point{X: len((*world)[0]) * multiplier, Y: len(*world) * multiplier}})

	for y, line := range *world {
		for x, block := range line {
			value := uint8(255 / (10 - block))
			for i := 0; i < multiplier; i++ {
				for j := 0; j < multiplier; j++ {
					im.SetRGBA(x*multiplier+j, y*multiplier+i, color.RGBA{A: 255, R: value, G: value, B: value})
				}
			}
		}
	}

	for {
		im.SetRGBA(winner.x*multiplier+1, winner.y*multiplier+1, color.RGBA{A: 255, R: 255})
		if winner.parent != "" {
			winner = checked[winner.parent].Value.(Pos)
		} else {
			break
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, im)
}

func partOne(filename string) string {
	return solve(filename, "one", 0, 3)
}
func solve(filename string, id string, minimum int, maximum int) string {
	checkList.Init()
	checked = make(map[string]*list.Element)

	lines := readLines(filename)
	world := make([][]int, len(lines))
	for i, line := range lines {
		world[i], _ = sliceAtoi(strings.Split(line, ""))
	}

	insertToPriorityPosition(makePos(&world, 0, 0, nil, -1))
	loops := 0

	// drawWorld(filename+"-"+id+"-initial.png", &world, checkList.Front().Value.(Pos))

	for {
		element := takeFirst()
		if element.x == len(world[0])-1 && element.y == len(world)-1 {
			if element.dirCount < minimum {
				continue
			}
			fmt.Println("FOUND GOAL!", loops)
			drawWorld(filename+"-"+id+"-final.png", &world, element)
			current := element
			total := 0
			for {
				if current.parent == "" {
					return strconv.Itoa(total)
				}
				total += world[current.y][current.x]
				current = checked[current.parent].Value.(Pos)
			}
		}
		for i := 0; i < 3; i++ {
			dir := 1
			if element.dir >= 0 {
				dir = element.dir
			}
			if element.dirCount < maximum {
				insertToPriorityPosition(createPosForDir(&world, &element, dir))
			}
			if element.dirCount >= minimum || element.dirCount == 0 {
				insertToPriorityPosition(createPosForDir(&world, &element, (dir+1)%4))
				insertToPriorityPosition(createPosForDir(&world, &element, (dir+3)%4))
			}
		}

		/*
			for e := checkList.Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
		*/
		loops++
		/*
			if loops > 2000 {
				for e := checkList.Front(); e != nil; e = e.Next() {
					fmt.Println(e.Value)
				}
				printWorld(world, checkList.Front().Value.(Pos))
				return "fail"
			}
		*/
	}
}

func partTwo(filename string) string {
	return solve(filename, "two", 4, 10)
}
