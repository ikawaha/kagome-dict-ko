package ko

import(
	"os"
	"time"
)


func dictKodictBpBytes() ([]byte, error) {
	return _dictKodictBp, nil
}

func dictKodictBp() (*asset, error) {
	bytes, err := dictKodictBpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
