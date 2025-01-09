package files

import (
	"fmt"
	"os"
)

/*
	This code create the file if does not exit, or truncates the existing one
	before writing the content
	data is persistent only when you call fsync(fp.sync)
*/

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Sync()
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tem.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer func() {
		fp.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	err = fp.Sync()
	if err != nil {
		return err
	}

	return os.Rename(tmp, path)
}
