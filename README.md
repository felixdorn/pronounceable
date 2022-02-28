# Pronounceable

[![Go Report Card](https://goreportcard.com/badge/github.com/felixdorn/pronounceable)](https://goreportcard.com/report/github.com/felixdorn/pronounceable)

Pronounceable returns a score between 0 and 1 for how pronounceable a string is.

## Installation

```bash
go get github.com/felixdorn/pronounceable
```

## Usage

### Datasets

> A dataset is just a list of words separated by newlines.

We provide a few datasets of various size.

* [10k words](datasets/10k.txt)
* [50k words](datasets/50k.txt)
* [100k words](datasets/100k.txt) **(recommended)**
* [333k words](datasets/333k.txt) (not recommended)

You may download them to include in your project.

### Score Interpretation

* `score > 0.6`: good
* `0.5 > score < 0.6`: meh, it depends
* `score < 0.5`: terrible

Be aware that the larger the dataset, the higher the minimum score. See [how it works](#how-it-works).

```go
package main

import "github.com/felixdorn/pronounceable"

dataset, _ := pronounceable.NewDatasetFromFile("wordlist.txt")

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

For each of these n-grams, we compute their score.

```
log(dataset[n][ngram] / len(dataset[n])) * (5+n)
```

Where `n` is the n-gram length and `dataset[n]` is the map of n-grams.

> We multiply the score by `5 + n` to reward longer n-grams proportionally to their length.

We can then calculate the score for a given word by summing the scores for each n-gram.

```
sum(scores) / 1.5 * len(word)
```

> We multiply the score by 1.5 to penalize longer words.

**Pronounceable** was created by [@afelixdorn](https://twitter.com/afelixdorn) under the [MIT license](LICENSE).
