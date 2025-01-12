package component

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"testing"
)

type User struct {
	Name      string
	FansCount int64
}

func TestTemp(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := []byte("320a8f00-cf99-408c-bd33-f0126faa") // 32 bytes key for AES-256
			plaintext := []byte("{\"version\":1,\"componentName\":\"tchousex\",\"eventId\":\"53cded2a-184c-437e-8d47-f4fb7515abe9\",\"timestamp\":1735202587,\"interface\":{\"interfaceName\":\"logic.cam.sigAndAuth\",\"para\":{\"version\":\"2.0\",\"url\":\"tchousex.tencentcloudapi.com/?\",\"header\":\"content-type:application/json\\nhost:tchousex.tencentcloudapi.com\\nx-tc-action:ConnectInstance\\nauthorization:q-sign-algorithm=api_sha256_v4&q-ak=AKIDPLaV6fkYomRQVA6FTgW212O9r1W3TraCmWUsfHmPuxO3c1uW51o1MKdV3Os8_KDu&q-sign-time=1735202287;1735202887&q-key-time=1735202287;1735202887&q-header-list=content-type;host;x-tc-action&q-url-param-list=&q-signature=a5064540764b17c9dbf156113ada7ae17aa1481d94ff17489c368f5459942ea6&q-sign-time-api-v4=1735202587;2024-12-26\\nq-payload:44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a\",\"method\":\"POST\",\"mode\":0,\"action\":\"tchousex:ConnectInstance\",\"resource\":\"*\",\"condition\":{},\"q-token\":\"r5J6BP30FHDwNPqd3gdqqGR7gHnjgqtHb03de3e22bcc40bdd9bb3ffefc55119aRHjfLo8uwTtyXU_Bn8No1W7TcK9qnj5SmgONmSh2w97d96007mmUSEqMVcxo70NiOyLzrELsb3TKPPC0JC_GvVq77o2T5jPJBwVsLIFxsYYXRDO_vEn2IZAyYIwhN8H-1f2_4HCtubd3oNN3VMNrPDG7xiU-vaWfTQehOuLKzH_o-rTRgdQlvJcJg-4LxT4thd07nJW62Xd3qq2d1iHeYfnIybeMs4cXaVDG5fmHrZMF2-6viTVjSC-2LpHgxaYul6oG1z9vaY-RFOps-ifIwuL_7axlj8zk2HWzsnhi-ACdafy3cRXKjaWvODI4VtCK\",\"camContext\":{\"clientIP\":\"192.168.1.1\"}}}}")

			// Encrypt
			ciphertext, err := AESEncrypt(plaintext, key)
			if err != nil {
				fmt.Println("Error encrypting:", err)
				return
			}
			fmt.Println("Encrypted:", ciphertext)

			// Decrypt
			decrypted, err := AESDecrypt("MzIwYThmMDAtY2Y5OS00MDmA5+Ze2+2YFGojmMRbrU5KOWnx6ni+BSmLABXWiyj9Z1ZnvPyxIh+dllDpgGHvqbchoX1SfFQVZzyUd3VkHeaQpM068wJrxbUkQPBo8bpuTeBmcUyEs1C8IUHs4G7zD3Ptu1wjLzxQJRTNkIvp2eM4GyhNLEnNlstsIR2ukUcjEyBNW9ZXyzx37xy43TS8fUsqEIE8EPnUBBNRNscewpDqeMuWPYf5G32mwoRVQfOZkbv3lFrfGhiwjC/voOcWOS8CSjHi2UcEVJU2S7Z+Z+PDcXyZFHCX2CrWwYF+8DdyLx8pHWZvbDtuG449BZ/FZ15uLDwUlHjl/wSfRtLrGVEeQizfWZXH6g0UKG66zR+HItlMVGYV0PCVfDdq+IOXWHLQhcLE1YiOqO+uVaVz3IT0eSwmkAX8IOLiUOm2W+oJb/PhtAiGa3ZQp/jgTvZpTes63GKZRkKCdUWpmIkXV1P0mtxucm9yo8E1jmqtQaoPUXLXQXJ41fRVRI13OYP7dw6e0JxFO1tnPXXxb5Z2uc6pwH2bFIhy5kC/4ASD3y2ZOQ5/CmvhmKh5BCBpacadJtwvppxlA4utCYOOqY8/FSK8ZVCtG0nXasNepCD1Ygkzv1TPb33gNR+LDqClZv29d7jU+4e0iqX4gkbpgRgZ8lWY8yVlcU01RJK9GE+qBcRQKhsniCfftbtrNDX7t5ZDv1kxr/rs2ahuUx5mwpJY3DaYIbSCOg8AYFKmbo3tvLcuTo4WlLHhycZBm37vWB0gWTOMBSpFdO70Jw7f78gEeDs3i4U3DYrCeULOXgoDHx2aEVsCO0sawXpNr+htp6vwals410smo8X8Rx3ZSUUVmVMlQoU7De5S59lqU0NEgmtoWmDjaZ3HCKXxDamEH3G7WMuScklbi/T+sgis9+D+DYd8XuYqkjy1h3MMeIVU9ZNmd3qBEBUoYPsfUzV2TuahDngLBHtIu2B+Fle/cT+srkxCZZ0Jm4SE8AgNvm7uuNhTttnkns/7dQbUyTO9BN/FdeVIollTPHFiBvj1LuEAAmAlJPVInaIvnzoc7ImXPwxUPVVndz7BHDMYHelBEs8AO1Yhoe6ewci4cB++IWKZEGNP11CLWt+7c3CAQAeX9Io1KUjdvS6VEoP/WeL11fNxny2t1hFFMjX6Pv6o+BlbXi5sy25O14tuxEAaI4KcxlDK1ta6xmH4MqfoT4sPVGDy5MvfBUlIuBPmZAHDyFbocMuLrVuFkE0dlQH5jXoWvRdj2E2+hUPMFbQZIIBRWJCs1bqPVRICTve9qCOD0SsEnWODG04ZdHfI0DI7w2y8gUepF8Xk7lZJiffVxuqUGgORC/UHNB6+LOAW9p1YZK2s6+YxGi+QMnc/P+QfOylQuUXv2a9YJ3U6gSQaWnMy6SkNIQPxS65T6t0Xd9TD0gSBCD+CDoXidmhakHHZc+3Ykutx4kNlgUzP9hEORK/7JRzxEZKuzTDJ811ov+9k4xl/puKWoF4LSPygUkgb5eQYbzgDgq3kHzkM1aHK+yHMg9o8Xvs9di3VjbWo7YGZFJG9G5Un610E6ddLtFpVkyqvXr1CWbuyO9J9/AFeIfSIyDgYcVQDCrXRzBAT7KGd1Uib5MyKf78CZN8O+/xSAGDExRmPdTqnwH8qGG9zgdhEYMPSApAASaunS4o5djawsXpsb0Ocq4xcttHmapil3+/JzGdUi1wcbVQcbgLEhSLABkTtDYgp5IUiVjB0dUxfkOiTZnPk1P6UZ1EOVOlZU7WY0UKc", key)
			if err != nil {
				fmt.Println("Error decrypting:", err)
				return
			}
			fmt.Println("Decrypted:", string(decrypted))
		})
	}
}

type Blog struct {
	BlogId  string `mapstructure:"blogId"`
	Title   string `mapstructrue:"title"`
	Content string `mapstructure:"content"`
	Uid     string `mapstructure:"uid"`
	State   string `mapstructure:"state"`
}

type Event struct {
	Type     string              `json:"type"`
	Database string              `json:"database"`
	Table    string              `json:"table"`
	Data     []map[string]string `json:"data"`
}

func TestMapStructure(t *testing.T) {
	e := Event{}
	msg := []byte(`{
     "type": "UPDATE",
     "database": "blog",
     "table": "blog",
     "data": [
          {
               "blogId": "100001",
               "title": "title",
               "content": "this is a blog",
               "uid": "1000012",
               "state": "1"
          },
          {
               "blogId": "100002",
               "title": "title",
               "content": "this is a blog",
               "uid": "1000012",
               "state": "2"
          }
     ]
}`)
	if err := json.Unmarshal(msg, &e); err != nil {
		panic(err)
	}
	if e.Table == "blog" {
		var blogs []Blog
		if err := mapstructure.Decode(e.Data, &blogs); err != nil {
			panic(err)
		}
		//fmt.Printf("%+v\n", blogs)
	}

	type Address struct {
		City    string
		Country string
	}
	type Person struct {
		Name    string
		Age     int
		Address *Address
	}

	var i interface{}
	i = &Person{Name: "Alice", Age: 30, Address: &Address{City: "New York", Country: "USA"}}
	fmt.Printf("%+v\n", i) // 输出: {Name:Alice Age:30 Address:{City:New York Country:USA}}
}

// PKCS7Padding pads the plaintext to be a multiple of the block size
func PKCS7Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// PKCS7UnPadding removes the padding from the plaintext
func PKCS7UnPadding(plaintext []byte) ([]byte, error) {
	length := len(plaintext)
	if length == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	unpadding := int(plaintext[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("invalid padding size")
	}
	return plaintext[:(length - unpadding)], nil
}

// AESEncrypt encrypts plaintext using AES-256-CBC with the given key
func AESEncrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext = PKCS7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecrypt decrypts ciphertext using AES-256-CBC with the given key
func AESDecrypt(ciphertext string, key []byte) ([]byte, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	if len(ciphertextBytes)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	plaintext, err := PKCS7UnPadding(ciphertextBytes)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
