package store

import (
	"fmt"
	"github.com/89z/googleplay"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func DownloadGooglePlayAPK(email, password, appId, tmpDir, tmpFile string) error {
	fmt.Printf("Downloading APK %s\n", appId)

	cache, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(cache, "googleplay"), 0755)
	if err != nil {
		return err
	}

	device, err := googleplay.OpenDevice(cache, "googleplay/device.json")
	if err != nil {
		device, err = googleplay.DefaultConfig.Checkin()
		if err != nil {
			return err
		}
		time.Sleep(googleplay.Sleep)
		err = device.Create(cache, "googleplay/device.json")
		if err != nil {
			return err
		}
	}

	token, err := googleplay.OpenToken(cache, "googleplay/token.json")
	if err != nil {
		token, err = googleplay.NewToken(email, password)
		if err != nil {
			return err
		}
		err = token.Create(cache, "googleplay/token.json")
		if err != nil {
			return err
		}
	}

	header, err := token.SingleAPK(device)
	if err != nil {
		return err
	}

	_ = header.Purchase(appId)

	details, err := header.Details(appId)
	if err != nil {
		return err
	}

	delivery, err := header.Delivery(appId, int64(details.VersionCode))
	if err != nil {
		return err
	}

	out, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(delivery.DownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
