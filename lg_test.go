package lg

import (
	"os"
	"testing"
)

func TestLg_print(t *testing.T) {
	t.Run("print", func(t *testing.T) {
		lg := loggerCreateNew(os.Stdout, "", "")
		if err := lg.print("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}
func TestLg_println(t *testing.T) {
	t.Run("println", func(t *testing.T) {
		lg := loggerCreateNew(os.Stdout, "", "")
		if err := lg.println("test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}
func TestLg_printf(t *testing.T) {
	t.Run("printf", func(t *testing.T) {
		lg := loggerCreateNew(os.Stdout, "", "")
		if err := lg.printf("%v,%v,%v,%v,%v,%v", "test", 1, true, struct{}{}, [0]int{}, 1.1); err != nil {
			t.Error(err)
		}
	})
}
