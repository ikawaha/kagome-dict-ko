package ko

import(
	"os"
	"time"
)


func dictKodictAdBytes() ([]byte, error) {
	return _dictKodictAd, nil
}

func dictKodictAd() (*asset, error) {
	bytes, err := dictKodictAdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ad", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
