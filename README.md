# Pronounceable

Pronounceable returns a score between 0 and 1 for how pronounceable a string is.

It is pretty good, but not perfect.

## Installation

```bash
go get github.com/felixdorn/pronounceable
```

## Usage

### Datasets

We provide a few datasets of various size.

* [10k words](10k.txt)
* [50k words](50k.txt)
* [100k words](100k.txt) **(recommended)**
* [333k words](333k.txt) (not recommended)

You may download them to include in your project.

To create your own, a dataset is just a list of words separated by newlines.

### Score Interpretation

The score is mainly for ranking purposes as a pronounceability of 66% does not really mean anything.

* `score > 0.6`: good
* `0.5 > score < 0.6`: meh, it depends
* `score < 0.5`: terrible

Be aware that the larger the dataset, the higher the minimum score. See [how it works](#how-it-works).

```go
package main

import "github.com/felixdorn/pronounceable"

dataset := pronounceable.NewDataset("wordlist.txt")

dataset.Score("incomprehensibilities") // ~0.66 with 100k words dataset
dataset.Score("hello") // ~0.79 with 100k words dataset
```

## How it works

Given the following dataset:

```
hello
world
```

We generate mono-grams, bi-grams and tri-grams for each word.

We end up with a map that looks like this:

```yaml
monograms:
  h: 1
  e: 1
  l: 3
  o: 2
  w: 1
  r: 1
  d: 1
bigrams:
  he: 1
  el: 1
  ll: 1
  lo: 1
  wo: 1
  or: 1
  rl: 1
  ld: 1
trigrams:
  hel: 1
  ell: 1
  llo: 1
  wor: 1
  orl: 1
  rld: 1
  ...
```

Then for a given word, let's say `chicken`, we generate its mono-grams, bi-grams and tri-grams.

```yaml
monograms:
  c: 2
  h: 1
  i: 1
  k: 1
  e: 1
  n: 1
bigrams:
  ch: 1
  hi: 1
  ik: 1
  ke: 1
  en: 1
trigrams:
  chi: 1
  hic: 1
  ick: 1
  cke: 1
  ken: 1
```

For each of these n-grams, we add the following to the score:

```
log(dataset[n][ngram] / len(dataset[n])) * (5+n)
```

Where `n` is the n-gram length and `dataset[n]` is the map of n-grams.

The score is then divided by 1.5 times the length of the word.

As you can see there are two heuristics (1.5 and 5+n).

The first one is to make sure that the n-grams are in the dataset. The second one is to reward longer n-grams by
increasing the score proportionally to their length.

**Pronounceable** was created by [@afelixdorn](https://twitter.com/afelixdorn) under the [MIT license](LICENSE).