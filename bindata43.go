package ko

import(
	"os"
	"time"
)


func dictKodictBrBytes() ([]byte, error) {
	return _dictKodictBr, nil
}

func dictKodictBr() (*asset, error) {
	bytes, err := dictKodictBrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.br", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
