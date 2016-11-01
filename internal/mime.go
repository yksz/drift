package internal

import "mime"

func init() {
	// audio
	mime.AddExtensionType(".mp3", "audio/mpeg")
	mime.AddExtensionType(".wav", "audio/wav")
	// video
	mime.AddExtensionType(".mp4", "video/mp4")
	// others
	mime.AddExtensionType(".zip", "application/zip")
}
