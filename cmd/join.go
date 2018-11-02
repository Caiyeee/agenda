// Copyright © 2018 Zongyuan Zhan <zhanzy5@mail2.sysu.edu.cn>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhanzongyuan/agenda/agenda"
	"github.com/zhanzongyuan/agenda/cmd/utils"
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Command join other user to a certain meeting you initate.",
	Long: `Please login your account before you  
join other in agenda system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Parse flag
		flag := cmd.Flags()
		title, err := flag.GetString("title")
		if err != nil {
			cmd.Help()
			log.Fatal(err)
		}
		username, err := flag.GetString("username")
		if err != nil {
			cmd.Help()
			log.Fatal(err)
		}

		// Title
		utils.Scan(&title, "[Title]: ", "[Title]: ")

		// Username
		utils.Scan(&username, "[Username]: ", "[Username]: ")

		if err := agenda.JoinUser(title, username); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("User '%s' join meeting '%s' successfully!", username, title)
		}
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")
	joinCmd.Flags().StringP("title", "t", "", "Meeting title you initial")
	joinCmd.Flags().StringP("username", "u", "", "User's name to be joined")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
