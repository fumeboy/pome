package wrong

import (
	"fmt"
)

type E struct {
	Code    string
	Message string
}

func (this *E) Error() string {
	return fmt.Sprintf("pomi error, code:%v message:%v", this.Code, this.Message)
}
