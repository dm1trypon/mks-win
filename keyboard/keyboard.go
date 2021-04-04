package keyboard

/*
extern void keyDown(short key);
extern void keyUp(short key);
extern short getKeyCode(char key);
*/
import "C"

func KeyPress(key string) {
	if len(key) < 1 {
		return
	}

	code := GetKeyCode(key)

	C.keyDown(C.short(code))
	C.keyUp(C.short(code))
}

func ComboKeyPress(keys []string) {
	for _, key := range keys {
		if len(key) < 1 {
			continue
		}

		code, ok := codes[key]
		if !ok {
			code = GetKeyCode(key)
		}

		C.keyDown(C.short(code))
	}

	for _, key := range keys {
		code, ok := codes[key]
		if !ok {
			code = GetKeyCode(key)
		}

		C.keyUp(C.short(code))
	}
}

func GetKeyCode(key string) int16 {
	return int16(C.getKeyCode(*C.CString(string([]rune(key)[0]))))
}
