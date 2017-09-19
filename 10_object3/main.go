package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte
type Box struct {
	width, hight, depth float64
	color               Color
}
type BoxList []Box //a slice of boxes
func (b Box) Volume() float64 {
	return b.width * b.hight * b.depth
}
func (b *Box) SetColor(c Color) {
	b.color = c
}
func (b1 BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}
func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}
func (c Color) string() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}
func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of first one is ", boxes[0].Volume())
	fmt.Println("The color of the last one is ", boxes[len(boxes)-1])
	fmt.Println("The biggest one is ", boxes.BiggestColor().string())
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is ", boxes[1].color.string())
	fmt.Println("Obviously, now,the biggest one is ", boxes.BiggestColor())

}
