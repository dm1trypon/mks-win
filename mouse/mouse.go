package mouse

/*
extern void move(int x, int y);
extern void leftPress(int x, int y);
extern void leftRelease(int x, int y);
extern void rightPress(int x, int y);
extern void rightRelease(int x, int y);
*/
import "C"

const (
	LeftButton  = "LEFT_BUTTON"
	RightButton = "RIGHT_BUTTON"
)

func Move(posX, posY int) {
	C.move(C.int(posX), C.int(posY))
}

func Press(posX, posY int, button string) {
	if button == LeftButton {
		C.leftPress(C.int(posX), C.int(posY))
	} else if button == RightButton {
		C.rightPress(C.int(posX), C.int(posY))
	}
}

func Release(posX, posY int, button string) {
	if button == LeftButton {
		C.leftRelease(C.int(posX), C.int(posY))
	} else if button == RightButton {
		C.rightRelease(C.int(posX), C.int(posY))
	}
}
