/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"fmt"
	"image/jpeg"
	"os"
	"path"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/gen2brain/go-fitz"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-pdf2image/lib"
)

var (
	av = aw.NewArgVars()
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert PDF to image",
	Run:   runConvertCmd,
}

func ErrorHandle(err error) {
	av.Var("error", err.Error())
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args to Alfred: %v", err)
	}
}

func runConvertCmd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		ErrorHandle(fmt.Errorf("page range is required, e.g. \"2, 5-8, 20-, -20, -\""))
		return
	}

	file, _ := cmd.Flags().GetString("file")
	doc, err := fitz.New(file)
	if err != nil {
		ErrorHandle(fmt.Errorf("failed to open PDF: %v", err))
		return
	}
	defer doc.Close()

	ranges, err := lib.ParsePageNumber(args[0], doc.NumPage())
	if err != nil {
		ErrorHandle(fmt.Errorf("failed to parse page range: %v", err))
		return
	}

	for i, r := range *ranges {
		dir := filepath.Dir(file)
		fn := strings.TrimSuffix(path.Base(file), filepath.Ext(path.Base(file)))

		for j := r.Start; j <= r.End; j++ {
			img, err := doc.Image(j - 1)
			if err != nil {
				ErrorHandle(fmt.Errorf("cannot convert page %d: %v", j, err))
				return
			}
			f, err := os.Create(filepath.Join(dir, fmt.Sprintf("%s-part%d-%03d.jpg", fn, i+1, j)))
			if err != nil {
				ErrorHandle(fmt.Errorf("cannot create file: %v", err))
				return
			}

			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			if err != nil {
				ErrorHandle(fmt.Errorf("cannot encode image: %v", err))
				return
			}

			f.Close()
		}
	}

	av.Var("success", "true")
	av.Var("file", path.Base(file))
	av.Var("page", args[0])
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args to Alfred: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.PersistentFlags().StringP("file", "f", "", "file path")
}
