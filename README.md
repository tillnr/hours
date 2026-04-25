# hours

`hours` is a file format and program suite to continously track hours. This might be 
working, personal, project, practicing or really any other kind of hours

## Didfile

`hours` revolves around the Didfile.

The Didfile is a plain text file consisting of entries consisting of a RFC3339 timestamp
in UTC, a colon, the entry text itself and a newline character.
The timestamps must be going forward in time.

The entry text may be completely freeform. Depending on your needs, you can invent a system 
to make more sense of your entries.

If you work on a project and ticket basis for example, your didfile might look something 
like this:

```
2026-04-21T17:01:23Z:break
2026-04-22T08:05:47Z:comms
2026-04-22T08:25:11Z:projectA#11
2026-04-22T10:01:29Z:projectA#scrum
```

If you use `hours` to track the time you spent at home, it might look more like this:

```
2026-04-23T22:35:53Z:sleep
2026-04-24T07:48:00Z:work
2026-04-24T15:15:00Z:spouse
2026-04-25T11:15:05Z:chores
2026-04-25T12:11:21Z:project
2026-04-25T13:44:58Z:break
```

### Location

The Didfile is located in `$XDG_DATA_HOME` if set, defaulting to the user's home or an
equivalent directory based on system defaults.
If no location can be determined, `did` and `report` will not run succesfully.

You can overwrite the location by setting `$DIDFILE` to your preferred location.

## Starting out

To start using a Didfile to track your hours, just create it:

```
$ touch $DIDFILE
```

After finishing the first activity you want to track (for example some research you did on
a specific topic right after creating your didfile), simply do:

```
did research#golang
```

## did

`did` is the tool for making entries into the Didfile.

It reads the last modified time of your Didfile and uses it as the timestamp for the
entry you are making.
Use it like so:

```
$ did something or the other
```

Presuming you last modified your Didfile on 22.04.2026, 13:52:11 CEST, it will have a new
entry reading `2026-04-22T11:52:11Z:something or the other`.

You can also edit the Didfile by hand, but since `did` uses the last modified time of the 
file to date your entries, be sure to
    
- remember and reset the last modified time after your edit OR
- edit the file AND make your new entry at the same time.

## report

`report` is a tool for making sense of your didfile.

It reads the didfile, erroring if it is not sorted by timestamp or an entry does not
conform to the entry format described above.
Then it sums the time of activity per day and prints an line consisting of day, activity,
and aggregated time spent.

## Recipes

As `report` is a very basic tool right now, it must be enhanced via the command line.

### Sorting

```
$ report | sort -k1
```

### Filtering

Filtering for time spent on dancing

```
$ report | grep "dancing"
```

Filtering out all breaks:

```
$ report | grep -v "break"
```

