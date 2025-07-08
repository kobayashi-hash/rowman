package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func makeDirs() error {
	var errs []error
	if err := os.MkdirAll("completions/bash", 0755); err != nil {
		errs = append(errs, err)
	}
	if err := os.MkdirAll("completions/zsh", 0755); err != nil {
		errs = append(errs, err)
	}
	if err := os.MkdirAll("completions/fish", 0755); err != nil {
		errs = append(errs, err)
	}
	if err := os.MkdirAll("completions/ps1", 0755); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func generate(command *cobra.Command, commandName string) error {
	var errs []error
	if err := command.GenBashCompletionFileV2(filepath.Join("completions/bash", commandName), true); err != nil {
		errs = append(errs, err)
	}
	if err := command.GenZshCompletionFile(filepath.Join("completions/zsh", commandName)); err != nil {
		errs = append(errs, err)
	}
	if err := command.GenFishCompletionFile(filepath.Join("completions/fish", commandName), true); err != nil {
		errs = append(errs, err)
	}
	if err := command.GenPowerShellCompletionFile(filepath.Join("completions/ps1", commandName)); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func generateCompletions(flag *pflag.FlagSet, commandName string) error {
	command := &cobra.Command{
		Use: "completions",
	}
	command.Flags().AddFlagSet(flag)
	if err := makeDirs(); err != nil {
		return err
	}
	if err := generate(command, commandName); err != nil {
		return err
	}
	return nil
}