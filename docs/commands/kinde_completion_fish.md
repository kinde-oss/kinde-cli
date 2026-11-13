## kinde completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	kinde completion fish | source

To load completions for every new session, execute once:

	kinde completion fish > ~/.config/fish/completions/kinde.fish

You will need to start a new shell for this setup to take effect.


```
kinde completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde completion](kinde_completion.md)	 - Generate the autocompletion script for the specified shell

