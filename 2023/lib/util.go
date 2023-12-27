package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
)

func readLines(filename string) []string {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	lines := []string{}

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	return lines
}

func readLinesAsByteArray(filename string) [][]byte {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	lines := [][]byte{}

	for fileScanner.Scan() {
		lines = append(lines, []byte(fileScanner.Text()))
	}

	readFile.Close()

	return lines
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

type Coord struct {
	x int
	y int
}

func shoelace(polygon []Coord) int {
	red := 0
	blue := 0
	for i := 0; i < len(polygon); i++ {
		red += (polygon[i].x * polygon[(i+1)%len(polygon)].y)
		blue += (polygon[i].y * polygon[(i+1)%len(polygon)].x)
	}

	return max(red-blue, blue-red) / 2
}

type Buffer struct {
	data   []byte
	width  int
	height int
	offset Coord
}

func drawDot(img *image.RGBA, pos Coord, col color.RGBA, multiplier int) {
	margin := 0
	if multiplier > 2 {
		margin = 1
	}

	for y := 0; y < multiplier; y++ {
		for x := 0; x < multiplier; x++ {
			if x < margin || y < margin {
				img.SetRGBA(pos.x*multiplier+x, pos.y*multiplier+y, palette[BLACK])
			} else {
				img.SetRGBA(pos.x*multiplier+x, pos.y*multiplier+y, col)
			}
		}
	}
}

func drawLine(buffer Buffer, p1 Coord, p2 Coord, col byte) {
	xdiff := max(p2.x-p1.x, p1.x-p2.x)
	ydiff := max(p2.y-p1.y, p1.y-p2.y)

	xt := float64(p2.x-p1.x) / float64(max(xdiff, ydiff))
	yt := float64(p2.y-p1.y) / float64(max(xdiff, ydiff))

	for i := 0; i < max(xdiff, ydiff); i++ {
		x := p1.x + int(math.Round(float64(i)*xt))
		y := p1.y + int(math.Round(float64(i)*yt))
		//img.SetRGBA(x, y, col)
		//drawDot(img, Coord{x, y}, col, multiplier)
		buffer.data[y*buffer.width+x] = col
	}
}

func mapOffset(orig Coord, offset Coord) Coord {
	return Coord{orig.x + offset.x, orig.y + offset.y}
}

func findDimensions(corners []Coord) (int, int, Coord) {
	minimum := Coord{0, 0}
	maximum := Coord{0, 0}

	for _, c := range corners {
		minimum.x = min(minimum.x, c.x)
		minimum.y = min(minimum.y, c.y)
		maximum.x = max(maximum.x, c.x)
		maximum.y = max(maximum.y, c.y)
	}

	width := maximum.x - minimum.x + 1
	height := maximum.y - minimum.y + 1
	offset := Coord{-minimum.x, -minimum.y}
	return width, height, offset
}

func renderToBuffer(corners []Coord) Buffer {

	width, height, offset := findDimensions(corners)

	data := make([]byte, width*height)
	for i := 0; i < width*height; i++ {
		data[i] = 0
	}

	buffer := Buffer{data, width, height, offset}

	for i := 0; i < len(corners); i++ {
		p1 := mapOffset(corners[i], offset)
		p2 := mapOffset(corners[(i+1)%len(corners)], offset)
		drawLine(buffer, p1, p2, WHITE)
	}
	return buffer
}

func checkBlock(buffer Buffer, pos Coord, color byte) bool {
	return buffer.data[pos.y*buffer.width+pos.x] == color
}

func setBlock(buffer Buffer, pos Coord, color byte) {
	buffer.data[pos.y*buffer.width+pos.x] = color
}

func floodFill(buffer Buffer, startPos Coord, target byte, replacement byte) {

	if !checkBlock(buffer, startPos, target) {
		return
	}

	setBlock(buffer, startPos, replacement)

	stack := []Coord{startPos}

	for len(stack) > 0 {
		l := len(stack) - 1
		item := stack[l]
		stack = stack[:l]

		for i := 0; i < 4; i++ {
			dir := dirList[i]
			pos := Coord{dir.x + item.x, dir.y + item.y}
			if checkBlock(buffer, pos, target) {
				setBlock(buffer, pos, replacement)
				stack = append(stack, pos)
			}
		}
	}
}
func drawWorld(filename string, buffer Buffer, multiplier int) {
	im := image.NewRGBA(image.Rectangle{Max: image.Point{X: buffer.width * multiplier, Y: buffer.height * multiplier * expansion}})

	for x := 0; x < buffer.width; x++ {
		for y := 0; y < buffer.height; y++ {
			drawDot(im, Coord{x, y}, palette[buffer.data[y*buffer.width+x]], multiplier)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png.Encode(f, im)
}

func countNonBlack(buffer Buffer) int {
	total := 0
	for _, c := range buffer.data {
		if c > 0 {
			total += 1
		}
	}
	return total
}

var palette = []color.RGBA{
	{A: 255},
	{A: 255, R: 255},
	{A: 255, G: 255},
	{A: 255, B: 255},
	{A: 255, R: 128, G: 128, B: 128},
	{A: 255, R: 255, G: 255, B: 255},
}

const BLACK = 0
const RED = 1
const GREEN = 2
const BLUE = 3
const GRAY = 4
const WHITE = 5
