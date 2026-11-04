## kinde completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(kinde completion bash)

To load completions for every new session, execute once:

#### Linux:

	kinde completion bash > /etc/bash_completion.d/kinde

#### macOS:

	kinde completion bash > $(brew --prefix)/etc/bash_completion.d/kinde

You will need to start a new shell for this setup to take effect.


```
kinde completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde completion](kinde_completion.md)	 - Generate the autocompletion script for the specified shell

