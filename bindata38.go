package ko

import(
	"os"
	"time"
)


func dictKodictBmBytes() ([]byte, error) {
	return _dictKodictBm, nil
}

func dictKodictBm() (*asset, error) {
	bytes, err := dictKodictBmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
