package subcommand

import (
  "fmt"
  "os"
  "strconv"
  "github.com/spf13/cobra"
  "tgl_cli/mylogger"
  "github.com/joho/godotenv"
)

var logger = mylogger.GetLogger()
var verbose bool

type MyToggl struct {
  ApiToken string 
  WorkSpaceId int
}

func readConfig() (me *MyToggl, err error) {
  _ = godotenv.Load()
  fmt.Println(os.Getenv("TGL_APITOKEN"))
  fmt.Println(os.Getenv("TGL_WORKSPACEID"))
  workspaceId, err := strconv.Atoi(os.Getenv("TGL_WORKSPACEID"))
  if err != nil { return nil, err }

  return &MyToggl{ApiToken: os.Getenv("TGL_APITOKEN"), WorkSpaceId: workspaceId}, nil
}

func NewRootCommand() *cobra.Command {
  me := &cobra.Command{
    Use: "tgl_cli",
    Long:  `hogehoge`,
    RunE: func(cmd *cobra.Command, args []string) error {
      return cmd.Help()
    },
    SilenceUsage: true,
		SilenceErrors: true,
  }

  me.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "開発者用デバッグ出力")

  cobra.OnInitialize(func() {  
    if verbose {
      mylogger.SetLevelDebug()
    }
  })

  versionCmd := newVersionCommand()
  me.AddCommand(versionCmd)

  dayCmd := newDayCommand()
  me.AddCommand(dayCmd)
  
  return me
} 

func Execute() {
  rootCmd := NewRootCommand()
  rootCmd.SetOutput(os.Stdout)
  Exit(rootCmd.Execute())
}

func Exit(err error) {
  code := 0
  if err != nil {
    logger.Error(err)
    code = 1
  }
  os.Exit(code)
}