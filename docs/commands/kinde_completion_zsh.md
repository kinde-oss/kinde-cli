## kinde completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(kinde completion zsh)

To load completions for every new session, execute once:

#### Linux:

	kinde completion zsh > "${fpath[1]}/_kinde"

#### macOS:

	kinde completion zsh > $(brew --prefix)/share/zsh/site-functions/_kinde

You will need to start a new shell for this setup to take effect.


```
kinde completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde completion](kinde_completion.md)	 - Generate the autocompletion script for the specified shell

