package draw

import (
	"context"
	"fmt"
	"strings"
)

func HandleDispatch(ctx context.Context, endpoint string, payload []byte) error {
	if strings.TrimSpace(endpoint) == "" {
		return fmt.Errorf("empty Draw endpoint")
	}
	return PostOutbound(context.Background(), endpoint, payload)
}
