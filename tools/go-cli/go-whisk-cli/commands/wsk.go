/*
 * Copyright 2015-2016 IBM Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
    "github.com/spf13/cobra"
)

// WskCmd defines the entry point for the cli.
var WskCmd = &cobra.Command{
    Use:              "wsk",
    Short:            "Whisk cloud computing command line interface.",
    Long:             logoText(),
    SilenceUsage:     true,
    PersistentPreRunE:parseConfigFlags,
}




func init() {
    WskCmd.SetHelpTemplate(`{{with or .Long .Short }}{{.}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`)

    WskCmd.AddCommand(
        actionCmd,
        activationCmd,
        packageCmd,
        ruleCmd,
        triggerCmd,
        sdkCmd,
        propertyCmd,
        namespaceCmd,
        listCmd,
    )

    WskCmd.PersistentFlags().BoolVarP(&flags.global.verbose, "verbose", "v", false, "verbose output")
    WskCmd.PersistentFlags().BoolVarP(&flags.global.debug, "debug", "d", false, "debug level output")
    WskCmd.PersistentFlags().StringVarP(&flags.global.auth, "auth", "u", "", "authorization key")
    WskCmd.PersistentFlags().StringVar(&flags.global.apihost, "apihost", "", "whisk API host")
    WskCmd.PersistentFlags().StringVar(&flags.global.apiversion, "apiversion", "", "whisk API version")
    WskCmd.PersistentFlags().BoolVarP(&flags.global.insecure, "insecure", "i", false, "bypass certificate checking")
}
