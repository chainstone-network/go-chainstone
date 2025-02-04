// Copyright 2020 The go-chainstone Authors
// This file is part of go-chainstone.
//
// go-chainstone is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-chainstone is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-chainstone. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"

	"github.com/chainstone-network/go-chainstone/cmd/devp2p/internal/v4test"
	"github.com/chainstone-network/go-chainstone/internal/utesting"
	"github.com/chainstone-network/go-chainstone/log"
	"github.com/urfave/cli/v2"
)

var (
	testPatternFlag = &cli.StringFlag{
		Name:  "run",
		Usage: "Pattern of test suite(s) to run",
	}
	testTAPFlag = &cli.BoolFlag{
		Name:  "tap",
		Usage: "Output TAP",
	}
	// These two are specific to the discovery tests.
	testListen1Flag = &cli.StringFlag{
		Name:  "listen1",
		Usage: "IP address of the first tester",
		Value: v4test.Listen1,
	}
	testListen2Flag = &cli.StringFlag{
		Name:  "listen2",
		Usage: "IP address of the second tester",
		Value: v4test.Listen2,
	}
)

func runTests(ctx *cli.Context, tests []utesting.Test) error {
	// Filter test cases.
	if ctx.IsSet(testPatternFlag.Name) {
		tests = utesting.MatchTests(tests, ctx.String(testPatternFlag.Name))
	}
	// Disable logging unless explicitly enabled.
	if !ctx.IsSet("verbosity") && !ctx.IsSet("vmodule") {
		log.Root().SetHandler(log.DiscardHandler())
	}
	// Run the tests.
	var run = utesting.RunTests
	if ctx.Bool(testTAPFlag.Name) {
		run = utesting.RunTAP
	}
	results := run(tests, os.Stdout)
	if utesting.CountFailures(results) > 0 {
		os.Exit(1)
	}
	return nil
}
