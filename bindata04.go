package ko

import(
	"os"
	"time"
)


func dictKodictAeBytes() ([]byte, error) {
	return _dictKodictAe, nil
}

func dictKodictAe() (*asset, error) {
	bytes, err := dictKodictAeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ae", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
