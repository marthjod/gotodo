// Code generated by "stringer -type=Priority"; DO NOT EDIT.

package priority

import "fmt"

const _Priority_name = "ABCDEFGHIJKLMNOPQRSTUVWXYZNone"

var _Priority_index = [...]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 30}

func (i Priority) String() string {
	if i < 0 || i >= Priority(len(_Priority_index)-1) {
		return fmt.Sprintf("Priority(%d)", i)
	}
	return _Priority_name[_Priority_index[i]:_Priority_index[i+1]]
}
