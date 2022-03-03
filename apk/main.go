package apk

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/kamushadenes/apkrash/store"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func ParseAndroidManifest(manifestText []byte) (*Manifest, error) {
	var manifest Manifest
	err := xml.Unmarshal(manifestText, &manifest)

	return &manifest, err
}

func ParseAPKInput(arg string, decompile bool, email string, password string) (*APK, error) {
	file, err := os.ReadFile(arg)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found:", arg)
			fmt.Println("Attempting to download from Google Play Store")

			tmpDir, err := ioutil.TempDir(os.TempDir(), "apkrash")
			if err != nil {
				return nil, err
			}

			tmpFile := path.Join(tmpDir, fmt.Sprintf("%s.apk", arg))

			err = store.DownloadGooglePlayAPK(email, password, arg, tmpDir, tmpFile)
			if err != nil {
				return nil, err
			}

			file, err = os.ReadFile(arg)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	mtype := mimetype.Detect(file)

	var apk APK

	switch strings.Split(mtype.String(), ";")[0] {
	case "text/xml":
		err = apk.ParseManifest(file)
		if err != nil {
			return nil, err
		}
	case "application/zip", "application/jar":
		fname := arg
		if strings.HasSuffix(fname, ".aab") {
			fmt.Println("Building APK using bundletool")

			tmpDir, err := ioutil.TempDir(os.TempDir(), "apkrash")
			if err != nil {
				return nil, err
			}
			defer os.RemoveAll(tmpDir)

			apksFile := path.Join(tmpDir, "aab.apks")

			var spec BundleToolSpec
			spec.SupportedABIs = []string{"armeabi-v7a"}
			spec.SupportedLocales = []string{"en", "pt-BR"}
			spec.ScreenDensity = 640
			spec.SdkVersion = 27

			specFile := path.Join(tmpDir, "spec.json")

			b, err := json.Marshal(&spec)
			if err != nil {
				return nil, err
			}

			err = ioutil.WriteFile(specFile, b, 0644)
			if err != nil {
				return nil, err
			}

			command := exec.Command("bundletool", "build-apks",
				fmt.Sprintf("--device-spec=%s", specFile),
				fmt.Sprintf("--bundle=%s", fname),
				fmt.Sprintf("--output=%s", apksFile))
			command.Stdout = os.Stderr
			command.Stderr = os.Stderr

			err = command.Run()
			if err != nil {
				return nil, err
			}

			archive, err := zip.OpenReader(apksFile)
			if err != nil {
				return nil, err
			}
			defer archive.Close()

			apkFile := path.Join(tmpDir, "base.apk")
			fname = apkFile

			for _, f := range archive.File {
				if strings.HasSuffix(f.Name, "-master.apk") {
					zipReader, err := f.Open()
					defer zipReader.Close()
					if err != nil {
						return nil, err
					}
					bb, err := ioutil.ReadAll(zipReader)
					if err != nil {
						return nil, err
					}
					err = ioutil.WriteFile(apkFile, bb, 0644)
					if err != nil {
						return nil, err
					}
					break
				}
			}
		}

		apk.Filename = fname
		err = apk.ParseZIP()
		if err != nil {
			return nil, err
		}
		if decompile {
			err = apk.Decompile()
			if err != nil {
				return nil, err
			}
			err = apk.ParseSources()
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("unsupported file type: %s", mtype.String())
	}

	return &apk, nil
}
