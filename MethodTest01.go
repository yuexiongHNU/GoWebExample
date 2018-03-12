package main

import . "fmt"

const (
	WHITE = iota //WHITE = 0 每次const出现时初始化为0，const中每新增一行，iota新增1
	BLACK //BLACK = 1
	BLUE //BLUE = 2
	RED //RED = 3
	YELLOW
)

type Color byte

type Box struct{
	width, height, depth float64
	color Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (b1 BoxList) BiggestsColor() Color {
	v := 0.0
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
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

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
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
	Printf("%d \n", len(boxes))
	Printf("\nFirst box's volume: %f", boxes[0].Volume())
	Printf("\nLast one's color: %s", boxes[len(boxes)-1].color.String())
	Printf("\nBiggest one is: %s", boxes.BiggestsColor().String())

	boxes.PaintItBlack()
	Printf("\n2nd box's color: %s", boxes[1].color.String())

	Printf("\nNow, biggest one is: %s", boxes.BiggestsColor().String())
}