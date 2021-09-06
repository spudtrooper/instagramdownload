# Download all instagram photos

1. Open someone's instagram page
2. Run `download.js` until it stops scrolling
3. Copy the console output to a file `some-file.log`
4. Run the following:

```
go run download.go --infile some-file.log
```

The images will appear in `data`.

Or just run (I though I'd do something with the notes)

```
grep 640w some-file.log | awk '{print $3}' | sort | uniq | xargs wget
```