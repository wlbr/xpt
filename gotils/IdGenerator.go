package gotils

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

func LastSID(namespace string) string {
	id, success := sidstore[namespace]
	if !success {
		id = 0
	}
	return fmt.Sprintf("%s%d", namespace, id)
}

func SetLastSID(namespace string, id uint) {
	sidstore[namespace] = id
}

func ResetSID(namespace string) {
	delete(sidstore, namespace)
}

func ResetAllSID() {
	sidstore = make(map[string]uint)
}
