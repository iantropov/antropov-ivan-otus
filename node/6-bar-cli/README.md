# bar-cli

This app allows you to manage your own bar!

## Installation

```
npm run build
npm install --global .
```

## Using

To add/update/delete a drink you must provide your credentials.

You can get list of drinks:

```
➜  6-bar-cli git:(6-bar-cli) ✗ bar-cli get-drinks
┌──────────┬────────────────────┬──────────┐
│ #        │ Name               │ Volume   │
├──────────┼────────────────────┼──────────┤
│ 1        │ cola               │ 0.5      │
├──────────┼────────────────────┼──────────┤
│ 2        │ new-drink-2        │ 6        │
├──────────┼────────────────────┼──────────┤
│ 3        │ new-drink-3        │ 5        │
└──────────┴────────────────────┴──────────┘
➜  6-bar-cli git:(6-bar-cli) ✗
```

You can check a drink:
```
➜  6-bar-cli git:(6-bar-cli) ✗ bar-cli get-drink cola
┌──────────┬────────────────────┬──────────┐
│ #        │ Name               │ Volume   │
├──────────┼────────────────────┼──────────┤
│ 1        │ cola               │ 0.5      │
└──────────┴────────────────────┴──────────┘
➜  6-bar-cli git:(6-bar-cli) ✗
```

You can a new drink (bar owner's credentials are required):
```
➜  6-bar-cli git:(6-bar-cli) ✗ bar-cli add-drink -c iam:boss -n new-drink-2 -v 5
Here you are! This is your new drink!
┌──────────┬────────────────────┬──────────┐
│ #        │ Name               │ Volume   │
├──────────┼────────────────────┼──────────┤
│ 1        │ new-drink-2        │ 5        │
└──────────┴────────────────────┴──────────┘
➜  6-bar-cli git:(6-bar-cli) ✗
```

You can update an existing drink (bar owner's credentials are required):
```
➜  6-bar-cli git:(6-bar-cli) ✗ bar-cli update-drink 'new-drink-2' -c iam:boss -n old-drink -v 6
Here you are! This is your updated drink!
┌──────────┬────────────────────┬──────────┐
│ #        │ Name               │ Volume   │
├──────────┼────────────────────┼──────────┤
│ 1        │ old-drink          │ 6        │
└──────────┴────────────────────┴──────────┘
➜  6-bar-cli git:(6-bar-cli) ✗
```

Finally, you can delete a drink from the bar (bar owner's credentials are required):
```
➜  6-bar-cli git:(6-bar-cli) ✗ bar-cli delete-drink new-drink-3 -c iam:boss
Here you are! These are your remain drinks:
┌──────────┬────────────────────┬──────────┐
│ #        │ Name               │ Volume   │
├──────────┼────────────────────┼──────────┤
│ 1        │ cola               │ 0.5      │
├──────────┼────────────────────┼──────────┤
│ 2        │ old-drink          │ 6        │
└──────────┴────────────────────┴──────────┘
➜  6-bar-cli git:(6-bar-cli) ✗
```
