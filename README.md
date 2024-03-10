Skriveleif
========

Skriveleif is a small application to search for Norwegian words and open their definition on [ordbøkene.no](https://ordbøkene.no), a service by _Språkrådet_ and _Universitetet i Bergen_.

Skriveleif is inspired by [SRLookup](https://github.com/veloek/srlookup), but written from scratch in Go using the GUI toolkit [Fyne](https://fyne.io/).

Build
-----

Make sure you have the Go tool chain installed, then simply run:
```
$ go build
```
Then you should have an executable named `skriveleif` ready to run.

Run
---

Simply run the executable file `skriveleif`.

Usage
-----

The application client is fairly simple. You just start typing the word and at the 2nd character the application will start fetching suggestions that are presented in the list below. If you press enter the focus will be moved from the input field to the list so you can select the word you're looking for. Pressing enter while selecting a word will open a web browser directed to ordbøkene.no presenting the definition of the word.

Disclaimer
-------

All the information provided in this application is fetched from [ord.uib.no](https://ord.uib.no). I take no responsibility for the availablility nor the content received and presented from their API.

Credits
-------

This application was written mainly for personal use, but I could see that it may have value also for others and therefore I will publish it online.

Copyright 2024 Vegard Løkken <vegard@loekken.org>
