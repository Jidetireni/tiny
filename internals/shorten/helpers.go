package shorten

// shuffled base62 alphabet — reorder this to change the output appearance.
// Sequential IDs (0,1,2…) will map to visually non-sequential codes.
const alphabet = "7fGHkLmNpQrStUvWxYzAaBbCcDdEeFgHhIiJjKlMnOoPqRsT2u3V4w5X6y"

func encode(id int64) string {
	base := int64(len(alphabet))
	if id == 0 {
		return string(alphabet[0])
	}
	result := make([]byte, 0, 7)
	for id > 0 {
		result = append([]byte{alphabet[id%base]}, result...)
		id /= base
	}
	return string(result)
}
