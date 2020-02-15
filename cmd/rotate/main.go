package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-image-decode"
	"github.com/aaronland/go-image-rotate"
	"log"
	"os"
)

func main() {

	orientation := flag.String("orientation", "auto", "")
	flag.Parse()

	ctx := context.Background()

	dec, err := decode.NewDecoder(ctx, "image://")

	if err != nil {
		log.Fatal(err)
	}

	uris := flag.Args()

	for _, path := range uris {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer fh.Close()

		im, format, err := dec.Decode(ctx, fh)

		if err != nil {
			log.Fatal(err)
		}

		if format == "jpeg" && *orientation == "auto" {

			_, err := fh.Seek(0, 0)

			if err != nil {
				log.Fatal(err)
			}

			o, err := rotate.GetImageOrientation(ctx, fh)

			if err != nil {
				log.Fatal(err)
			}

			*orientation = o
		}

		im, err = rotate.RotateImageWithOrientation(ctx, im, *orientation)

		if err != nil {
			log.Fatal(err)
		}
	}

}
