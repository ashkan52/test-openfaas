package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("ashkan52 - ba ba blackship ....9 You said: %s", string(req))
}
