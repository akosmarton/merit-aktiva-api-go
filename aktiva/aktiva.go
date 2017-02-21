package aktiva

import (
	"fmt"
	"time"
)

const URLEstonia = "https://aktiva.merit.ee/api/v1/"
const URLFinland = "https://aktiva.meritaktiva.fi/api/v1/"
const URLPoland = "https://program.360ksiegowosc.pl/api/v1/"

type Aktiva struct {
	apiUrl string
	apiId  string
	apiKey string
}

func NewAktiva(apiUrl, apiId string, apiKey string) *Aktiva {
	return &Aktiva{
		apiUrl: apiUrl,
		apiId:  apiId,
		apiKey: apiKey,
	}
}

func TimeToString(t time.Time) string {
	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}
