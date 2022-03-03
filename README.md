<h1 align="center">APKrash</h1>
<p>
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>


<p align="center">
  APKrash is an Android APK security analysis toolkit focused on comparing APKs to detect tampering and repackaging.
</p>

<hr>

## Features

- Able to analyze pure Android Manifests, APKs, AABs and JARs.
- Downloads APKs from Google Play Store to perform analysis.
- Analyzes and detects differences on permissions, activities, services, receivers, providers, features and source code.
- With optional dependencies, supports APK extraction, decompiling and conversion to JAR.
- Outputs results as plain text, tables and JSON.

<hr>

## Install

```shell
git clone https://github.com/kamushadenes/apkrash.git
cd apkrash/cmd
go build -o apkrash
```

## Dependencies

Those are optional non-Go dependencies that enable certain features.

### apktool

For the `extract` command

### bundletool

To support `.aab` files

### dex2jar

For the `jar` command

### jadx

For the `decompile` command and for using the `-l` flag to compare source code files

## Usage

```shell
apkrash help
```

```
Android APK security analysis toolkit

Usage:
  apkrash [command]

Available Commands:
  analyze     Analyze an APK or Manifest
  compare     Compares two APKs or Manifests
  completion  Generate the autocompletion script for the specified shell
  decompile   Decompile APK into Java code using jadx
  extract     Extract APK using apktool
  help        Help about any command
  jar         Convert APK to JAR using dex2jar

Flags:
  -c, --color             Output with color (only valid for text mode)
  -e, --email string      Email to use for downloading APKs from Google Play
  -o, --format string     Output format, one of text, json, json_pretty, table (default "text")
  -h, --help              help for apkrash
  -d, --onlyDiffs         Output only diffs (only valid for text mode)
  -w, --password string   Password to use for downloading APKs from Google Play

Use "apkrash [command] --help" for more information about a command.
```

### Analyze an APK or Manifest

```shell
apkrash analyze <file.apk or AndroidManifest.xml>
```

### Compare two APKs

```shell
apkrash compare <file1.apk or AndroidManifest1.xml> <file2.apk or AndroidManifest2.xml>
```

### Decompile an APK using jadx

```shell
apkrash decompile <file.apk> [output_dir]
```

### Extract an APK using apktool

```shell
apkrash extract <file.apk> [output_dir]
```

### Convert APK to JAR using dex2jar

```shell
apkrash jar <file.apk> [output_dir]
```

## Examples

### Compare two APKs showing only diffs with colored output

```shell
apkrash compare -c -d apk1.apk apk2.apk
```

![](.github/images/compare_example.png)

### Analyze an APK and output to JSON (pretty), including files and statistics

```shell
apkrash analyze -o json_pretty -f apk.apk
```

### Compare two APKs and their source code, outputting to JSON

*Note: this may take a few minutes as the APK needs to be decompiled using jadx*

```shell
apkrash compare -o json -f -l apk1.apk apk2.apk
```

## Roadmap

- [x] Add support for AndroidManifest.xml
- [x] Add support for APKs
- [x] Add support for JARs
- [x] Add support for AABs
- [x] Add support for downloading APKs from Play Store
- [ ] Add support for downloading APKs from other stores

## Credits

- Inspired by [AndroCompare](https://github.com/harismuneer/AndroCompare)
- Google Play support provided by [@89z](https://github.com/89z/googleplay)
- Binary Android Manifest support provided by [@shogo82148](https://github.com/shogo82148/androidbinary)

## Show your support

Give a ⭐️ if this project helped you!