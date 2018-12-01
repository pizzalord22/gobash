package main

// set the text to show when the help command is ran
func setHelpCommands() map[string]string {
	var hc = make(map[string]string)
	hc["help"] = "help\n----------\nshows this help window\n"
	hc["cd"] = "cd\n----------\nuse cd file/dir to go to a directory, use ../ to go up a file\n"
	hc["ls"] = "ls\n----------\nlist all files and directories in current path\n"
	hc["exec"] = "ls\n----------\nuse this just like how you would use a command line\n"
	hc["exit"] = "exit\n----------\nexits the simulated\n"
	return hc
}
