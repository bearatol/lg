// I set the color of the text.
// The colour will not change in the Windows operating system.
package color

import "runtime"

type color struct {
	Reset,
	Red,
	RedBold,
	RedBoldUnderline,
	Green,
	Yellow,
	Blue,
	BlueUnderline,
	Purple,
	PurpleUnderline,
	Cyan,
	CyanBold,
	Gray,
	GrayBold,
	GrayUnderline string
}

var Color color

func init() {
	if runtime.GOOS != "windows" {
		Color = color{
			Reset:           "\033[0;0;0m",
			Red:             "\033[0;0;31m",
			Yellow:          "\033[0;0;33m",
			Purple:          "\033[0;0;35m",
			PurpleUnderline: "\033[0;4;35m",
			Cyan:            "\033[0;0;36m",
			Gray:            "\033[0;0;37m",
			GrayUnderline:   "\033[0;4;37m",
		}
	}
}
