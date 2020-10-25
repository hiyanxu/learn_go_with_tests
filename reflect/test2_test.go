package reflect

import (
	"fmt"
	"testing"
)

func TestGetKind(t *testing.T) {
	ks := getKind()

	// 输出：[int ptr struct string]
	fmt.Println(ks)
}
