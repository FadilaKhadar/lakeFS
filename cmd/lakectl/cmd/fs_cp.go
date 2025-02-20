package cmd

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/treeverse/lakefs/pkg/api/apigen"
)

var fsCpCmd = &cobra.Command{
	Use:               "cp <source path URI> <destination path URI>",
	Short:             "Copy object",
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: ValidArgsRepository,
	Run: func(cmd *cobra.Command, args []string) {
		srcPathURI := MustParsePathURI("source path URI", args[0])
		destPathURI := MustParsePathURI("destination path URI", args[1])

		client := getClient()

		resp, err := client.CopyObject(cmd.Context(), destPathURI.Repository, destPathURI.Ref,
			&apigen.CopyObjectParams{
				DestPath: *destPathURI.Path,
			}, apigen.CopyObjectJSONRequestBody{
				SrcPath: *srcPathURI.Path,
			})

		DieOnErrorOrUnexpectedStatusCode(resp, err, http.StatusCreated)
		os.Exit(0)
	},
}

//nolint:gochecknoinits
func init() {
	withPresignFlag(fsCpCmd)
	fsCmd.AddCommand(fsCpCmd)
}
