package cmd

import (
	"fmt"
	"github.com/dgryski/go-farm"
	"strings"
)

type RedisCmd struct {
	RequestID uint32
	Cmd       string
	Args      []string
}

type RedisCmds struct {
	Cmds      []*RedisCmd
	RequestID uint32
}

// GetFingerprint returns a 32-bit fingerprint of the command and its arguments.
func (cmd *RedisCmd) GetFingerprint() uint32 {
	return farm.Fingerprint32([]byte(fmt.Sprintf("%s-%s", cmd.Cmd, strings.Join(cmd.Args, " "))))
}

// GetKey Returns the key which the command operates on.
//
// TODO: This is a naive implementation which assumes that the first argument is the key.
// This is not true for all commands, however, for now this is only used by the watch manager,
// which as of now only supports a small subset of commands (all of which fit this implementation).
func (cmd *RedisCmd) GetKey() string {
	return cmd.Args[0]
}
