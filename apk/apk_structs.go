package apk

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kamushadenes/apkrash/utils"
	"github.com/shogo82148/androidbinary"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type APKZipFile struct {
	Path             string `json:"path"`
	Filename         string `json:"filename"`
	CompressedSize   uint64 `json:"compressedSize"`
	UncompressedSize uint64 `json:"uncompressedSize"`
	CRC32            uint32 `json:"crc32"`
}

type APKSource struct {
	FullPath string   `json:"fullPath"`
	Path     string   `json:"path"`
	Filename string   `json:"filename"`
	Hash     string   `json:"hash"`
	Package  string   `json:"package"`
	Imports  []string `json:"imports"`
	Size     int64    `json:"size"`
}

func (f *APKZipFile) GetFullPath() string {
	if f.Path == "" {
		return f.Filename
	}
	return f.Path + "/" + f.Filename
}

func (f *APKSource) GetFullPath() string {
	if f.Path == "" {
		return f.Filename
	}
	return f.Path + "/" + f.Filename
}

type APK struct {
	TmpDir         string         `json:"-"`
	Filename       string         `json:"filename"`
	Manifest       *Manifest      `json:"manifest"`
	Files          []*APKZipFile  `json:"files"`
	FileStatistics map[string]int `json:"fileStatistics"`
	Sources        []*APKSource   `json:"sources"`
	Permissions    []string       `json:"permissions"`
	Services       []string       `json:"services"`
	Activities     []string       `json:"activities"`
	Receivers      []string       `json:"receivers"`
	Providers      []string       `json:"providers"`
	Features       []string       `json:"features"`
	FileSize       int64          `json:"fileSize"`
}

func (a *APK) Decompile() error {

	if a.TmpDir == "" {
		a.TmpDir, _ = ioutil.TempDir(os.TempDir(), "apkrash")
	}

	cmd := exec.Command("jadx", "-d", a.TmpDir, "--deobf", a.Filename)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (a *APK) ParseSources() error {
	files, err := utils.ListFiles(strings.Join([]string{a.TmpDir, "sources"}, "/"))
	if err != nil {
		return err
	}

	for k := range files {
		f := files[k]

		var source APKSource
		source.FullPath = f.FullPath
		source.Path = f.Path
		source.Filename = f.Name
		source.Size = f.Size

		source.Hash = f.Hash

		b, err := ioutil.ReadFile(f.FullPath)
		if err != nil {
			return err
		}
		content := string(b)

		pkg := utils.PackageRegex.FindStringSubmatch(content)
		if pkg != nil {
			source.Package = pkg[1]
		}

		imports := utils.ImportRegex.FindAllStringSubmatch(content, -1)
		if imports != nil {
			for i := range imports {
				source.Imports = append(source.Imports, imports[i][1])
			}
		}

		a.Sources = append(a.Sources, &source)

	}
	return nil
}

func (a *APK) ParseManifest(data []byte) error {
	manifest, err := ParseAndroidManifest(data)
	if err != nil {
		return err
	}

	a.Manifest = manifest

	return nil
}

func (a *APK) GetFiles() []string {
	var files []string
	for _, file := range a.Files {
		files = append(files, file.GetFullPath())
	}
	return files
}

func (a *APK) GetSources() []string {
	var sources []string
	for _, source := range a.Sources {
		sources = append(sources, source.GetFullPath())
	}
	return sources
}

func (a *APK) Analyze() {
	a.Features = a.Manifest.GetFeatures()
	a.Permissions = a.Manifest.GetPermissions()
	a.Services = a.Manifest.GetServices()
	a.Activities = a.Manifest.GetActivities()
	a.Receivers = a.Manifest.GetReceivers()
	a.Providers = a.Manifest.GetProviders()
}

func (a *APK) Compare(b *APK) *APKComparison {
	var comparison APKComparison
	comparison.APK1 = a
	comparison.APK2 = b

	comparison.CompareAll()

	return &comparison
}

func (a *APK) ParseZIP() error {
	archive, err := zip.OpenReader(a.Filename)
	if err != nil {
		return err
	}
	defer archive.Close()

	a.FileStatistics = make(map[string]int)

	for _, f := range archive.File {
		var file APKZipFile

		fields := strings.Split(f.Name, "/")

		file.Path = strings.Join(fields[:len(fields)-1], "/")
		file.Filename = fields[len(fields)-1]
		file.CRC32 = f.CRC32
		file.CompressedSize = f.CompressedSize64
		file.UncompressedSize = f.UncompressedSize64

		if f.Name == "AndroidManifest.xml" {
			zipReader, err := f.Open()
			defer zipReader.Close()
			if err != nil {
				return err
			}
			b, err := ioutil.ReadAll(zipReader)
			if err != nil {
				return err
			}

			readerAt := bytes.NewReader(b)

			xml, _ := androidbinary.NewXMLFile(readerAt)
			reader := xml.Reader()

			nb, err := ioutil.ReadAll(reader)
			if err != nil {
				return err
			}

			err = a.ParseManifest(nb)
			if err != nil {
				return err
			}
		}

		a.Files = append(a.Files, &file)

		fields = strings.Split(file.Filename, ".")
		if len(fields) >= 2 {
			ext := fields[len(fields)-1]
			if ext != "" {
				if _, ok := a.FileStatistics[ext]; ok {
					a.FileStatistics[ext]++
				} else {
					a.FileStatistics[ext] = 1
				}
			}
		}
	}

	return nil
}

func (a *APK) GetAnalysis(format string, includeFiles bool) (string, error) {
	a.Analyze()

	switch format {
	case "json":
		b, err := json.Marshal(a)
		return string(b), err
	case "json_pretty":
		b, err := json.MarshalIndent(a, "", "  ")
		return string(b), err
	case "table":
		t := table.NewWriter()
		t.AppendHeader(table.Row{"Field", "Value"})
		t.AppendRows([]table.Row{{"Package Name", a.Manifest.Package}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Platform Build Version", a.Manifest.PlatformBuildVersionCode}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Main Activity", a.Manifest.GetMainActivity()}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Features", strings.Join(a.Manifest.GetFeatures(), "\n")}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Permissions", strings.Join(a.Manifest.GetPermissions(), "\n")}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Services", strings.Join(a.Manifest.GetServices(), "\n")}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Providers", strings.Join(a.Manifest.GetProviders(), "\n")}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Activities", strings.Join(a.Manifest.GetActivities(), "\n")}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Receivers", strings.Join(a.Manifest.GetReceivers(), "\n")}})
		t.SetStyle(table.StyleLight)

		return t.Render(), nil
	case "text":
		var writer bytes.Buffer

		writer.WriteString(white(bold("Package: ")))
		writer.WriteString(a.Manifest.GetPackageName())
		writer.WriteString("\n")
		writer.WriteString("\n")

		writer.WriteString(white(bold("Features\n")))
		for _, v := range a.Features {
			writer.WriteString("* ")
			writer.WriteString(v)
			writer.WriteString("\n")
		}
		writer.WriteString("\n")

		writer.WriteString(white(bold("Permissions\n")))
		for _, v := range a.Permissions {
			writer.WriteString("* ")
			writer.WriteString(v)
			writer.WriteString("\n")
		}
		writer.WriteString("\n")

		writer.WriteString(white(bold("Services\n")))
		for _, v := range a.Services {
			writer.WriteString("* ")
			writer.WriteString(v)
			writer.WriteString("\n")
		}
		writer.WriteString("\n")

		writer.WriteString(white(bold("Activities\n")))
		for _, v := range a.Activities {
			writer.WriteString("* ")
			writer.WriteString(v)
			writer.WriteString("\n")
		}
		writer.WriteString("\n")

		writer.WriteString(white(bold("Receivers\n")))
		for _, v := range a.Receivers {
			writer.WriteString("* ")
			writer.WriteString(v)
			writer.WriteString("\n")
		}
		writer.WriteString("\n")

		writer.WriteString(white(bold("Providers\n")))
		for _, v := range a.Providers {
			writer.WriteString("* ")
			writer.WriteString(v)
			writer.WriteString("\n")
		}
		writer.WriteString("\n")

		if includeFiles {
			writer.WriteString(white(bold("Files\n")))
			for _, v := range a.Files {
				writer.WriteString("* ")
				writer.WriteString(v.GetFullPath())
				writer.WriteString("\n")
			}
			writer.WriteString("\n")

			writer.WriteString(white(bold("File Statistics\n")))
			for k, v := range a.FileStatistics {
				writer.WriteString(bold(fmt.Sprintf("%s: ", k)))
				writer.WriteString(fmt.Sprintf("%d", v))
				writer.WriteString("\n")
			}
			writer.WriteString("\n")

			writer.WriteString(white(bold("Sources\n")))
			for _, v := range a.Sources {
				writer.WriteString("* ")
				writer.WriteString(v.GetFullPath())
				writer.WriteString("\n")
			}
			writer.WriteString("\n")
		}

		return writer.String(), nil
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
