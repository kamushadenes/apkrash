package apk

import "encoding/xml"

func ParseAndroidManifest(manifestText []byte) (*Manifest, error) {
	var manifest Manifest
	err := xml.Unmarshal(manifestText, &manifest)

	return &manifest, err
}
