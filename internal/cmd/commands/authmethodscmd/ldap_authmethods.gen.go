// Code generated by "make cli"; DO NOT EDIT.
package authmethodscmd

import (
	"errors"
	"fmt"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/authmethods"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/cmd/common"
	"github.com/hashicorp/go-secure-stdlib/strutil"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

func initLdapFlags() {
	flagsOnce.Do(func() {
		extraFlags := extraLdapActionsFlagsMapFunc()
		for k, v := range extraFlags {
			flagsLdapMap[k] = append(flagsLdapMap[k], v...)
		}
	})
}

var (
	_ cli.Command             = (*LdapCommand)(nil)
	_ cli.CommandAutocomplete = (*LdapCommand)(nil)
)

type LdapCommand struct {
	*base.Command

	Func string

	plural string

	extraLdapCmdVars
}

func (c *LdapCommand) AutocompleteArgs() complete.Predictor {
	initLdapFlags()
	return complete.PredictAnything
}

func (c *LdapCommand) AutocompleteFlags() complete.Flags {
	initLdapFlags()
	return c.Flags().Completions()
}

func (c *LdapCommand) Synopsis() string {
	if extra := extraLdapSynopsisFunc(c); extra != "" {
		return extra
	}

	synopsisStr := "auth method"

	synopsisStr = fmt.Sprintf("%s %s", "ldap-type", synopsisStr)

	return common.SynopsisFunc(c.Func, synopsisStr)
}

func (c *LdapCommand) Help() string {
	initLdapFlags()

	var helpStr string
	helpMap := common.HelpMap("auth method")

	switch c.Func {

	default:

		helpStr = c.extraLdapHelpFunc(helpMap)

	}

	// Keep linter from complaining if we don't actually generate code using it
	_ = helpMap
	return helpStr
}

var flagsLdapMap = map[string][]string{

	"create": {"scope-id", "name", "description"},

	"update": {"id", "name", "description", "version"},
}

func (c *LdapCommand) Flags() *base.FlagSets {
	if len(flagsLdapMap[c.Func]) == 0 {
		return c.FlagSet(base.FlagSetNone)
	}

	set := c.FlagSet(base.FlagSetHTTP | base.FlagSetClient | base.FlagSetOutputFormat)
	f := set.NewFlagSet("Command Options")
	common.PopulateCommonFlags(c.Command, f, "ldap-type auth method", flagsLdapMap, c.Func)

	extraLdapFlagsFunc(c, set, f)

	return set
}

func (c *LdapCommand) Run(args []string) int {
	initLdapFlags()

	switch c.Func {
	case "":
		return cli.RunResultHelp

	}

	c.plural = "ldap-type auth method"
	switch c.Func {
	case "list":
		c.plural = "ldap-type auth methods"
	}

	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}

	if strutil.StrListContains(flagsLdapMap[c.Func], "id") && c.FlagId == "" {
		c.PrintCliError(errors.New("ID is required but not passed in via -id"))
		return base.CommandUserError
	}

	var opts []authmethods.Option

	if strutil.StrListContains(flagsLdapMap[c.Func], "scope-id") {
		switch c.Func {

		case "create":
			if c.FlagScopeId == "" {
				c.PrintCliError(errors.New("Scope ID must be passed in via -scope-id or BOUNDARY_SCOPE_ID"))
				return base.CommandUserError
			}

		}
	}

	client, err := c.Client()
	if c.WrapperCleanupFunc != nil {
		defer func() {
			if err := c.WrapperCleanupFunc(); err != nil {
				c.PrintCliError(fmt.Errorf("Error cleaning kms wrapper: %w", err))
			}
		}()
	}
	if err != nil {
		c.PrintCliError(fmt.Errorf("Error creating API client: %w", err))
		return base.CommandCliError
	}
	authmethodsClient := authmethods.NewClient(client)

	switch c.FlagName {
	case "":
	case "null":
		opts = append(opts, authmethods.DefaultName())
	default:
		opts = append(opts, authmethods.WithName(c.FlagName))
	}

	switch c.FlagDescription {
	case "":
	case "null":
		opts = append(opts, authmethods.DefaultDescription())
	default:
		opts = append(opts, authmethods.WithDescription(c.FlagDescription))
	}

	switch c.FlagRecursive {
	case true:
		opts = append(opts, authmethods.WithRecursive(true))
	}

	if c.FlagFilter != "" {
		opts = append(opts, authmethods.WithFilter(c.FlagFilter))
	}

	var version uint32

	switch c.Func {

	case "update":
		switch c.FlagVersion {
		case 0:
			opts = append(opts, authmethods.WithAutomaticVersioning(true))
		default:
			version = uint32(c.FlagVersion)
		}

	}

	if ok := extraLdapFlagsHandlingFunc(c, f, &opts); !ok {
		return base.CommandUserError
	}

	var resp *api.Response
	var item *authmethods.AuthMethod

	var createResult *authmethods.AuthMethodCreateResult

	var updateResult *authmethods.AuthMethodUpdateResult

	switch c.Func {

	case "create":
		createResult, err = authmethodsClient.Create(c.Context, "ldap", c.FlagScopeId, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = createResult.GetResponse()
		item = createResult.GetItem()

	case "update":
		updateResult, err = authmethodsClient.Update(c.Context, c.FlagId, version, opts...)
		if exitCode := c.checkFuncError(err); exitCode > 0 {
			return exitCode
		}
		resp = updateResult.GetResponse()
		item = updateResult.GetItem()

	}

	resp, item, err = executeExtraLdapActions(c, resp, item, err, authmethodsClient, version, opts)
	if exitCode := c.checkFuncError(err); exitCode > 0 {
		return exitCode
	}

	output, err := printCustomLdapActionOutput(c)
	if err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}
	if output {
		return base.CommandSuccess
	}

	switch c.Func {

	}

	switch base.Format(c.UI) {
	case "table":
		c.UI.Output(printItemTable(item, resp))

	case "json":
		if ok := c.PrintJsonItem(resp); !ok {
			return base.CommandCliError
		}
	}

	return base.CommandSuccess
}

func (c *LdapCommand) checkFuncError(err error) int {
	if err == nil {
		return 0
	}
	if apiErr := api.AsServerError(err); apiErr != nil {
		c.PrintApiError(apiErr, fmt.Sprintf("Error from controller when performing %s on %s", c.Func, c.plural))
		return base.CommandApiError
	}
	c.PrintCliError(fmt.Errorf("Error trying to %s %s: %s", c.Func, c.plural, err.Error()))
	return base.CommandCliError
}

var (
	extraLdapActionsFlagsMapFunc = func() map[string][]string { return nil }
	extraLdapSynopsisFunc        = func(*LdapCommand) string { return "" }
	extraLdapFlagsFunc           = func(*LdapCommand, *base.FlagSets, *base.FlagSet) {}
	extraLdapFlagsHandlingFunc   = func(*LdapCommand, *base.FlagSets, *[]authmethods.Option) bool { return true }
	executeExtraLdapActions      = func(_ *LdapCommand, inResp *api.Response, inItem *authmethods.AuthMethod, inErr error, _ *authmethods.Client, _ uint32, _ []authmethods.Option) (*api.Response, *authmethods.AuthMethod, error) {
		return inResp, inItem, inErr
	}
	printCustomLdapActionOutput = func(*LdapCommand) (bool, error) { return false, nil }
)