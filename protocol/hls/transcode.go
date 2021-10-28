package hls

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/lbryio/transcoder/encoder"
	"github.com/lbryio/transcoder/formats"
	log "github.com/sirupsen/logrus"
)

// transcode checks the given source key and will transcode the input source
// based on the transcoding preset found in the source key.
func transcode(filename string, source []byte) ([]byte, error) {
	key, file := filepath.Split(filename)

	splitFn := func(c rune) bool {
		return c == '/'
	}
	root := strings.FieldsFunc(key, splitFn)

	var targetFormats []formats.Format
	switch root[0] {
	case "144":
		targetFormats = append(targetFormats, formats.Format{
			Resolution: formats.SD144,
			Bitrate:    formats.Bitrate{FPS30: 100, FPS60: 160},
		})
	case "480":
		targetFormats = append(targetFormats, formats.Format{
			Resolution: formats.SD480,
			Bitrate:    formats.Bitrate{FPS30: 900, FPS60: 1700},
		})
	default:
		// no transcoding preset found, return the source as it is
		return source, nil
	}

	cfg := encoder.Configure()

	target := encoder.Target{
		Formats: targetFormats,
		Type:    encoder.TargetTypeTS,
	}
	enc, err := encoder.NewEncoder(cfg.Target(target))
	if err != nil {
		return nil, err
	}

	inputDir := filepath.Join(os.TempDir(), key)
	if err := os.MkdirAll(inputDir, os.ModePerm); err != nil {
		return nil, err
	}
	f, err := os.Create(filepath.Join(inputDir, file))
	if _, err := f.Write(source); err != nil {
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())

	outputDir := filepath.Join(inputDir, "out")
	res, err := enc.Encode(f.Name(), outputDir)
	if err != nil {
		return nil, err
	}

	for range res.Progress {
		// drain the channel until its closed
	}
	log.Debug("transcoding done")

	outputFile := filepath.Join(outputDir, file)
	transcodedSourceFile, err := os.Open(outputFile)
	if err != nil {
		return nil, err
	}
	transcodedSource, err := ioutil.ReadAll(transcodedSourceFile)
	if err != nil {
		return nil, err
	}

	return transcodedSource, nil
}
