package ko

import(
	"os"
	"time"
)


func dictKodictAjBytes() ([]byte, error) {
	return _dictKodictAj, nil
}

func dictKodictAj() (*asset, error) {
	bytes, err := dictKodictAjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.aj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
