package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/mipmip/dirty-repo-scanner/src/output"
	"github.com/mipmip/dirty-repo-scanner/src/scanner"
	"github.com/mipmip/dirty-repo-scanner/src/ui"
)

func getDefaultConfigPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	return filepath.Join(configDir, "dirty-repo-scanner", "config.yml")
}

//go:embed config.yml
var defaultConfig string

//go:embed VERSION
var version string

func init() {
	version = strings.TrimSpace(version)
}

func main() {
	app := cli.NewApp()
	app.Name = "dirty-repo-scanner"
	app.Version = version
	app.Usage = "Finds git repos in need of commitment"
	app.EnableBashCompletion = true
	app.CommandNotFound = func(c *cli.Context, cmd string) {
		fmt.Printf("ERROR: Unknown command '%s'\n", cmd)
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Location of config file",
			Value:   getDefaultConfigPath(),
		},

		&cli.BoolFlag{
			Name:    "ignore_dir_errors",
			Aliases: []string{"i"},
			Value:   true,
			Usage:   "Don't halt on errors while finding dirs",
		},
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "show debug output instead of UI",
		},
		&cli.BoolFlag{
			Name:  "list",
			Usage: "print dirty repository paths, one per line, and exit",
		},
		&cli.BoolFlag{
			Name:  "json",
			Usage: "print dirty repositories as a JSON array and exit",
		},
		&cli.BoolFlag{
			Name:  "multiplex",
			Usage: "multiplex mode: Enter runs switch_command and quits (for tmux popups)",
		},
	}
	app.Action = func(c *cli.Context) error {
		listMode := c.Bool("list")
		jsonMode := c.Bool("json")
		debugMode := c.Bool("debug")
		multiplex := c.Bool("multiplex")

		modes := 0
		for _, enabled := range []bool{listMode, jsonMode, debugMode, multiplex} {
			if enabled {
				modes++
			}
		}
		if modes > 1 {
			return fmt.Errorf("choose only one mode: --list, --json, --debug, or --multiplex")
		}

		config, err := scanner.ParseConfigFile(c.String("config"), defaultConfig)
		if c.Args().Len() > 0 {

			fmt.Fprintln(os.Stderr, "Arguments given, skipping config")
			if config == nil {
				config = &scanner.Config{}
			}
			config.ScanDirs.Include = c.Args().Slice()

		} else {
			if err != nil {
				return err
			}
			for i := range config.ScanDirs.Include {
				config.ScanDirs.Include[i] = os.ExpandEnv(config.ScanDirs.Include[i])
			}
			for i := range config.ScanDirs.Exclude {
				config.ScanDirs.Exclude[i] = os.ExpandEnv(config.ScanDirs.Exclude[i])
			}
		}

		if listMode || jsonMode {
			mgs, err := scanner.Scan(config, true)
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			if listMode {
				err = output.List(&buf, mgs)
			} else {
				err = output.JSON(&buf, mgs)
			}
			if err != nil {
				return err
			}
			_, err = os.Stdout.Write(buf.Bytes())
			return err
		}

		if debugMode {
			mgs, err := scanner.Scan(config, c.Bool("ignore_dir_errors"))
			if err != nil {
				return err
			}

			for r, st := range mgs {
				fmt.Printf("%-40s %v\n", r, st.ScanTime)
			}
			return nil
		}
		if multiplex && strings.TrimSpace(config.SwitchCommand) == "" {
			return fmt.Errorf("--multiplex requires 'switch_command' to be set in the config")
		}

		err = ui.Run(config, c.Bool("ignore_dir_errors"), version, multiplex)
		if err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
