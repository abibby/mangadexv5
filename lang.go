package mangadexv5

type LangMap map[string]string

func (l LangMap) String() string {
	return l["en"]
}
