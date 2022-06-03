package gobotapi

import (
	"errors"
	"os"
)

func (ctx *Client) DownloadFile(fileId, filePath string) error {
	bytes, err := ctx.DownloadBytes(fileId)
	if err != nil {
		return err
	}
	if res, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("path not found")
	} else {
		_ = os.WriteFile(filePath, bytes, res.Mode().Perm())
	}
	return nil
}
