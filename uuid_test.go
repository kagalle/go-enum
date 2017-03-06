package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUuid(t *testing.T) {
	uuid1, err := NewUUID()
	if assert.NoError(t, err) {
		uuid2, err := NewUUID()
		if assert.NoError(t, err) {
			fmt.Printf("%s\n%s\n", uuid1, uuid2)
			assert.NotEqual(t, uuid1, uuid2)
		}
	}
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
