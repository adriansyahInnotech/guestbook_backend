package utils

import (
	"encoding/json"
	"strings"
	"time"
)

const layoutISO = "2-1-2006"

// CustomDate adalah tipe tanggal kustom untuk parsing "2006-01-02"
type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		cd.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Format(layoutISO))
}
