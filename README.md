# yaig

Yet another index generator

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/yaig
```

## Usage

Yaig creates an (inverted index)[https://xlinux.nist.gov/dads/HTML/invertedIndex.html] for a given `"*.txt"` input

## Example

Given a sample input:
The Road Not Taken 
BY ROBERT FROST
Two roads diverged in a yellow wood,
And sorry I could not travel both
And be one traveler, long I stood
And looked down one as far as I could
To where it bent in the undergrowth;

Then took the other, as just as fair,
And having perhaps the better claim,
Because it was grassy and wanted wear;
Though as for that the passing there
Had worn them really about the same,

And both that morning equally lay
In leaves no step had trodden black.
Oh, I kept the first for another day!
Yet knowing how way leads on to way,
I doubted if I should ever come back.

I shall be telling this with a sigh
Somewhere ages and ages hence:
Two roads diverged in a wood, and I—
I took the one less traveled by,
And that has made all the difference.

```bash
map[
    ages:[{1 125} {1 127}]
    bent:[{1 40}] 
    better:[{1 56}]
    black:[{1 91}]
    claim:[{1 57}]
    come:[{1 114}]
    day:[{1 99}]
    difference:[{1 150}]
    diverged:[{1 9} {1 131}]
    doubted:[{1 109}]
    equally:[{1 83}]
    fair:[{1 51}]
    far:[{1 33}]
    frost:[{1 6}]
    grassy:[{1 61}]
    having:[{1 53}]
    just:[{1 49}]
    kept:[{1 94}]
    knowing:[{1 101}]
    lay:[{1 84}]
    leads:[{1 104}]
    leaves:[{1 86}]
    long:[{1 25}]
    looked:[{1 29}]
    morning:[{1 82}]
    oh:[{1 92}]
    passing:[{1 70}]
    really:[{1 75}]
    road:[{1 1}]
    roads:[{1 8} {1 130}]
    robert:[{1 5}]
    shall:[{1 117}]
    sigh:[{1 123}]
    sorry:[{1 15}]
    step:[{1 88}]
    stood:[{1 27}]
    taken:[{1 3}]
    telling:[{1 119}]
    took:[{1 45} {1 138}]
    travel:[{1 19}]
    traveled:[{1 142}]
    traveler:[{1 24}]
    trodden:[{1 90}]
    undergrowth:[{1 43}]
    wanted:[{1 63}]
    way:[{1 103} {1 107}]
    wear:[{1 64}]
    wood:[{1 13} {1 134}]
    worn:[{1 73}]
    yellow:[{1 12}]
]
```

### Testing

``make test``