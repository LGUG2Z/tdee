package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tdee",
	Short: "Calculate an estimate of your total daily energy expenditure",
	Long: `tdee is a simple command line tool to calculate your total daily energy
expenditure by averaging the basal metabolic rate estimates given by
three of the most commonly used formulas: Mifflin-St. Jeor, the original
Harris-Benedict formula and the revised Harris-Benedict formula.

The raw output of this tool is designed to be piped to gainit and loseit
to calculate a surplus for bulking or a deficit for cutting.

Valid lifestyle modifiers are:

1.2   : Sedentary
1.375 : Lightly Active
1.55  : Moderately Active
1.7   : Very Active
1.9   : Extremely Active
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := Root(f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Root(f Flags) error {
	u, err := determineUnits(f.Imperial, f.Metric)
	if err != nil {
		return err
	}

	if !hasRequiredInformation(f.Height, f.Weight, f.Age, f.Lifestyle, f.Sex) {
		return ErrMissingInformation
	}

	var i Information
	if err := i.FromInput(f.Height, f.Weight, f.Age, f.Lifestyle, f.Sex); err != nil {
		return err
	}

	tdee := i.CalculateTDEE(u)

	if f.Raw {
		fmt.Print(tdee)
	} else {
		fmt.Printf("%v kcal", tdee)
	}

	return nil
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Flags struct {
	Height, Weight, Age, Lifestyle float64
	Sex                            string
	Raw, Metric, Imperial          bool
}

var f Flags

func init() {
	RootCmd.Flags().Float64Var(&f.Height, "height", 0.0, "height")
	RootCmd.Flags().Float64Var(&f.Weight, "weight", 0.0, "weight")
	RootCmd.Flags().Float64Var(&f.Age, "age", 0.0, "age")
	RootCmd.Flags().Float64Var(&f.Lifestyle, "lifestyle", 0.0, "lifestyle modifier")
	RootCmd.Flags().StringVar(&f.Sex, "sex", "", "sex")

	RootCmd.Flags().BoolVarP(&f.Raw, "raw", "r", false, "provide raw output")
	RootCmd.Flags().BoolVarP(&f.Metric, "metric", "m", false, "use metric units (cm/kg)")
	RootCmd.Flags().BoolVarP(&f.Imperial, "imperial", "i", false, "use imperial units (ft/lb)")
}
