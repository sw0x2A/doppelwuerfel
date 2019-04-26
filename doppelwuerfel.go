package doppelwuerfel

type ColumnarTransposition struct {
	psk []byte
}

func NewColumnarTransposition(psk []byte) *ColumnarTransposition {
	return &ColumnarTransposition{psk}
}

// Encipher enciphers string using ROT cipher with alphabet according to key.
func (r *ColumnarTransposition) Encrypt(text []byte) []byte {
	return encrypt(text, r.psk)
}

// Decipher deciphers string using ROT cipher with alphabet according to key.
func (r *ColumnarTransposition) Decrypt(text []byte) []byte {
	return decrypt(text, r.psk)
}

type Doppelwürfel struct {
	psk1ct *ColumnarTransposition
	psk2ct *ColumnarTransposition
}

func NewDoppelwürfel(psk1 []byte, psk2 []byte) *Doppelwürfel {
	return &Doppelwürfel{NewColumnarTransposition(psk1), NewColumnarTransposition(psk2)}
}

func (r *Doppelwürfel) Encrypt(text []byte) []byte {
	return r.psk2ct.Encrypt(r.psk1ct.Encrypt(text))
}

func (r *Doppelwürfel) Decrypt(text []byte) []byte {
	return r.psk1ct.Decrypt(r.psk2ct.Decrypt(text))
}

func permutation(s []byte) []int {
	p := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		p[i] = 0
		for j := 0; j < len(s); j++ {
			if s[j] < s[i] || (s[j] == s[i] && j < i) {
				p[i]++
			}
		}
	}
	return reverse(p)
}

func reverse(perm []int) []int {
	p := make([]int, len(perm))
	for i := 0; i < len(p); i++ {
		p[perm[i]] = i
	}
	return p
}

func encrypt(message []byte, psk []byte) []byte {
	sortedIdx := permutation(psk)

	output := make([]byte, 0, len(message))
	for pski := 0; pski < len(psk); pski++ {
		for j := 0; j <= int(len(message)/len(psk)); j++ {
			idx := j*len(psk) + sortedIdx[pski]
			if idx < len(message) {
				output = append(output, message[idx])
			}

		}
	}
	return output
}

func decrypt(cipher []byte, psk []byte) []byte {
	sortedIdx := permutation(psk)

	output := make([]byte, len(cipher))

	columnCount := len(psk)
	rowCount := len(cipher) / columnCount

	counter := 0

	for i := 0; i < columnCount; i++ {
		for j := 0; j <= rowCount; j++ {
			idx := j*columnCount + sortedIdx[i]
			if idx < len(cipher) {
				output[idx] = cipher[counter]
				counter++
			}
		}
	}

	return output
}
