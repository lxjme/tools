package color_output

import "testing"

func TestPrintColor(t *testing.T) {
	Colorful.WithBackColor("green").WithFrontColor("red").Println("焦喜亮")
}
