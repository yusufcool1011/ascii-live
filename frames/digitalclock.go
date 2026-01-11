package frames

import "time"

var DigitalClock = FrameType{
	GetFrame:  getDigitalClockFrame,
	GetLength: getDigitalClockLength,
	GetSleep:  DefaultGetSleep(),
}

var digits = map[rune][]string{
	'0': {" █████ ", "██   ██", "██   ██", "██   ██", " █████ "},
	'1': {"   ██  ", " ████  ", "   ██  ", "   ██  ", " █████ "},
	'2': {" █████ ", "     ██", " █████ ", "██     ", "███████"},
	'3': {"██████ ", "     ██", " █████ ", "     ██", "██████ "},
	'4': {"██   ██", "██   ██", "███████", "     ██", "     ██"},
	'5': {"███████", "██     ", "██████ ", "     ██", "██████ "},
	'6': {" █████ ", "██     ", "██████ ", "██   ██", " █████ "},
	'7': {"███████", "     ██", "    ██ ", "   ██  ", "   ██  "},
	'8': {" █████ ", "██   ██", " █████ ", "██   ██", " █████ "},
	'9': {" █████ ", "██   ██", " ██████", "     ██", " █████ "},
	':': {"       ", "   ██  ", "       ", "   ██  ", "       "},
}

func getDigitalClockFrame(i int) string {
	now := time.Now().Format("15:04:05")
	lines := make([]string, 5)

	for _, ch := range now {
		glyph := digits[ch]
		for i := 0; i < 5; i++ {
			lines[i] += glyph[i] + "  "
		}
	}

	out := "\n"
	for _, line := range lines {
		out += line + "\n"
	}
	return out
}

func getDigitalClockLength() int {
	return 0 // infinite
}
