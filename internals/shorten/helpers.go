package shorten

// shuffled base62 alphabet — reorder this to change the output appearance.
// Sequential IDs (0,1,2…) will map to visually non-sequential codes.
// Exactly 62 unique characters, no duplicates, no missing indices
const alphabet = "7fGHkLmNpQrStUvWxYzAaBbCcDdEeFghIiJjKlMnOoPqRsT2u3V4w5X6y8901"

func base62Encode(id uint64) string {
	if id == 0 {
		return string(alphabet[0])
	}

	var hashByte []byte
	for id > 0 {
		hashByte = append(hashByte, alphabet[id%62])
		id /= 62
	}

	return string(hashByte)
}
