package ko

import(
	"os"
	"time"
)


func dictKodictArBytes() ([]byte, error) {
	return _dictKodictAr, nil
}

func dictKodictAr() (*asset, error) {
	bytes, err := dictKodictArBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ar", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
