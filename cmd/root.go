/*
    EPOS Open Source - Local installation with Docker
    Copyright (C) 2022  EPOS ERIC

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
//file: ./cmd/root.go
package cmd

import (
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = & cobra.Command {
    Use: "[command]",
    Short: "A CLI calculator",
    Long: `A CLI calculator that can add and subtract two numbers.`,
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(deployCmd)
    rootCmd.AddCommand(deleteCmd)
    rootCmd.AddCommand(populateCmd)
}