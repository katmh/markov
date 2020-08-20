# markov

Generates text using a Markov chain

Thanks to this [explainer article](https://mb-14.github.io/tech/2018/10/24/gomarkov.html) and [the `gomarkov` library](https://github.com/mb-14/gomarkov)

I made this to train it on emails on [MIT's `eecs-jobs-announce` mailing list](https://lists.csail.mit.edu/mailman/listinfo/eecs-jobs-announce), so I'll add my data file once I scrub it of potentially personally identifying info. It's just a text file with one entry per line.

My first non-trivial Go project!

## Usage

In the root of this project:

```
# compile into an executable
go build

# (install any packages you don't have)
go get package_name

# train model on data
./markov -train

# get phrases of varying levels of funny
./markov
```

## A few highlights

```
Funded MEng position within the UA Committee on scheduling

Hiring App Developer for pay at the real world? Use your own (including one or for Start-Up!

HackMIT 2020: Building the end.   Virtual.

We are extremely passionate about the semester

The new project is looking for engineers and not required.
```