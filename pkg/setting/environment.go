package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

func SetEnvironmentInit() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Println("Error loading config file:", err)
		os.Exit(1)
	}

	// Loop through sections
	for _, section := range cfg.Sections() {
		// Loop through keys in each section
		for _, key := range section.Keys() {
			// Set environment variable
			err := os.Setenv(fmt.Sprintf("%s_%s", section.Name(), key.Name()), key.Value())
			if err != nil {
				fmt.Println("Error setting environment variable:", err)
				os.Exit(1)
			}
		}
	}

}
