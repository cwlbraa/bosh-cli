package cmd

import (
	"errors"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"
)

type DiffConfigCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewDiffConfigCmd(ui boshui.UI, director boshdir.Director) DiffConfigCmd {
	return DiffConfigCmd{ui: ui, director: director}
}

func (c DiffConfigCmd) Run(opts DiffConfigOpts) error {
	configDiff, err := c.director.DiffConfigByID(opts.FromID, opts.FromContent.Bytes, opts.ToID, opts.ToContent.Bytes)
	if err != nil {
		return err
	}

	diff := NewDiff(configDiff.Diff)

	var headers []boshtbl.Header
	headers = append(headers, boshtbl.NewHeader("From ID"))
	headers = append(headers, boshtbl.NewHeader("To ID"))
	headers = append(headers, boshtbl.NewHeader("Diff"))

	table := boshtbl.Table{
		Content: "",
		Header:  headers,
		Notes:   []string{},

		FillFirstColumn: true,

		Transpose: true,
	}

	var result []boshtbl.Value
	result = append(result, boshtbl.NewValueString(opts.FromID))
	result = append(result, boshtbl.NewValueString(opts.ToID))
	result = append(result, boshtbl.NewValueString(""))
	table.Rows = append(table.Rows, result)

	c.ui.PrintTable(table)

	diff.Print(c.ui)
	return nil
}

func (c DiffConfigCmd) CheckInput(opts DiffConfigOpts) error {
	return errors.New("not implemented")
}
