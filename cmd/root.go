package cmd

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/goeuro/myke/core"
	"github.com/jessevdk/go-flags"
	"io"
	"os"
	"path/filepath"
)

type mykeOpts struct {
	Verbose  int    `short:"v" long:"verbose" description:"verbosity level, <=0 nothing, =3 info, >=5 everything" default:"3"`
	File     string `short:"f" long:"file" description:"yml file to load" default:"myke.yml"`
	DryRun   bool   `short:"n" long:"dry-run" description:"print tasks without running them"`
	Version  bool   `long:"version" description:"print myke version"`
	Template string `long:"template" description:"template file to render"`
	License  bool   `long:"license" description:"show open source licenses"`
	Writer   io.Writer
}

// Exec is CLI entrypoint
func Exec(_args []string) error {
	args := _args[:0]
	for _, x := range _args {
		if len(x) > 0 {
			args = append(args, x)
		}
	}

	var mykeOpts mykeOpts
	mykeOpts.Writer = os.Stdout
	parser := flags.NewNamedParser("myke", flags.IgnoreUnknown|flags.HelpFlag|flags.PassAfterNonOption)
	parser.AddGroup("myke options", "myke options", &mykeOpts)
	parser.Usage = "[--myke-options] [tag/]task [--task-options] ..."
	tasks, err := parser.ParseArgs(args)

	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			fmt.Fprintln(mykeOpts.Writer, flagsErr.Message)
			return nil
		}
		return err
	} else {
		return Action(&mykeOpts, tasks)
	}
}

// Action runs using parsed args
func Action(opts *mykeOpts, tasks []string) error {
	log.SetHandler(&cli.Handler{Writer: opts.Writer, Padding: 0})
	if opts.Verbose <= 0 {
		log.SetLevel(log.FatalLevel)
	} else if opts.Verbose == 1 {
		log.SetLevel(log.ErrorLevel)
	} else if opts.Verbose == 2 {
		log.SetLevel(log.WarnLevel)
	} else if opts.Verbose == 3 {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	if opts.Version {
		return Version(opts)
	} else if opts.License {
		return License(opts)
	} else if len(opts.Template) > 0 {
		return Template(opts)
	} else if len(tasks) > 0 {
		return Run(opts, tasks)
	}

	return List(opts)
}

func loadWorkspace(path string) (core.Workspace, error) {
	if !filepath.IsAbs(path) {
		cwd, _ := os.Getwd()
		path = filepath.Join(cwd, path)
	}
	return core.ParseWorkspace(path)
}
