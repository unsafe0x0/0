package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select download option:")
	fmt.Println("1. Download whole playlist as MP3")
	fmt.Println("2. Download whole playlist as MP4")
	fmt.Println("3. Download single video as MP3")
	fmt.Println("4. Download single video as MP4")
	fmt.Print("Enter choice (1-4): ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	var choice int
	switch choiceStr {
	case "1":
		choice = 1
	case "2":
		choice = 2
	case "3":
		choice = 3
	case "4":
		choice = 4
	default:
		fmt.Println("Invalid choice. Exiting.")
		return
	}

	fmt.Print("Enter the YouTube URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)
	if url == "" {
		fmt.Println("URL cannot be empty. Exiting.")
		return
	}

	outDir := "downloads"
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "cannot create downloads dir:", err)
		return
	}

	outputTemplate := filepath.Join(outDir, "%(title)s.%(ext)s")
	var args []string

	switch choice {
	case 1:
		args = []string{
			"--yes-playlist", "-x", "--audio-format", "mp3", "--audio-quality", "0",
			"--output", outputTemplate,
			"--embed-metadata",
			"--add-metadata",
			"--no-mtime",
			url,
		}
	case 2:
		args = []string{
			"--yes-playlist", "--format", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/mp4",
			"--output", outputTemplate,
			"--embed-metadata",
			"--add-metadata",
			"--no-mtime",
			url,
		}
	case 3:
		args = []string{
			"--no-playlist", "-x", "--audio-format", "mp3", "--audio-quality", "0",
			"--output", outputTemplate,
			"--embed-metadata",
			"--add-metadata",
			"--no-mtime",
			url,
		}
	case 4:
		args = []string{
			"--no-playlist", "--format", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/mp4",
			"--output", outputTemplate,
			"--embed-metadata",
			"--add-metadata",
			"--no-mtime",
			url,
		}
	}

	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Starting download...")
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "yt-dlp failed:", err)
		return
	}
	fmt.Println("Download finished. Check the 'downloads' folder.")
}
