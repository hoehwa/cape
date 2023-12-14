package constants

import "os"

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = userHomeDir + "/cape"

// github username and repo in order to fetch snippets.
var Owner = "jipilmuk"
var Repository = "background"

// urls for hosted websites.
const GardenBaseurl = "https://mindulle.github.io/garden"
