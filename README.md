# tdee
tdee is a simple command line tool to calculate your total daily energy expenditure.

## Requirements
* [Go](https://github.com/golang/go)

## Install
The latest version of bfm can be installed using `go get`.

```
go get -u github.com/LGUG2Z/tdee
```

Make sure `$GOPATH` is set correctly that and that `$GOPATH/bin` is in your `$PATH`.

The `tdee` executable will be installed under the `$GOPATH/bin` directory.

## Overview
`tdee` calculates your total daily energy expenditure by averaging the 
basal metabolic rate estimates given by three of the most commonly 
used formulas: Mifflin-St. Jeor, the original Harris-Benedict formula 
and the revised Harris-Benedict formula.


The raw output of this tool is designed to be piped to [gainit](https://github.com/lgug2z/gainit)
and [loseit](https://github.com/lgug2z/loseit) to calculate a surplus for bulking or a deficit for cutting.

## Usage
Certain information is required for this tool to be able to calculate an estimate of
your total daily energy expenditure: height, weight, age, sex and a lifestyle modifier.
Height and weight can be given in either cm/kg or ft/lb by using either the `--metric`
or `--imperial` flags. The `--height`, `--weight`, `--age` and `--sex` flags are used
to input the required information.

Valid lifestyle modifiers are:
```
1.2   : Sedentary
1.375 : Lightly Active
1.55  : Moderately Active
1.7   : Very Active
1.9   : Extremely Active
```

Valid sex options are:
```
male
female
```

## Examples

Basic use:
```
❯ tdee --metric \
       --height 172 \
       --weight 63.7 \
       --age 29 \
       --sex male \
       --lifestyle 1.375
       
2191 kcal
```

Raw numerical output:
```
❯ tdee --metric \
       --height 172 \
       --weight 63.7 \
       --age 29 \
       --sex male \
       --lifestyle 1.375 \
       --raw
       
2191
```

Raw numerical output piped to `gainit`
```
❯ tdee --metric \
       --height 172 \
       --weight 63.7 \
       --age 29 \
       --sex male \
       --lifestyle 1.375 \
       --raw \
       | xargs gainit --gain 0.5 --kg
       
2741 kcal
```

Raw numerical output piped to `loseit`
```
❯ tdee --metric \
       --height 172 \
       --weight 63.7 \
       --age 29 \
       --sex male \
       --lifestyle 1.375 \
       --raw \
       | xargs loseit --lose 0.4 --kg
       
1751 kcal
```

## Scripting
Given that weight is the only variable that is likely to change regularly, these tools can
be scripted in order to be less verbose:

```bash
# Add this function to your shell rc
gain() {
  tdee --metric \
    --height 172 \
    --weight $1 \
    --age 29 \
    --sex male \
    --lifestyle 1.375 \
    --raw \
    | xargs gainit --gain 0.3 --kg
}

# Use it
❯ gain 63.7

2521 kcal
```

```bash
# Add this function to your shell rc
shred() {
  tdee --metric \
    --height 172 \
    --weight $1 \
    --age 29 \
    --sex male \
    --lifestyle 1.375 \
    --raw \
    | xargs loseit --lose 0.5 --kg
}

# Use it
❯ shred 70.6

1758 kcal
```