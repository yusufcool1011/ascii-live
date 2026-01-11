package frames

import (
	"strings"
	"time"
)

var DigitalClock = FrameType{
	GetFrame:  getDigitalClockFrame,
	GetLength: getDigitalClockLength,
	GetSleep:  DefaultGetSleep(),
}

var digits = map[rune][]string{
	'0': {
		" █████ ",
		"██   ██",
		"██   ██",
		"██   ██",
		" █████ ",
	},
	'1': {
		"   ██  ",
		" ████  ",
		"   ██  ",
		"   ██  ",
		" █████ ",
	},
	'2': {
		" █████ ",
		"     ██",
		" █████ ",
		"██     ",
		"███████",
	},
	'3': {
		"██████ ",
		"     ██",
		" █████ ",
		"     ██",
		"██████ ",
	},
	'4': {
		"██   ██",
		"██   ██",
		"███████",
		"     ██",
		"     ██",
	},
	'5': {
		"███████",
		"██     ",
		"██████ ",
		"     ██",
		"██████ ",
	},
	'6': {
		" █████ ",
		"██     ",
		"██████ ",
		"██   ██",
		" █████ ",
	},
	'7': {
		"███████",
		"     ██",
		"    ██ ",
		"   ██  ",
		"   ██  ",
	},
	'8': {
		" █████ ",
		"██   ██",
		" █████ ",
		"██   ██",
		" █████ ",
	},
	'9': {
		" █████ ",
		"██   ██",
		" ██████",
		"     ██",
		" █████ ",
	},
	':': {
		"       ",
		"   ██  ",
		"       ",
		"   ██  ",
		"       ",
	},
}

func getDigitalClockFrame(i int) string {
	// Eastern Time (handles EST/EDT automatically)
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		loc = time.UTC
	}

	now := time.Now().In(loc).Format("15:04:05")

	lines := make([]string, 5)
	for _, ch := range now {
		glyph := digits[ch]
		for j := 0; j < 5; j++ {
			lines[j] += glyph[j] + "  "
		}
	}

	// Simple, terminal-safe padding
	topPadding := 5
	leftPadding := 4

	var out strings.Builder
	for i := 0; i < topPadding; i++ {
		out.WriteString("\n")
	}

	for _, line := range lines {
		out.WriteString(strings.Repeat(" ", leftPadding))
		out.WriteString(line)
		out.WriteString("\n")
	}

	// Optional label
	out.WriteString("\n")
	out.WriteString(strings.Repeat(" ", leftPadding+10))
	out.WriteString("ET\n")

	return out.String()
}

func getDigitalClockLength() int {
	return 0 // infinite animation
}
