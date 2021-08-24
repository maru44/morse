package config

const (
	TYPING_INTERVAL = 400
	SINGLE_PING     = "j"
	TRIPLE_PING     = "k"
	QUIT_PING       = "l"

	SINGLE_LETTER   = "."
	TRIPLE_LETTER   = "-"
	INTERVAL_LETTER = " "
	QUIT_LETTER     = "QUIT"

	DEFAULT_FILE_NAME = "morse"
	DEFAULT_FILE_PATH = "storage/"

	INTERVAL_MODE = "NORMAL"
	/*
		SPEED
		// A INTERVAL_LETTER means character spacing.
		// 3 INTERVAL_LETTER will be converted to space.

		NORMAL
		// 3 INTERVAL_LETTERs means character spacing.
		// 7 INTERVAL_LETTERs will be converted to space.
	*/
)
