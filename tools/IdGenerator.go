package tools

import (
	"fmt"
)

var sidstore = make(map[string]uint)

func GenSID(namespace string) string {
	id, success := sidstore[namespace]
	if !success {
		sidstore[namespace] = 1
		id = 1
	} else {
		id++
		sidstore[namespace] = id
	}
	return fmt.Sprintf("%s%d", namespace, id)
}

func SetLastSID(namespace string, id uint) {
	sidstore[namespace] = id
}

func TestID(namespace string, id uint) {

}

func TestSID(namespace string, id uint) {
}
