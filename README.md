# 0

A simple Go CLI tool to download YouTube videos and playlists as MP3 or MP4 files using yt-dlp.

## Features
- Download whole playlists as MP3 (audio only, highest quality)
- Download whole playlists as MP4 (video, highest available quality)
- Download single videos as MP3 (audio only, highest quality)
- Download single videos as MP4 (video, highest available quality)
- Interactive menu for easy selection
- Downloads saved in the `downloads` folder

## Requirements
- Go 1.18+
- [yt-dlp](https://github.com/yt-dlp/yt-dlp) installed and available in your PATH

## Usage
1. Build the program:
   ```sh
   go build -o 0 main.go
   ```
2. Run the program:
   ```sh
   ./0
   ```
3. Follow the prompts:
   - Select the download option (playlist/video, mp3/mp4)
   - Enter the YouTube URL
   - Downloads will be saved in the `downloads` folder

## Example
```
$ ./0
Select download option:
1. Download whole playlist as MP3
2. Download whole playlist as MP4
3. Download single video as MP3
4. Download single video as MP4
Enter choice (1-4): 4
Enter the YouTube URL: https://www.youtube.com/watch?v=example
Starting download...
Download finished. Check the 'downloads' folder.
```

## License
MIT
