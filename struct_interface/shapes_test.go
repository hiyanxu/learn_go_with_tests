package struct_interface

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 20.0)
	want := 200.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestPerimeterForStruct(t *testing.T) {
	t.Run("perimeter for rectangle", func(t *testing.T) {
		rectangle := Rectangle{height: 10.0, width: 20.0}
		got := PerimeterForStruct(rectangle)
		want := 60.0

		if got != want {
			t.Errorf("got %.2f, but want %.2f", got, want)
		}
	})
}

func TestAreaForStruct(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f, but want %.2f", got, want)
		}
	}

	t.Run("area for rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			width:  10.0,
			height: 20.0,
		}
		//got := rectangle.Area()
		want := 200.0
		//
		//if got != want {
		//	t.Errorf("got %.2f, but want %.2f", got, want)
		//}
		checkArea(t, rectangle, want)
	})

	t.Run("area for circle", func(t *testing.T) {
		circle := Circle{radius: 10}
		//got := circle.Area()
		want := 314.16
		//if got != want {
		//	t.Errorf("got %.2f, but want %.2f", got, want)
		//}
		checkArea(t, circle, want)
	})
}

func TestAreaForStruct2(t *testing.T) {
	// 表格驱动测试：将测试数据和测试代码分离。
	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{shape: Rectangle{width: 10, height: 20}, want: 200, name: "Rectangle"},
		{shape: Circle{radius: 10}, want: 314.15, name: "Circle"},
		{shape: Triangle{12.0, 6.0}, want: 36.0, name: "Triangle"},
	}

	for _, tt := range areaTests {
		//got := tt.shape.Area()
		//if got != tt.want {
		//	t.Errorf("got %.2f, but want %.2f", got, tt.want)
		//}
		// %#v输出go语言语法格式的值
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v, got %.2f, but want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}
