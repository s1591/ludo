package playlists

import (
	"bufio"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/libretro/ludo/utils"
)

// PlaylistEntry represents a game in a playlist.
type PlaylistEntry struct {
	Path  string // Absolute path of the game on the filesystem
	Name  string // Human readable name of the game, comes from the RDB
	CRC32 uint32 // Checksum of the game, used for deduplication
	LPL   string
}

// Playlist is a list of games, result of scanning for games on the filesystem.
type Playlist []PlaylistEntry

// Playlists is a map of playlists organized per system.
var Playlists = map[string]Playlist{}

// LoadPlaylists loops over lpl files in ~/.ludo/playlists and loads them into
// memory.
func LoadPlaylists() {
	usr, _ := user.Current()
	paths, _ := filepath.Glob(usr.HomeDir + "/.ludo/playlists/*.lpl")

	Playlists = map[string]Playlist{}
	for _, path := range paths {
		path := path

		system := utils.Filename(path)

		file, _ := os.Open(path)
		defer file.Close()
		scanner := bufio.NewScanner(file)

		playlist := Playlist{}
		for {
			more := scanner.Scan()
			if !more {
				break
			}
			var entry PlaylistEntry
			entry.Path = scanner.Text() // path
			scanner.Scan()
			entry.Name = scanner.Text()
			scanner.Scan() // unused
			scanner.Scan() // unused
			scanner.Scan() // CRC
			u64, _ := strconv.ParseUint(strings.Replace(scanner.Text(), "|crc", "", -1), 16, 32)
			entry.CRC32 = uint32(u64)
			scanner.Scan()             // LPL
			entry.LPL = scanner.Text() // LPL

			playlist = append(playlist, entry)
		}
		Playlists[system] = playlist
	}
}

// ExistsInPlaylist checks if a game is already in a playlist.
func ExistsInPlaylist(system, path string, CRC32 uint32) bool {
	for _, entry := range Playlists[system] {
		if entry.Path == path || entry.CRC32 == CRC32 {
			return true
		}
	}
	return false
}